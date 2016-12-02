// Package stringsh is helper for system package strings.
package stringsh

import (
	"golang.org/x/text/unicode/norm"
	"strings"
)

// Len returns number of glyph in UTF-8 encoded string.
func Len(s string) (l int) {
	var ia norm.Iter
	ia.InitString(norm.NFKD, s)
	for !ia.Done() {
		l++
		ia.Next()
	}
	return
}

// GetLine returns first line from s and position in s of remaining part.
// Line delimiter may be "\n" or "\r\n". In both cases delimiter does not include nor in line nor in pos (pos points to first byte of second line).
// If s does not contain delimiter than GetLine returns s as first line.
// If there is no second line in s (s does not contain delimiter or there is no other bytes after delimiter) than pos will be point to non existing position in s.
func GetLine(s string) (line string, pos int) {
	i := strings.Index(s, "\n")
	if i == -1 {
		return s, len(s)
	}

	if i > 0 && s[i-1] == '\r' {
		return s[:i-1], i + 1
	}

	return s[:i], i + 1
}

// GetFirstLine is a shortcut for GetLine but returning only first line.
// As line delimiter does not include in result it may be hard to manipulate with remaining string.
func GetFirstLine(s string) string {
	line, _ := GetLine(s)
	return line
}

// ExtractLine returns first line from s and remaining part of s.
// Line delimiter is the same as in GetLine.
// Also as in GetLine delimiter does not include nor in line nor in rem.
func ExtractLine(s string) (line, rem string) {
	line, pos := GetLine(s)
	if pos < len(s) {
		rem = s[pos:]
	}
	return
}

// IndexMulti returns the index of the first instance of any seps in s and index of founded sep, or (-1,-1) if seps are not present in s.
// Checking for seps in each position of s proceeds in order of they are passed, so if seps[5] and seps[7] both present in s at the same position (let it be 29) then result will be (29; 5), not (29; 7).
// Empty string (as seps) presents at position 0 in any string.
func IndexMulti(s string, seps ...string) (i int, sep int) {
	if len(s) == 0 { // catch case with empty string and empty sep
		for j := range seps {
			if seps[j] == "" {
				return 0, j
			}
		}
		return -1, -1
	}

	for i = range s {
		for j := range seps {
			if strings.HasPrefix(s[i:], seps[j]) {
				return i, j
			}
		}
	}

	return -1, -1
}

// ReplaceMulti returns a copy of the string s with non-overlapping instances of old elements replaced by corresponding new elements.
// If len(old) != len(new) => panic
// if len(old[i])==0 => panic
// ReplaceMulti returns s as-is if old is empty.
func ReplaceMulti(s string, old, new []string) (r string) {
	if len(old) != len(new) {
		panic("number of old elemnts and new elemnts should be the same")
	}

	for i := range old {
		if len(old[i]) == 0 {
			panic("no one of old elements can be empty string")
		}
	}

	if len(old) == 0 {
		return s
	}

	for i, j := IndexMulti(s, old...); i != -1; i, j = IndexMulti(s, old...) {
		r += s[:i] + new[j]
		s = s[i+len(old[j]):]
	}
	r += s
	return
}
