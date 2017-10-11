package boundaryh

import "github.com/apaxa-go/helper/unicodeh"

type gClass uint8

const (
	gClassOther        gClass = iota
	gClassControl      gClass = iota
	gClassHangulL      gClass = iota
	gClassHangulV      gClass = iota
	gClassHangulT      gClass = iota
	gClassHangulLV     gClass = iota
	gClassHangulLVT    gClass = iota
	gClassExtend       gClass = iota
	gClassSpacingMark  gClass = iota
	gClassPrepend      gClass = iota
	gClassZWJ          gClass = iota
	gClassEBase        gClass = iota
	gClassEBG          gClass = iota
	gClassEModifier    gClass = iota
	gClassGlueAfterZWJ gClass = iota
	gClassRI           gClass = iota
)

// TODO
func gGetClass(r rune) gClass {
	switch {
	default:
		return gClassOther
	case unicodeh.IsGraphemeClusterBreakControl(r) || r == crRune || r == lfRune:
		return gClassControl
	case unicodeh.IsGraphemeClusterBreakL(r):
		return gClassHangulL
	case unicodeh.IsGraphemeClusterBreakV(r):
		return gClassHangulV
	case unicodeh.IsGraphemeClusterBreakT(r):
		return gClassHangulT
	case unicodeh.IsGraphemeClusterBreakLV(r):
		return gClassHangulLV
	case unicodeh.IsGraphemeClusterBreakLVT(r):
		return gClassHangulLVT
	case unicodeh.IsGraphemeClusterBreakExtend(r):
		return gClassExtend
	case unicodeh.IsGraphemeClusterBreakSpacingMark(r):
		return gClassSpacingMark
	case unicodeh.IsGraphemeClusterBreakPrepend(r):
		return gClassPrepend
	case unicodeh.IsGraphemeClusterBreakZWJ(r):
		return gClassZWJ
	case unicodeh.IsGraphemeClusterBreakEBase(r):
		return gClassEBase
	case unicodeh.IsGraphemeClusterBreakEBaseGAZ(r):
		return gClassEBG
	case unicodeh.IsGraphemeClusterBreakEModifier(r):
		return gClassEModifier
	case unicodeh.IsGraphemeClusterBreakGlueAfterZwj(r):
		return gClassGlueAfterZWJ
	case unicodeh.IsGraphemeClusterBreakRegionalIndicator(r):
		return gClassRI
	}
}
