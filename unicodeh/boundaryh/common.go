package boundaryh

import "unicode/utf8"

//replacer:ignore
// TODO replace windows path separator
//go:generate go run $GOPATH\src\github.com\apaxa-go\generator\replacer\main.go -- $GOFILE

const (
	crRune = '\u000D'
	lfRune = '\u000A'
)

// InvalidPos indicates invalid result position.
// This means that arguments to function is invalid (i.e. looking for word in empty string or looking for grapheme cluster at outside of string).
const InvalidPos = -1

// Boundary represents subslice by its indexes.
// Subslice can be retrieved by "s[b.From:b.To]".
type Boundary struct {
	From int // index of first element
	To   int // (index of last element) + 1
}

// Invalid returns "Invalid" boundary.
func Invalid() Boundary { return Boundary{InvalidPos, InvalidPos} }

// Len returns length of subslice represented by boundary.
// Len of "Invalid" boundary is 0. Valid boundary also may have zero length.
func (b Boundary) Len() int { return b.To - b.From }

// IsInvalid checks if boundary is invalid in terms of "Invalid" constructor (and does not check anything else).
// Warning: !b.IsValid() is not equal to b.IsInvalid().
func (b Boundary) IsInvalid() bool { return b == Invalid() }

// IsValid checks if boundary is valid in general.
// Warning: !b.IsValid() is not equal to b.IsInvalid().
func (b Boundary) IsValid() bool { return b.From >= 0 && b.To >= b.From }

//replacer:replace
//replacer:old "func g"	gClass gGetClass
//replacer:new "func l"	lClass lGetClass
//replacer:new "func s"	sClass sGetClass
//replacer:new "func w" wClass wGetClass

//
// Group of functions to retrieve gClass of the first or last rune in runes, string or bytes.
//
func gFirstClassInRunes(runes []rune) (c gClass, pos int) {
	return gGetClass(runes[0]), 1
}

func gLastClassInRunes(runes []rune) (c gClass, pos int) {
	l := len(runes) - 1
	return gGetClass(runes[l]), l
}

func gFirstClassInString(s string) (c gClass, pos int) {
	r, pos := utf8.DecodeRuneInString(s)
	return gGetClass(r), pos
}

func gLastClassInString(s string) (c gClass, pos int) {
	r, pos := utf8.DecodeLastRuneInString(s)
	return gGetClass(r), len(s) - pos
}

// ...InBytes
func gFirstClass(b []byte) (c gClass, pos int) {
	r, pos := utf8.DecodeRune(b)
	return gGetClass(r), pos
}

// ...InBytes
func gLastClass(b []byte) (c gClass, pos int) {
	r, pos := utf8.DecodeLastRune(b)
	return gGetClass(r), len(b) - pos
}

//replacer:ignore

//
// Group of functions for going to the beginning of current or next rune.
// This is important if passed by user index in string or []byte (not for []rune) points to the middle or end of multibyte rune.
//
func toRuneBeginInRunes(_ []rune, pos int) int { return pos }
func toRuneBeginInString(s string, pos int) int {
	for pos > 0 && !utf8.RuneStart(s[pos]) {
		pos--
	}
	return pos
}

// ...InBytes
func toRuneBegin(b []byte, pos int) int {
	for pos > 0 && !utf8.RuneStart(b[pos]) {
		pos--
	}
	return pos
}

func toNextRuneInRunes(_ []rune, pos int) int { return pos + 1 }
func toNextRuneInString(s string, pos int) int {
	pos++
	for pos < len(s) && !utf8.RuneStart(s[pos]) {
		pos++
	}
	return pos
}

// ...InBytes
func toNextRune(b []byte, pos int) int {
	pos++
	for pos < len(b) && !utf8.RuneStart(b[pos]) {
		pos++
	}
	return pos
}
