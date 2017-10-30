//replacer:generated-file

package boundaryh

// Then looking end of sentence at custom position it may be required to parse some s before.
// First of all going back to beginning of SB5 group is required.
// Moreover, custom position may be in the middle (Close or Sp) of (possible) SB11.
// Or if custom point is ATerm then class of previous SB5 group required for SB7 checking.
// Returns position at which it is safe to begin analysis.
func sForwardSafePositionInString(s string, pos int) (r int) {
	if pos == 0 {
		return 0
	}

	pos = toNextRuneInString(s, pos) // TODO make name similar
	c, pos := sLastSequenceInString(s[:pos])
	if pos == 0 {
		return 0
	}
	r = pos

	switch c {
	case sClassSp, sClassClose: // in the middle of possible SB11
	case sClassATerm: // possible SB7 - require <prevClass>
		if pos > 0 {
			_, pos = sLastSequenceInString(s[:pos])
		}
		return pos
	default:
		return
	}

	for c == sClassSp && pos > 0 {
		c, pos = sLastSequenceInString(s[:pos])
	}
	for c == sClassClose && pos > 0 {
		c, pos = sLastSequenceInString(s[:pos])
	}
	switch c {
	case sClassATerm, sClassSTerm:
		return pos
	default:
		return // Not SB11 sequence
	}
}

func sBackwardSafePositionInString(s string, pos int) int {
	pos = toNextRuneInString(s, pos)
	_, pos = sLastSequenceInString(s[:pos])
	return pos
}

// Computes first sSequence.
// Returns:
// 	"c"   - sSequence class (see "sSequence rules").
// 	"pos" - point to first rune of next sequence (in other words "pos" is length of current sSequence).
func sFirstSequenceInString(s string) (c sClass, pos int) {
	l := len(s)
	if l == 0 {
		c = sClassOther
		return
	}

	if s[0] == crRune { // SB4
		if l > 1 && s[1] == lfRune { // SB3
			return sClassSep, 2
		}
		return sClassSep, 1
	}

	if s[0] == lfRune { // SB4
		return sClassSep, 1
	}

	c, pos = sFirstClassInString(s)

	if c == sClassSep { // SB4
		return
	}

	for pos < len(s) { // SB5
		c0, delta := sFirstClassInString(s[pos:])
		if !c0.isSkip() {
			break
		}
		pos += delta
	}

	return
}

