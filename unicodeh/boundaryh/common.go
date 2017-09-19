package boundaryh

const (
	crRune rune = '\u000D'
	lfRune rune = '\u000A'
)

type Boundary struct {
	From, To int
}

func (b Boundary) Len() int { return b.To - b.From }
