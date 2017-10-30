//replacer:generated-file

package boundaryh

// Computes first wSequence.
// Returns:
// 	"c"   - wSequence class (see "wSequence rules").
// 	"pos" - point to first rune of next sequence (in other words "pos" is length of current wSequence).
func wFirstSequenceInString(s string) (c wClass, pos int) {
	l := len(s)
	if l == 0 {
		return wClassOther, 0
	}

	if l > 1 && s[0] == crRune && s[1] == lfRune { // WB3
		return wClassNewline, 2
	}

	c, pos = wFirstClassInString(s)
	if l == 1 {
		return
	}

	if c == wClassZWJ {
		c1, delta := wFirstClassInString(s[pos:])
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
		l0, delta := wFirstClassInString(s[pos:])

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
// Analogue to wFirstSequenceInString.
// "pos" points to first rune in sequence.
func wLastSequenceInString(s string) (c wClass, pos int) {
	l := len(s)
	if l == 0 {
		return wClassOther, 0
	}

	if l > 1 && s[l-2] == crRune && s[l-1] == lfRune { // WB3
		return wClassNewline, l - 2
	}

	c, pos = wLastClassInString(s)
	if pos == 0 {
		return
	}

	for pos > 0 && c.isWB4() { // WB4
		c, pos = wLastClassInString(s[:pos])
	}

	if (c == wClassGlueAfterZWJ || c == wClassEBG) && pos > 0 { // WB3c
		c1, pos1 := wLastClassInString(s[:pos])
		if c1 == wClassZWJ {
			pos = pos1
		}
	}

	return
}

// Returns position at which it is safe to begin analysis.
func wSequenceBeginInString(s string, pos int) int {
	pos = toRuneBeginInString(s, pos)

	if pos == 0 {
		return 0
	}

	c, cDelta := wFirstClassInString(s[pos:])
	l0, l0Pos := wLastClassInString(s[:pos])
	r0 := wClassOther
	if pos+cDelta < len(s) {
		r0, _ = wFirstClassInString(s[pos+cDelta:])
	}

	switch {
	case s[pos-1] == crRune && s[pos] == lfRune: // WB3
		pos--
	case c == wClassZWJ && (r0 == wClassGlueAfterZWJ || r0 == wClassEBG): // WB3c
	case l0 == wClassZWJ && (c == wClassGlueAfterZWJ || c == wClassEBG): // WB3c
		pos = l0Pos
	case c.isWB4(): // WB4
		for pos > 0 {
			c, l0Pos = wLastClassInString(s[:pos])
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

// True if l0 is RI and it opens RI sequence in string <s..., l0, ...> (may be joined with next RI).
func wIsOpenRIInString(s string, l1, l0 wClass) (r bool) {
	r = l0 == wClassRI
	if !r {
		return
	}
	r = l1 != wClassRI
	if r {
		return
	}
	for len(s) > 0 {
		c, pos := wLastSequenceInString(s)
		if c != wClassRI {
			break
		}
		r = !r
		s = s[:pos]
	}
	return
}

// s must be valid (len>1).
// l0Pos must be valid (in s; really begin of sequence).
func wordEndInString(s string, l0Pos int) int {
	l := len(s)

	l1, l1Pos := wLastSequenceInString(s[:l0Pos])
	l0, r0Delta := wFirstSequenceInString(s[l0Pos:])
	lOddRI := wIsOpenRIInString(s[:l1Pos], l1, l0)
	r0Pos := l0Pos + r0Delta
	r0, r1Delta := wFirstSequenceInString(s[r0Pos:])

	for r0Pos < l {
		r1, r2Delta := wFirstSequenceInString(s[r0Pos+r1Delta:])
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

// WordEndInString computes word which contains pos-th rune.
// Returns (index of word's last rune)+1.
// In other words, returns first word's boundary on the right of pos-th rune.
func WordEndInString(s string, pos int) int {
	l := len(s)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = wSequenceBeginInString(s, pos)

	return wordEndInString(s, pos)
}

// s must be valid (len>1).
// r0Pos must be valid (in s; really begin of sequence).
func wordBeginInString(s string, r0Pos int) int {
	r0, r1Delta := wFirstSequenceInString(s[r0Pos:])
	r1, _ := wFirstSequenceInString(s[r0Pos+r1Delta:])
	l0, l0Pos := wLastSequenceInString(s[:r0Pos])

	for r0Pos >= 0 {
		l1, l1Pos := wLastSequenceInString(s[:l0Pos])
		lOddRI := wIsOpenRIInString(s[:l1Pos], l1, l0)
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

// WordBeginInString computes word which contains pos-th rune.
// Returns word's first rune index.
// In other words, returns first word's boundary on the left of pos-th rune.
func WordBeginInString(s string, pos int) int {
	l := len(s)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	pos = wSequenceBeginInString(s, pos)

	return wordBeginInString(s, pos)
}

// FirstWordInString computes first word.
// Returns (index of word's last rune)+1.
// Result also may be treated as length of the first word.
// First word may retrieved by "s[:r]".
func FirstWordInString(s string) int {
	return WordEndInString(s, 0)
}

// LastWordInString computes last word.
// Returns index of word's first rune.
// Last word may retrieved by "s[r:]".
func LastWordInString(s string) int {
	return WordBeginInString(s, len(s)-1)
}

// WordAtInString computes word which contains pos-th rune and return their boundary.
// word may retrieved by "s[r.From:r.To]".
func WordAtInString(s string, pos int) Boundary {
	return Boundary{WordBeginInString(s, pos), WordEndInString(s, pos)}
}

// WordsInString computes all words and returns theirs boundaries.
func WordsInString(s string) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(s)) // TODO memory efficient
	for pos := 0; pos < len(s); {
		length := FirstWordInString(s[pos:])
		boundaries = append(boundaries, Boundary{pos, pos + length})
		pos += length
	}
	return
}

// WordBreaksInString computes all words and returns all breaks.
func WordBreaksInString(s string) (breaks []int) {
	l := len(s)
	if l == 0 {
		return
	}
	breaks = make([]int, 1, len(s)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstWordInString(s[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}

// Computes first wSequence.
// Returns:
// 	"c"   - wSequence class (see "wSequence rules").
// 	"pos" - point to first rune of next sequence (in other words "pos" is length of current wSequence).
func wFirstSequence(bytes []byte) (c wClass, pos int) {
	l := len(bytes)
	if l == 0 {
		return wClassOther, 0
	}

	if l > 1 && bytes[0] == crRune && bytes[1] == lfRune { // WB3
		return wClassNewline, 2
	}

	c, pos = wFirstClass(bytes)
	if l == 1 {
		return
	}

	if c == wClassZWJ {
		c1, delta := wFirstClass(bytes[pos:])
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
		l0, delta := wFirstClass(bytes[pos:])

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
// Analogue to wFirstSequence.
// "pos" points to first rune in sequence.
func wLastSequence(bytes []byte) (c wClass, pos int) {
	l := len(bytes)
	if l == 0 {
		return wClassOther, 0
	}

	if l > 1 && bytes[l-2] == crRune && bytes[l-1] == lfRune { // WB3
		return wClassNewline, l - 2
	}

	c, pos = wLastClass(bytes)
	if pos == 0 {
		return
	}

	for pos > 0 && c.isWB4() { // WB4
		c, pos = wLastClass(bytes[:pos])
	}

	if (c == wClassGlueAfterZWJ || c == wClassEBG) && pos > 0 { // WB3c
		c1, pos1 := wLastClass(bytes[:pos])
		if c1 == wClassZWJ {
			pos = pos1
		}
	}

	return
}

// Returns position at which it is safe to begin analysis.
func wSequenceBegin(bytes []byte, pos int) int {
	pos = toRuneBegin(bytes, pos)

	if pos == 0 {
		return 0
	}

	c, cDelta := wFirstClass(bytes[pos:])
	l0, l0Pos := wLastClass(bytes[:pos])
	r0 := wClassOther
	if pos+cDelta < len(bytes) {
		r0, _ = wFirstClass(bytes[pos+cDelta:])
	}

	switch {
	case bytes[pos-1] == crRune && bytes[pos] == lfRune: // WB3
		pos--
	case c == wClassZWJ && (r0 == wClassGlueAfterZWJ || r0 == wClassEBG): // WB3c
	case l0 == wClassZWJ && (c == wClassGlueAfterZWJ || c == wClassEBG): // WB3c
		pos = l0Pos
	case c.isWB4(): // WB4
		for pos > 0 {
			c, l0Pos = wLastClass(bytes[:pos])
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

// True if l0 is RI and it opens RI sequence in string <bytes..., l0, ...> (may be joined with next RI).
func wIsOpenRI(bytes []byte, l1, l0 wClass) (r bool) {
	r = l0 == wClassRI
	if !r {
		return
	}
	r = l1 != wClassRI
	if r {
		return
	}
	for len(bytes) > 0 {
		c, pos := wLastSequence(bytes)
		if c != wClassRI {
			break
		}
		r = !r
		bytes = bytes[:pos]
	}
	return
}

// bytes must be valid (len>1).
// l0Pos must be valid (in bytes; really begin of sequence).
func wordEnd(bytes []byte, l0Pos int) int {
	l := len(bytes)

	l1, l1Pos := wLastSequence(bytes[:l0Pos])
	l0, r0Delta := wFirstSequence(bytes[l0Pos:])
	lOddRI := wIsOpenRI(bytes[:l1Pos], l1, l0)
	r0Pos := l0Pos + r0Delta
	r0, r1Delta := wFirstSequence(bytes[r0Pos:])

	for r0Pos < l {
		r1, r2Delta := wFirstSequence(bytes[r0Pos+r1Delta:])
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
func WordEnd(bytes []byte, pos int) int {
	l := len(bytes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = wSequenceBegin(bytes, pos)

	return wordEnd(bytes, pos)
}

// bytes must be valid (len>1).
// r0Pos must be valid (in bytes; really begin of sequence).
func wordBegin(bytes []byte, r0Pos int) int {
	r0, r1Delta := wFirstSequence(bytes[r0Pos:])
	r1, _ := wFirstSequence(bytes[r0Pos+r1Delta:])
	l0, l0Pos := wLastSequence(bytes[:r0Pos])

	for r0Pos >= 0 {
		l1, l1Pos := wLastSequence(bytes[:l0Pos])
		lOddRI := wIsOpenRI(bytes[:l1Pos], l1, l0)
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
func WordBegin(bytes []byte, pos int) int {
	l := len(bytes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	pos = wSequenceBegin(bytes, pos)

	return wordBegin(bytes, pos)
}

// FirstWord computes first word.
// Returns (index of word's last rune)+1.
// Result also may be treated as length of the first word.
// First word may retrieved by "bytes[:r]".
func FirstWord(bytes []byte) int {
	return WordEnd(bytes, 0)
}

// LastWord computes last word.
// Returns index of word's first rune.
// Last word may retrieved by "bytes[r:]".
func LastWord(bytes []byte) int {
	return WordBegin(bytes, len(bytes)-1)
}

// WordAt computes word which contains pos-th rune and return their boundary.
// word may retrieved by "bytes[r.From:r.To]".
func WordAt(bytes []byte, pos int) Boundary {
	return Boundary{WordBegin(bytes, pos), WordEnd(bytes, pos)}
}

// Words computes all words and returns theirs boundaries.
func Words(bytes []byte) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(bytes)) // TODO memory efficient
	for pos := 0; pos < len(bytes); {
		length := FirstWord(bytes[pos:])
		boundaries = append(boundaries, Boundary{pos, pos + length})
		pos += length
	}
	return
}

// WordBreaks computes all words and returns all breaks.
func WordBreaks(bytes []byte) (breaks []int) {
	l := len(bytes)
	if l == 0 {
		return
	}
	breaks = make([]int, 1, len(bytes)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstWord(bytes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}
