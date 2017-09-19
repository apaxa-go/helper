package boundaryh

/*

import (
	"github.com/apaxa-go/helper/unicodeh"
)

// This functions calculates extended grapheme cluster length as described Table 1b, 1c & 2 (regexp notation).
// Regexp notation is incomplete and does not pass all test.
// There is another function based on rules GB1-GB999 from 3.1.1 which passed all tests.

func FirstBoundaryRunes(runes []rune) int {
	l := len(runes)

	// Check for empty slice
	if l == 0 {
		return 0
	}

	const fallbackLen = 1

	// Short path for CRLF
	if l >= 2 && runes[0] == crRune && runes[1] == lfRune {
		return 2
	}

	pos := 0

	//	Prepend
	for ; pos < l && unicodeh.IsGraphemeClusterBreakPrepend(runes[pos]); pos++ {
	}

	// Base (RI-sequence | Hangul-Syllable | !Control)
	if pos >= l {
		return fallbackLen
	}

	{
		r := runes[pos]
		if unicodeh.IsGraphemeClusterBreakRegionalIndicator(r) { // RI-sequence
			pos++
			for ; pos < l && unicodeh.IsGraphemeClusterBreakRegionalIndicator(runes[pos]); pos++ {
			}
		} else if !unicodeh.IsHangulSyllableTypeNotApplicable(r) { // Hangul-Syllable
			pos += solveHangulSyllable(runes[pos:])
		} else if !unicodeh.IsGraphemeClusterBreakControl(r) && r!=crRune && r!=lfRune { // !Control
			pos++
		} else {
			return fallbackLen
		}
	}

	// Suffix (Grapheme_Extend | SpacingMark)
	for ; pos < l && (unicodeh.IsGraphemeClusterBreakExtend(runes[pos])||unicodeh.IsGraphemeClusterBreakSpacingMark(runes[pos])); pos++ {
	}

	return pos
}

func solveHangulSyllable(str []rune)(pos int){
	// L* V+ T*
	// | L* LV V* T*
	// | L* LVT T*
	// | L+
	// | T+

	l:=len(str)
	if unicodeh.IsHangulSyllableTypeTrailingJamo(str[0]) { // "T+"
		pos++
		for ; pos<l && unicodeh.IsHangulSyllableTypeTrailingJamo(str[pos]); pos++{}
		return
	}

	// just skip hsL s and update firstType
	if unicodeh.IsHangulSyllableTypeLeadingJamo(str[0]) {
		pos++
		for ; pos<l && unicodeh.IsHangulSyllableTypeLeadingJamo(str[pos]); pos++{}
		if pos>=l{ return }
	}

	switch {
	case unicodeh.IsHangulSyllableTypeVowelJamo(str[pos]), unicodeh.IsHangulSyllableTypeLVSyllable(str[pos]): // "L* V+ T*" (without "L*") or "L* LV V* T*" (without "L*")
		pos++
		for ; pos<l && unicodeh.IsHangulSyllableTypeVowelJamo(str[pos]); pos++{}
		for ; pos<l && unicodeh.IsHangulSyllableTypeTrailingJamo(str[pos]); pos++{}
	case unicodeh.IsHangulSyllableTypeLVTSyllable(str[pos]): // "L* LVT T*" (without "L*")
		pos++
		for ; pos<l && unicodeh.IsHangulSyllableTypeTrailingJamo(str[pos]); pos++{}
	}
	return
}
*/
