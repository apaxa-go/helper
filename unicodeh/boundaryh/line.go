package boundaryh

// lSequence rules:
// 	     CR lF     => BK
// 	   (CR | lF)   => BK
// 	IS (CM | ZWJ)* => SY (custom rule; it simplify decision)
// 	 X (CM | ZWJ)* => X (where X is defined by LB9)
// 	   (CM | ZWJ)  => AL
// 	       X       => X

// NoLineBreak indicates that where is no possible line break.
// TR14 says, that there is no break at start of string (this is required to avoid breaking line after zero runes).
const NoLineBreak int = 0

// Computes first lSequence.
// Returns:
// 	"c"      - lSequence class (see "lSequence rules").
// 	"rawIS"  - is sequence primary class is IS. This information required by some rules which require exactly IS or SY, but IS will be replaced with SY.
// 	"rawZWJ" - is sequence just single ZWJ. This information required bee LB8a, but "c" in such cases will be AL.
// 	"pos"    - point to first rune of next sequence (in other words "pos" is length of current lSequence).
func lFirstSequence(runes []rune) (c lClass, rawIS, rawZWJ bool, pos int) {
	l := len(runes)
	if l == 0 {
		c = lClassXX
		return
	}

	if runes[0] == crRune { // LB4
		if l > 1 && runes[1] == lfRune { // LB5
			return lClassBK, false, false, 2
		}
		return lClassBK, false, false, 1
	}

	if runes[0] == lfRune { // LB4
		return lClassBK, false, false, 1
	}

	c = lGetClass(runes[0])

	if c.isSkip() { // LB10
		return lClassAL, false, c == lClassZWJ, 1
	} else if !c.isSkipBase() { // LB9
		return c, false, false, 1
	}

	if c == lClassIS {
		c = lClassSY
		rawIS = true
	}

	for pos = 1; pos < len(runes); pos++ { // LB9
		if !lGetClass(runes[pos]).isSkip() {
			break
		}
	}

	return
}

// Computes last lSequence.
// Analogue to lFirstSequence.
// "pos" points to first rune in sequence.
func lLastSequence(runes []rune) (c lClass, rawIS, rawZWJ bool, pos int) {
	l := len(runes)
	if l == 0 {
		c = lClassXX
		return
	}

	if runes[l-1] == lfRune { // LB4
		if l > 1 && runes[l-2] == crRune { // LB5
			return lClassBK, false, false, l - 2
		}
		return lClassBK, false, false, l - 1
	}

	if runes[l-1] == crRune { // LB4
		return lClassBK, false, false, l - 1
	}

	pos = l - 1
	c = lGetClass(runes[pos])

	if !c.isSkip() { // not LB9
		if c == lClassIS {
			c = lClassSY
			rawIS = true
		}
		return
	}

	rawZWJ = c == lClassZWJ

	for pos--; pos >= 0; pos-- {
		c0 := lGetClass(runes[pos])
		switch {
		case c0.isSkip(): // LB9
		case c0.isSkipBase(): // LB9
			c = c0
			rawZWJ = false
			if c == lClassIS {
				c = lClassSY
				rawIS = true
			}
			return
		default:
			return lClassAL, false, rawZWJ, l - 1 // LB10
		}
	}

	return lClassAL, false, rawZWJ, l - 1 // LB10
}

// Returns position at which it is safe to begin analysis.
func lSequenceBegin(runes []rune) int {
	l := len(runes)
	if l <= 1 {
		return 0
	}

	if runes[l-2] == crRune && runes[l-1] == lfRune { // LB5
		return l - 2
	}

	if !lGetClass(runes[l-1]).isSkip() { // not LB9
		return l - 1
	}

	for pos := l - 2; pos >= 0; pos-- {
		c := lGetClass(runes[pos])
		switch {
		case c.isSkip(): // LB9
		case c.isSkipBase(): // LB9
			return pos
		default: // LB10
			return l - 1
		}
	}

	return l - 1 // LB10
}

