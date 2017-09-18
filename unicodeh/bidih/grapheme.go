package bidih

import (
	"github.com/apaxa-go/helper/unicodeh/boundaryh"
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
	"sort"
	"github.com/apaxa-go/helper/unicodeh"
	"log"
)

func initClassesGraphemes(runes []rune, graphemes []boundaryh.Boundary, bases []int)(classes []bidi.Class){
	classes = make([]bidi.Class, len(graphemes))
	for graphemeI, g := range graphemes {
		classes[graphemeI] = bidi.GetClass(runes[g.From+bases[graphemeI]])
	}
	return classes
}

func implicitGraphemes(runes []rune, graphemes []boundaryh.Boundary, classes []bidi.Class, paragraphLevel EmbeddingLevel, levels []EmbeddingLevel, directions []bidi.Class, pairs map[int]pairElement) {
	sequences := computeRunSequences(classes, paragraphLevel, levels, directions, pairs)
	for _, sequence := range sequences {
		w1(sequence, directions)
		w23(sequence, directions)
		w4(sequence, directions)
		w5(sequence, directions)
		w6(sequence, directions)
		w7(sequence, directions)
		n0Graphemes(sequence, runes, graphemes, classes, levels, directions)
		n12(sequence, levels, directions)
	}
	i12(levels, directions)
}

func n0ComputeBracketPairsGraphemes(sequence runSequence, runes []rune, graphemes []boundaryh.Boundary) (bracketPairs []bracketPairElement) {
	type bracketStackElement struct {
		start sequenceState
		pair  rune
	}
	bracketStack := make([]bracketStackElement, 0, maxBracketStack)

	for sequence.start(); !sequence.isEnd(); sequence.next() {
		bracketPairShift := bidi.GetBracketPairShift(runes[graphemes[sequence.runeI].From])
		switch {
		case bracketPairShift > 0:
			if len(bracketStack) < maxBracketStack {
				bracketStack = append(bracketStack, bracketStackElement{sequence.saveState(), runes[graphemes[sequence.runeI].From] + bracketPairShift})
			} else {
				return
			}
		case bracketPairShift < 0:
			for stackI := len(bracketStack) - 1; stackI >= 0; stackI-- {
				if n0IsExpectedCloseBracket(bracketStack[stackI].pair, runes[graphemes[sequence.runeI].From]) {
					bracketPairs = append(bracketPairs, bracketPairElement{bracketStack[stackI].start, sequence.saveState()})
					bracketStack = bracketStack[:stackI]
					break
				}
			}
		}
	}

	sort.Slice(bracketPairs, func(i, j int) bool { return bracketPairs[i].start.runeI < bracketPairs[j].start.runeI })
	return
}


func n0Graphemes(sequence runSequence, runes []rune, graphemes []boundaryh.Boundary, classes []bidi.Class, levels []EmbeddingLevel, directions []bidi.Class) {
	bracketPairs := n0ComputeBracketPairsGraphemes(sequence, runes, graphemes)
	embeddingDirection := levels[sequence.ranges[0].from].direction()
	for _, bracketPair := range bracketPairs {
		sequence.loadState(bracketPair.start)
		replace, replaceWith := n0Decide(sequence, bracketPair.end.runeI, embeddingDirection, directions)
		if replace {
			directions[bracketPair.start.runeI] = replaceWith
			directions[bracketPair.end.runeI] = replaceWith
			sequence.loadState(bracketPair.start)
			n0FollowingNSM(sequence, classes, directions, replaceWith)
			sequence.loadState(bracketPair.end)
			n0FollowingNSM(sequence, classes, directions, replaceWith)
		}
	}
}


func parseParagraphGraphemes(runes []rune, graphemes []boundaryh.Boundary, classes []bidi.Class, paragraphLevel EmbeddingLevel, lineBreaker LineBreaker, order []int)(r []rune){
	if paragraphLevel == AutoParagraphDirection {
		paragraphLevel = p23(classes)
	}
	levels, directions, pairs, segmentSeparators := explicit(classes, paragraphLevel)
	//log.Println("Levels 10", levels)
	//log.Println("Directions 10", directions)
	implicitGraphemes(runes,graphemes, classes, paragraphLevel, levels, directions, pairs)
	//log.Println("Levels 11", levels)
	//log.Println("Directions 11", directions)
	shape()
	lineLengths := l1ComputeLines(runes, lineBreaker) // TODO pass graphemes to line breaker?
	reorderGraphemes(runes,graphemes, classes, paragraphLevel, levels, directions, lineLengths, segmentSeparators, order)
	//log.Println("Levels 12", levels)
	//log.Println("Directions 12", directions)
	//log.Println("Oreder 12", order)
	//
	// Reorder runes
	//
	r=make([]rune,len(runes))
	runeI:=0
	for graphemeI,g:=range graphemes{
		//log.Println("R 0", r)
		copy(r[runeI:],runes[g.From:g.To])
		//log.Println("R 1", r)
		if directions[order[graphemeI]]==bidi.RightToLeft{
			//log.Println("R reverse",g,g.Len())
			l2Reverse(r[runeI:runeI+g.Len()],nil)
		}
		//log.Println("R 2", r)
		runeI+=g.Len()
	}
	return
}

