// Package stringsh is helper for system package strings.
package stringsh

import "golang.org/x/text/unicode/norm"

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