// Computes last sSequence.
// Analogue to sFirstSequenceInString.
// "pos" points to first rune in sequence.
func sLastSequenceInString(s string) (c sClass, pos int) {
	l := len(s)
	if l == 0 {
		c = sClassOther
		return
	}

	pos = l - 1
	if s[pos] == lfRune { // SB4
		c = sClassSep
		if pos > 0 && s[pos-1] == crRune { // SB3
			pos--
		}
		return
	}

	if s[pos] == crRune { // SB4
		c = sClassSep
		return
	}

	c, pos = sLastClassInString(s)
	if !c.isSkip() {
		return
	}

	for pos > 0 { // SB5
		newC, pos1 := sLastClassInString(s[:pos])
		if newC == sClassSep || s[pos-1] == crRune || s[pos-1] == lfRune { // SB4
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
func sParseAfterSATermInString(s string) (closeSp bool, nextClass sClass, pos1, pos2 int) {
	l := len(s)
	for pos1 < l {
		if c, delta := sFirstSequenceInString(s[pos1:]); c == sClassClose {
			closeSp = true
			pos1 += delta
		} else {
			break
		}
	}
	for pos1 < l {
		if c, delta := sFirstSequenceInString(s[pos1:]); c == sClassSp {
			closeSp = true
			pos1 += delta
		} else {
			break
		}
	}
	nextClass, delta := sFirstSequenceInString(s[pos1:])
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
func sDecideAtATermInString(l0 sClass, s string) (stop bool, pos int, newL0 sClass) {
	closeSp, newL0, pos1, pos := sParseAfterSATermInString(s)
	switch {
	case newL0 == sClassSep: // SB11
		stop = true
	case !closeSp && newL0 == sClassNumeric: // SB6
	case (l0 == sClassUpper || l0 == sClassLower) && !closeSp && newL0 == sClassUpper: // SB7
	case newL0 == sClassSContinue || newL0 == sClassATerm || newL0 == sClassSTerm: // SB8a
	case newL0 == sClassLower: // SB8 (part 1)
	case newL0.isSB8(): // SB8 (part 2)
		for pos < len(s) {
			var deltaI int
			newL0, deltaI = sFirstSequenceInString(s[pos:])
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
func sDecideAtSTermInString(s string) (stop bool, pos int, newL0 sClass) {
	_, newL0, i1, pos := sParseAfterSATermInString(s)
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

// SentenceEndInString computes sentence which contains pos-th rune.
// Returns (index of sentence's last rune)+1.
// In other words, returns first sentence's boundary on the right of pos-th rune.
func SentenceEndInString(s string, pos int) int {
	l := len(s)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = sForwardSafePositionInString(s, pos) // go backward to analysis-safe position

	l1 := sClassOther
	for pos < l {
		l0, deltaPos := sFirstSequenceInString(s[pos:])
		pos += deltaPos

		if pos >= l { // end of s
			return l
		}

		switch l0 {
		case sClassSep: // SB4
			return pos
		case sClassATerm:
			stop, deltaPos, newL0 := sDecideAtATermInString(l1, s[pos:])
			pos += deltaPos
			if stop {
				return pos
			}
			l1 = newL0
		case sClassSTerm:
			stop, deltaPos, newL0 := sDecideAtSTermInString(s[pos:])
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

// SentenceBeginInString computes sentence which contains pos-th rune.
// Returns sentence's first rune index.
// In other words, returns first sentence's boundary on the left of pos-th rune.
func SentenceBeginInString(s string, pos int) int {
	l := len(s)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	origPos := pos
	pos = sBackwardSafePositionInString(s, pos)
	//_, pos = sLastSequenceInString(s[:pos+1])
	for pos > 0 {
		c, newPos := sLastSequenceInString(s[:pos])
		switch c {
		case sClassSep:
			return pos
		case sClassATerm:
			l0, _ := sLastSequenceInString(s[:newPos])
			stop, deltaI, _ := sDecideAtATermInString(l0, s[pos:])
			if stop && pos+deltaI <= origPos { // avoid cases with passed "pos" in the middle of SB11 (in such cases boundary may be founded on the right side instead of left).
				return pos + deltaI
			}
		case sClassSTerm:
			stop, deltaI, _ := sDecideAtSTermInString(s[pos:])
			if stop && pos+deltaI <= origPos { // avoid cases with passed "pos" in the middle of SB11 (in such cases boundary may be founded on the right side instead of left).
				return pos + deltaI
			}
		}
		pos = newPos
	}
	return 0
}

// SentenceAtInString computes sentence which contains pos-th rune and return their boundary.
// Sentence may retrieved by "s[r.From:r.To]".
func SentenceAtInString(s string, pos int) Boundary {
	return Boundary{SentenceBeginInString(s, pos), SentenceEndInString(s, pos)}
}

// LastSentenceInString computes last sentence.
// Returns index of sentence's first rune.
// Last sentence may retrieved by "s[r:]".
func LastSentenceInString(s string) int {
	return SentenceBeginInString(s, len(s)-1)
}

// FirstSentenceInString computes first sentence.
// Returns (index of sentence's last rune)+1.
// Result also may be treated as length of the first sentence.
// First sentence may retrieved by "s[:r]".
func FirstSentenceInString(s string) int {
	return SentenceEndInString(s, 0)
}

// SentencesInString computes all sentences and returns theirs boundaries.
func SentencesInString(s string) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(s)) // TODO memory efficient
	for i := 0; i < len(s); {
		length := FirstSentenceInString(s[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}

// SentenceBreaksInString computes all sentences and returns all breaks.
func SentenceBreaksInString(s string) (breaks []int) {
	l := len(s)
	if l == 0 {
		return
	}
	breaks = make([]int, 1, len(s)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstSentenceInString(s[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}

// Then looking end of sentence at custom position it may be required to parse some bytes before.
// First of all going back to beginning of SB5 group is required.
// Moreover, custom position may be in the middle (Close or Sp) of (possible) SB11.
// Or if custom point is ATerm then class of previous SB5 group required for SB7 checking.
// Returns position at which it is safe to begin analysis.
func sForwardSafePosition(bytes []byte, pos int) (r int) {
	if pos == 0 {
		return 0
	}

	pos = toNextRune(bytes, pos) // TODO make name similar
	c, pos := sLastSequence(bytes[:pos])
	if pos == 0 {
		return 0
	}
	r = pos

	switch c {
	case sClassSp, sClassClose: // in the middle of possible SB11
	case sClassATerm: // possible SB7 - require <prevClass>
		if pos > 0 {
			_, pos = sLastSequence(bytes[:pos])
		}
		return pos
	default:
		return
	}

	for c == sClassSp && pos > 0 {
		c, pos = sLastSequence(bytes[:pos])
	}
	for c == sClassClose && pos > 0 {
		c, pos = sLastSequence(bytes[:pos])
	}
	switch c {
	case sClassATerm, sClassSTerm:
		return pos
	default:
		return // Not SB11 sequence
	}
}

func sBackwardSafePosition(bytes []byte, pos int) int {
	pos = toNextRune(bytes, pos)
	_, pos = sLastSequence(bytes[:pos])
	return pos
}

// Computes first sSequence.
// Returns:
// 	"c"   - sSequence class (see "sSequence rules").
// 	"pos" - point to first rune of next sequence (in other words "pos" is length of current sSequence).
func sFirstSequence(bytes []byte) (c sClass, pos int) {
	l := len(bytes)
	if l == 0 {
		c = sClassOther
		return
	}

	if bytes[0] == crRune { // SB4
		if l > 1 && bytes[1] == lfRune { // SB3
			return sClassSep, 2
		}
		return sClassSep, 1
	}

	if bytes[0] == lfRune { // SB4
		return sClassSep, 1
	}

	c, pos = sFirstClass(bytes)

	if c == sClassSep { // SB4
		return
	}

	for pos < len(bytes) { // SB5
		c0, delta := sFirstClass(bytes[pos:])
		if !c0.isSkip() {
			break
		}
		pos += delta
	}

	return
}

// Computes last sSequence.
// Analogue to sFirstSequence.
// "pos" points to first rune in sequence.
func sLastSequence(bytes []byte) (c sClass, pos int) {
	l := len(bytes)
	if l == 0 {
		c = sClassOther
		return
	}

	pos = l - 1
	if bytes[pos] == lfRune { // SB4
		c = sClassSep
		if pos > 0 && bytes[pos-1] == crRune { // SB3
			pos--
		}
		return
	}

	if bytes[pos] == crRune { // SB4
		c = sClassSep
		return
	}

	c, pos = sLastClass(bytes)
	if !c.isSkip() {
		return
	}

	for pos > 0 { // SB5
		newC, pos1 := sLastClass(bytes[:pos])
		if newC == sClassSep || bytes[pos-1] == crRune || bytes[pos-1] == lfRune { // SB4
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
func sParseAfterSATerm(bytes []byte) (closeSp bool, nextClass sClass, pos1, pos2 int) {
	l := len(bytes)
	for pos1 < l {
		if c, delta := sFirstSequence(bytes[pos1:]); c == sClassClose {
			closeSp = true
			pos1 += delta
		} else {
			break
		}
	}
	for pos1 < l {
		if c, delta := sFirstSequence(bytes[pos1:]); c == sClassSp {
			closeSp = true
			pos1 += delta
		} else {
			break
		}
	}
	nextClass, delta := sFirstSequence(bytes[pos1:])
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
func sDecideAtATerm(l0 sClass, bytes []byte) (stop bool, pos int, newL0 sClass) {
	closeSp, newL0, pos1, pos := sParseAfterSATerm(bytes)
	switch {
	case newL0 == sClassSep: // SB11
		stop = true
	case !closeSp && newL0 == sClassNumeric: // SB6
	case (l0 == sClassUpper || l0 == sClassLower) && !closeSp && newL0 == sClassUpper: // SB7
	case newL0 == sClassSContinue || newL0 == sClassATerm || newL0 == sClassSTerm: // SB8a
	case newL0 == sClassLower: // SB8 (part 1)
	case newL0.isSB8(): // SB8 (part 2)
		for pos < len(bytes) {
			var deltaI int
			newL0, deltaI = sFirstSequence(bytes[pos:])
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
func sDecideAtSTerm(bytes []byte) (stop bool, pos int, newL0 sClass) {
	_, newL0, i1, pos := sParseAfterSATerm(bytes)
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

// SentenceEnd computes sentence which contains pos-th rune.
// Returns (index of sentence's last rune)+1.
// In other words, returns first sentence's boundary on the right of pos-th rune.
func SentenceEnd(bytes []byte, pos int) int {
	l := len(bytes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = sForwardSafePosition(bytes, pos) // go backward to analysis-safe position

	l1 := sClassOther
	for pos < l {
		l0, deltaPos := sFirstSequence(bytes[pos:])
		pos += deltaPos

		if pos >= l { // end of bytes
			return l
		}

		switch l0 {
		case sClassSep: // SB4
			return pos
		case sClassATerm:
			stop, deltaPos, newL0 := sDecideAtATerm(l1, bytes[pos:])
			pos += deltaPos
			if stop {
				return pos
			}
			l1 = newL0
		case sClassSTerm:
			stop, deltaPos, newL0 := sDecideAtSTerm(bytes[pos:])
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

// SentenceBegin computes sentence which contains pos-th rune.
// Returns sentence's first rune index.
// In other words, returns first sentence's boundary on the left of pos-th rune.
func SentenceBegin(bytes []byte, pos int) int {
	l := len(bytes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	origPos := pos
	pos = sBackwardSafePosition(bytes, pos)
	//_, pos = sLastSequence(bytes[:pos+1])
	for pos > 0 {
		c, newPos := sLastSequence(bytes[:pos])
		switch c {
		case sClassSep:
			return pos
		case sClassATerm:
			l0, _ := sLastSequence(bytes[:newPos])
			stop, deltaI, _ := sDecideAtATerm(l0, bytes[pos:])
			if stop && pos+deltaI <= origPos { // avoid cases with passed "pos" in the middle of SB11 (in such cases boundary may be founded on the right side instead of left).
				return pos + deltaI
			}
		case sClassSTerm:
			stop, deltaI, _ := sDecideAtSTerm(bytes[pos:])
			if stop && pos+deltaI <= origPos { // avoid cases with passed "pos" in the middle of SB11 (in such cases boundary may be founded on the right side instead of left).
				return pos + deltaI
			}
		}
		pos = newPos
	}
	return 0
}

// SentenceAt computes sentence which contains pos-th rune and return their boundary.
// Sentence may retrieved by "bytes[r.From:r.To]".
func SentenceAt(bytes []byte, pos int) Boundary {
	return Boundary{SentenceBegin(bytes, pos), SentenceEnd(bytes, pos)}
}

// LastSentence computes last sentence.
// Returns index of sentence's first rune.
// Last sentence may retrieved by "bytes[r:]".
func LastSentence(bytes []byte) int {
	return SentenceBegin(bytes, len(bytes)-1)
}

// FirstSentence computes first sentence.
// Returns (index of sentence's last rune)+1.
// Result also may be treated as length of the first sentence.
// First sentence may retrieved by "bytes[:r]".
func FirstSentence(bytes []byte) int {
	return SentenceEnd(bytes, 0)
}

// Sentences computes all sentences and returns theirs boundaries.
func Sentences(bytes []byte) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(bytes)) // TODO memory efficient
	for i := 0; i < len(bytes); {
		length := FirstSentence(bytes[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}

// SentenceBreaks computes all sentences and returns all breaks.
func SentenceBreaks(bytes []byte) (breaks []int) {
	l := len(bytes)
	if l == 0 {
		return
	}
	breaks = make([]int, 1, len(bytes)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstSentence(bytes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}
