package boundaryh

import (
	"github.com/apaxa-go/helper/unicodeh"
)

///go:generate go run ./internal/maketables/maketables.go

type gCClass uint8

const (
	gCClassOther        gCClass = iota
	gCClassControl      gCClass = iota
	gCClassHangulL      gCClass = iota
	gCClassHangulV      gCClass = iota
	gCClassHangulT      gCClass = iota
	gCClassHangulLV     gCClass = iota
	gCClassHangulLVT    gCClass = iota
	gCClassExtend       gCClass = iota
	gCClassSpacingMark  gCClass = iota
	gCClassPrepend      gCClass = iota
	gCClassZWJ          gCClass = iota
	gCClassEBase        gCClass = iota
	gCClassEBG          gCClass = iota
	gCClassEModifier    gCClass = iota
	gCClassGlueAfterZWJ gCClass = iota
	gCClassRI           gCClass = iota
)

// TODO
func getGCClass(r rune) gCClass {
	switch {
	default:
		return gCClassOther
	case unicodeh.IsGraphemeClusterBreakControl(r) || r==crRune || r==lfRune:
		return gCClassControl
	case unicodeh.IsGraphemeClusterBreakL(r):
		return gCClassHangulL
	case unicodeh.IsGraphemeClusterBreakV(r):
		return gCClassHangulV
	case unicodeh.IsGraphemeClusterBreakT(r):
		return gCClassHangulT
	case unicodeh.IsGraphemeClusterBreakLV(r):
		return gCClassHangulLV
	case unicodeh.IsGraphemeClusterBreakLVT(r):
		return gCClassHangulLVT
	case unicodeh.IsGraphemeClusterBreakExtend(r):
		return gCClassExtend
	case unicodeh.IsGraphemeClusterBreakSpacingMark(r):
		return gCClassSpacingMark
	case unicodeh.IsGraphemeClusterBreakPrepend(r):
		return gCClassPrepend
	case unicodeh.IsGraphemeClusterBreakZWJ(r):
		return gCClassZWJ
	case unicodeh.IsGraphemeClusterBreakEBase(r):
		return gCClassEBase
	case unicodeh.IsGraphemeClusterBreakEBaseGAZ(r):
		return gCClassEBG
	case unicodeh.IsGraphemeClusterBreakEModifier(r):
		return gCClassEModifier
	case unicodeh.IsGraphemeClusterBreakGlueAfterZwj(r):
		return gCClassGlueAfterZWJ
	case unicodeh.IsGraphemeClusterBreakRegionalIndicator(r):
		return gCClassRI
	}
}

func FirstGraphemeCluster(runes []rune) int {
	l := len(runes)
	if l == 0 {
		return 0
	}

	if l >= 2 && runes[0] == crRune && runes[1] == lfRune {	// GB3
		return 2
	}

	prevPrevClass := gCClassOther
	prevPrevDifferentClass := gCClassOther
	prevClass := getGCClass(runes[0])
	if prevClass == gCClassControl { // GB4
		return 1
	}
	for i := 1; i < l; i++ {
		curClass := getGCClass(runes[i])
		switch {
		case curClass == gCClassControl: // GB5
			return i
		case prevClass == gCClassHangulL && (curClass == gCClassHangulL || curClass == gCClassHangulV || curClass == gCClassHangulLV || curClass == gCClassHangulLVT): // GB6
		case (prevClass == gCClassHangulLV || prevClass == gCClassHangulV) && (curClass == gCClassHangulV || curClass == gCClassHangulT): // GB7
		case (prevClass == gCClassHangulLVT || prevClass == gCClassHangulT) && curClass == gCClassHangulT: // GB8
		case curClass == gCClassExtend || curClass == gCClassZWJ: // GB9
		case curClass == gCClassSpacingMark: //GB9a
		case prevClass == gCClassPrepend: //GB9b
		case ((prevClass == gCClassEBase || prevClass == gCClassEBG) || (prevClass == gCClassExtend && (prevPrevDifferentClass == gCClassEBase || prevPrevDifferentClass == gCClassEBG))) && curClass == gCClassEModifier: // GB10
		case prevClass == gCClassZWJ && (curClass == gCClassGlueAfterZWJ || curClass == gCClassEBG): // GB11
		case prevPrevClass != gCClassRI && prevClass == gCClassRI && curClass == gCClassRI: // GB12 & GB13
		default:
			return i
		}
		prevPrevClass = prevClass
		prevClass = curClass
		if prevPrevClass != prevClass {
			prevPrevDifferentClass = prevPrevClass
		}
	}
	return len(runes)
}

func GraphemeClusters(runes []rune) (boundaries []Boundary) {
	boundaries = make([]Boundary, 0, len(runes)) // TODO memory efficient
	for i := 0; i < len(runes); {
		//length:=FirstBoundaryRunes(runes[i:])
		length := FirstGraphemeCluster(runes[i:])
		boundaries = append(boundaries, Boundary{i, i + length})
		i += length
	}
	return
}
