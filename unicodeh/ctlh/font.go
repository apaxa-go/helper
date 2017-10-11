package ctlh

/*
import "github.com/apaxa-go/helper/unicodeh/boundaryh"

// TODO move this file to package ~font

type glyphBoundary struct {
	pos      int // not including
	runes boundaryh.Boundary // in storage order
}

// Single dimension (x axis) mapper coordinate to range of runes.
// len must be always > 0
// [0] is pseudo element with empty runes. It replaces "from".
// TODO is [0] ("from") always 0?
type GlyphBoundaries []glyphBoundary

// [From; To)
func (b GlyphBoundaries)From()int{return b[0].pos}
func (b GlyphBoundaries)To()int{return b[len(b)-1].pos}
func (b *GlyphBoundaries)Width()int{return b.To()-b.From()}

// TODO special value if nothing ?
// pos is x coordinate.
// Result is in storage order.
func (b GlyphBoundaries)At(pos int)boundaryh.Boundary {
	if pos<b.From() || pos>=b.To(){ return boundaryh.Invalid() }
	// TODO This must be done more effective by predicting result
	for i:=1; i<len(b); i++{
		if pos>=b[i-1].pos && pos<b[i].pos{
			return b[i].runes
		}
	}
	panic("Must be unreachable")
}

// runes is in storage order.
// order is visual order.
// positions is in storage order.
func computeVisualisation(runes []rune, order []int)(drawData interface{}, positions GlyphBoundaries){

}
