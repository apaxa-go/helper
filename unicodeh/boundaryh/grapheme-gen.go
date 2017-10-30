//replacer:generated-file

package boundaryh

// Returns class of last rune in s which is not equal to l0.
func gLastNotEqualToInString(s string, l0 gClass) gClass {
	for len(s) > 0 {
		c, pos := gLastClassInString(s)
		if c != l0 {
			return c
		}
		s = s[:pos]
	}

	return gClassOther
}

// True if l0 is RI and it opens RI sequence in string <s..., l0, ...> (may be joined with next RI).
func gIsOpenRIInString(s string, l0 gClass) (res bool) {
	if l0 != gClassRI {
		return
	}
	res = true

	for len(s) > 0 {
		c, pos := gLastClassInString(s)
		if c != gClassRI {
			return
		}
		res = !res
		s = s[:pos]
	}

	return
}

// s must be valid (len>1).
// l0Pos must be valid (in s).
func graphemeClusterEndInString(s string, l0Pos int) int {
	l := len(s)

	if l0Pos+1 < l && s[l0Pos] == crRune && s[l0Pos+1] == lfRune { // GB3
		return l0Pos + 2
	}

	l0Pos = toRuneBeginInString(s, l0Pos) // TODO do it only on external call
	l0, r0Delta := gFirstClassInString(s[l0Pos:])
	l1Diff := gLastNotEqualToInString(s[:l0Pos], l0)
	lOddRI := gIsOpenRIInString(s[:l0Pos], l0)

	for l0Pos+r0Delta < l {
		r0, r1Delta := gFirstClassInString(s[l0Pos+r0Delta:])
		if gDecision(l1Diff, l0, lOddRI, r0) {
			return l0Pos + r0Delta
		}
		if l0 != r0 {
			l1Diff = l0
		}
		l0 = r0
		lOddRI = l0 == gClassRI && !lOddRI
		l0Pos += r0Delta
		r0Delta = r1Delta
	}
	return l
}

// GraphemeClusterEndInString computes grapheme cluster which contains pos-th rune.
// Returns (index of grapheme cluster's last rune)+1.
// In other words, returns first grapheme cluster's boundary on the right of pos-th rune.
func GraphemeClusterEndInString(s string, pos int) int {
	l := len(s)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	return graphemeClusterEndInString(s, pos)
}

// s must be valid (len>1).
// r0Pos must be valid (in s).
func graphemeClusterBeginInString(s string, r0Pos int) int {
	if r0Pos >= 1 && s[r0Pos-1] == crRune && s[r0Pos] == lfRune { // GB3
		return r0Pos - 1
	}

	r0Pos = toRuneBeginInString(s, r0Pos) // TODO do it only on external call
	r0, _ := gFirstClassInString(s[r0Pos:])

	for r0Pos > 0 {
		l0, l0Pos := gLastClassInString(s[:r0Pos])
		l1Diff := gLastNotEqualToInString(s[:l0Pos], l0)
		lOddRI := gIsOpenRIInString(s[:l0Pos], l0)

		if gDecision(l1Diff, l0, lOddRI, r0) {
			return r0Pos
		}

		r0 = l0
		r0Pos = l0Pos
	}
	return 0
}

// GraphemeClusterBeginInString computes grapheme cluster which contains pos-th rune.
// Returns grapheme cluster's first rune index.
// In other words, returns first grapheme cluster's boundary on the left of pos-th rune.
func GraphemeClusterBeginInString(s string, pos int) int {
	l := len(s)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	return graphemeClusterBeginInString(s, pos)
}

// GraphemeClusterAtInString computes grapheme clusters which contains pos-th rune and return their boundary.
// Grapheme cluster may retrieved by "s[r.From:r.To]".
func GraphemeClusterAtInString(s string, pos int) Boundary {
	return Boundary{GraphemeClusterBeginInString(s, pos), GraphemeClusterEndInString(s, pos)}
}

// FirstGraphemeClusterInString computes first grapheme cluster.
// Returns (index of cluster's last rune)+1.
// Result also may be treated as length of the first grapheme cluster.
// First grapheme cluster may retrieved by "s[:r]".
func FirstGraphemeClusterInString(s string) (r int) {
	return GraphemeClusterEndInString(s, 0)
}

// LastGraphemeClusterInString computes last grapheme cluster.
// Returns index of cluster's first rune.
// Last grapheme cluster may retrieved by "s[r:]".
func LastGraphemeClusterInString(s string) (r int) {
	return GraphemeClusterBeginInString(s, len(s)-1)
}

