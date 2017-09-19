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
func (c wClass) isWB4()bool{return c == wClassFormat || c == wClassExtend || c == wClassZWJ}

// Returns first not WB4 class and number of skipped WB4 runes.
func firstNotWB4(runes []rune)(c wClass, skip int){
	c=wClassOther
	for skip=0; skip<len(runes); skip++{
		c=getWClass(runes[skip])
		if !c.isWB4(){
			break
		}
	}
	return
}

// Skip WB4 and init classes.
// Skipping here because at the beginning of word WB4 must be saved in classes, but not in main loop (in the middle of word).
// len(runes) must be >=2
func firstWordInit(runes []rune)(prevPrevClass, prevClass, curClass wClass, pos int){
	prevPrevClass=wClassOther
	prevClass = getWClass(runes[0])
	curClass=getWClass(runes[1])
	pos = 1
	if prevClass.isWB4() {
		for pos < len(runes)-1 && curClass.isWB4() {
			pos++
			prevPrevClass = prevClass
			prevClass = curClass
			curClass = getWClass(runes[pos])
		}
	}
	return
}

func FirstWord(runes []rune) int {
	l := len(runes)
	if l <= 1 {
		return l
	}

	if runes[0] == crRune && runes[1] == lfRune { // WB3
		return 2
	}

	prevPrevClass,prevClass,curClass,i:=firstWordInit(runes)

	if prevClass == wClassNewline { // WB3a
		return 1
	}

	for ; i < l; i++ {
		nextClass,skip:= firstNotWB4(runes[i+1:])
		ignore := false
		switch {
		case curClass == wClassNewline: // WB3b
			return i
		case prevClass == wClassZWJ && (curClass == wClassGlueAfterZWJ || curClass == wClassEBG): // WB3c
		case curClass == wClassFormat || curClass == wClassExtend || curClass == wClassZWJ: // WB4 (part 1)
			ignore = true
		case prevClass.isAHLetter() && curClass.isAHLetter(): // WB5
		case prevClass.isAHLetter() && (curClass == wClassMidLetter || curClass.isMidNumLetQ()) && nextClass.isAHLetter(): // WB6
		case prevPrevClass.isAHLetter() && (prevClass == wClassMidLetter || prevClass.isMidNumLetQ()) && curClass.isAHLetter(): // WB7 // TODO combine with WB6
		case prevClass == wClassHebrewLetter && curClass == wClassSingleQuote: // WB7a
		case prevClass == wClassHebrewLetter && curClass == wClassDoubleQuote && nextClass == wClassHebrewLetter: // WB7b
		case prevPrevClass == wClassHebrewLetter && prevClass == wClassDoubleQuote && curClass == wClassHebrewLetter: // WB7b // TODO combine with WB7b
		case prevClass == wClassNumeric && curClass == wClassNumeric: // WB8
		case prevClass.isAHLetter() && curClass == wClassNumeric: // WB9
		case prevClass == wClassNumeric && curClass.isAHLetter(): // WB10
		case prevPrevClass == wClassNumeric && (prevClass == wClassMidNum || prevClass.isMidNumLetQ()) && curClass == wClassNumeric: // WB11
		case prevClass == wClassNumeric && (curClass == wClassMidNum || curClass.isMidNumLetQ()) && nextClass == wClassNumeric: // WB12 // TODO combine with WB11
		case prevClass == wClassKatakana && curClass == wClassKatakana: // WB13
		case (prevClass.isAHLetter() || prevClass == wClassNumeric || prevClass == wClassKatakana || prevClass == wCLassExtendNumLet) && curClass == wCLassExtendNumLet: // WB13a
		case prevClass == wCLassExtendNumLet && (curClass.isAHLetter() || curClass == wClassNumeric || curClass == wClassKatakana): // WB13b
		case (prevClass == wClassEBase || prevClass == wClassEBG) && curClass == wClassEModifier: // WB14
		case prevPrevClass != wClassRI && prevClass == wClassRI && curClass == wClassRI: // WB15 & WB16
		default:
			return i
		}
		if !ignore { // WB4 (part 2)
			prevPrevClass = prevClass
			prevClass = curClass
		}
		curClass = nextClass
		i+=skip
	}
	return l
}

func Words(runes []rune) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(runes)) // TODO memory efficient
	for i := 0; i < len(runes); {
		length := FirstWord(runes[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}
