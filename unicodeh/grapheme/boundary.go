package grapheme

import (
	"github.com/apaxa-go/helper/unicodeh"
	"unicode"
)

//go:generate go run ./internal/maketables/maketables.go

const (
	crRune rune = '\u000D'
	lfRune rune = '\u000A'
)

// 0 - if r is empty
func FirstBoundaryRunes(ruens []rune) int {
	l := len(ruens)

	// Check for empty slice
	if l == 0 {
		return 0
	}

	const fallbackLen = 1

	// Short path for CRLF
	if l >= 2 && ruens[0] == crRune && ruens[1] == lfRune {
		return 2
	}

	pos := 0

	//	Prepend
	for ; pos < l && unicode.Is(prependTable, ruens[pos]); pos++ {
	}

	// Base (RI-sequence | Hangul-Syllable | !Control)
	if pos >= l {
		return fallbackLen
	}

	{
		r := ruens[pos]
		if unicodeh.IsRegionalIndicatorYes(r) { // RI-sequence
			pos++
			for ; pos < l && unicodeh.IsRegionalIndicatorYes(ruens[pos]); pos++ {
			}
		} else if !unicodeh.IsHangulSyllableTypeNotApplicable(r) { // Hangul-Syllable
			pos+=solveHangulSyllable(ruens[pos:])
		} else if !unicode.Is(controlTable, r) { // !Control
			pos++
		} else {
			return fallbackLen
		}
	}

	// Suffix (Grapheme_Extend | SpacingMark)
	for ; pos < l && unicode.Is(suffixTable, ruens[pos]); pos++ {
	}

	return pos
}

func BoundariesRunes(runes []rune)(boundaries []ClusterBoundary){
	boundaries=make([]ClusterBoundary,0,len(runes)) // TODO memory efficient
	for i:=0; i<len(runes); i++{
		length:=FirstBoundaryRunes(runes[i:])
		boundaries=append(boundaries,ClusterBoundary{i,i+length})
		i+=length
	}
	return
}