package bidih

import (
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/testdata"
	"reflect"
	"testing"
)

//go:generate go run ./internal/testdatagen/main.go c:\\ucd-10.0.0

func removeX9(runes []rune) []rune {
	for runeI := 0; runeI < len(runes); runeI++ {
		if isX9(bidi.GetClass(runes[runeI])) {
			copy(runes[runeI:], runes[runeI+1:])
			runes = runes[:len(runes)-1]
			runeI--
		}
	}
	return runes
}

func TestBidiCharacterTest(t *testing.T) {
	var failed []int
	displayed := false
	for testI, test := range testdata.BidiCharacterTests {
		runes := make([]rune, len(test.Runes))
		copy(runes, test.Runes)
		//runes, levels := ParseParagraph(runes, EmbeddingLevel(test.ParagraphDirection), nil)
		ParseParagraph(runes, EmbeddingLevel(test.ParagraphDirection), nil)
		runes = removeX9(runes)
		//if expected := test.Result(); !reflect.DeepEqual(runes, expected) || !compareEL(test.Levels, levels) {
		if expected := test.Result(); !reflect.DeepEqual(runes, expected) {
			if !displayed {
				//t.Errorf("%v:\ninput  %v\nP direction %v\nexpect %v\ngot    %v\nexp EL %v\ngot EL %v\norder  %v\n", testI, test.Runes, test.ParagraphDirection, expected, runes, test.Levels, levels, test.Orders)
				t.Errorf("%v:\ninput  %v\nP direction %v\nexpect %v\ngot    %v\nexp EL %v\norder  %v\n", testI, test.Runes, test.ParagraphDirection, expected, runes, test.Levels, test.Orders)
				displayed = true
			}
			failed = append(failed, testI)
		}
	}
	t.Log("Failed", len(failed), "of", len(testdata.BidiCharacterTests), "\n", failed)
}
