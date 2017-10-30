package bidih

import (
	"github.com/apaxa-go/helper/unicodeh"
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
)

// = FSI, LRI, RLI, and PDI
// and LRE, RLE, LRO, RLO, PDF and BN - for unprintable runes // TODO keep PDF here???
func isL1Replacebale(class bidi.Class) bool {
	return class == bidi.WhiteSpace || class == bidi.FirstStrongIsolate || class == bidi.LeftToRightIsolate || class == bidi.RightToLeftIsolate || class == bidi.PopDirectionalIsolate || class == bidi.LeftToRightEmbedding || class == bidi.RightToLeftEmbedding || class == bidi.LeftToRightOverride || class == bidi.RightToLeftOverride || class == bidi.PopDirectionalFormat || class == bidi.BoundaryNeutral
}

func l1ReplacePrevious(runeI int, classes []bidi.Class, paragraphLevel EmbeddingLevel, levels []EmbeddingLevel) {
	for runeI--; runeI >= 0 && isL1Replacebale(classes[runeI]); runeI-- {
		levels[runeI] = paragraphLevel
	}
}

func l1(classes []bidi.Class, paragraphLevel EmbeddingLevel, levels []EmbeddingLevel, lineLengths, segmentSeparators []int) {
	{
		runeI := 0
		for _, lineLength := range lineLengths {
			runeI += lineLength
			l1ReplacePrevious(runeI, classes, paragraphLevel, levels)
		}
	}
	for _, runeI := range segmentSeparators {
		levels[runeI] = paragraphLevel
		l1ReplacePrevious(runeI, classes, paragraphLevel, levels)
	}

	{
		runeI := len(classes) - 1
		if classes[runeI] == bidi.ParagraphSeparator {
			levels[runeI] = paragraphLevel
			l1ReplacePrevious(runeI, classes, paragraphLevel, levels)
		} else {
			l1ReplacePrevious(runeI+1, classes, paragraphLevel, levels)
		}
	}
}

func isX9(class bidi.Class) bool {
	return class == bidi.RightToLeftEmbedding || class == bidi.LeftToRightEmbedding || class == bidi.LeftToRightOverride || class == bidi.RightToLeftOverride || class == bidi.PopDirectionalFormat || class == bidi.BoundaryNeutral
}

func l1x9(classes []bidi.Class, paragraphLevel EmbeddingLevel, levels []EmbeddingLevel) {
	for runeI, rawClass := range classes {
		if isX9(rawClass) {
			levels[runeI] = paragraphLevel
		} else {
			paragraphLevel = levels[runeI]
		}
	}
}

func l1ComputeLines(runes []rune, lineBreaker LineBreaker) (lineLenths []int) {
	if lineBreaker == nil { // lineLengths == nil => no split to lines
		return nil
	}
	return lineBreaker(runes)
}

func l2Reverse(runes []rune, order []int) {
	if order == nil {
		for runeI, l := 0, len(runes); runeI < l/2; runeI++ {
			runes[runeI], runes[l-1-runeI] = runes[l-1-runeI], runes[runeI]
		}
	} else {
		for runeI, l := 0, len(runes); runeI < l/2; runeI++ {
			runes[runeI], runes[l-1-runeI] = runes[l-1-runeI], runes[runeI]
			order[runeI], order[l-1-runeI] = order[l-1-runeI], order[runeI]
		}
	}
}

func l2Line(runes []rune, levels []EmbeddingLevel, order []int) {
	type embeddingLevelStatus struct {
		start int
		level EmbeddingLevel
	}
	//
	if len(runes) == 0 {
		return
	}
	stack := make([]embeddingLevelStatus, 0, MaxDepth)
	status := embeddingLevelStatus{0, levels[0]}
	for runeI := range runes {
		embeddingLevel := levels[runeI]
		if embeddingLevel > status.level {
			stack = append(stack, status)
			status = embeddingLevelStatus{runeI, embeddingLevel}
		} else if embeddingLevel < status.level {
			// reverse
			for status.level > embeddingLevel {
				if len(stack) == 0 || stack[len(stack)-1].level < embeddingLevel {
					// down to embeddingLevel directly
					if (status.level-embeddingLevel)%2 == 1 {
						l2Reverse(runes[status.start:runeI], sliceOrder(order, status.start, runeI))
					}
					status.level = embeddingLevel
				} else {
					// down to stack
					if (status.level-stack[len(stack)-1].level)%2 == 1 {
						l2Reverse(runes[status.start:runeI], sliceOrder(order, status.start, runeI))
					}
					status = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	// Unpack stack
	for len(stack) > 0 {
		if (status.level-stack[len(stack)-1].level)%2 == 1 {
			l2Reverse(runes[status.start:], sliceOrder(order, status.start, len(order)))
		}
		status = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	// Final reordering
	if status.level%2 == 1 {
		l2Reverse(runes[status.start:], sliceOrder(order, status.start, len(order)))
	}
}

func initOrder(length int) (order []int) {
	order = make([]int, length)
	for i := range order {
		order[i] = i
	}
	return
}

func sliceOrder(order []int, from, to int) []int {
	if order == nil {
		return nil
	}
	return order[from:to]
}

func l2(runes []rune, levels []EmbeddingLevel, lineLengths []int, order []int) {
	// TODO init order in func does not work!!!
	//if order!=nil{
	//	order=initOrder(len(runes))
	//}
	if lineLengths == nil {
		l2Line(runes, levels, order)
	} else {
		runeI := 0
		for _, lineLength := range lineLengths {
			nextRuneI := runeI + lineLength
			l2Line(runes[runeI:nextRuneI], levels[runeI:nextRuneI], sliceOrder(order, runeI, nextRuneI))
			runeI = nextRuneI
		}
	}
}

func l3() {
	// TODO
}

func l4(runes []rune, directions []bidi.Class) {
	for runeI, direction := range directions {
		if direction == bidi.RightToLeft {
			if mirror, ok := unicodeh.BidiMirroringGlyph[runes[runeI]]; ok {
				runes[runeI] = mirror
			}
		}
	}
}

func reorder(runes []rune, classes []bidi.Class, paragraphLevel EmbeddingLevel, levels []EmbeddingLevel, directions []bidi.Class, lineLengths, segmentSeparators []int, order []int) {
	l1(classes, paragraphLevel, levels, lineLengths, segmentSeparators)
	l1x9(classes, paragraphLevel, levels)
	l2(runes, levels, lineLengths, order)
	l3()
	//l4(runes,directions) // TODO skip this only in tests
}
