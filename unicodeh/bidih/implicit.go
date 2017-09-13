package bidih

import (
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
)

// BD8
// = LRI, RLI, or FSI
func isIsolateInitiator(class bidi.Class) bool {
	return class == bidi.LeftToRightIsolate || class == bidi.RightToLeftIsolate || class == bidi.FirstStrongIsolate
}

// = L, R, or AL
func isStrong(c bidi.Class) bool {
	return c == bidi.LeftToRight || c == bidi.RightToLeft || c == bidi.ArabicLetter
}

func isRunAlreadyInSequence(runeI int, pairs map[int]pairElement) bool {
	prev, ok := pairs[runeI]
	return ok && prev.pair < runeI
}

func nextRunInSequence(runeI int, directions []bidi.Class, pairs map[int]pairElement) int {
	next, ok := pairs[runeI]
	// If pair is following, not preceding
	if !ok || next.pair <= runeI {
		return -1
	}
	// If pair is isolate - no other checking required
	if next.isolate {
		return next.pair
	}
	// Validate non isolate pair
	for runeI++; runeI < next.pair; runeI++ {
		if directions[runeI] != bidi.BoundaryNeutral {
			delete(pairs, next.pair)
			delete(pairs, runeI)
			return -1
		}
	}
	return next.pair
}

// BD7
func computeRunEnd(levels []EmbeddingLevel, from int) int {
	level := levels[from]
	for from++; from < len(levels) && levels[from] == level; from++ {
	}
	return from
}

// BD13 & BD7
func computeRunSequencesRanges(levels []EmbeddingLevel, directions []bidi.Class, pairs map[int]pairElement) (sequences []runSequence) {
	l := len(levels)
	for runeI := 0; runeI < l; {
		if isRunAlreadyInSequence(runeI, pairs) {
			// Skip run
			runeI = computeRunEnd(levels, runeI)
		} else {
			// Get first run in sequence
			from := runeI
			runeI = computeRunEnd(levels, runeI)
			to := runeI
			sequence := makeRunSequence(from, to)
			// Get all other runs in sequence
			for from := nextRunInSequence(to-1, directions, pairs); from >= 0; from = nextRunInSequence(to-1, directions, pairs) {
				to = computeRunEnd(levels, from)
				sequence.ranges = append(sequence.ranges, struct{ from, to int }{from, to})
			}
			// Save sequence
			sequences = append(sequences, sequence)
		}
	}
	return
}

func computeRunSequencesBorderTypes(sequences []runSequence, classes []bidi.Class, levels []EmbeddingLevel, directions []bidi.Class, paragraphLevel EmbeddingLevel) {
	for sequenceI := range sequences {
		// SOS
		sequences[sequenceI].start()
		runeI := sequences[sequenceI].runeI
		bos := levels[runeI]
		for runeI--; runeI >= 0 && directions[runeI] == bidi.BoundaryNeutral; runeI-- {
		}

		if runeI < 0 {
			if paragraphLevel > bos {
				bos = paragraphLevel
			}
		} else {
			if runeEmbeddingLevel := levels[runeI]; runeEmbeddingLevel > bos {
				bos = runeEmbeddingLevel
			}
		}
		sequences[sequenceI].sos = bos.direction()
		// EOS
		sequences[sequenceI].end()
		runeI = sequences[sequenceI].runeI
		bos = levels[runeI]
		if isIsolateInitiator(classes[runeI]) { // TODO isIsolateInitiator - now not only isolate??????
			if paragraphLevel > bos {
				bos = paragraphLevel
			}
		} else {
			for runeI++; runeI < len(levels) && directions[runeI] == bidi.BoundaryNeutral; runeI++ {
			}
			if runeI >= len(levels) {
				if paragraphLevel > bos {
					bos = paragraphLevel
				}
			} else {
				if runeEmbeddingLevel := levels[runeI]; runeEmbeddingLevel > bos {
					bos = runeEmbeddingLevel
				}
			}
		}
		sequences[sequenceI].eos = bos.direction()
	}
}

// By tr9 it should be isolate run sequence, but it is isolate and/or overriding and/or embedding.
func computeRunSequences(classes []bidi.Class, paragraphLevel EmbeddingLevel, levels []EmbeddingLevel, directions []bidi.Class, pairs map[int]pairElement) (sequences []runSequence) {
	sequences = computeRunSequencesRanges(levels, directions, pairs)
	computeRunSequencesBorderTypes(sequences, classes, levels, directions, paragraphLevel)
	return
}

// Rules I1 & I2
func i12(levels []EmbeddingLevel, directions []bidi.Class) {
	for runeI := range levels {
		switch levels[runeI] % 2 {
		case 0:
			switch directions[runeI] {
			case bidi.RightToLeft:
				levels[runeI]++
			case bidi.ArabicNumber, bidi.EuropeanNumber:
				levels[runeI] += 2
			}
		case 1:
			switch directions[runeI] {
			case bidi.LeftToRight, bidi.ArabicNumber, bidi.EuropeanNumber:
				levels[runeI]++
			}
		}
	}
}

// TODO store embeddingLevel/embeddingDirection in sequence???
func implicit(runes []rune, classes []bidi.Class, paragraphLevel EmbeddingLevel, levels []EmbeddingLevel, directions []bidi.Class, pairs map[int]pairElement) {
	sequences := computeRunSequences(classes, paragraphLevel, levels, directions, pairs)
	for _, sequence := range sequences {
		w1(sequence, directions)
		w23(sequence, directions)
		w4(sequence, directions)
		w5(sequence, directions)
		w6(sequence, directions)
		w7(sequence, directions)
		n0(sequence, runes, classes, levels, directions)
		n12(sequence, levels, directions)
	}
	i12(levels, directions)
}
