package bidih

import (
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
)

// Maximum explicit embedding level.
const MaxDepth = 125

// Primary it is isolate pairs, but may also be overriding or embedding pair.
// By tr9 it should be only isolates.
type pairElement struct {
	pair    int  // position of pair
	isolate bool // if pair is isolate
}

func nextRTLLevel(level EmbeddingLevel) EmbeddingLevel {
	return level + 1 + level%2
}
func nextLTRLevel(level EmbeddingLevel) EmbeddingLevel {
	return level + 2 - level%2
}

// Similar to rules p2 & p3, but stops on related PDI.
// Returns 0 or 1.
func computeIsolateEmbeddingLevel(estimatedClasses []bidi.Class) EmbeddingLevel {
	isolateLevel := 0
	for _, class := range estimatedClasses {
		if isIsolateInitiator(class) {
			isolateLevel++
		} else if class == bidi.PopDirectionalIsolate {
			if isolateLevel == 0 {
				break
			}
			isolateLevel--
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

// Returns runes embedding levels & "effective" classes.
func explicit(classes []bidi.Class, paragraphLevel EmbeddingLevel) (levels []EmbeddingLevel, directions []bidi.Class, pairs map[int]pairElement, segmentSeparators []int) {
	//
	// Init result
	//
	levels = make([]EmbeddingLevel, len(classes)) // embedding levels
	directions = make([]bidi.Class, len(classes)) // aka "current class", "effective class" and "direction"
	pairs = make(map[int]pairElement)             // "isolate" pairs

	//
	// Define internal types and constants
	//
	type directionalOverrideStatus uint8
	const (
		overrideStatusNeutral directionalOverrideStatus = iota
		overrideStatusRTL     directionalOverrideStatus = iota
		overrideStatusLTR     directionalOverrideStatus = iota
	)
	type directionalStatus struct {
		embeddingLevel EmbeddingLevel
		overrideStatus directionalOverrideStatus
		isolateStatus  bool
		runeI          int // for pairs only
	}

	//
	// Define runtime internal variables
	//
	var (
		overflowIsolateCount   int
		overflowEmbeddingCount int
		validIsolateCount      int
		stack                  []directionalStatus
	)
	stack = make([]directionalStatus, 1, MaxDepth+2)
	stack[0] = directionalStatus{paragraphLevel, overrideStatusNeutral, false, -1}

	//
	// Run algorithm
	//
	for runeI, class := range classes {
		//
		// Compute segments separators for future uses.
		//
		if class == bidi.SegmentSeparator {
			segmentSeparators = append(segmentSeparators, runeI)
		}
		//
		// Apply rules X2-X8
		//
		var level EmbeddingLevel
		switch class {
		case bidi.RightToLeftEmbedding: // X2
			level = stack[len(stack)-1].embeddingLevel // For hidden characters only
			class = bidi.BoundaryNeutral               // For hidden characters only (instead of rule X9)
			newLevel := nextRTLLevel(level)
			if newLevel <= MaxEmbeddingLevel && overflowIsolateCount == 0 && overflowEmbeddingCount == 0 {
				stack = append(stack, directionalStatus{newLevel, overrideStatusNeutral, false, runeI})
			} else if overflowIsolateCount == 0 {
				overflowEmbeddingCount++
			}
		case bidi.LeftToRightEmbedding: // X3
			level = stack[len(stack)-1].embeddingLevel // For hidden characters only
			class = bidi.BoundaryNeutral               // For hidden characters only (instead of rule X9)
			newLevel := nextLTRLevel(level)
			if newLevel <= MaxEmbeddingLevel && overflowIsolateCount == 0 && overflowEmbeddingCount == 0 {
				stack = append(stack, directionalStatus{newLevel, overrideStatusNeutral, false, runeI})
			} else if overflowIsolateCount == 0 {
				overflowEmbeddingCount++
			}
		case bidi.RightToLeftOverride: // X4
			level = stack[len(stack)-1].embeddingLevel // For hidden characters only
			class = bidi.BoundaryNeutral               // For hidden characters only (instead of rule X9)
			newLevel := nextRTLLevel(level)
			if newLevel <= MaxEmbeddingLevel && overflowIsolateCount == 0 && overflowEmbeddingCount == 0 {
				stack = append(stack, directionalStatus{newLevel, overrideStatusRTL, false, runeI})
			} else if overflowIsolateCount == 0 {
				overflowEmbeddingCount++
			}
		case bidi.LeftToRightOverride: // X5
			level = stack[len(stack)-1].embeddingLevel // For hidden characters only
			class = bidi.BoundaryNeutral               // For hidden characters only (instead of rule X9)
			newLevel := nextLTRLevel(level)
			if newLevel <= MaxEmbeddingLevel && overflowIsolateCount == 0 && overflowEmbeddingCount == 0 {
				stack = append(stack, directionalStatus{newLevel, overrideStatusLTR, false, runeI})
			} else if overflowIsolateCount == 0 {
				overflowEmbeddingCount++
			}
		case bidi.RightToLeftIsolate: // X5a
			level = stack[len(stack)-1].embeddingLevel
			switch stack[len(stack)-1].overrideStatus {
			case overrideStatusLTR:
				class = bidi.LeftToRight
			case overrideStatusRTL:
				class = bidi.RightToLeft
			}
			newLevel := nextRTLLevel(level)
			if newLevel <= MaxEmbeddingLevel && overflowIsolateCount == 0 && overflowEmbeddingCount == 0 {
				validIsolateCount++
				stack = append(stack, directionalStatus{newLevel, overrideStatusNeutral, true, runeI})
			} else {
				overflowIsolateCount++
			}
		case bidi.LeftToRightIsolate: // X5b
			level = stack[len(stack)-1].embeddingLevel
			switch stack[len(stack)-1].overrideStatus {
			case overrideStatusLTR:
				class = bidi.LeftToRight
			case overrideStatusRTL:
				class = bidi.RightToLeft
			}
			newLevel := nextLTRLevel(level)
			if newLevel <= MaxEmbeddingLevel && overflowIsolateCount == 0 && overflowEmbeddingCount == 0 {
				validIsolateCount++
				stack = append(stack, directionalStatus{newLevel, overrideStatusNeutral, true, runeI})
			} else {
				overflowIsolateCount++
			}
		case bidi.FirstStrongIsolate: // X5c
			level = stack[len(stack)-1].embeddingLevel
			switch stack[len(stack)-1].overrideStatus {
			case overrideStatusLTR:
				class = bidi.LeftToRight
			case overrideStatusRTL:
				class = bidi.RightToLeft
			}
			newLevel := level + 2 - (level+computeIsolateEmbeddingLevel(classes[runeI+1:]))%2 // computeIsolateEmbeddingLevel: 0 => nextLTRLevel; 1 => nextRTLLevel
			if newLevel <= MaxEmbeddingLevel && overflowIsolateCount == 0 && overflowEmbeddingCount == 0 {
				validIsolateCount++
				stack = append(stack, directionalStatus{newLevel, overrideStatusNeutral, true, runeI})
			} else {
				overflowIsolateCount++
			}
		default: // X6
			level = stack[len(stack)-1].embeddingLevel
			switch stack[len(stack)-1].overrideStatus {
			case overrideStatusLTR:
				class = bidi.LeftToRight
			case overrideStatusRTL:
				class = bidi.RightToLeft
			}
		case bidi.BoundaryNeutral: // Not by tr9 5.2 (required full X6), but looks reasonable (fix, for example, see 119.1072.0).
			// BoundaryNeutral here for hidden characters only.
			level = stack[len(stack)-1].embeddingLevel
		case bidi.PopDirectionalIsolate: // X6a
			if overflowIsolateCount > 0 {
				overflowIsolateCount--
			} else if validIsolateCount == 0 {
			} else {
				overflowEmbeddingCount = 0
				for !stack[len(stack)-1].isolateStatus {
					stack = stack[:len(stack)-1]
				}
				// save isolatePair
				isolatePair := stack[len(stack)-1].runeI
				pairs[runeI] = pairElement{isolatePair, true}
				pairs[isolatePair] = pairElement{runeI, true}
				//
				stack = stack[:len(stack)-1]
				validIsolateCount--
			}
			level = stack[len(stack)-1].embeddingLevel
			switch stack[len(stack)-1].overrideStatus {
			case overrideStatusLTR:
				class = bidi.LeftToRight
			case overrideStatusRTL:
				class = bidi.RightToLeft
			}
		case bidi.PopDirectionalFormat: // X7
			class = bidi.BoundaryNeutral // For hidden characters only (instead of rule X9)
			if overflowIsolateCount == 0 {
				if overflowEmbeddingCount > 0 {
					overflowEmbeddingCount--
				} else if !stack[len(stack)-1].isolateStatus && len(stack) >= 2 {
					// save pair
					pair := stack[len(stack)-1].runeI
					pairs[runeI] = pairElement{pair, false}
					pairs[pair] = pairElement{runeI, false}
					//
					stack = stack[:len(stack)-1]
				}
			}
			level = stack[len(stack)-1].embeddingLevel
		case bidi.ParagraphSeparator: // X8
			level = paragraphLevel
		}
		levels[runeI] = level
		directions[runeI] = class
	}
	return
}
