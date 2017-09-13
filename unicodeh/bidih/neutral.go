package bidih

import (
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
	"sort"
)

const maxBracketStack = 63

// = B, S, WS, ON, FSI, LRI, RLI or PDI
func isNeutralOrIsolate(class bidi.Class) bool {
	return class == bidi.ParagraphSeparator || class == bidi.SegmentSeparator || class == bidi.WhiteSpace || class == bidi.OtherNeutral || class == bidi.FirstStrongIsolate || class == bidi.LeftToRightIsolate || class == bidi.RightToLeftIsolate || class == bidi.PopDirectionalIsolate
}

// = NI or BN
func isPossibleNeutralOrIsolate(class bidi.Class) bool {
	return isNeutralOrIsolate(class) || class == bidi.BoundaryNeutral
}

// = L, R, AL, EN or AN
func isN0Strong(class bidi.Class) bool {
	return isStrong(class) || class == bidi.EuropeanNumber || class == bidi.ArabicNumber
}

func n0Direction(class bidi.Class) bidi.Class {
	if class == bidi.ArabicLetter || class == bidi.EuropeanNumber || class == bidi.ArabicNumber {
		class = bidi.RightToLeft
	}
	return class
}

type bracketPairElement struct {
	start sequenceState
	end   sequenceState
}

func n0IsExpectedCloseBracket(expectedRune, gotRune rune) bool {
	// TODO assumption about canonical equivalent brackets
	return expectedRune == gotRune || (expectedRune == 0x232a && gotRune == 0x3009) || (expectedRune == 0x3009 && gotRune == 0x232a) // Special case for conical equivalent brackets.
}

func n0ComputeBracketPairs(sequence runSequence, runes []rune) (bracketPairs []bracketPairElement) {
	type bracketStackElement struct {
		start sequenceState
		pair  rune
	}
	bracketStack := make([]bracketStackElement, 0, maxBracketStack)

	for sequence.start(); !sequence.isEnd(); sequence.next() {
		bracketPairShift := bidi.GetBracketPairShift(runes[sequence.runeI])
		switch {
		case bracketPairShift > 0:
			if len(bracketStack) < maxBracketStack {
				bracketStack = append(bracketStack, bracketStackElement{sequence.saveState(), runes[sequence.runeI] + bracketPairShift})
			} else {
				return
			}
		case bracketPairShift < 0:
			for stackI := len(bracketStack) - 1; stackI >= 0; stackI-- {
				if n0IsExpectedCloseBracket(bracketStack[stackI].pair, runes[sequence.runeI]) {
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

// Inspects enclosed runes for strong (as N0 defines) direction.
// (sequence.runeI; end)
func n0InspectEnclosed(sequence runSequence, end int, embeddingDirection bidi.Class, directions []bidi.Class) (strongFound bool, toEmbeddingDirection bool) {
	for sequence.next(); !sequence.isEnd() && sequence.runeI < end; sequence.next() {
		if isN0Strong(directions[sequence.runeI]) {
			strongFound = true
		}
		if n0Direction(directions[sequence.runeI]) == embeddingDirection {
			return true, true
		}
	}
	return
}

// Returns previous strong (as N0 defines) direction.
// [sos; sequence.runeI)
func n0InspectPrevious(sequence runSequence, directions []bidi.Class) (previousStrongDirection bidi.Class) {
	for sequence.prev(); !sequence.isStart() && !isN0Strong(directions[sequence.runeI]); sequence.prev() {
	}
	if sequence.isStart() {
		return sequence.sos
	}
	return n0Direction(directions[sequence.runeI])
}

func n0Decide(sequence runSequence, end int, embeddingDirection bidi.Class, directions []bidi.Class) (replace bool, replaceWith bidi.Class) {
	strongFound, toEmbeddingDirection := n0InspectEnclosed(sequence, end, embeddingDirection, directions)
	if toEmbeddingDirection {
		return true, embeddingDirection
	}
	if !strongFound {
		return false, 0
	}
	return true, n0InspectPrevious(sequence, directions)
}

func n0FollowingNSM(sequence runSequence, classes []bidi.Class, directions []bidi.Class, replaceWith bidi.Class) {
	for sequence.next(); !sequence.isEnd() && classes[sequence.runeI] == bidi.NonSpacingMark; sequence.next() {
		directions[sequence.runeI] = replaceWith
	}
}

// rule N0
func n0(sequence runSequence, runes []rune, classes []bidi.Class, levels []EmbeddingLevel, directions []bidi.Class) {
	bracketPairs := n0ComputeBracketPairs(sequence, runes)
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

func n1GetSequence(sequence *runSequence, directions []bidi.Class) bool {
	if directions[sequence.runeI] == bidi.BoundaryNeutral {
		for sequence.next(); !sequence.isEnd() && directions[sequence.runeI] == bidi.BoundaryNeutral; sequence.next() {
		}
		if sequence.isEnd() || !isNeutralOrIsolate(directions[sequence.runeI]) {
			return false
		}
	}
	for sequence.next(); !sequence.isEnd() && isPossibleNeutralOrIsolate(directions[sequence.runeI]); sequence.next() {
	}
	return true
}

// rule N1 & N2
func n12(sequence runSequence, levels []EmbeddingLevel, directions []bidi.Class) {
	previousDirection := sequence.sos
	sequence.start()
	for !sequence.isEnd() {
		if isPossibleNeutralOrIsolate(directions[sequence.runeI]) {
			// Compute sequence
			start := sequence.runeI
			if !n1GetSequence(&sequence, directions) {
				continue // TODO prevDir
			}
			// Compute nextDirection
			var nextDirection bidi.Class
			if sequence.isEnd() {
				nextDirection = sequence.eos
			} else {
				nextDirection = directions[sequence.runeI]
			}
			// Replace sequence with surrounded strong direction, or with embedding direction
			if isN0Strong(previousDirection) && isN0Strong(nextDirection) && (previousDirection == bidi.LeftToRight) == (nextDirection == bidi.LeftToRight) {
				var newDirection bidi.Class
				if previousDirection == bidi.LeftToRight {
					newDirection = bidi.LeftToRight
				} else {
					newDirection = bidi.RightToLeft
				}
				state := sequence.saveState()
				for sequence.prev(); sequence.runeI >= start && !sequence.isStart(); sequence.prev() {
					directions[sequence.runeI] = newDirection
				}
				sequence.loadState(state)
			} else {
				state := sequence.saveState()
				for sequence.prev(); sequence.runeI >= start && !sequence.isStart(); sequence.prev() {
					directions[sequence.runeI] = levels[sequence.runeI].direction()
				}
				sequence.loadState(state)
			}
		} else {
			sequence.next()
		}
		previousDirection = directions[sequence.runeI-1]
	}
}
