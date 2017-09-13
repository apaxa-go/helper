package testdata

type BidiCharacterTest struct {
	// Input
	Runes              []rune
	ParagraphDirection uint8// EmbeddingLevel
	// Output
	ParagraphLevel uint8   // EmbeddingLevel
	Levels         []uint8 // []EmbeddingLevel; maxUint8 => skip
	Orders         []int
}

func (t *BidiCharacterTest) Result() []rune {
	runes := make([]rune, len(t.Orders))
	for runeI, o := range t.Orders {
		runes[runeI] = t.Runes[o]
	}
	return runes
}

// Test data.
// Generated
var BidiCharacterTests []BidiCharacterTest