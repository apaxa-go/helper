package boundaryh

// Returns class of last rune in runes which is not equal to l0.
func gLastNotEqualTo(runes []rune, l0 gClass) gClass {
	for pos := len(runes) - 1; pos >= 0; pos-- {
		if r := gGetClass(runes[pos]); r != l0 {
			return r
		}
	}
	return gClassOther
}

// True if l0 is RI and it opens RI sequence in string <runes..., l0, ...> (may be joined with next RI).
func gIsOpenRI(runes []rune, l0 gClass) (r bool) {
	if l0 != gClassRI {
		return
	}
	r = true
	for pos := len(runes) - 1; pos >= 0 && gGetClass(runes[pos]) == gClassRI; pos-- {
		r = !r
	}
	return
}

// Returns if there is a break between l0 and r0.
func gDecision(l1Diff, l0 gClass, lOddRI bool, r0 gClass) bool {
	// TODO translate to single boolean expression?
	switch {
	case l0 == gClassControl || r0 == gClassControl: // GB4 & GB5
		return true
	case l0 == gClassHangulL && (r0 == gClassHangulL || r0 == gClassHangulV || r0 == gClassHangulLV || r0 == gClassHangulLVT): // GB6
	case (l0 == gClassHangulLV || l0 == gClassHangulV) && (r0 == gClassHangulV || r0 == gClassHangulT): // GB7
	case (l0 == gClassHangulLVT || l0 == gClassHangulT) && r0 == gClassHangulT: // GB8
	case r0 == gClassExtend || r0 == gClassZWJ: // GB9
	case r0 == gClassSpacingMark: //GB9a
	case l0 == gClassPrepend: //GB9b
	case ((l0 == gClassEBase || l0 == gClassEBG) || (l0 == gClassExtend && (l1Diff == gClassEBase || l1Diff == gClassEBG))) && r0 == gClassEModifier: // GB10
	case l0 == gClassZWJ && (r0 == gClassGlueAfterZWJ || r0 == gClassEBG): // GB11
	case lOddRI && r0 == gClassRI: // GB12 & GB13
	default:
		return true
	}
	return false
}

// runes must be valid (len>1).
// l0Pos must be valid (in runes).
func graphemeClusterEnd(runes []rune, l0Pos int) int {
	l := len(runes)

	if l0Pos+1 < l && runes[l0Pos] == crRune && runes[l0Pos+1] == lfRune { // GB3
		return l0Pos + 2
	}

	l0 := gGetClass(runes[l0Pos])
	l1Diff := gLastNotEqualTo(runes[:l0Pos], l0)
	lOddRI := gIsOpenRI(runes[:l0Pos], l0)

	for ; l0Pos+1 < l; l0Pos++ {
		r0 := gGetClass(runes[l0Pos+1])

		if gDecision(l1Diff, l0, lOddRI, r0) {
			return l0Pos + 1
		}

		if l0 != r0 {
			l1Diff = l0
		}
		l0 = r0
		lOddRI = l0 == gClassRI && !lOddRI
	}
	return l
}

// GraphemeClusterEnd computes grapheme cluster which contains pos-th rune.
// Returns (index of grapheme cluster's last rune)+1.
// In other words, returns first grapheme cluster's boundary on the right of pos-th rune.
func GraphemeClusterEnd(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	return graphemeClusterEnd(runes, pos)
}

// runes must be valid (len>1).
// r0Pos must be valid (in runes).
func graphemeClusterBegin(runes []rune, r0Pos int) int {
	if r0Pos >= 1 && runes[r0Pos-1] == crRune && runes[r0Pos] == lfRune { // GB3
		return r0Pos - 1
	}

	r0 := gGetClass(runes[r0Pos])

	for ; r0Pos > 0; r0Pos-- {
		l0 := gGetClass(runes[r0Pos-1])
		l1Diff := gLastNotEqualTo(runes[:r0Pos-1], l0)
		lOddRI := gIsOpenRI(runes[:r0Pos-1], l0)

		if gDecision(l1Diff, l0, lOddRI, r0) {
			return r0Pos
		}

		r0 = l0
	}
	return 0
}

// GraphemeClusterBegin computes grapheme cluster which contains pos-th rune.
// Returns grapheme cluster's first rune index.
// In other words, returns first grapheme cluster's boundary on the left of pos-th rune.
func GraphemeClusterBegin(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	return graphemeClusterBegin(runes, pos)
}

// GraphemeClusterAt computes grapheme clusters which contains pos-th rune and return their boundary.
// Grapheme cluster may retrieved by "runes[r.From:r.To]".
func GraphemeClusterAt(runes []rune, pos int) Boundary {
	return Boundary{GraphemeClusterBegin(runes, pos), GraphemeClusterEnd(runes, pos)}
}

// FirstGraphemeCluster computes first grapheme cluster.
// Returns (index of cluster's last rune)+1.
// Result also may be treated as length of the first grapheme cluster.
// First grapheme cluster may retrieved by "runes[:r]".
func FirstGraphemeCluster(runes []rune) (r int) {
	return GraphemeClusterEnd(runes, 0)
}

// LastGraphemeCluster computes last grapheme cluster.
// Returns index of cluster's first rune.
// Last grapheme cluster may retrieved by "runes[r:]".
func LastGraphemeCluster(runes []rune) (r int) {
	return GraphemeClusterBegin(runes, len(runes)-1)
}

// GraphemeClusters computes all grapheme clusters and returns theirs boundaries.
func GraphemeClusters(runes []rune) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(runes)) // TODO memory efficient
	for i := 0; i < len(runes); {
		length := FirstGraphemeCluster(runes[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}

// GraphemeClusterBreaks computes all grapheme clusters and returns all breaks.
func GraphemeClusterBreaks(runes []rune) (breaks []int) {
	l := len(runes)
	if l == 0 {
		return // []int{0}
	}
	breaks = make([]int, 1, len(runes)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstGraphemeCluster(runes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}
