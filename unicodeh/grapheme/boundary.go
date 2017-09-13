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
func FirstBoundaryInRunes(str []rune) int {
	l := len(str)

	// Check for empty slice
	if l == 0 {
		return 0
	}

	const fallbackLen = 1

	// Short path for CRLF
	if l >= 2 && str[0] == crRune && str[1] == lfRune {
		return 2
	}

	pos := 0

	//	Prepend
	for ; pos < l && unicode.Is(prependTable, str[pos]); pos++ {
	}

	// Base (RI-sequence | Hangul-Syllable | !Control)
	if pos >= l {
		return fallbackLen
	}

	{
		r := str[pos]
		if unicodeh.IsRegionalIndicatorYes(r) { // RI-sequence
			pos++
			for ; pos < l && unicodeh.IsRegionalIndicatorYes(str[pos]); pos++ {
			}
		} else if !unicodeh.IsHangulSyllableTypeNotApplicable(r) { // Hangul-Syllable
			pos+=solveHangulSyllable(str[pos:])
		} else if !unicode.Is(controlTable, r) { // !Control
			pos++
		} else {
			return fallbackLen
		}
	}

	// Suffix (Grapheme_Extend | SpacingMark)
	for ; pos < l && unicode.Is(suffixTable, str[pos]); pos++ {
	}

	return pos
}