// Returns if there is a break between l0 and r0.
func lDecision(l1, l0, r0, r1, l2Diff, l1Diff lClass, l0IS, l0ZWJ, lOddRI bool) bool {
	// TODO convert to single boolean expression?
	switch {
	case l0 == lClassNL || l0 == lClassBK: // LB4 & LB5
		return true
	case r0 == lClassNL || r0 == lClassBK: // LB6
	case r0 == lClassSP || r0 == lClassZW: // LB7
	case l0 == lClassZW: // LB8 (part 1)
		return true
	case l1Diff == lClassZW && l0 == lClassSP: // LB8 (part 2)
		return true
	case l0ZWJ && (r0 == lClassID || r0 == lClassEB || r0 == lClassEM): // LB8a
	case l0 == lClassWJ || r0 == lClassWJ: // LB 11
	case l0 == lClassGL: // LB12
	case (l0 != lClassSP && l0 != lClassBA && l0 != lClassHY) && r0 == lClassGL: // LB12a
	case r0 == lClassCL || r0 == lClassCP || r0 == lClassEX || r0 == lClassSY: // LB13
	case l0 == lClassOP: // LB14 (part 1)
	case l1Diff == lClassOP && l0 == lClassSP: // LB14 (part 2)
	case l0 == lClassQU && r0 == lClassOP: // LB15 (part 1)
	case l1Diff == lClassQU && l0 == lClassSP && r0 == lClassOP: // LB15 (part 2)
	case (l0 == lClassCL || l0 == lClassCP) && r0 == lClassNS: // LB16 (part 1)
	case (l1Diff == lClassCL || l1Diff == lClassCP) && l0 == lClassSP && r0 == lClassNS: // LB16 (part 2)
	case l0 == lClassB2 && r0 == lClassB2: // LB17 (part 1)
	case l1Diff == lClassB2 && l0 == lClassSP && r0 == lClassB2: // LB17 (part 2)
	case l0 == lClassSP: // LB18
		return true
	case l0 == lClassQU || r0 == lClassQU: // LB19
	case l0 == lClassCB || r0 == lClassCB: // LB20
		return true
	case r0 == lClassBA || r0 == lClassHY || r0 == lClassNS || l0 == lClassBB: // LB21
	case l1 == lClassHL && (l0 == lClassHY || l0 == lClassBA): // LB21a
	case l0 == lClassSY && !l0IS && r0 == lClassHL: // LB21b
	case (l0 == lClassAL || l0 == lClassHL || l0 == lClassEX || l0 == lClassID || l0 == lClassEB || l0 == lClassEM || l0 == lClassIN || l0 == lClassNU) && r0 == lClassIN: // LB22
	case (l0 == lClassAL || l0 == lClassHL) && r0 == lClassNU: // LB23 (part 1)
	case l0 == lClassNU && (r0 == lClassAL || r0 == lClassHL): // LB23 (part 2)
	case l0 == lClassPR && (r0 == lClassID || r0 == lClassEB || r0 == lClassEM): // LB23a (part 1)
	case (l0 == lClassID || l0 == lClassEB || l0 == lClassEM) && r0 == lClassPO: // LB23a (part 2)
	case (l0 == lClassPR || l0 == lClassPO) && (r0 == lClassAL || r0 == lClassHL): // LB24 (part 1)
	case (l0 == lClassAL || l0 == lClassHL) && (r0 == lClassPR || r0 == lClassPO): // LB24 (part 2)
		/*
			// This is the default implementation of LB25.
			// Regexp "( PR | PO) ? ( OP | HY ) ? NU (NU | SY | IS) * ( CL | CP ) ? ( PR | PO) ?" suggested instead.
			// Also LineBreakTest.txt use this regexp, not default implementation.
			case (l0 == lClassCL && r0 == lClassPO) ||
				(l0 == lClassCP && r0 == lClassPO) ||
				(l0 == lClassCL && r0 == lClassPR) ||
				(l0 == lClassCP && r0 == lClassPR) ||
				(l0 == lClassNU && r0 == lClassPO) ||
				(l0 == lClassNU && r0 == lClassPR) ||
				(l0 == lClassPO && r0 == lClassOP) ||
				(l0 == lClassPO && r0 == lClassNU) ||
				(l0 == lClassPR && r0 == lClassOP) ||
				(l0 == lClassPR && r0 == lClassNU) ||
				(l0 == lClassHY && r0 == lClassNU) ||
				(l0 == lClassIS && r0 == lClassNU) ||
				(l0 == lClassNU && r0 == lClassNU) ||
				(l0 == lClassSY && r0 == lClassNU): // LB25
		*/

		// Regexp "( PR | PO) ? ( OP | HY ) ? NU (NU | SY | IS) * ( CL | CP ) ? ( PR | PO) ?" implementation of LB25.
		// See TR14 8.2 Example 7 for more information.
		// Rules:
		// 	CLB 1: (PR | PO) × (OP | HY)? NU - replaced with:
		//	CLB 1.1: (PR | PO) × NU
		// 	CLB 1.2: (PR | PO) × (OP | HY) NU
		// 	CLB 2: ( OP | HY ) × NU
		// 	CLB 3: NU × (NU | SY | IS)
		// 	CLB 4: NU (NU | SY | IS)* × (NU | SY | IS | CL | CP ) - replaced with:
		// 	CLB 4.1: NU × (NU | SY | IS | CL | CP )
		// 	CLB 4.2: NU (NU | SY | IS)+ × (NU | SY | IS | CL | CP )
		// 	CLB 5: NU (NU | SY | IS)* (CL | CP)? × (PO | PR) - replaced with:
		// 	CLB 5.1: NU × (PO | PR)
		// 	CLB 5.2: NU (CL | CP) × (PO | PR)
		// 	CLB 5.3: NU (SY | IS)+ × (PO | PR)
		// 	CLB 5.4: NU (SY | IS)+ (CL | CP) × (PO | PR)
	case ((l0 == lClassPR || l0 == lClassPO) && r0 == lClassNU) || // CLB1.1
		((l0 == lClassPR || l0 == lClassPO) && (r0 == lClassOP || r0 == lClassHY) && r1 == lClassNU) || // CLB1.2
		((l0 == lClassOP || l0 == lClassHY) && r0 == lClassNU) || // CLB2
		(l0 == lClassNU && (r0 == lClassNU || r0 == lClassSY)) || // CLB3
		(l0 == lClassNU && (r0 == lClassNU || r0 == lClassSY || r0 == lClassCL || r0 == lClassCP)) || // CLB4.1
		(l1Diff == lClassNU && l0 == lClassSY && (r0 == lClassNU || r0 == lClassSY || r0 == lClassCL || r0 == lClassCP)) || // CLB4.2
		(l0 == lClassNU && (r0 == lClassPO || r0 == lClassPR)) || // CLB5.1
		(l1 == lClassNU && (l0 == lClassCL || l0 == lClassCP) && (r0 == lClassPO || r0 == lClassPR)) || // CLB5.2
		(l1Diff == lClassNU && l0 == lClassSY && (r0 == lClassPO || r0 == lClassPR)) || // CLB5.3
		(l2Diff == lClassNU && l1 == lClassSY && (l0 == lClassCL || l0 == lClassCP) && (r0 == lClassPO || r0 == lClassPR)): // CLB5.4
	case l0 == lClassJL && (r0 == lClassJL || r0 == lClassJV || r0 == lClassH2 || r0 == lClassH3): // LB26 (part 1)
	case (l0 == lClassJV || l0 == lClassH2) && (r0 == lClassJV || r0 == lClassJT): // LB26 (part 2)
	case (l0 == lClassJT || l0 == lClassH3) && r0 == lClassJT: // LB26 (part 3)
	case (l0 == lClassJL || l0 == lClassJV || l0 == lClassJT || l0 == lClassH2 || l0 == lClassH3) && (r0 == lClassIN || r0 == lClassPO): // LB27 (part 1)
	case l0 == lClassPR && (r0 == lClassJL || r0 == lClassJV || r0 == lClassJT || r0 == lClassH2 || r0 == lClassH3): // LB27 (part 2)
	case (l0 == lClassAL || l0 == lClassHL || l0IS) && (r0 == lClassAL || r0 == lClassHL): // LB28 & LB29
	case (l0 == lClassAL || l0 == lClassHL || l0 == lClassNU) && r0 == lClassOP: // LB30 (part 1)
	case l0 == lClassCP && (r0 == lClassAL || r0 == lClassHL || r0 == lClassNU): // LB30 (part 2)
	case lOddRI && r0 == lClassRI: // LB30a
	case l0 == lClassEB && r0 == lClassEM: // LB30b
	default: // LB31
		return true
	}
	return false
}

