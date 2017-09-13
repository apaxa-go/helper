package testdata

import "github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"

type BidiTest struct {
	Classes          []bidi.Class
	ParagraphDirections []uint8// EmbeddingLevel
}

func (t *BidiTest) Runes() []rune {
	runes := make([]rune, len(t.Classes))
	for runeI, rawClass := range t.Classes {
		runes[runeI] = runeSamples[rawClass]
	}
	return runes
}

type BidiTestGroup struct {
	// Output
	Levels []uint8 // EmbeddingLevel; maxUint8 => skip
	Orders []int
	// Input
	Tests []BidiTest
}

func (t *BidiTestGroup) Result(tRunes []rune) []rune {
	runes := make([]rune, len(t.Orders))
	for runeI, o := range t.Orders {
		runes[runeI] = tRunes[o]
	}
	return runes
}

// Test data.
// Generated.
var BidiTests []BidiTestGroup