package boundaryh

const (
	crRune rune = '\u000D'
	lfRune rune = '\u000A'
)

// InvalidPos indicates invalid result position.
// This means that arguments to function is invalid (i.e. looking for word in empty string or looking for grapheme cluster at outside of string).
const InvalidPos = -1

// Boundary represents subslice by its indexes.
// Subslice can be retrieved by "s[b.From:b.To]".
type Boundary struct {
	From int // index of first element
	To   int // (index of last element) + 1
}

// Invalid returns invalid boundary.
func Invalid() Boundary { return Boundary{InvalidPos, InvalidPos} }

// Len returns length of subslice represented by boundary.
// Len of invalid boundary is 0. Valid boundary also may have zero length.
func (b Boundary) Len() int { return b.To - b.From }

// IsInvalid checks if boundary is invalid in terms of "Invalid" constructor (and does not check anything else).
// Warning: !b.IsValid() is not equal to b.IsInvalid().
func (b Boundary) IsInvalid() bool { return b == Invalid() }

// IsValid checks if boundary is valid in general.
// Warning: !b.IsValid() is not equal to b.IsInvalid().
func (b Boundary) IsValid() bool { return b.From >= 0 && b.To >= b.From }
