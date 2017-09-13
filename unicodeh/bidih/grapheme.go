package bidih

import (
	"github.com/apaxa-go/helper/unicodeh/grapheme"
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
)

func initClassesGraphemes(runes []rune, graphemes []grapheme.ClusterBoundary)(classes []bidi.Class){
	classes = make([]bidi.Class, len(graphemes))
	for graphemeI, g := range graphemes {
		classes[graphemeI] = bidi.GetClass(runes[g.From])
	}
	return classes
}

func ParseParagraphGraphemes(runes []rune, paragraphLevel EmbeddingLevel, lineBreaker LineBreaker){
	graphemes:=grapheme.BoundariesRunes(runes)
	classes:=initClassesGraphemes(runes,graphemes)
	parseParagraph(runes, classes, paragraphLevel, lineBreaker,nil)
}
