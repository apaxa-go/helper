// Package boundaryh implements boundary Unicode algorithms:
// - Unicode Line Breaking Algorithm (TR14) - finds possible line breaking position,
// - Unicode Text Segmentation (TR29) - finds grapheme clusters, words and sentences boundary.
// For more information see http://unicode.org/reports/tr14 and http://unicode.org/reports/tr29 .
// Passing empty slice (or nil) or position outside of passed slice causes "Invalid", "InvalidPos" or empty result (depending on result type).
// Line breaking related functions also can return "NoLineBreak" result if there is no possible breaking. This is because line breaking algorithm does not allow breaking before 0-th rune (while other algorithms allow).
// In fact "NoLineBreak" is just named "0", but usually it must be treated specially.
package boundaryh

// Naming convention for private variable:
//	l - length
//	l0 - class of first rune on left side of possible break
//  l1 - class of second rune ...
//	...
//	r0 - class of first rune on right side ...
//	...
//	l<smth> - describe some state/property on left side of possible break
//	r<smth> - describe some state/property on right side ...
// Naming convention for private function:
//	g... - function related to grapheme clusters
//	s... - function related to sentences
//	w... - function related to words
//	l... - function related to lines (in terms of line breaking)
