package bidih

import (
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
)

// Rules P2 & P3
func p23(classes []bidi.Class) EmbeddingLevel {
	isolateLevel := 0
	for _, class := range classes {
		if isIsolateInitiator(class) {
			isolateLevel++
		} else if class == bidi.PopDirectionalIsolate {
			if isolateLevel > 0 {
				isolateLevel--
			}
		} else if isolateLevel == 0 {
			switch class {
			case bidi.ArabicLetter, bidi.RightToLeft:
				return RightToLeftParagraph
			case bidi.LeftToRight:
				return LeftToRightParagraph
			}
		}
	}
	return LeftToRightParagraph
}

type LineBreaker func(runes []rune) (lineLengths []int)

// order == nil => do not save order; order != nil (with any length) => save order.
func parseParagraph(runes []rune, classes []bidi.Class, paragraphLevel EmbeddingLevel, lineBreaker LineBreaker, order []int){
	if paragraphLevel == AutoParagraphDirection {
		paragraphLevel = p23(classes)
	}
	levels, directions, pairs, segmentSeparators := explicit(classes, paragraphLevel)
	implicit(runes, classes, paragraphLevel, levels, directions, pairs)
	shape()
	lineLengths := l1ComputeLines(runes, lineBreaker)
	reorder(runes, classes, paragraphLevel, levels, directions, lineLengths, segmentSeparators, order)
}

func initClasses(runes []rune)(classes []bidi.Class){
	classes = make([]bidi.Class, len(runes))
	for runeI, r := range runes {
		classes[runeI] = bidi.GetClass(r)
	}
	return classes
}

func ParseParagraph(runes []rune, paragraphLevel EmbeddingLevel, lineBreaker LineBreaker){
	classes:=initClasses(runes)
	parseParagraph(runes, classes, paragraphLevel, lineBreaker,nil)
}

func ParseParagraphExtended(runes []rune, paragraphLevel EmbeddingLevel, lineBreaker LineBreaker)(order []int){
	classes:=initClasses(runes)
	order=[]int{}	// real initialization will done automatically, here we just must pass non-nil value.
	parseParagraph(runes, classes, paragraphLevel, lineBreaker,order)
	return
}