/*
This functions is useful in case of IS does not mapped to SY.

// Returns is <runes..., l1, l0> ends with regexp "NU (SY IS)+".
func lIsNusyis(runes []rune, l1, l0 lClass) bool {
	if l0 != lClassSY && l0 != lClassIS {
		return false
	}
	switch l1 {
	case lClassNU:
		return true
	case lClassSY, lClassIS:
	default:
		return false
	}
	for len(runes) > 0 {
		c, _, pos := lLastSequence(runes)
		switch c {
		case lClassNU:
			return true
		case lClassSY, lClassIS:
		default:
			return false
		}
		runes = runes[:pos]
	}
	return false
}

// Same as without suffix "1" but with only one predefined class.
func lIsNusyis1(runes []rune, l0 lClass) bool {
	l1, _, pos := lLastSequence(runes)
	return lIsNusyis(runes[:pos], l1, l0)
}
*/

// Returns class of last rune in <runes..., l1> which is not equal to l0.
func lLastNotEqualTo(runes []rune, l1, l0 lClass) lClass {
	if l1 != l0 {
		return l1
	}
	for len(runes) > 0 {
		c, _, _, pos := lLastSequence(runes)
		if c != l0 {
			return c
		}
		runes = runes[:pos]
	}
	return lClassXX
}

