package boundaryh

import "github.com/apaxa-go/helper/unicodeh"

type lClass uint8

const (
	lClassAI  lClass = iota
	lClassAL  lClass = iota
	lClassB2  lClass = iota
	lClassBA  lClass = iota
	lClassBB  lClass = iota
	lClassBK  lClass = iota
	lClassCB  lClass = iota
	lClassCJ  lClass = iota
	lClassCL  lClass = iota
	lClassCM  lClass = iota
	lClassCP  lClass = iota
	lClassEB  lClass = iota
	lClassEM  lClass = iota
	lClassEX  lClass = iota
	lClassGL  lClass = iota
	lClassH2  lClass = iota
	lClassH3  lClass = iota
	lClassHL  lClass = iota
	lClassHY  lClass = iota
	lClassID  lClass = iota
	lClassIN  lClass = iota
	lClassIS  lClass = iota
	lClassJL  lClass = iota
	lClassJT  lClass = iota
	lClassJV  lClass = iota
	lClassNL  lClass = iota
	lClassNS  lClass = iota
	lClassNU  lClass = iota
	lClassOP  lClass = iota
	lClassPO  lClass = iota
	lClassPR  lClass = iota
	lClassQU  lClass = iota
	lClassRI  lClass = iota
	lClassSA  lClass = iota
	lClassSG  lClass = iota
	lClassSP  lClass = iota
	lClassSY  lClass = iota
	lClassWJ  lClass = iota
	lClassXX  lClass = iota
	lClassZW  lClass = iota
	lClassZWJ lClass = iota
)

func getLClassRaw(r rune) lClass {
	switch {
	default:
		return lClassXX
	case unicodeh.IsLineBreakAmbiguous(r):
		return lClassAI
	case unicodeh.IsLineBreakAlphabetic(r):
		return lClassAL
	case unicodeh.IsLineBreakBreakBoth(r):
		return lClassB2
	case unicodeh.IsLineBreakBreakAfter(r):
		return lClassBA
	case unicodeh.IsLineBreakBreakBefore(r):
		return lClassBB
	case unicodeh.IsLineBreakMandatoryBreak(r):
		return lClassBK
	case unicodeh.IsLineBreakContingentBreak(r):
		return lClassCB
	case unicodeh.IsLineBreakConditionalJapaneseStarter(r):
		return lClassCJ
	case unicodeh.IsLineBreakClosePunctuation(r):
		return lClassCL
	case unicodeh.IsLineBreakCombiningMark(r):
		return lClassCM
	case unicodeh.IsLineBreakCloseParenthesis(r):
		return lClassCP
	case unicodeh.IsLineBreakEBase(r):
		return lClassEB
	case unicodeh.IsLineBreakEModifier(r):
		return lClassEM
	case unicodeh.IsLineBreakExclamation(r):
		return lClassEX
	case unicodeh.IsLineBreakGlue(r):
		return lClassGL
	case unicodeh.IsLineBreakH2(r):
		return lClassH2
	case unicodeh.IsLineBreakH3(r):
		return lClassH3
	case unicodeh.IsLineBreakHebrewLetter(r):
		return lClassHL
	case unicodeh.IsLineBreakHyphen(r):
		return lClassHY
	case unicodeh.IsLineBreakIdeographic(r):
		return lClassID
	case unicodeh.IsLineBreakInseparable(r):
		return lClassIN
	case unicodeh.IsLineBreakInfixNumeric(r):
		return lClassIS
	case unicodeh.IsLineBreakJL(r):
		return lClassJL
	case unicodeh.IsLineBreakJT(r):
		return lClassJT
	case unicodeh.IsLineBreakJV(r):
		return lClassJV
	case unicodeh.IsLineBreakNextLine(r):
		return lClassNL
	case unicodeh.IsLineBreakNonstarter(r):
		return lClassNS
	case unicodeh.IsLineBreakNumeric(r):
		return lClassNU
	case unicodeh.IsLineBreakOpenPunctuation(r):
		return lClassOP
	case unicodeh.IsLineBreakPostfixNumeric(r):
		return lClassPO
	case unicodeh.IsLineBreakPrefixNumeric(r):
		return lClassPR
	case unicodeh.IsLineBreakQuotation(r):
		return lClassQU
	case unicodeh.IsLineBreakRegionalIndicator(r):
		return lClassRI
	case unicodeh.IsLineBreakComplexContext(r):
		return lClassSA
	case unicodeh.IsLineBreakSurrogate(r):
		return lClassSG
	case unicodeh.IsLineBreakSpace(r):
		return lClassSP
	case unicodeh.IsLineBreakBreakSymbols(r):
		return lClassSY
	case unicodeh.IsLineBreakWordJoiner(r):
		return lClassWJ
	case unicodeh.IsLineBreakZWSpace(r):
		return lClassZW
	case unicodeh.IsLineBreakZWJ(r):
		return lClassZWJ
	}
}