// GraphemeClustersInString computes all grapheme clusters and returns theirs boundaries.
func GraphemeClustersInString(s string) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(s)) // TODO memory efficient
	for i := 0; i < len(s); {
		length := FirstGraphemeClusterInString(s[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}

// GraphemeClusterBreaksInString computes all grapheme clusters and returns all breaks.
func GraphemeClusterBreaksInString(s string) (breaks []int) {
	l := len(s)
	if l == 0 {
		return // []int{0}
	}
	breaks = make([]int, 1, len(s)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstGraphemeClusterInString(s[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}

// Returns class of last rune in bytes which is not equal to l0.
func gLastNotEqualTo(bytes []byte, l0 gClass) gClass {
	for len(bytes) > 0 {
		c, pos := gLastClass(bytes)
		if c != l0 {
			return c
		}
		bytes = bytes[:pos]
	}

	return gClassOther
}

// True if l0 is RI and it opens RI sequence in string <bytes..., l0, ...> (may be joined with next RI).
func gIsOpenRI(bytes []byte, l0 gClass) (res bool) {
	if l0 != gClassRI {
		return
	}
	res = true

	for len(bytes) > 0 {
		c, pos := gLastClass(bytes)
		if c != gClassRI {
			return
		}
		res = !res
		bytes = bytes[:pos]
	}

	return
}

// bytes must be valid (len>1).
// l0Pos must be valid (in bytes).
func graphemeClusterEnd(bytes []byte, l0Pos int) int {
	l := len(bytes)

	if l0Pos+1 < l && bytes[l0Pos] == crRune && bytes[l0Pos+1] == lfRune { // GB3
		return l0Pos + 2
	}

	l0Pos = toRuneBegin(bytes, l0Pos) // TODO do it only on external call
	l0, r0Delta := gFirstClass(bytes[l0Pos:])
	l1Diff := gLastNotEqualTo(bytes[:l0Pos], l0)
	lOddRI := gIsOpenRI(bytes[:l0Pos], l0)

	for l0Pos+r0Delta < l {
		r0, r1Delta := gFirstClass(bytes[l0Pos+r0Delta:])
		if gDecision(l1Diff, l0, lOddRI, r0) {
			return l0Pos + r0Delta
		}
		if l0 != r0 {
			l1Diff = l0
		}
		l0 = r0
		lOddRI = l0 == gClassRI && !lOddRI
		l0Pos += r0Delta
		r0Delta = r1Delta
	}
	return l
}

// GraphemeClusterEnd computes grapheme cluster which contains pos-th rune.
// Returns (index of grapheme cluster's last rune)+1.
// In other words, returns first grapheme cluster's boundary on the right of pos-th rune.
func GraphemeClusterEnd(bytes []byte, pos int) int {
	l := len(bytes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	return graphemeClusterEnd(bytes, pos)
}

// bytes must be valid (len>1).
// r0Pos must be valid (in bytes).
func graphemeClusterBegin(bytes []byte, r0Pos int) int {
	if r0Pos >= 1 && bytes[r0Pos-1] == crRune && bytes[r0Pos] == lfRune { // GB3
		return r0Pos - 1
	}

	r0Pos = toRuneBegin(bytes, r0Pos) // TODO do it only on external call
	r0, _ := gFirstClass(bytes[r0Pos:])

	for r0Pos > 0 {
		l0, l0Pos := gLastClass(bytes[:r0Pos])
		l1Diff := gLastNotEqualTo(bytes[:l0Pos], l0)
		lOddRI := gIsOpenRI(bytes[:l0Pos], l0)

		if gDecision(l1Diff, l0, lOddRI, r0) {
			return r0Pos
		}

		r0 = l0
		r0Pos = l0Pos
	}
	return 0
}

// GraphemeClusterBegin computes grapheme cluster which contains pos-th rune.
// Returns grapheme cluster's first rune index.
// In other words, returns first grapheme cluster's boundary on the left of pos-th rune.
func GraphemeClusterBegin(bytes []byte, pos int) int {
	l := len(bytes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	return graphemeClusterBegin(bytes, pos)
}

// GraphemeClusterAt computes grapheme clusters which contains pos-th rune and return their boundary.
// Grapheme cluster may retrieved by "bytes[r.From:r.To]".
func GraphemeClusterAt(bytes []byte, pos int) Boundary {
	return Boundary{GraphemeClusterBegin(bytes, pos), GraphemeClusterEnd(bytes, pos)}
}

// FirstGraphemeCluster computes first grapheme cluster.
// Returns (index of cluster's last rune)+1.
// Result also may be treated as length of the first grapheme cluster.
// First grapheme cluster may retrieved by "bytes[:r]".
func FirstGraphemeCluster(bytes []byte) (r int) {
	return GraphemeClusterEnd(bytes, 0)
}

// LastGraphemeCluster computes last grapheme cluster.
// Returns index of cluster's first rune.
// Last grapheme cluster may retrieved by "bytes[r:]".
func LastGraphemeCluster(bytes []byte) (r int) {
	return GraphemeClusterBegin(bytes, len(bytes)-1)
}

// GraphemeClusters computes all grapheme clusters and returns theirs boundaries.
func GraphemeClusters(bytes []byte) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(bytes)) // TODO memory efficient
	for i := 0; i < len(bytes); {
		length := FirstGraphemeCluster(bytes[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}

// GraphemeClusterBreaks computes all grapheme clusters and returns all breaks.
func GraphemeClusterBreaks(bytes []byte) (breaks []int) {
	l := len(bytes)
	if l == 0 {
		return // []int{0}
	}
	breaks = make([]int, 1, len(bytes)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstGraphemeCluster(bytes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}
