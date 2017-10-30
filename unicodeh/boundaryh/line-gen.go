//replacer:generated-file

package boundaryh

// Computes first lSequence.
// Returns:
// 	"c"      - lSequence class (see "lSequence rules").
// 	"rawIS"  - is sequence primary class is IS. This information required by some rules which require exactly IS or SY, but IS will be replaced with SY.
// 	"rawZWJ" - is sequence just single ZWJ. This information required bee LB8a, but "c" in such cases will be AL.
// 	"pos"    - point to first rune of next sequence (in other words "pos" is length of current lSequence).
func lFirstSequenceInString(s string) (c lClass, rawIS, rawZWJ bool, pos int) {
	l := len(s)
	if l == 0 {
		c = lClassXX
		return
	}

	if s[0] == crRune { // LB4
		if l > 1 && s[1] == lfRune { // LB5
			return lClassBK, false, false, 2
		}
		return lClassBK, false, false, 1
	}

	if s[0] == lfRune { // LB4
		return lClassBK, false, false, 1
	}

	c, pos = lFirstClassInString(s)

	if c.isSkip() { // LB10
		return lClassAL, false, c == lClassZWJ, pos
	} else if !c.isSkipBase() { // LB9
		return c, false, false, pos
	}

	if c == lClassIS {
		c = lClassSY
		rawIS = true
	}

	for pos < len(s) { // LB9
		c0, delta := lFirstClassInString(s[pos:])
		if !c0.isSkip() {
			break
		}
		pos += delta
	}

	return
}

// Computes last lSequence.
// Analogue to lFirstSequenceInString.
// "pos" points to first rune in sequence.
func lLastSequenceInString(s string) (c lClass, rawIS, rawZWJ bool, pos int) {
	l := len(s)
	if l == 0 {
		c = lClassXX
		return
	}

	if s[l-1] == lfRune { // LB4
		if l > 1 && s[l-2] == crRune { // LB5
			return lClassBK, false, false, l - 2
		}
		return lClassBK, false, false, l - 1
	}

	if s[l-1] == crRune { // LB4
		return lClassBK, false, false, l - 1
	}

	c, pos = lLastClassInString(s)

	if !c.isSkip() { // not LB9
		if c == lClassIS {
			c = lClassSY
			rawIS = true
		}
		return
	}

	rawZWJ = c == lClassZWJ

	for pos0 := pos; pos0 > 0; {
		c, pos0 = lLastClassInString(s[:pos0])
		switch {
		case c.isSkip(): // LB9
		case c.isSkipBase(): // LB9
			if c == lClassIS {
				c = lClassSY
				rawIS = true
			}
			return c, rawIS, false, pos0
		default:
			return lClassAL, false, rawZWJ, pos // LB10
		}
	}

	return lClassAL, false, rawZWJ, pos // LB10
}

// Returns position at which it is safe to begin analysis.
// pos must be in s: [0; len(s)-1] .
func lSequenceBeginInString(s string, pos int) int {
	if pos == 0 {
		return 0
	}

	if s[pos-1] == crRune && s[pos] == lfRune { // LB5
		return pos - 1
	}

	pos = toRuneBeginInString(s, pos) // TODO make name similar

	c, _ := lFirstClassInString(s[pos:]) // not LB9
	if !c.isSkip() {
		return pos
	}

	for pos0 := pos; pos0 > 0; {
		c, pos0 = lLastClassInString(s[:pos0])
		switch {
		case c.isSkip(): // LB9
		case c.isSkipBase(): // LB9
			return pos0
		default: // LB10
			return pos
		}
	}

	return pos
}

// Returns class of last rune in <s..., l1> which is not equal to l0.
func lLastNotEqualToInString(s string, l1, l0 lClass) lClass {
	if l1 != l0 {
		return l1
	}
	for len(s) > 0 {
		c, _, _, pos := lLastSequenceInString(s)
		if c != l0 {
			return c
		}
		s = s[:pos]
	}
	return lClassXX
}

// Same as without suffix "1" but with only one predefined class (looking in <s...>, not <s..., l1>).
func lLastNotEqualTo0InString(s string, l0 lClass) lClass {
	l1, _, _, pos := lLastSequenceInString(s)
	return lLastNotEqualToInString(s[:pos], l1, l0)
}

