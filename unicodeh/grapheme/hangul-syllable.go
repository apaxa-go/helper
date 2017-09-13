package grapheme

import (
	"github.com/apaxa-go/helper/unicodeh"
)

func solveHangulSyllable(str []rune)(pos int){
	/*
	L* V+ T*
	| L* LV V* T*
	| L* LVT T*
	| L+
	| T+
	 */
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