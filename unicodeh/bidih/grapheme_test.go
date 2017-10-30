package bidih

/*

import (
	"testing"
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/testdata"
	"reflect"
)

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
*/