// True if l0 is RI and it opens RI sequence in string <s..., l1, l0, ...> (may be joined with next RI).
func lIsOpenRIInString(s string, l1, l0 lClass) (r bool) {
	r = l0 == lClassRI
	if !r {
		return
	}
	r = l1 != lClassRI
	if r {
		return
	}
	for len(s) > 0 {
		c, _, _, pos := lLastSequenceInString(s)
		if c != lClassRI {
			break
		}
		r = !r
		s = s[:pos]
	}
	return
}

// s must be valid (len>1).
// l0Pos must be valid (in s; really begin of sequence).
func lineBreakAfterInString(s string, l0Pos int) int {
	l := len(s)

	l1, _, _, l1Pos := lLastSequenceInString(s[:l0Pos])
	l0, l0IS, l0ZWJ, r0Delta := lFirstSequenceInString(s[l0Pos:])
	r0Pos := l0Pos + r0Delta
	r0, r0IS, r0ZWJ, r1Delta := lFirstSequenceInString(s[r0Pos:])
	l2Diff := lLastNotEqualTo0InString(s[:l1Pos], l1)
	l1Diff := lLastNotEqualToInString(s[:l1Pos], l1, l0)
	lOddRI := lIsOpenRIInString(s[:l1Pos], l1, l0)

	for r0Pos < l {
		r1, r1IS, r1ZWJ, r2Delta := lFirstSequenceInString(s[r0Pos+r1Delta:])

		if lDecision(l1, l0, r0, r1, l2Diff, l1Diff, l0IS, l0ZWJ, lOddRI) {
			return r0Pos
		}

		r0Pos += r1Delta
		r1Delta = r2Delta
		l0ZWJ = r0ZWJ
		r0ZWJ = r1ZWJ
		if l1 != l0 {
			l2Diff = l1
		}
		if l0 != r0 {
			l1Diff = l0
		}
		l1 = l0
		l0 = r0
		r0 = r1
		l0IS = r0IS
		r0IS = r1IS
		lOddRI = l0 == lClassRI && !lOddRI
	}
	return l
}

// LineBreakAfterInString returns first possible line break on the right side of pos-th rune.
func LineBreakAfterInString(s string, pos int) int {
	l := len(s)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = lSequenceBeginInString(s, pos)

	return lineBreakAfterInString(s, pos)
}

// s must be valid (len>1).
// r0Pos must be valid (in s; really begin of sequence).
func lineBreakBeforeInString(s string, r0Pos int) int {
	r0, _, _, r1Delta := lFirstSequenceInString(s[r0Pos:])
	r1, _, _, _ := lFirstSequenceInString(s[r0Pos+r1Delta:])
	l0, l0IS, l0ZWJ, l0Pos := lLastSequenceInString(s[:r0Pos])
	l1Diff := lLastNotEqualTo0InString(s[:l0Pos], l0)
	l2Diff := l1Diff

	for r0Pos > 0 {
		l1, l1IS, l1ZWJ, l1Pos := lLastSequenceInString(s[:l0Pos])
		lOddRI := lIsOpenRIInString(s[:l1Pos], l1, l0)
		if l2Diff == l1 {
			l2Diff = lLastNotEqualTo0InString(s[:l1Pos], l1)
		}

		if lDecision(l1, l0, r0, r1, l2Diff, l1Diff, l0IS, l0ZWJ, lOddRI) {
			return r0Pos
		}

		if l1Diff == l1 {
			l1Diff = l2Diff
		}
		r0Pos = l0Pos
		l0Pos = l1Pos

		l0ZWJ = l1ZWJ

		r1 = r0
		r0 = l0
		l0 = l1
		l0IS = l1IS
	}
	return NoLineBreak
}

// LineBreakBeforeInString returns first (nearest) possible line break on the left side of pos-th rune.
func LineBreakBeforeInString(s string, pos int) int {
	l := len(s)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return NoLineBreak
	}

	pos = lSequenceBeginInString(s, pos)

	return lineBreakBeforeInString(s, pos)
}

// FirstLineBreakInString returns first possible line break.
func FirstLineBreakInString(s string) int {
	return LineBreakAfterInString(s, 0)
}

// LastLineBreakInString returns last possible line break (except line break at end of string).
func LastLineBreakInString(s string) int {
	return LineBreakBeforeInString(s, len(s)-1)
}

