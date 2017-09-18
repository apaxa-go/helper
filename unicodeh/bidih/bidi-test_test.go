package bidih

import (
	"github.com/apaxa-go/helper/mathh"
	"reflect"
	"testing"
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/testdata"
)

func compareEL(expectedLevels []uint8, levels []EmbeddingLevel) bool {
	if len(expectedLevels) != len(levels) {
		return false
	}
	for runeI := range expectedLevels {
		if expectedLevels[runeI] != mathh.MaxUint8 && EmbeddingLevel(expectedLevels[runeI]) != levels[runeI] {
			return false
		}
	}
	return true
}

func TestBidiTest(t *testing.T) {
	var failed int
	var total int
	displayed := false
	for groupI, group := range testdata.BidiTests {
		for testI, test := range group.Tests {
			for _, paragraphDirection := range test.ParagraphDirections {
				runes := test.Runes()
				//runes, levels := ParseParagraph(runes, EmbeddingLevel(paragraphDirection), nil)
				ParseParagraph(runes, EmbeddingLevel(paragraphDirection), nil)
				runes = removeX9(runes)
				//if expected := group.Result(test.Runes()); !reflect.DeepEqual(runes, expected) || !compareEL(group.Levels, levels) {
				if expected := group.Result(test.Runes()); !reflect.DeepEqual(runes, expected) {
					if !displayed {
						//t.Errorf("%v.%v.%v:\nclasses %v\ninput  %v\nexpect %v\ngot    %v\nexp EL %v\ngot EL %v\norder  %v\n", groupI, testI, paragraphDirection, test.Classes, test.Runes(), expected, runes, group.Levels, levels, group.Orders)
						t.Errorf("%v.%v.%v:\nclasses %v\ninput  %v\nexpect %v\ngot    %v\nexp EL %v\norder  %v\n", groupI, testI, paragraphDirection, test.Classes, test.Runes(), expected, runes, group.Levels, group.Orders)
						displayed = true
					}
					failed++
				}
				total++
			}
		}
	}
	t.Log("Failed", failed, "of", total)
}

func TestBidiTestGraphemes(t *testing.T) {
	var failed int
	var total int
	displayed := false
	for groupI, group := range testdata.BidiTests {
		for testI, test := range group.Tests {
			for _, paragraphDirection := range test.ParagraphDirections {
				runes := test.Runes()
				//runes, levels := ParseParagraph(runes, EmbeddingLevel(paragraphDirection), nil)
				runes = ParseParagraphGraphemes(runes, EmbeddingLevel(paragraphDirection), nil)
				runes = removeX9(runes)
				//if expected := group.Result(test.Runes()); !reflect.DeepEqual(runes, expected) || !compareEL(group.Levels, levels) {
				if expected := group.Result(test.Runes()); !reflect.DeepEqual(runes, expected) {
					if !displayed {
						//t.Errorf("%v.%v.%v:\nclasses %v\ninput  %v\nexpect %v\ngot    %v\nexp EL %v\ngot EL %v\norder  %v\n", groupI, testI, paragraphDirection, test.Classes, test.Runes(), expected, runes, group.Levels, levels, group.Orders)
						t.Errorf("%v.%v.%v:\nclasses %v\ninput  %v\nexpect %v\ngot    %v\nexp EL %v\norder  %v\n", groupI, testI, paragraphDirection, test.Classes, test.Runes(), expected, runes, group.Levels, group.Orders)
						displayed = true
					}
					failed++
				}
				total++
			}
		}
	}
	t.Log("Failed", failed, "of", total)
}

// For debug purpose: test single case.
func TestSomeOne(t *testing.T){
	groupI:=0
	testI:=0
	paragraphDirection:= EmbeddingLevel(0)

	group:=testdata.BidiTests[groupI]
	test:=group.Tests[testI]

	runes := test.Runes()
	//runes, levels := ParseParagraph(runes, paragraphDirection, nil)
	ParseParagraph(runes, paragraphDirection, nil)
	runes = removeX9(runes)
	//if expected := group.Result(test.Runes()); !reflect.DeepEqual(runes, expected) || !compareEL(group.Levels, levels) {
	if expected := group.Result(test.Runes()); !reflect.DeepEqual(runes, expected) {
		//t.Errorf("%v.%v.%v:\nclasses %v\ninput  %v\nexpect %v\ngot    %v\nexp EL %v\ngot EL %v\norder  %v\n", groupI,testI,paragraphDirection,test.Classes, test.Runes(), expected, runes, group.Levels, levels, group.Orders)
		t.Errorf("%v.%v.%v:\nclasses %v\ninput  %v\nexpect %v\ngot    %v\nexp EL %v\norder  %v\n", groupI,testI,paragraphDirection,test.Classes, test.Runes(), expected, runes, group.Levels, group.Orders)
	}
}

// For debug purpose: test single case.
func TestSomeOneGraphemes(t *testing.T){
	groupI:=186
	testI:=11
	paragraphDirection:= EmbeddingLevel(0)

	group:=testdata.BidiTests[groupI]
	test:=group.Tests[testI]

	runes := test.Runes()
	//runes, levels := ParseParagraph(runes, paragraphDirection, nil)
	runes = ParseParagraphGraphemes(runes, paragraphDirection, nil)
	runes = removeX9(runes)
	//if expected := group.Result(test.Runes()); !reflect.DeepEqual(runes, expected) || !compareEL(group.Levels, levels) {
	if expected := group.Result(test.Runes()); !reflect.DeepEqual(runes, expected) {
		//t.Errorf("%v.%v.%v:\nclasses %v\ninput  %v\nexpect %v\ngot    %v\nexp EL %v\ngot EL %v\norder  %v\n", groupI,testI,paragraphDirection,test.Classes, test.Runes(), expected, runes, group.Levels, levels, group.Orders)
		t.Errorf("%v.%v.%v:\nclasses %v\ninput  %v\nexpect %v\ngot    %v\nexp EL %v\norder  %v\n", groupI,testI,paragraphDirection,test.Classes, test.Runes(), expected, runes, group.Levels, group.Orders)
	}
}