func getLClass(r rune) lClass { // LB1
	switch res := getLClassRaw(r); res {
	case lClassAI, lClassSG, lClassXX:
		return lClassAL
	case lClassSA:
		if unicodeh.IsGeneralCategoryNonspacingMark(r) || unicodeh.IsGeneralCategorySpacingMark(r) {
			return lClassCM
		}
		return lClassAL
	case lClassCJ:
		return lClassNS
	default:
		return res
	}

}

func (c lClass) isNotSkipBase() bool {
	return c == lClassBK || c == lClassNL || c == lClassSP || c == lClassZW
}                             // LB9
func (c lClass) isSkip() bool { return c == lClassCM || c == lClassZWJ } // LB9

// "c" contains resolved class after applying LB9 & LB10.
// rawZWJ required for LB8a. For simplify LB10 applied out-of-order and "c" unable say is class ZWJ before applying LB10.
func getLClassSkip(runes []rune) (c lClass, rawZWJ bool, i int) {
	l := len(runes)
	if l == 0 {
		c = lClassXX
		return
	}

	if runes[0] == crRune { // LB4
		if l > 1 && runes[1] == lfRune { // LB5
			return lClassBK, false, 2
		}
		return lClassBK, false, 1
	}

	if runes[0] == lfRune { // LB4
		return lClassBK, false, 1
	}

	c = getLClass(runes[0])

	if c.isSkip() { // LB10
		return lClassAL, c == lClassZWJ, 1
	} else if c.isNotSkipBase() { // LB9
		return c, false, 1
	}

	for i = 1; i < len(runes); i++ { // LB9
		if !getLClass(runes[i]).isSkip() {
			break
		}
	}

	return
}