// Same as without suffix "1" but with only one predefined class (looking in <runes...>, not <runes..., l1>).
func lLastNotEqualTo0(runes []rune, l0 lClass) lClass {
	l1, _, _, pos := lLastSequence(runes)
	return lLastNotEqualTo(runes[:pos], l1, l0)
}

// True if l0 is RI and it opens RI sequence in string <runes..., l1, l0, ...> (may be joined with next RI).
func lIsOpenRI(runes []rune, l1, l0 lClass) (r bool) {
	r = l0 == lClassRI
	if !r {
		return
	}
	r = l1 != lClassRI
	if r {
		return
	}
	for len(runes) > 0 {
		c, _, _, pos := lLastSequence(runes)
		if c != lClassRI {
			break
		}
		r = !r
		runes = runes[:pos]
	}
	return
}

// runes must be valid (len>1).
// l0Pos must be valid (in runes; really begin of sequence).
func lineBreakAfter(runes []rune, l0Pos int) int {
	l := len(runes)

	l1, _, _, l1Pos := lLastSequence(runes[:l0Pos])
	l0, l0IS, l0ZWJ, r0Delta := lFirstSequence(runes[l0Pos:])
	r0Pos := l0Pos + r0Delta
	r0, r0IS, r0ZWJ, r1Delta := lFirstSequence(runes[r0Pos:])
	l2Diff := lLastNotEqualTo0(runes[:l1Pos], l1)
	l1Diff := lLastNotEqualTo(runes[:l1Pos], l1, l0)
	lOddRI := lIsOpenRI(runes[:l1Pos], l1, l0)

	for r0Pos < l {
		r1, r1IS, r1ZWJ, r2Delta := lFirstSequence(runes[r0Pos+r1Delta:])

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
func LineBreakAfter(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	pos = lSequenceBegin(runes[:pos+1])

	return lineBreakAfter(runes, pos)
}

// runes must be valid (len>1).
// r0Pos must be valid (in runes; really begin of sequence).
func lineBreakBefore(runes []rune, r0Pos int) int {
	r0, _, _, r1Delta := lFirstSequence(runes[r0Pos:])
	r1, _, _, _ := lFirstSequence(runes[r0Pos+r1Delta:])
	l0, l0IS, l0ZWJ, l0Pos := lLastSequence(runes[:r0Pos])
	l1Diff := lLastNotEqualTo0(runes[:l0Pos], l0)
	l2Diff := l1Diff

	for r0Pos > 0 {
		l1, l1IS, l1ZWJ, l1Pos := lLastSequence(runes[:l0Pos])
		lOddRI := lIsOpenRI(runes[:l1Pos], l1, l0)
		if l2Diff == l1 {
			l2Diff = lLastNotEqualTo0(runes[:l1Pos], l1)
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
func LineBreakBefore(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return NoLineBreak
	}

	pos = lSequenceBegin(runes[:pos+1])

	return lineBreakBefore(runes, pos)
}

// FirstLineBreak returns first possible line break.
func FirstLineBreak(runes []rune) int {
	return LineBreakAfter(runes, 0)
}

// LastLineBreak returns last possible line break (except line break at end of string).
func LastLineBreak(runes []rune) int {
	return LineBreakBefore(runes, len(runes)-1)
}

// LineBreaks returns all possible line breaks.
func LineBreaks(runes []rune) (breaks []int) {
	l := len(runes)
	if l == 0 {
		return // []int{NoLineBreak}
	}
	breaks = make([]int, 0, l) // TODO memory efficient
	for pos := 0; pos < l; {
		length := FirstLineBreak(runes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}