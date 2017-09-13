package bidih

import (
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
)

// rule W1
func w1(sequence runSequence, directions []bidi.Class) {
	previousDirection := sequence.sos
	for sequence.start(); !sequence.isEnd(); sequence.next() {
		direction := directions[sequence.runeI]
		if direction == bidi.NonSpacingMark {
			if isIsolateInitiator(previousDirection) || previousDirection == bidi.PopDirectionalIsolate {
				directions[sequence.runeI] = bidi.OtherNeutral
				previousDirection = bidi.OtherNeutral
			} else {
				directions[sequence.runeI] = previousDirection
			}
		} else if direction != bidi.BoundaryNeutral {
			previousDirection = direction
		}
	}
}

// rule W2 & rule W3
func w23(sequence runSequence, directions []bidi.Class) {
	previousStrongDirection := sequence.sos
	for sequence.start(); !sequence.isEnd(); sequence.next() {
		direction := directions[sequence.runeI]
		// W2
		if direction == bidi.EuropeanNumber && previousStrongDirection == bidi.ArabicLetter {
			directions[sequence.runeI] = bidi.ArabicNumber
		} else if isStrong(direction) {
			previousStrongDirection = direction
		}
		// W3
		if direction == bidi.ArabicLetter {
			directions[sequence.runeI] = bidi.RightToLeft
		}
	}
}

// rule W4
func w4(sequence runSequence, directions []bidi.Class) {
	previousDirection := sequence.sos // Skip BN (BoundaryNeutral)
	for sequence.start(); !sequence.isEnd(); sequence.next() {
		direction := directions[sequence.runeI]
		if (direction == bidi.EuropeanSeparator && previousDirection == bidi.EuropeanNumber) || (direction == bidi.CommonSeparator && (previousDirection == bidi.EuropeanNumber || previousDirection == bidi.ArabicNumber)) {
			state := sequence.saveState()
			for sequence.next(); !sequence.isEnd() && directions[sequence.runeI] == bidi.BoundaryNeutral; sequence.next() {
			}
			if !sequence.isEnd() && directions[sequence.runeI] == previousDirection {
				directions[state.runeI] = previousDirection
			}
			sequence.loadState(state)
			// TODO skip next???
		}
		if direction != bidi.BoundaryNeutral {
			previousDirection = direction
		}
	}
}

// rule W5
func w5(sequence runSequence, directions []bidi.Class) {
	sequence.start()
	for !sequence.isEnd() {
		//
		// skip all not ET & not BN
		//
		for !sequence.isEnd() && directions[sequence.runeI] != bidi.EuropeanTerminator && directions[sequence.runeI] != bidi.BoundaryNeutral {
			sequence.next()
		}
		if sequence.isEnd() {
			return
		}
		//
		// beginning of ET and/or BN
		//
		// check previous direction
		state := sequence.saveState() // also store beginning of ET and/or BN sequence
		sequence.prev()
		change := !sequence.isStart() && directions[sequence.runeI] == bidi.EuropeanNumber
		sequence.loadState(state)
		// go over sequence of ET and/or BN
		for ; !sequence.isEnd() && (directions[sequence.runeI] == bidi.EuropeanTerminator || directions[sequence.runeI] == bidi.BoundaryNeutral); sequence.next() {
			if change {
				directions[sequence.runeI] = bidi.EuropeanNumber
			}
		}
		// reach end of ET and/or BN sequence
		// change sequence if it is not already changed and it is required
		if !change && !sequence.isEnd() && directions[sequence.runeI] == bidi.EuropeanNumber {
			endState := sequence.saveState()
			for sequence.prev(); sequence.runeI >= state.runeI; sequence.prev() {
				directions[sequence.runeI] = bidi.EuropeanNumber
			}
			sequence.loadState(endState)
		}
	}
}

// rule W6
func w6(sequence runSequence, directions []bidi.Class) {
	for sequence.start(); !sequence.isEnd(); {
		direction := directions[sequence.runeI]
		if direction == bidi.EuropeanSeparator || direction == bidi.CommonSeparator || direction == bidi.EuropeanTerminator {
			directions[sequence.runeI] = bidi.OtherNeutral
			// Change adjacent BN (on left)
			state := sequence.saveState()
			for ; !sequence.isStart() && directions[sequence.runeI] == bidi.BoundaryNeutral; sequence.prev() {
				directions[sequence.runeI] = bidi.OtherNeutral
			}
			sequence.loadState(state)
			// Change adjacent BN (on right)
			for sequence.next(); !sequence.isEnd() && directions[sequence.runeI] == bidi.BoundaryNeutral; sequence.next() {
				directions[sequence.runeI] = bidi.OtherNeutral
			}
		} else {
			sequence.next()
		}
	}
}

// rule W7
func w7(sequence runSequence, directions []bidi.Class) {
	previousStrongDirection := sequence.sos
	for sequence.start(); !sequence.isEnd(); sequence.next() {
		direction := directions[sequence.runeI]
		if direction == bidi.EuropeanNumber && previousStrongDirection == bidi.LeftToRight {
			directions[sequence.runeI] = bidi.LeftToRight
		} else if (direction == bidi.LeftToRight || direction == bidi.RightToLeft) && directions[sequence.runeI] != bidi.BoundaryNeutral {
			previousStrongDirection = direction
		}
	}
}
