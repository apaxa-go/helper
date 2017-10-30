package boundaryh

//replacer:ignore
// TODO replace windows path separator
//go:generate go run $GOPATH\src\github.com\apaxa-go\generator\replacer\main.go -- $GOFILE

// wSequence rules:
// -         CR LF              => NewLine
// -        NewLine             => NewLine
// -   ZWJ Glue_After_Zwj       => Glue_After_Zwj
// -   ZWJ EBG                  => EBG
// - X (Extend | Format | ZWJ)* => X
// wSequence allow avoid rules WB3, WB3c & WB4.

// Returns if there is a break between l0 and r0.
func wDecision(l1, l0 wClass, lOddRI bool, r0, r1 wClass) bool {
	// TODO translate to single boolean expression?
	switch {
	case l0 == wClassNewline || r0 == wClassNewline: // WB3a && WB3b
		return true
		//case l0 == wClassZWJ && (r0 == wClassGlueAfterZWJ || r0 == wClassEBG): // WB3c, but now covered by GetClassSkip
	case r0.isWB4(): // WB4 (part 1)
	case l0.isAHLetter() && r0.isAHLetter(): // WB5
	case l0.isAHLetter() && (r0 == wClassMidLetter || r0.isMidNumLetQ()) && r1.isAHLetter(): // WB6
	case l1.isAHLetter() && (l0 == wClassMidLetter || l0.isMidNumLetQ()) && r0.isAHLetter(): // WB7
	case l0 == wClassHebrewLetter && r0 == wClassSingleQuote: // WB7a
	case l0 == wClassHebrewLetter && r0 == wClassDoubleQuote && r1 == wClassHebrewLetter: // WB7b
	case l1 == wClassHebrewLetter && l0 == wClassDoubleQuote && r0 == wClassHebrewLetter: // WB7b
	case l0 == wClassNumeric && r0 == wClassNumeric: // WB8
	case l0.isAHLetter() && r0 == wClassNumeric: // WB9
	case l0 == wClassNumeric && r0.isAHLetter(): // WB10
	case l1 == wClassNumeric && (l0 == wClassMidNum || l0.isMidNumLetQ()) && r0 == wClassNumeric: // WB11
	case l0 == wClassNumeric && (r0 == wClassMidNum || r0.isMidNumLetQ()) && r1 == wClassNumeric: // WB12
	case l0 == wClassKatakana && r0 == wClassKatakana: // WB13
	case (l0.isAHLetter() || l0 == wClassNumeric || l0 == wClassKatakana || l0 == wCLassExtendNumLet) && r0 == wCLassExtendNumLet: // WB13a
	case l0 == wCLassExtendNumLet && (r0.isAHLetter() || r0 == wClassNumeric || r0 == wClassKatakana): // WB13b
	case (l0 == wClassEBase || l0 == wClassEBG) && r0 == wClassEModifier: // WB14
	case lOddRI && r0 == wClassRI: // WB15 & WB16
	default:
		return true
	}
	return false
}

//replacer:replace
//replacer:old InRunes	[]rune	runes
//replacer:new InString	string	s
//replacer:new ""		[]byte	bytes

// Computes first wSequence.
// Returns:
// 	"c"   - wSequence class (see "wSequence rules").
// 	"pos" - point to first rune of next sequence (in other words "pos" is length of current wSequence).
func wFirstSequenceInRunes(runes []rune) (c wClass, pos int) {
	l := len(runes)
	if l == 0 {
		return wClassOther, 0
	}

	if l > 1 && runes[0] == crRune && runes[1] == lfRune { // WB3
		return wClassNewline, 2
	}

	c, pos = wFirstClassInRunes(runes)
	if l == 1 {
		return
	}

	if c == wClassZWJ {
		c1, delta := wFirstClassInRunes(runes[pos:])
		if c1 == wClassGlueAfterZWJ || c1 == wClassEBG { // WB3c
			c = c1
			pos += delta
		}
	}

	if c == wClassNewline {
		return
	}

	l1 := wClassOther
	l1Pos := 0
	for pos < l { // WB4
		l0, delta := wFirstClassInRunes(runes[pos:])

		switch {
		case l1 == wClassZWJ && (l0 == wClassGlueAfterZWJ || l0 == wClassEBG): // preserve following WB3c
			pos = l1Pos
			return
		case !l0.isWB4():
			return
		}
		l1 = l0
		l1Pos = pos
		pos += delta
	}

	return
}

// Computes last wSequence.
// Analogue to wFirstSequenceInRunes.
// "pos" points to first rune in sequence.
func wLastSequenceInRunes(runes []rune) (c wClass, pos int) {
	l := len(runes)
	if l == 0 {
		return wClassOther, 0
	}

	if l > 1 && runes[l-2] == crRune && runes[l-1] == lfRune { // WB3
		return wClassNewline, l - 2
	}

	c, pos = wLastClassInRunes(runes)
	if pos == 0 {
		return
	}

	for pos > 0 && c.isWB4() { // WB4
		c, pos = wLastClassInRunes(runes[:pos])
	}

	if (c == wClassGlueAfterZWJ || c == wClassEBG) && pos > 0 { // WB3c
		c1, pos1 := wLastClassInRunes(runes[:pos])
		if c1 == wClassZWJ {
			pos = pos1
		}
	}

	return
}

