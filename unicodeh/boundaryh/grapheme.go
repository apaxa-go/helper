package boundaryh

//replacer:ignore
// TODO replace windows path separator
//go:generate go run $GOPATH\src\github.com\apaxa-go\generator\replacer\main.go -- $GOFILE

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

//replacer:replace
//replacer:old InRunes	[]rune	runes
//replacer:new InString	string	s
//replacer:new ""		[]byte	bytes

// Returns class of last rune in runes which is not equal to l0.
func gLastNotEqualToInRunes(runes []rune, l0 gClass) gClass {
	for len(runes) > 0 {
		c, pos := gLastClassInRunes(runes)
		if c != l0 {
			return c
		}
		runes = runes[:pos]
	}

	return gClassOther
}

// True if l0 is RI and it opens RI sequence in string <runes..., l0, ...> (may be joined with next RI).
func gIsOpenRIInRunes(runes []rune, l0 gClass) (res bool) {
	if l0 != gClassRI {
		return
	}
	res = true

	for len(runes) > 0 {
		c, pos := gLastClassInRunes(runes)
		if c != gClassRI {
			return
		}
		res = !res
		runes = runes[:pos]
	}

	return
}

// runes must be valid (len>1).
// l0Pos must be valid (in runes).
func graphemeClusterEndInRunes(runes []rune, l0Pos int) int {
	l := len(runes)

	if l0Pos+1 < l && runes[l0Pos] == crRune && runes[l0Pos+1] == lfRune { // GB3
		return l0Pos + 2
	}

	l0Pos = toRuneBeginInRunes(runes, l0Pos) // TODO do it only on external call
	l0, r0Delta := gFirstClassInRunes(runes[l0Pos:])
	l1Diff := gLastNotEqualToInRunes(runes[:l0Pos], l0)
	lOddRI := gIsOpenRIInRunes(runes[:l0Pos], l0)

	for l0Pos+r0Delta < l {
		r0, r1Delta := gFirstClassInRunes(runes[l0Pos+r0Delta:])
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

// GraphemeClusterEndInRunes computes grapheme cluster which contains pos-th rune.
// Returns (index of grapheme cluster's last rune)+1.
// In other words, returns first grapheme cluster's boundary on the right of pos-th rune.
func GraphemeClusterEndInRunes(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == l-1 {
		return l
	}

	return graphemeClusterEndInRunes(runes, pos)
}

// runes must be valid (len>1).
// r0Pos must be valid (in runes).
func graphemeClusterBeginInRunes(runes []rune, r0Pos int) int {
	if r0Pos >= 1 && runes[r0Pos-1] == crRune && runes[r0Pos] == lfRune { // GB3
		return r0Pos - 1
	}

	r0Pos = toRuneBeginInRunes(runes, r0Pos) // TODO do it only on external call
	r0, _ := gFirstClassInRunes(runes[r0Pos:])

	for r0Pos > 0 {
		l0, l0Pos := gLastClassInRunes(runes[:r0Pos])
		l1Diff := gLastNotEqualToInRunes(runes[:l0Pos], l0)
		lOddRI := gIsOpenRIInRunes(runes[:l0Pos], l0)

		if gDecision(l1Diff, l0, lOddRI, r0) {
			return r0Pos
		}

		r0 = l0
		r0Pos = l0Pos
	}
	return 0
}

// GraphemeClusterBeginInRunes computes grapheme cluster which contains pos-th rune.
// Returns grapheme cluster's first rune index.
// In other words, returns first grapheme cluster's boundary on the left of pos-th rune.
func GraphemeClusterBeginInRunes(runes []rune, pos int) int {
	l := len(runes)
	if pos < 0 || pos >= l {
		return InvalidPos
	}
	if pos == 0 {
		return 0
	}

	return graphemeClusterBeginInRunes(runes, pos)
}

// GraphemeClusterAtInRunes computes grapheme clusters which contains pos-th rune and return their boundary.
// Grapheme cluster may retrieved by "runes[r.From:r.To]".
func GraphemeClusterAtInRunes(runes []rune, pos int) Boundary {
	return Boundary{GraphemeClusterBeginInRunes(runes, pos), GraphemeClusterEndInRunes(runes, pos)}
}

// FirstGraphemeClusterInRunes computes first grapheme cluster.
// Returns (index of cluster's last rune)+1.
// Result also may be treated as length of the first grapheme cluster.
// First grapheme cluster may retrieved by "runes[:r]".
func FirstGraphemeClusterInRunes(runes []rune) (r int) {
	return GraphemeClusterEndInRunes(runes, 0)
}

// LastGraphemeClusterInRunes computes last grapheme cluster.
// Returns index of cluster's first rune.
// Last grapheme cluster may retrieved by "runes[r:]".
func LastGraphemeClusterInRunes(runes []rune) (r int) {
	return GraphemeClusterBeginInRunes(runes, len(runes)-1)
}

// GraphemeClustersInRunes computes all grapheme clusters and returns theirs boundaries.
func GraphemeClustersInRunes(runes []rune) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(runes)) // TODO memory efficient
	for i := 0; i < len(runes); {
		length := FirstGraphemeClusterInRunes(runes[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}

// GraphemeClusterBreaksInRunes computes all grapheme clusters and returns all breaks.
func GraphemeClusterBreaksInRunes(runes []rune) (breaks []int) {
	l := len(runes)
	if l == 0 {
		return // []int{0}
	}
	breaks = make([]int, 1, len(runes)) // TODO memory efficient
	breaks[0] = 0
	for pos := 0; pos < l; {
		length := FirstGraphemeClusterInRunes(runes[pos:])
		pos += length
		breaks = append(breaks, pos)
	}
	return
}