func FirstLine(runes []rune) int {
	l := len(runes)
	if l <= 1 {
		return l
	}

	prevPrevPrevDifferentL25NU := false // Catch "_NU_ (NU | SY | IS)+ (CL | CP) × ..."
	prevPrevDifferentL25NU := false     // Catch "_NU_ (NU | SY | IS)+ × ..."
	prevPrevDifferentClass := lClassXX
	prevPrevClass := lClassXX
	prevClass, prevZWJ, i := getLClassSkip(runes)
	curClass, curZWJ, deltaI := getLClassSkip(runes[i:])

	for i < l {
		nextClass, nextZWJ, nextDeltaI := getLClassSkip(runes[i+deltaI:])
		switch {
		case prevClass == lClassNL || prevClass == lClassBK: // LB4 & LB5
			return i
		case curClass == lClassNL || curClass == lClassBK: // LB6
		case curClass == lClassSP || curClass == lClassZW: // LB7
		case prevClass == lClassZW: // LB8 (part 1)
			return i
		case prevPrevDifferentClass == lClassZW && prevClass == lClassSP: // LB8 (part 2)
			return i
		case prevZWJ && (curClass == lClassID || curClass == lClassEB || curClass == lClassEM): // LB8a
		case prevClass == lClassWJ || curClass == lClassWJ: // LB 11
		case prevClass == lClassGL: // LB12
		case (prevClass != lClassSP && prevClass != lClassBA && prevClass != lClassHY) && curClass == lClassGL: // LB12a
		case curClass == lClassCL || curClass == lClassCP || curClass == lClassEX || curClass == lClassIS || curClass == lClassSY: // LB13
		case prevClass == lClassOP: // LB14 (part 1)
		case prevPrevDifferentClass == lClassOP && prevClass == lClassSP: // LB14 (part 2)
		case prevClass == lClassQU && curClass == lClassOP: // LB15 (part 1)
		case prevPrevDifferentClass == lClassQU && prevClass == lClassSP && curClass == lClassOP: // LB15 (part 2)
		case (prevClass == lClassCL || prevClass == lClassCP) && curClass == lClassNS: // LB16 (part 1)
		case (prevPrevDifferentClass == lClassCL || prevPrevDifferentClass == lClassCP) && prevClass == lClassSP && curClass == lClassNS: // LB16 (part 2)
		case prevClass == lClassB2 && curClass == lClassB2: // LB17 (part 1)
		case prevPrevDifferentClass == lClassB2 && prevClass == lClassSP && curClass == lClassB2: // LB17 (part 2)
		case prevClass == lClassSP: // LB18
			return i
		case prevClass == lClassQU || curClass == lClassQU: // LB19
		case prevClass == lClassCB || curClass == lClassCB: // LB20 // TODO condition check!!!
			return i
		case curClass == lClassBA || curClass == lClassHY || curClass == lClassNS || prevClass == lClassBB: // LB21
		case prevPrevClass == lClassHL && (prevClass == lClassHY || prevClass == lClassBA): // LB21a
		case prevClass == lClassSY && curClass == lClassHL: // LB21b
		case (prevClass == lClassAL || prevClass == lClassHL || prevClass == lClassEX || prevClass == lClassID || prevClass == lClassEB || prevClass == lClassEM || prevClass == lClassIN || prevClass == lClassNU) && curClass == lClassIN: // LB22
		case (prevClass == lClassAL || prevClass == lClassHL) && curClass == lClassNU: // LB23 (part 1)
		case prevClass == lClassNU && (curClass == lClassAL || curClass == lClassHL): // LB23 (part 2)
		case prevClass == lClassPR && (curClass == lClassID || curClass == lClassEB || curClass == lClassEM): // LB23a (part 1)
		case (prevClass == lClassID || prevClass == lClassEB || prevClass == lClassEM) && curClass == lClassPO: // LB23a (part 2)
		case (prevClass == lClassPR || prevClass == lClassPO) && (curClass == lClassAL || curClass == lClassHL): // LB24 (part 1)
		case (prevClass == lClassAL || prevClass == lClassHL) && (curClass == lClassPR || curClass == lClassPO): // LB24 (part 2)
		/*
			// This is the default implementation of LB25.
			// Regexp "( PR | PO) ? ( OP | HY ) ? NU (NU | SY | IS) * ( CL | CP ) ? ( PR | PO) ?" suggested instead.
			// Also LineBreakTest.txt use this regexp, not default implementation.
			case (prevClass == lClassCL && curClass == lClassPO) ||
				(prevClass == lClassCP && curClass == lClassPO) ||
				(prevClass == lClassCL && curClass == lClassPR) ||
				(prevClass == lClassCP && curClass == lClassPR) ||
				(prevClass == lClassNU && curClass == lClassPO) ||
				(prevClass == lClassNU && curClass == lClassPR) ||
				(prevClass == lClassPO && curClass == lClassOP) ||
				(prevClass == lClassPO && curClass == lClassNU) ||
				(prevClass == lClassPR && curClass == lClassOP) ||
				(prevClass == lClassPR && curClass == lClassNU) ||
				(prevClass == lClassHY && curClass == lClassNU) ||
				(prevClass == lClassIS && curClass == lClassNU) ||
				(prevClass == lClassNU && curClass == lClassNU) ||
				(prevClass == lClassSY && curClass == lClassNU): // LB25
		*/

		// Regexp "( PR | PO) ? ( OP | HY ) ? NU (NU | SY | IS) * ( CL | CP ) ? ( PR | PO) ?" implementation of LB25.
		// See TR14 8.2 Example 7 for more information.
		case ((prevClass == lClassPR || prevClass == lClassPO) && curClass == lClassNU) ||
			((prevClass == lClassPR || prevClass == lClassPO) && (curClass == lClassOP || curClass == lClassHY) && nextClass == lClassNU) ||
			((prevClass == lClassOP || prevClass == lClassHY) && curClass == lClassNU) ||
			(prevClass == lClassNU && (curClass == lClassNU || curClass == lClassSY || curClass == lClassIS)) ||
			(prevClass == lClassNU && (curClass == lClassNU || curClass == lClassSY || curClass == lClassIS || curClass == lClassCL || curClass == lClassCP)) || // "NU × (NU | SY | IS | CL | CP )"
			(prevPrevDifferentL25NU && (prevClass == lClassNU || prevClass == lClassSY || prevClass == lClassIS) && (curClass == lClassNU || curClass == lClassSY || curClass == lClassIS || curClass == lClassCL || curClass == lClassCP)) || // "NU (NU | SY | IS)+ × (NU | SY | IS | CL | CP )"
			(prevClass == lClassNU && (curClass == lClassPO || curClass == lClassPR)) || // "NU × (PO | PR)"
			(prevPrevClass == lClassNU && (prevClass == lClassCL || prevClass == lClassCP) && (curClass == lClassPO || curClass == lClassPR)) || // "NU (CL | CP) × (PO | PR)"
			(prevPrevDifferentL25NU && (prevClass == lClassNU || prevClass == lClassSY || prevClass == lClassIS) && (curClass == lClassPO || curClass == lClassPR)) || // "NU (NU | SY | IS)+ × (PO | PR)"
			(prevPrevPrevDifferentL25NU && (prevPrevClass == lClassNU || prevPrevClass == lClassSY || prevPrevClass == lClassIS) && (prevClass == lClassCL || prevClass == lClassCP) && (curClass == lClassPO || curClass == lClassPR)): // "NU (NU | SY | IS)+ (CL | CP) × (PO | PR)" ; LB25
		// End of LB25
		case prevClass == lClassJL && (curClass == lClassJL || curClass == lClassJV || curClass == lClassH2 || curClass == lClassH3): // LB26 (part 1)
		case (prevClass == lClassJV || prevClass == lClassH2) && (curClass == lClassJV || curClass == lClassJT): // LB26 (part 2)
		case (prevClass == lClassJT || prevClass == lClassH3) && curClass == lClassJT: // LB26 (part 3)
		case (prevClass == lClassJL || prevClass == lClassJV || prevClass == lClassJT || prevClass == lClassH2 || prevClass == lClassH3) && (curClass == lClassIN || curClass == lClassPO): // LB27 (part 1)
		case prevClass == lClassPR && (curClass == lClassJL || curClass == lClassJV || curClass == lClassJT || curClass == lClassH2 || curClass == lClassH3): // LB27 (part 2)
		case (prevClass == lClassAL || prevClass == lClassHL || prevClass == lClassIS) && (curClass == lClassAL || curClass == lClassHL): // LB28 & LB29
		case (prevClass == lClassAL || prevClass == lClassHL || prevClass == lClassNU) && curClass == lClassOP: // LB30 (part 1)
		case prevClass == lClassCP && (curClass == lClassAL || curClass == lClassHL || curClass == lClassNU): // LB30 (part 2)
		case prevPrevClass != lClassRI && prevClass == lClassRI && curClass == lClassRI: // LB30a
		case prevClass == lClassEB && curClass == lClassEM: // LB30b
		default: // LB31
			return i
		}

		i += deltaI
		deltaI = nextDeltaI

		prevZWJ = curZWJ
		curZWJ = nextZWJ

		prevPrevPrevDifferentL25NU = prevPrevDifferentL25NU && (curClass == lClassCL || curClass == lClassCP)
		prevPrevDifferentL25NU = (prevPrevDifferentL25NU || prevClass == lClassNU) && (curClass == lClassNU || curClass == lClassSY || curClass == lClassIS)

		prevPrevClass = prevClass
		prevClass = curClass
		if prevPrevClass != prevClass {
			prevPrevDifferentClass = prevPrevClass
		}
		curClass = nextClass
	}
	return l
}

func Lines(runes []rune) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(runes)) // TODO memory efficient
	for i := 0; i < len(runes); {
		length := FirstLine(runes[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}
