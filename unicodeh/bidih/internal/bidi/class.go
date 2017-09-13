package bidi

type Class uint8

const (
	ArabicLetter            Class = iota // 0
	ArabicNumber          Class = iota   // 1
	ParagraphSeparator    Class = iota   // 2
	BoundaryNeutral       Class = iota   // 3
	CommonSeparator       Class = iota   // 4
	EuropeanNumber        Class = iota   // 5
	EuropeanSeparator     Class = iota   // 6
	EuropeanTerminator    Class = iota   // 7
	FirstStrongIsolate    Class = iota   // 8
	LeftToRight           Class = iota   // 9
	LeftToRightEmbedding  Class = iota   // 10
	LeftToRightIsolate    Class = iota   // 11
	LeftToRightOverride   Class = iota   // 12
	NonSpacingMark        Class = iota   // 13
	OtherNeutral          Class = iota   // 14
	PopDirectionalFormat  Class = iota   // 15
	PopDirectionalIsolate Class = iota   // 16
	RightToLeft           Class = iota   // 17
	RightToLeftEmbedding  Class = iota   // 18
	RightToLeftIsolate    Class = iota   // 19
	RightToLeftOverride   Class = iota   // 20
	SegmentSeparator      Class = iota   // 21
	WhiteSpace            Class = iota   // 22
	Count                       = iota
)
