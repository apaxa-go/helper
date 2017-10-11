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

// TODO
func lGetClassRaw(r rune) lClass {
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
	case unicodeh.IsLineBreakMandatoryBreak(r) || r == crRune || r == lfRune:
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

func lGetClass(r rune) lClass { // LB1
	switch res := lGetClassRaw(r); res {
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

func (c lClass) isSkip() bool { return c == lClassCM || c == lClassZWJ } // LB9
func (c lClass) isSkipBase() bool { // LB9
	return c != lClassBK && c != lClassNL && c != lClassSP && c != lClassZW
}