// LineBreaksInString returns all possible line breaks.
func LineBreaksInString(s string) (breaks []int) {
	l := len(s)
	if l == 0 {
		return
	}
	breaks = make([]int, 0, l) // TODO memory efficient
	for pos := 0; pos < l; {
		length := FirstLineBreakInString(s[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}

// Computes first lSequence.
// Returns:
// 	"c"      - lSequence class (see "lSequence rules").
// 	"rawIS"  - is sequence primary class is IS. This information required by some rules which require exactly IS or SY, but IS will be replaced with SY.
// 	"rawZWJ" - is sequence just single ZWJ. This information required bee LB8a, but "c" in such cases will be AL.
// 	"pos"    - point to first rune of next sequence (in other words "pos" is length of current lSequence).
func lFirstSequence(bytes []byte) (c lClass, rawIS, rawZWJ bool, pos int) {
	l := len(bytes)
	if l == 0 {
		c = lClassXX
		return
	}

	if bytes[0] == crRune { // LB4
		if l > 1 && bytes[1] == lfRune { // LB5
			return lClassBK, false, false, 2
		}
		return lClassBK, false, false, 1
	}

	if bytes[0] == lfRune { // LB4
		return lClassBK, false, false, 1
	}

	c, pos = lFirstClass(bytes)

	if c.isSkip() { // LB10
		return lClassAL, false, c == lClassZWJ, pos
	} else if !c.isSkipBase() { // LB9
		return c, false, false, pos
	}

	if c == lClassIS {
		c = lClassSY
		rawIS = true
	}

	for pos < len(bytes) { // LB9
		c0, delta := lFirstClass(bytes[pos:])
		if !c0.isSkip() {
			break
		}
		pos += delta
	}

	return
}

// Computes last lSequence.
// Analogue to lFirstSequence.
// "pos" points to first rune in sequence.
func lLastSequence(bytes []byte) (c lClass, rawIS, rawZWJ bool, pos int) {
	l := len(bytes)
	if l == 0 {
		c = lClassXX
		return
	}

	if bytes[l-1] == lfRune { // LB4
		if l > 1 && bytes[l-2] == crRune { // LB5
			return lClassBK, false, false, l - 2
		}
		return lClassBK, false, false, l - 1
	}

	if bytes[l-1] == crRune { // LB4
		return lClassBK, false, false, l - 1
	}

	c, pos = lLastClass(bytes)

	if !c.isSkip() { // not LB9
		if c == lClassIS {
			c = lClassSY
			rawIS = true
		}
		return
	}

	rawZWJ = c == lClassZWJ

	for pos0 := pos; pos0 > 0; {
		c, pos0 = lLastClass(bytes[:pos0])
		switch {
		case c.isSkip(): // LB9
		case c.isSkipBase(): // LB9
			if c == lClassIS {
				c = lClassSY
				rawIS = true
			}
			return c, rawIS, false, pos0
		default:
			return lClassAL, false, rawZWJ, pos // LB10
		}
	}

	return lClassAL, false, rawZWJ, pos // LB10
}

// Returns position at which it is safe to begin analysis.
// pos must be in bytes: [0; len(bytes)-1] .
func lSequenceBegin(bytes []byte, pos int) int {
	if pos == 0 {
		return 0
	}

	if bytes[pos-1] == crRune && bytes[pos] == lfRune { // LB5
		return pos - 1
	}

	pos = toRuneBegin(bytes, pos) // TODO make name similar

	c, _ := lFirstClass(bytes[pos:]) // not LB9
	if !c.isSkip() {
		return pos
	}

	for pos0 := pos; pos0 > 0; {
		c, pos0 = lLastClass(bytes[:pos0])
		switch {
		case c.isSkip(): // LB9
		case c.isSkipBase(): // LB9
			return pos0
		default: // LB10
			return pos
		}
	}

	return pos
}

// Returns class of last rune in <bytes..., l1> which is not equal to l0.
func lLastNotEqualTo(bytes []byte, l1, l0 lClass) lClass {
	if l1 != l0 {
		return l1
	}
	for len(bytes) > 0 {
		c, _, _, pos := lLastSequence(bytes)
		if c != l0 {
			return c
		}
		bytes = bytes[:pos]
	}
	return lClassXX
}

// Same as without suffix "1" but with only one predefined class (looking in <bytes...>, not <bytes..., l1>).
func lLastNotEqualTo0(bytes []byte, l0 lClass) lClass {
	l1, _, _, pos := lLastSequence(bytes)
	return lLastNotEqualTo(bytes[:pos], l1, l0)
}

// True if l0 is RI and it opens RI sequence in string <bytes..., l1, l0, ...> (may be joined with next RI).
func lIsOpenRI(bytes []byte, l1, l0 lClass) (r bool) {
	r = l0 == lClassRI
	if !r {
		return
	}
	r = l1 != lClassRI
	if r {
		return
	}
	for len(bytes) > 0 {
		c, _, _, pos := lLastSequence(bytes)
		if c != lClassRI {
			break
		}
		r = !r
		bytes = bytes[:pos]
	}
	return
}

// bytes must be valid (len>1).
// l0Pos must be valid (in bytes; really begin of sequence).
func lineBreakAfter(bytes []byte, l0Pos int) int {
	l := len(bytes)

	l1, _, _, l1Pos := lLastSequence(bytes[:l0Pos])
	l0, l0IS, l0ZWJ, r0Delta := lFirstSequence(bytes[l0Pos:])
	r0Pos := l0Pos + r0Delta
	r0, r0IS, r0ZWJ, r1Delta := lFirstSequence(bytes[r0Pos:])
	l2Diff := lLastNotEqualTo0(bytes[:l1Pos], l1)
	l1Diff := lLastNotEqualTo(bytes[:l1Pos], l1, l0)
	lOddRI := lIsOpenRI(bytes[:l1Pos], l1, l0)

	for r0Pos < l {
		r1, r1IS, r1ZWJ, r2Delta := lFirstSequence(bytes[r0Pos+r1Delta:])

		if lDecision(l1, l0, r0, r1, l2Diff, l1Diff, l0IS, l0ZWJ, lOddRI) {
			return r0Pos
		}

		r0Pos += r1Delta
		r1Delta = r2Delta
		l0ZWJ = r0ZWJ
		r0ZWJ = r1ZWJ
		if l1 != l0 {
			l2Diff = l1
		}
		if l0 != r0 {
			l1Diff = l0
		}
		l1 = l0
		l0 = r0
		r0 = r1
		l0IS = r0IS
		r0IS = r1IS
		lOddRI = l0 == lClassRI && !lOddRI
	}
	return l
}

// LineBreakAfter returns first possible line break on the right side of pos-th rune.
func LineBreakAfter(bytes []byte, pos int) int {
	l := len(bytes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = lSequenceBegin(bytes, pos)

	return lineBreakAfter(bytes, pos)
}

// bytes must be valid (len>1).
// r0Pos must be valid (in bytes; really begin of sequence).
func lineBreakBefore(bytes []byte, r0Pos int) int {
	r0, _, _, r1Delta := lFirstSequence(bytes[r0Pos:])
	r1, _, _, _ := lFirstSequence(bytes[r0Pos+r1Delta:])
	l0, l0IS, l0ZWJ, l0Pos := lLastSequence(bytes[:r0Pos])
	l1Diff := lLastNotEqualTo0(bytes[:l0Pos], l0)
	l2Diff := l1Diff

	for r0Pos > 0 {
		l1, l1IS, l1ZWJ, l1Pos := lLastSequence(bytes[:l0Pos])
		lOddRI := lIsOpenRI(bytes[:l1Pos], l1, l0)
		if l2Diff == l1 {
			l2Diff = lLastNotEqualTo0(bytes[:l1Pos], l1)
		}

		if lDecision(l1, l0, r0, r1, l2Diff, l1Diff, l0IS, l0ZWJ, lOddRI) {
			return r0Pos
		}

		if l1Diff == l1 {
			l1Diff = l2Diff
		}
		r0Pos = l0Pos
		l0Pos = l1Pos

		l0ZWJ = l1ZWJ

		r1 = r0
		r0 = l0
		l0 = l1
		l0IS = l1IS
	}
	return NoLineBreak
}

// LineBreakBefore returns first (nearest) possible line break on the left side of pos-th rune.
func LineBreakBefore(bytes []byte, pos int) int {
	l := len(bytes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return NoLineBreak
	}

	pos = lSequenceBegin(bytes, pos)

	return lineBreakBefore(bytes, pos)
}

// FirstLineBreak returns first possible line break.
func FirstLineBreak(bytes []byte) int {
	return LineBreakAfter(bytes, 0)
}

// LastLineBreak returns last possible line break (except line break at end of string).
func LastLineBreak(bytes []byte) int {
	return LineBreakBefore(bytes, len(bytes)-1)
}

// LineBreaks returns all possible line breaks.
func LineBreaks(bytes []byte) (breaks []int) {
	l := len(bytes)
	if l == 0 {
		return
	}
	breaks = make([]int, 0, l) // TODO memory efficient
	for pos := 0; pos < l; {
		length := FirstLineBreak(bytes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}
