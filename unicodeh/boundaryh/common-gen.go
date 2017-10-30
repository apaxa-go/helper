//replacer:generated-file

package boundaryh

import "unicode/utf8"

//
// Group of functions to retrieve lClass of the first or last rune in runes, string or bytes.
//
func lFirstClassInRunes(runes []rune) (c lClass, pos int) {
	return lGetClass(runes[0]), 1
}

func lLastClassInRunes(runes []rune) (c lClass, pos int) {
	l := len(runes) - 1
	return lGetClass(runes[l]), l
}

func lFirstClassInString(s string) (c lClass, pos int) {
	r, pos := utf8.DecodeRuneInString(s)
	return lGetClass(r), pos
}

func lLastClassInString(s string) (c lClass, pos int) {
	r, pos := utf8.DecodeLastRuneInString(s)
	return lGetClass(r), len(s) - pos
}

// ...InBytes
func lFirstClass(b []byte) (c lClass, pos int) {
	r, pos := utf8.DecodeRune(b)
	return lGetClass(r), pos
}

// ...InBytes
func lLastClass(b []byte) (c lClass, pos int) {
	r, pos := utf8.DecodeLastRune(b)
	return lGetClass(r), len(b) - pos
}

//
// Group of functions to retrieve sClass of the first or last rune in runes, string or bytes.
//
func sFirstClassInRunes(runes []rune) (c sClass, pos int) {
	return sGetClass(runes[0]), 1
}

func sLastClassInRunes(runes []rune) (c sClass, pos int) {
	l := len(runes) - 1
	return sGetClass(runes[l]), l
}

func sFirstClassInString(s string) (c sClass, pos int) {
	r, pos := utf8.DecodeRuneInString(s)
	return sGetClass(r), pos
}

func sLastClassInString(s string) (c sClass, pos int) {
	r, pos := utf8.DecodeLastRuneInString(s)
	return sGetClass(r), len(s) - pos
}

// ...InBytes
func sFirstClass(b []byte) (c sClass, pos int) {
	r, pos := utf8.DecodeRune(b)
	return sGetClass(r), pos
}

// ...InBytes
func sLastClass(b []byte) (c sClass, pos int) {
	r, pos := utf8.DecodeLastRune(b)
	return sGetClass(r), len(b) - pos
}

//
// Group of functions to retrieve wClass of the first or last rune in runes, string or bytes.
//
func wFirstClassInRunes(runes []rune) (c wClass, pos int) {
	return wGetClass(runes[0]), 1
}

func wLastClassInRunes(runes []rune) (c wClass, pos int) {
	l := len(runes) - 1
	return wGetClass(runes[l]), l
}

func wFirstClassInString(s string) (c wClass, pos int) {
	r, pos := utf8.DecodeRuneInString(s)
	return wGetClass(r), pos
}

func wLastClassInString(s string) (c wClass, pos int) {
	r, pos := utf8.DecodeLastRuneInString(s)
	return wGetClass(r), len(s) - pos
}

// ...InBytes
func wFirstClass(b []byte) (c wClass, pos int) {
	r, pos := utf8.DecodeRune(b)
	return wGetClass(r), pos
}

// ...InBytes
func wLastClass(b []byte) (c wClass, pos int) {
	r, pos := utf8.DecodeLastRune(b)
	return wGetClass(r), len(b) - pos
}
