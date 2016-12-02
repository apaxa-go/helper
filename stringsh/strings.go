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

// GetLine returns first line from s and position in s of estimated part.
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
// As line delimiter does not include in result it may be hard to manipulate with estimated string.
func GetFirstLine(s string) string {
	line, _ := GetLine(s)
	return line
}

// ExtractLine returns first line from s and estimated part of s.
// Line delimiter is the same as in GetLine.
// Also as in GetLine delimiter does not include nor in line nor in est.
func ExtractLine(s string) (line, est string) {
	line, pos := GetLine(s)
	if pos < len(s) {
		est = s[pos:]
	}
	return
}
