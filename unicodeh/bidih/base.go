package bidih

import (
	"github.com/apaxa-go/helper/mathh"
	"github.com/apaxa-go/helper/unicodeh/bidih/internal/bidi"
)

type EmbeddingLevel uint8

func (level EmbeddingLevel) direction() bidi.Class {
	if level%2 == 0 {
		return bidi.LeftToRight
	}
	return bidi.RightToLeft
}

const (
	LeftToRightParagraph   EmbeddingLevel = 0
	RightToLeftParagraph   EmbeddingLevel = 1
	MaxEmbeddingLevel      EmbeddingLevel = 125
	AutoParagraphDirection EmbeddingLevel = mathh.MaxUint8
	//internalUnset          EmbeddingLevel = mathh.MaxUint8
)

type Text struct {
	// User defined
	str                     string
	paragraphEmbeddingLevel EmbeddingLevel
	// Caches
	runes      []rune       // Next runes
	rawClasses []bidi.Class // Classes (as-is) of next runes
	// Result
	res []rune
}

func Parse(str string, paragraphEmbeddingLevel EmbeddingLevel) []rune {
	t := Text{
		// User defined
		str: str,
		paragraphEmbeddingLevel: paragraphEmbeddingLevel,
		// Caches
		// Result
		res: make([]rune, 0, len(str)), // TODO may be be more carefully about memory???
	}
	return t.res
}

/*
func (t *Text) parse() {
	nextStartNewParagraph := true

	// Paragraph dependent data
	var paragraphEL EmbeddingLevel
	for runeI, r := range t.str {
		var rawClass bidi.Class
		if len(t.runes) > 0 {
			t.runes = t.runes[1:]
			rawClass = t.rawClasses[0]
			t.rawClasses = t.rawClasses[1:]
		} else {
			rawClass = unicodeh.GetBidiClass(r)
		}

		if nextStartNewParagraph {
			if t.paragraphEmbeddingLevel == AutoParagraphDirection {
				// Paragraph embedding level auto detection
				paragraphEL = t.determineParagraphLevel(rawClass, runeI+utf8.RuneLen(r))
			} else {
				// Paragraph embedding level defined by user
				paragraphEL = t.paragraphEmbeddingLevel
			}
		}

		// Rule P1
		nextStartNewParagraph = rawClass == bidi.ParagraphSeparator

	}
}

// returns AutoParagraphDirection if passed class is not enough
func decideParagraphLevel(rawClass bidi.Class) EmbeddingLevel {
	switch rawClass {
	case bidi.ArabicLetter, bidi.RightToLeft:
		return RightToLeftParagraph
	case bidi.LeftToRight, bidi.ParagraphSeparator:
		return LeftToRightParagraph
	default:
		return AutoParagraphDirection
	}
}

// Rule P2 & P3
func (t *Text) determineParagraphLevel(rawClass bidi.Class, nextRunePos int) EmbeddingLevel {
	// By current rune
	if res := decideParagraphLevel(rawClass); res != AutoParagraphDirection {
		return res
	}
	// By cached runes
	for i, rawClass := range t.rawClasses {
		if res := decideParagraphLevel(rawClass); res != AutoParagraphDirection {
			return res
		}
		nextRunePos += utf8.RuneLen(t.runes[i])
	}
	// By estimated string
	for _, r := range t.str[nextRunePos:] {
		rawClass = unicodeh.GetBidiClass(r)
		// Cache
		t.runes = append(t.runes, r)
		t.rawClasses = append(t.rawClasses, rawClass)
		//
		if res := decideParagraphLevel(rawClass); res != AutoParagraphDirection {
			return res
		}
	}
	return LeftToRightParagraph
}
*/
