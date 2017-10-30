package boundaryh

//replacer:ignore
// TODO replace windows path separator
//go:generate go run $GOPATH\src\github.com\apaxa-go\generator\replacer\main.go -- $GOFILE

// sSequence rules:
// 	        CR LF        => Sep
// 	      (CR | LF)      => Sep
//	         Sep         => Sep
// 	X (Extend | Format)* => X (where X is not Sep)

//replacer:replace
//replacer:old InRunes	[]rune	runes
//replacer:new InString	string	s
//replacer:new ""		[]byte	bytes

// Then looking end of sentence at custom position it may be required to parse some runes before.
// First of all going back to beginning of SB5 group is required.
// Moreover, custom position may be in the middle (Close or Sp) of (possible) SB11.
// Or if custom point is ATerm then class of previous SB5 group required for SB7 checking.
// Returns position at which it is safe to begin analysis.
func sForwardSafePositionInRunes(runes []rune, pos int) (r int) {
	if pos == 0 {
		return 0
	}

	pos = toNextRuneInRunes(runes, pos) // TODO make name similar
	c, pos := sLastSequenceInRunes(runes[:pos])
	if pos == 0 {
		return 0
	}
	r = pos

	switch c {
	case sClassSp, sClassClose: // in the middle of possible SB11
	case sClassATerm: // possible SB7 - require <prevClass>
		if pos > 0 {
			_, pos = sLastSequenceInRunes(runes[:pos])
		}
		return pos
	default:
		return
	}

	for c == sClassSp && pos > 0 {
		c, pos = sLastSequenceInRunes(runes[:pos])
	}
	for c == sClassClose && pos > 0 {
		c, pos = sLastSequenceInRunes(runes[:pos])
	}
	switch c {
	case sClassATerm, sClassSTerm:
		return pos
	default:
		return // Not SB11 sequence
	}
}

func sBackwardSafePositionInRunes(runes []rune, pos int) int {
	pos = toNextRuneInRunes(runes, pos)
	_, pos = sLastSequenceInRunes(runes[:pos])
	return pos
}

// Computes first sSequence.
// Returns:
// 	"c"   - sSequence class (see "sSequence rules").
// 	"pos" - point to first rune of next sequence (in other words "pos" is length of current sSequence).
func sFirstSequenceInRunes(runes []rune) (c sClass, pos int) {
	l := len(runes)
	if l == 0 {
		c = sClassOther
		return
	}

	if runes[0] == crRune { // SB4
		if l > 1 && runes[1] == lfRune { // SB3
			return sClassSep, 2
		}
		return sClassSep, 1
	}

	if runes[0] == lfRune { // SB4
		return sClassSep, 1
	}

	c, pos = sFirstClassInRunes(runes)

	if c == sClassSep { // SB4
		return
	}

	for pos < len(runes) { // SB5
		c0, delta := sFirstClassInRunes(runes[pos:])
		if !c0.isSkip() {
			break
		}
		pos += delta
	}

	return
}

// Computes last sSequence.
// Analogue to sFirstSequenceInRunes.
// "pos" points to first rune in sequence.
func sLastSequenceInRunes(runes []rune) (c sClass, pos int) {
	l := len(runes)
	if l == 0 {
		c = sClassOther
		return
	}

	pos = l - 1
	if runes[pos] == lfRune { // SB4
		c = sClassSep
		if pos > 0 && runes[pos-1] == crRune { // SB3
			pos--
		}
		return
	}

	if runes[pos] == crRune { // SB4
		c = sClassSep
		return
	}

	c, pos = sLastClassInRunes(runes)
	if !c.isSkip() {
		return
	}

	for pos > 0 { // SB5
		newC, pos1 := sLastClassInRunes(runes[:pos])
		if newC == sClassSep || runes[pos-1] == crRune || runes[pos-1] == lfRune { // SB4
			//pos++
			return
		}
		c = newC
		pos = pos1
		if !newC.isSkip() {
			return
		}
	}
	return c, 0
}

// Computes required classes and possible positions for continue after founded ATerm or STerm.
func sParseAfterSATermInRunes(runes []rune) (closeSp bool, nextClass sClass, pos1, pos2 int) {
	l := len(runes)
	for pos1 < l {
		if c, delta := sFirstSequenceInRunes(runes[pos1:]); c == sClassClose {
			closeSp = true
			pos1 += delta
		} else {
			break
		}
	}
	for pos1 < l {
		if c, delta := sFirstSequenceInRunes(runes[pos1:]); c == sClassSp {
			closeSp = true
			pos1 += delta
		} else {
			break
		}
	}
	nextClass, delta := sFirstSequenceInRunes(runes[pos1:])
	pos2 = pos1 + delta
	return
}

