package ctlh

/*
import (
	"github.com/apaxa-go/helper/unicodeh/boundaryh"
	"github.com/apaxa-go/helper/unicodeh/bidih"
)

type Drawer interface {}

type Line struct {
	runes []rune
	//visualOrder []int
	drawData interface{} // TODO
	glyphBoundaries GlyphBoundaries
}

func (l *Line)computeVisualisation(){
	if len(l.runes)==0{
		l.resetVisialisation()
	}
	order :=bidih.ParseParagraph(l.runes,0,nil)                        // TODO use Parse* which returns only order // TODO define paragraphLevel // TODO what to do with line separators if any?
	l.drawData,l.glyphBoundaries= computeVisualisation(l.runes, order) // TODO this also may be performed at bidih
}

// Computes only if not yet computed
func (l *Line)computeVisualisationOnDemand(){
	if l.glyphBoundaries==nil && len(l.runes)!=0{
		l.computeVisualisation()
	}
}

func (l *Line)resetVisialisation(){
	l.drawData=nil
	l.glyphBoundaries=nil
}

func NewLineFromRunes(runes []rune)*Line{
	l:=&Line{runes, nil, nil}
	return l
}

func NewLineFromBytes(bytes []byte)*Line{
	return NewLine(string(bytes))
}

func NewLine(s string)*Line{
	return NewLineFromRunes([]rune(s))
}

func (l *Line)Runes() []rune{return l.runes}

func (l *Line)String()string{return string(l.Runes())}

func (l *Line)Bytes()[]byte{return []byte(l.String())}

func (l *Line)Draw(d Drawer){
	l.computeVisualisation()
	if l.drawData==nil{
		return
	}

	// TODO
}

// pos - x coordinate
// result in storage order
// TODO special value if pos out of string (StringAt must work)
// TODO what if glyph at "pos" contains multiple grapheme clusters?
// TODO looks like glyph must not contains breakable grapheme clusters sequence - this must be controlled.
func (l *Line)GraphemeClusterAt(pos int)boundaryh.Boundary{
	l.computeVisualisationOnDemand()
	b:=l.glyphBoundaries.At(pos)
	if b.IsInvalid() {
		return b
	}
	// TODO in boundaryh - go to both side
}

// pos - x coordinate
// result in storage order
func (l *Line)WordAt(pos int)boundaryh.Boundary{
	b:=l.GraphemeClusterAt(pos)
	if b.IsInvalid(){
			return b
	}
	// TODO  in boundaryh - go to both side
}

// pos - x coordinate
// result in storage order
func (l *Line)SentenceAt(pos int)boundaryh.Boundary{
	b:=l.GraphemeClusterAt(pos)
	if b.IsInvalid(){
		return b
	}
	// TODO  in boundaryh - go to both side
}

// [pos1; pos2] - x coordinate.
// result in storage order
func (l *Line)StringAt(pos1, pos2 int)boundaryh.Boundary{
	// TODO what if pos* out of string? Currently returns nothing (invalid boundary).
	gC1 :=l.GraphemeClusterAt(pos1)
	gC2 :=l.GraphemeClusterAt(pos2)
	if !gC1.IsInvalid() && !gC2.IsInvalid() {
		if gC1.From < gC2.To {
			return boundaryh.Boundary{gC1.From, gC2.To}
		} else if gC2.From < gC1.To {
			return boundaryh.Boundary{gC2.From, gC1.To}
		}
	}
	return boundaryh.Invalid()
}