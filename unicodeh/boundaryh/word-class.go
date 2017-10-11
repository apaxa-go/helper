package boundaryh

import "github.com/apaxa-go/helper/unicodeh"

type wClass uint8

const (
	wClassOther        wClass = iota
	wClassNewline      wClass = iota
	wClassZWJ          wClass = iota
	wClassGlueAfterZWJ wClass = iota
	wClassEBG          wClass = iota
	wClassExtend       wClass = iota
	wClassFormat       wClass = iota
	wClassALetter      wClass = iota
	wClassHebrewLetter wClass = iota
	wClassMidNumLet    wClass = iota
	wClassSingleQuote  wClass = iota
	wClassMidLetter    wClass = iota
	wClassDoubleQuote  wClass = iota
	wClassNumeric      wClass = iota
	wClassMidNum       wClass = iota
	wClassKatakana     wClass = iota
	wCLassExtendNumLet wClass = iota
	wClassEBase        wClass = iota
	wClassEModifier    wClass = iota
	wClassRI           wClass = iota
)

// TODO
func getWClass(r rune) wClass {
	switch {
	default:
		return wClassOther
	case unicodeh.IsWordBreakNewline(r) || r == crRune || r == lfRune:
		return wClassNewline
	case unicodeh.IsWordBreakZWJ(r):
		return wClassZWJ
	case unicodeh.IsWordBreakGlueAfterZwj(r):
		return wClassGlueAfterZWJ
	case unicodeh.IsWordBreakEBaseGAZ(r):
		return wClassEBG
	case unicodeh.IsWordBreakExtend(r):
		return wClassExtend
	case unicodeh.IsWordBreakFormat(r):
		return wClassFormat
	case unicodeh.IsWordBreakALetter(r):
		return wClassALetter
	case unicodeh.IsWordBreakHebrewLetter(r):
		return wClassHebrewLetter
	case unicodeh.IsWordBreakMidNumLet(r):
		return wClassMidNumLet
	case unicodeh.IsWordBreakSingleQuote(r):
		return wClassSingleQuote
	case unicodeh.IsWordBreakMidLetter(r):
		return wClassMidLetter
	case unicodeh.IsWordBreakDoubleQuote(r):
		return wClassDoubleQuote
	case unicodeh.IsWordBreakNumeric(r):
		return wClassNumeric
	case unicodeh.IsWordBreakMidNum(r):
		return wClassMidNum
	case unicodeh.IsWordBreakKatakana(r):
		return wClassKatakana
	case unicodeh.IsWordBreakExtendNumLet(r):
		return wCLassExtendNumLet
	case unicodeh.IsWordBreakEBase(r):
		return wClassEBase
	case unicodeh.IsWordBreakEModifier(r):
		return wClassEModifier
	case unicodeh.IsGraphemeClusterBreakRegionalIndicator(r):
		return wClassRI
	}
}

func (c wClass) isAHLetter() bool   { return c == wClassALetter || c == wClassHebrewLetter }
func (c wClass) isMidNumLetQ() bool { return c == wClassMidNumLet || c == wClassSingleQuote }
func (c wClass) isWB4() bool        { return c == wClassFormat || c == wClassExtend || c == wClassZWJ }
