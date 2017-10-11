package boundaryh

// wSequence rules:
// -         CR LF              => NewLine
// -        NewLine             => NewLine
// -   ZWJ Glue_After_Zwj       => Glue_After_Zwj
// -   ZWJ EBG                  => EBG
// - X (Extend | Format | ZWJ)* => X
// wSequence allow avoid rules WB3, WB3c & WB4.

// Computes first wSequence.
// Returns:
// 	"c"   - wSequence class (see "wSequence rules").
// 	"pos" - point to first rune of next sequence (in other words "pos" is length of current wSequence).
func wFirstSequence(runes []rune) (c wClass, pos int) {
	l := len(runes)
	if l == 0 {
		return wClassOther, 0
	}

	if l > 1 && runes[0] == crRune && runes[1] == lfRune { // WB3
		return wClassNewline, 2
	}

	c = getWClass(runes[0])
	if l == 1 {
		return c, 1
	}

	pos = 1
	if c == wClassZWJ {
		c1 := getWClass(runes[1])
		if c1 == wClassGlueAfterZWJ || c1 == wClassEBG { // WB3c
			c = c1
			pos = 2
		}
	}

	if c == wClassNewline {
		return
	}

	for l1 := wClassOther; pos < l; pos++ { // WB4
		l0 := getWClass(runes[pos])

		switch {
		case l1 == wClassZWJ && (l0 == wClassGlueAfterZWJ || l0 == wClassEBG): // preserve following WB3c
			pos--
			return
		case !l0.isWB4():
			return
		}
		l1 = l0
	}

	return
}

// Computes last wSequence.
// Analogue to wFirstSequence.
// "pos" points to first rune in sequence.
func wLastSequence(runes []rune) (c wClass, pos int) {
	l := len(runes)
	if l == 0 {
		return wClassOther, 0
	}

	if l > 1 && runes[l-2] == crRune && runes[l-1] == lfRune { // WB3
		return wClassNewline, l - 2
	}

	c = getWClass(runes[l-1])
	if l == 1 {
		return c, 0
	}

	pos = l - 1
	for pos > 0 && c.isWB4() { // WB4
		pos--
		c = getWClass(runes[pos])
	}

	if (c == wClassGlueAfterZWJ || c == wClassEBG) && pos > 0 && getWClass(runes[pos-1]) == wClassZWJ { // WB3c
		pos--
	}

	return
}

// Returns position at which it is safe to begin analysis.
func wSequenceBegin(runes []rune, pos int) int {
	if pos == 0 {
		return 0
	}

	c := getWClass(runes[pos])
	l0 := getWClass(runes[pos-1])
	r0 := wClassOther
	if pos+1 < len(runes) {
		r0 = getWClass(runes[pos+1])
	}

	switch {
	case runes[pos-1] == crRune && runes[pos] == lfRune: // WB3
		pos--
	case c == wClassZWJ && (r0 == wClassGlueAfterZWJ || r0 == wClassEBG): // WB3c
	case l0 == wClassZWJ && (c == wClassGlueAfterZWJ || c == wClassEBG): // WB3c
		pos--
	case c.isWB4(): // WB4
		for pos--; pos >= 0; pos-- {
			c = getWClass(runes[pos])
			if c == wClassNewline {
				pos++
				break
			}
			if !c.isWB4() {
				break
			}
		}
		if pos == -1 {
			pos = 0
		}
	}
	return pos
}

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

// True if l0 is RI and it opens RI sequence in string <runes..., l0, ...> (may be joined with next RI).
func wIsOpenRI(runes []rune, l1, l0 wClass) (r bool) {
	r = l0 == wClassRI
	if !r {
		return
	}
	r = l1 != wClassRI
	if r {
		return
	}
	for len(runes) > 0 {
		c, pos := wLastSequence(runes)
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
func wordEnd(runes []rune, l0Pos int) int {
	l := len(runes)

	l1, l1Pos := wLastSequence(runes[:l0Pos])
	l0, r0Delta := wFirstSequence(runes[l0Pos:])
	lOddRI := wIsOpenRI(runes[:l1Pos], l1, l0)
	r0Pos := l0Pos + r0Delta
	r0, r1Delta := wFirstSequence(runes[r0Pos:])

	for r0Pos < l {
		r1, r2Delta := wFirstSequence(runes[r0Pos+r1Delta:])
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

// WordEnd computes word which contains pos-th rune.
// Returns (index of word's last rune)+1.
// In other words, returns first word's boundary on the right of pos-th rune.
func WordEnd(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = wSequenceBegin(runes, pos)

	return wordEnd(runes, pos)
}

// runes must be valid (len>1).
// r0Pos must be valid (in runes; really begin of sequence).
func wordBegin(runes []rune, r0Pos int) int {
	r0, r1Delta := wFirstSequence(runes[r0Pos:])
	r1, _ := wFirstSequence(runes[r0Pos+r1Delta:])
	l0, l0Pos := wLastSequence(runes[:r0Pos])

	for r0Pos >= 0 {
		l1, l1Pos := wLastSequence(runes[:l0Pos])
		lOddRI := wIsOpenRI(runes[:l1Pos], l1, l0)
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

// WordBegin computes word which contains pos-th rune.
// Returns word's first rune index.
// In other words, returns first word's boundary on the left of pos-th rune.
func WordBegin(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	pos = wSequenceBegin(runes, pos)

	return wordBegin(runes, pos)
}

// FirstWord computes first word.
// Returns (index of word's last rune)+1.
// Result also may be treated as length of the first word.
// First word may retrieved by "runes[:r]".
func FirstWord(runes []rune) int {
	return WordEnd(runes, 0)
}

// LastWord computes last word.
// Returns index of word's first rune.
// Last word may retrieved by "runes[r:]".
func LastWord(runes []rune) int {
	return WordBegin(runes, len(runes)-1)
}

// WordAt computes word which contains pos-th rune and return their boundary.
// word may retrieved by "runes[r.From:r.To]".
func WordAt(runes []rune, pos int) Boundary {
	return Boundary{WordBegin(runes, pos), WordEnd(runes, pos)}
}

// Words computes all words and returns theirs boundaries.
func Words(runes []rune) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(runes)) // TODO memory efficient
	for pos := 0; pos < len(runes); {
		length := FirstWord(runes[pos:])
		boundaries = append(boundaries, Boundary{pos, pos + length})
		pos += length
	}
	return
}

// WordBreaks computes all words and returns all breaks.
func WordBreaks(runes []rune) (breaks []int) {
	l := len(runes)
	if l == 0 {
		return
	}
	breaks = make([]int, 1, len(runes)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstWord(runes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}