func l2Graphemes(graphemes []boundaryh.Boundary, levels []EmbeddingLevel, lineLengths []int, order []int) {
	// TODO init order in func does not work!!!
	//if order!=nil{
	//	order=initOrder(len(graphemes))
	//}
	if lineLengths == nil {
		l2LineGraphemes(graphemes, levels, order)
	} else {
		graphemeI := 0
		for _, lineLength := range lineLengths {
			nextRuneI := graphemeI + lineLength
			l2LineGraphemes(graphemes[graphemeI:nextRuneI], levels[graphemeI:nextRuneI], sliceOrder(order, graphemeI,nextRuneI))
			graphemeI = nextRuneI
		}
	}
}

func l2ReverseGraphemes(graphemes []boundaryh.Boundary, order []int) {
	if order==nil {
		for graphemeI, l := 0, len(graphemes); graphemeI < l/2; graphemeI++ {
			graphemes[graphemeI], graphemes[l-1-graphemeI] = graphemes[l-1-graphemeI], graphemes[graphemeI]
		}
	}else{
		for graphemeI, l := 0, len(graphemes); graphemeI < l/2; graphemeI++ {
			graphemes[graphemeI], graphemes[l-1-graphemeI] = graphemes[l-1-graphemeI], graphemes[graphemeI]
			order[graphemeI], order[l-1-graphemeI] = order[l-1-graphemeI], order[graphemeI]
		}
	}
}


func l2LineGraphemes(graphemes []boundaryh.Boundary, levels []EmbeddingLevel, order []int) {
	type embeddingLevelStatus struct {
		start int
		level EmbeddingLevel
	}
	//
	if len(graphemes) == 0 {
		return
	}
	stack := make([]embeddingLevelStatus, 0, MaxDepth)
	status := embeddingLevelStatus{0, levels[0]}
	for graphemeI := range graphemes {
		embeddingLevel := levels[graphemeI]
		if embeddingLevel > status.level {
			stack = append(stack, status)
			status = embeddingLevelStatus{graphemeI, embeddingLevel}
		} else if embeddingLevel < status.level {
			// reverse
			for status.level > embeddingLevel {
				if len(stack) == 0 || stack[len(stack)-1].level < embeddingLevel {
					// down to embeddingLevel directly
					if (status.level-embeddingLevel)%2 == 1 {
						l2ReverseGraphemes(graphemes[status.start:graphemeI], sliceOrder(order,status.start, graphemeI))
					}
					status.level = embeddingLevel
				} else {
					// down to stack
					if (status.level-stack[len(stack)-1].level)%2 == 1 {
						l2ReverseGraphemes(graphemes[status.start:graphemeI], sliceOrder(order,status.start, graphemeI))
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
			l2ReverseGraphemes(graphemes[status.start:], sliceOrder(order,status.start,len(order)))
		}
		status = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	// Final reordering
	if status.level%2 == 1 {
		l2ReverseGraphemes(graphemes[status.start:], sliceOrder(order,status.start,len(order)))
	}
}

func l4Graphemes(runes []rune, graphemes []boundaryh.Boundary, directions []bidi.Class) {
	for graphemeI, direction:=range directions{
		if direction == bidi.RightToLeft{
			for runeI:=range runes[graphemes[graphemeI].From:graphemes[graphemeI].To]{
				if mirror, ok := unicodeh.BidiMirroringGlyph[runes[runeI]]; ok {
					runes[runeI] = mirror
				}
			}
		}
	}
}

func reorderGraphemes(runes []rune, graphemes []boundaryh.Boundary, classes []bidi.Class, paragraphLevel EmbeddingLevel, levels []EmbeddingLevel, directions []bidi.Class, lineLengths, segmentSeparators []int, order []int) {
	l1(classes, paragraphLevel, levels, lineLengths, segmentSeparators)
	l1x9(classes, paragraphLevel, levels)
	l2Graphemes(graphemes, levels, lineLengths, order)
	l3()
	//l4Graphemes(runes, graphemes,directions) // TODO skip this in tests
}

func ParseParagraphGraphemes(runes []rune, paragraphLevel EmbeddingLevel, lineBreaker LineBreaker)(r []rune){
	graphemes,bases:= boundaryh.BoundariesRunesExtended(runes)
	log.Println("Graphemes 0",graphemes)
	log.Println("Bases 0",bases)
	classes:=initClassesGraphemes(runes,graphemes,bases)
	r=parseParagraphGraphemes(runes, graphemes, classes, paragraphLevel, lineBreaker,initOrder(len(graphemes)))
	//log.Println(runes,r)
	log.Println("Classes",classes)
	if len(runes)==-1{
		log.Println("")
	}
	return
}