// Returns position at which it is safe to begin analysis.
func wSequenceBeginInRunes(runes []rune, pos int) int {
	pos = toRuneBeginInRunes(runes, pos)

	if pos == 0 {
		return 0
	}

	c, cDelta := wFirstClassInRunes(runes[pos:])
	l0, l0Pos := wLastClassInRunes(runes[:pos])
	r0 := wClassOther
	if pos+cDelta < len(runes) {
		r0, _ = wFirstClassInRunes(runes[pos+cDelta:])
	}

	switch {
	case runes[pos-1] == crRune && runes[pos] == lfRune: // WB3
		pos--
	case c == wClassZWJ && (r0 == wClassGlueAfterZWJ || r0 == wClassEBG): // WB3c
	case l0 == wClassZWJ && (c == wClassGlueAfterZWJ || c == wClassEBG): // WB3c
		pos = l0Pos
	case c.isWB4(): // WB4
		for pos > 0 {
			c, l0Pos = wLastClassInRunes(runes[:pos])
			if c == wClassNewline {
				break
			}
			pos = l0Pos
			if !c.isWB4() {
				break
			}
		}
	}
	return pos
}

// True if l0 is RI and it opens RI sequence in string <runes..., l0, ...> (may be joined with next RI).
func wIsOpenRIInRunes(runes []rune, l1, l0 wClass) (r bool) {
	r = l0 == wClassRI
	if !r {
		return
	}
	r = l1 != wClassRI
	if r {
		return
	}
	for len(runes) > 0 {
		c, pos := wLastSequenceInRunes(runes)
		if c != wClassRI {
			break
		}
		r = !r
		runes = runes[:pos]
	}
	return
}

// runes must be valid (len>1).
// l0Pos must be valid (in runes; really begin of sequence).
func wordEndInRunes(runes []rune, l0Pos int) int {
	l := len(runes)

	l1, l1Pos := wLastSequenceInRunes(runes[:l0Pos])
	l0, r0Delta := wFirstSequenceInRunes(runes[l0Pos:])
	lOddRI := wIsOpenRIInRunes(runes[:l1Pos], l1, l0)
	r0Pos := l0Pos + r0Delta
	r0, r1Delta := wFirstSequenceInRunes(runes[r0Pos:])

	for r0Pos < l {
		r1, r2Delta := wFirstSequenceInRunes(runes[r0Pos+r1Delta:])
		if wDecision(l1, l0, lOddRI, r0, r1) {
			return r0Pos
		}
		l1 = l0
		l0 = r0
		r0 = r1
		r0Pos += r1Delta
		r1Delta = r2Delta
		lOddRI = l0 == wClassRI && !lOddRI
	}
	return l
}

// WordEndInRunes computes word which contains pos-th rune.
// Returns (index of word's last rune)+1.
// In other words, returns first word's boundary on the right of pos-th rune.
func WordEndInRunes(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = wSequenceBeginInRunes(runes, pos)

	return wordEndInRunes(runes, pos)
}

// runes must be valid (len>1).
// r0Pos must be valid (in runes; really begin of sequence).
func wordBeginInRunes(runes []rune, r0Pos int) int {
	r0, r1Delta := wFirstSequenceInRunes(runes[r0Pos:])
	r1, _ := wFirstSequenceInRunes(runes[r0Pos+r1Delta:])
	l0, l0Pos := wLastSequenceInRunes(runes[:r0Pos])

	for r0Pos >= 0 {
		l1, l1Pos := wLastSequenceInRunes(runes[:l0Pos])
		lOddRI := wIsOpenRIInRunes(runes[:l1Pos], l1, l0)
		if wDecision(l1, l0, lOddRI, r0, r1) {
			return r0Pos
		}
		r1 = r0
		r0 = l0
		l0 = l1
		r0Pos = l0Pos
		l0Pos = l1Pos
	}

	return 0
}

// WordBeginInRunes computes word which contains pos-th rune.
// Returns word's first rune index.
// In other words, returns first word's boundary on the left of pos-th rune.
func WordBeginInRunes(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	pos = wSequenceBeginInRunes(runes, pos)

	return wordBeginInRunes(runes, pos)
}

// FirstWordInRunes computes first word.
// Returns (index of word's last rune)+1.
// Result also may be treated as length of the first word.
// First word may retrieved by "runes[:r]".
func FirstWordInRunes(runes []rune) int {
	return WordEndInRunes(runes, 0)
}

// LastWordInRunes computes last word.
// Returns index of word's first rune.
// Last word may retrieved by "runes[r:]".
func LastWordInRunes(runes []rune) int {
	return WordBeginInRunes(runes, len(runes)-1)
}

// WordAtInRunes computes word which contains pos-th rune and return their boundary.
// word may retrieved by "runes[r.From:r.To]".
func WordAtInRunes(runes []rune, pos int) Boundary {
	return Boundary{WordBeginInRunes(runes, pos), WordEndInRunes(runes, pos)}
}

// WordsInRunes computes all words and returns theirs boundaries.
func WordsInRunes(runes []rune) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(runes)) // TODO memory efficient
	for pos := 0; pos < len(runes); {
		length := FirstWordInRunes(runes[pos:])
		boundaries = append(boundaries, Boundary{pos, pos + length})
		pos += length
	}
	return
}

// WordBreaksInRunes computes all words and returns all breaks.
func WordBreaksInRunes(runes []rune) (breaks []int) {
	l := len(runes)
	if l == 0 {
		return
	}
	breaks = make([]int, 1, len(runes)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstWordInRunes(runes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}