// newL0 used as l0 by parent function. It used only if stop==false and is undefined if stop==true.
//
//  Break?  |  <l0>   ATerm Close* Sp* <lastClass>        | Ruler
//   yes    |         ATerm Close* Sp* ParaSep            | SB11
//   no     |         ATerm -      -   Numeric            | SB6
//   no     | Upper   ATerm -      -   Upper              | SB7
//   no     | Lower   ATerm -      -   Upper              | SB7
//   no     |         ATerm Close* Sp* Lower              | SB8
//   no     |         ATerm Close* Sp* isSB8+       Lower | SB8
//   no     |         ATerm Close* Sp* SContinue          | SB8a
//   no     |         ATerm Close* Sp* ATerm              | SB8a
//   no     |         ATerm Close* Sp* STerm              | SB8a
//   yes    |         ATerm Close* Sp*                    | SB11
func sDecideAtATermInRunes(l0 sClass, runes []rune) (stop bool, pos int, newL0 sClass) {
	closeSp, newL0, pos1, pos := sParseAfterSATermInRunes(runes)
	switch {
	case newL0 == sClassSep: // SB11
		stop = true
	case !closeSp && newL0 == sClassNumeric: // SB6
	case (l0 == sClassUpper || l0 == sClassLower) && !closeSp && newL0 == sClassUpper: // SB7
	case newL0 == sClassSContinue || newL0 == sClassATerm || newL0 == sClassSTerm: // SB8a
	case newL0 == sClassLower: // SB8 (part 1)
	case newL0.isSB8(): // SB8 (part 2)
		for pos < len(runes) {
			var deltaI int
			newL0, deltaI = sFirstSequenceInRunes(runes[pos:])
			pos += deltaI
			if newL0 == sClassLower {
				return
			}
			if !newL0.isSB8() {
				break
			}
		}
		stop = true
		pos = pos1
	default: // SB11
		stop = true
		pos = pos1
	}
	return
}

// newL0 used as l0 by parent function. It used only if stop==false and is undefined if stop==true.
//
//  Break?  | STerm Close* Sp* some-type? | Ruler
//   yes    | STerm Close* Sp* ParaSep    | SB11
//   no     | STerm Close* Sp* SContinue  | SB8a
//   no     | STerm Close* Sp* ATerm      | SB8a
//   no     | STerm Close* Sp* STerm      | SB8a
//   yes    | STerm Close* Sp*            | SB11
func sDecideAtSTermInRunes(runes []rune) (stop bool, pos int, newL0 sClass) {
	_, newL0, i1, pos := sParseAfterSATermInRunes(runes)
	switch newL0 {
	case sClassSep: // SB11
		stop = true
	case sClassSContinue, sClassATerm, sClassSTerm: // SB8a
	default:
		stop = true // SB11
		pos = i1
	}
	return
}

// SentenceEndInRunes computes sentence which contains pos-th rune.
// Returns (index of sentence's last rune)+1.
// In other words, returns first sentence's boundary on the right of pos-th rune.
func SentenceEndInRunes(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = sForwardSafePositionInRunes(runes, pos) // go backward to analysis-safe position

	l1 := sClassOther
	for pos < l {
		l0, deltaPos := sFirstSequenceInRunes(runes[pos:])
		pos += deltaPos

		if pos >= l { // end of runes
			return l
		}

		switch l0 {
		case sClassSep: // SB4
			return pos
		case sClassATerm:
			stop, deltaPos, newL0 := sDecideAtATermInRunes(l1, runes[pos:])
			pos += deltaPos
			if stop {
				return pos
			}
			l1 = newL0
		case sClassSTerm:
			stop, deltaPos, newL0 := sDecideAtSTermInRunes(runes[pos:])
			pos += deltaPos
			if stop {
				return pos
			}
			l1 = newL0
		default:
			l1 = l0
		}
	}
	return l
}

// SentenceBeginInRunes computes sentence which contains pos-th rune.
// Returns sentence's first rune index.
// In other words, returns first sentence's boundary on the left of pos-th rune.
func SentenceBeginInRunes(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	origPos := pos
	pos = sBackwardSafePositionInRunes(runes, pos)
	//_, pos = sLastSequenceInRunes(runes[:pos+1])
	for pos > 0 {
		c, newPos := sLastSequenceInRunes(runes[:pos])
		switch c {
		case sClassSep:
			return pos
		case sClassATerm:
			l0, _ := sLastSequenceInRunes(runes[:newPos])
			stop, deltaI, _ := sDecideAtATermInRunes(l0, runes[pos:])
			if stop && pos+deltaI <= origPos { // avoid cases with passed "pos" in the middle of SB11 (in such cases boundary may be founded on the right side instead of left).
				return pos + deltaI
			}
		case sClassSTerm:
			stop, deltaI, _ := sDecideAtSTermInRunes(runes[pos:])
			if stop && pos+deltaI <= origPos { // avoid cases with passed "pos" in the middle of SB11 (in such cases boundary may be founded on the right side instead of left).
				return pos + deltaI
			}
		}
		pos = newPos
	}
	return 0
}

// SentenceAtInRunes computes sentence which contains pos-th rune and return their boundary.
// Sentence may retrieved by "runes[r.From:r.To]".
func SentenceAtInRunes(runes []rune, pos int) Boundary {
	return Boundary{SentenceBeginInRunes(runes, pos), SentenceEndInRunes(runes, pos)}
}

// LastSentenceInRunes computes last sentence.
// Returns index of sentence's first rune.
// Last sentence may retrieved by "runes[r:]".
func LastSentenceInRunes(runes []rune) int {
	return SentenceBeginInRunes(runes, len(runes)-1)
}

// FirstSentenceInRunes computes first sentence.
// Returns (index of sentence's last rune)+1.
// Result also may be treated as length of the first sentence.
// First sentence may retrieved by "runes[:r]".
func FirstSentenceInRunes(runes []rune) int {
	return SentenceEndInRunes(runes, 0)
}

// SentencesInRunes computes all sentences and returns theirs boundaries.
func SentencesInRunes(runes []rune) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(runes)) // TODO memory efficient
	for i := 0; i < len(runes); {
		length := FirstSentenceInRunes(runes[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}

// SentenceBreaksInRunes computes all sentences and returns all breaks.
func SentenceBreaksInRunes(runes []rune) (breaks []int) {
	l := len(runes)
	if l == 0 {
		return
	}
	breaks = make([]int, 1, len(runes)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstSentenceInRunes(runes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}
