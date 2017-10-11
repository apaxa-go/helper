package unicodeh

import "unicode"

// Unicode property "Bidi_Paired_Bracket_Type" (known as "bpt", "Bidi_Paired_Bracket_Type").
// Kind of property: "Enumerated".
// Based on file "BidiBrackets.txt".
var (
	BidiPairedBracketTypeClose = bidiPairedBracketTypeClose // Value "Close" (known as "c", "Close").
	BidiPairedBracketTypeNone  = bidiPairedBracketTypeNone  // Value "None" (known as "n", "None").
	BidiPairedBracketTypeOpen  = bidiPairedBracketTypeOpen  // Value "Open" (known as "o", "Open").
)

var (
	bidiPairedBracketTypeClose = &unicode.RangeTable{[]unicode.Range16{{0x29, 0x5d, 0x34}, {0x7d, 0xf3b, 0xebe}, {0xf3d, 0x169c, 0x75f}, {0x2046, 0x207e, 0x38}, {0x208e, 0x2309, 0x27b}, {0x230b, 0x232a, 0x1f}, {0x2769, 0x2775, 0x2}, {0x27c6, 0x27e7, 0x21}, {0x27e9, 0x27ef, 0x2}, {0x2984, 0x2998, 0x2}, {0x29d9, 0x29db, 0x2}, {0x29fd, 0x2e23, 0x426}, {0x2e25, 0x2e29, 0x2}, {0x3009, 0x3011, 0x2}, {0x3015, 0x301b, 0x2}, {0xfe5a, 0xfe5e, 0x2}, {0xff09, 0xff3d, 0x34}, {0xff5d, 0xff63, 0x3}}, nil, 1}
	bidiPairedBracketTypeNone  = &unicode.RangeTable{[]unicode.Range16{{0x0, 0x27, 0x1}, {0x2a, 0x5a, 0x1}, {0x5c, 0x5e, 0x2}, {0x5f, 0x7a, 0x1}, {0x7c, 0x7e, 0x2}, {0x7f, 0xf39, 0x1}, {0xf3e, 0x169a, 0x1}, {0x169d, 0x2044, 0x1}, {0x2047, 0x207c, 0x1}, {0x207f, 0x208c, 0x1}, {0x208f, 0x2307, 0x1}, {0x230c, 0x2328, 0x1}, {0x232b, 0x2767, 0x1}, {0x2776, 0x27c4, 0x1}, {0x27c7, 0x27e5, 0x1}, {0x27f0, 0x2982, 0x1}, {0x2999, 0x29d7, 0x1}, {0x29dc, 0x29fb, 0x1}, {0x29fe, 0x2e21, 0x1}, {0x2e2a, 0x3007, 0x1}, {0x3012, 0x3013, 0x1}, {0x301c, 0xfe58, 0x1}, {0xfe5f, 0xff07, 0x1}, {0xff0a, 0xff3a, 0x1}, {0xff3c, 0xff3e, 0x2}, {0xff3f, 0xff5a, 0x1}, {0xff5c, 0xff5e, 0x2}, {0xff61, 0xff64, 0x3}, {0xff65, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x10ffff, 0x1}}, 5}
	bidiPairedBracketTypeOpen  = &unicode.RangeTable{[]unicode.Range16{{0x28, 0x5b, 0x33}, {0x7b, 0xf3a, 0xebf}, {0xf3c, 0x169b, 0x75f}, {0x2045, 0x207d, 0x38}, {0x208d, 0x2308, 0x27b}, {0x230a, 0x2329, 0x1f}, {0x2768, 0x2774, 0x2}, {0x27c5, 0x27e6, 0x21}, {0x27e8, 0x27ee, 0x2}, {0x2983, 0x2997, 0x2}, {0x29d8, 0x29da, 0x2}, {0x29fc, 0x2e22, 0x426}, {0x2e24, 0x2e28, 0x2}, {0x3008, 0x3010, 0x2}, {0x3014, 0x301a, 0x2}, {0xfe59, 0xfe5d, 0x2}, {0xff08, 0xff3b, 0x33}, {0xff5b, 0xff5f, 0x4}, {0xff62, 0xff62, 0x1}}, nil, 1}
)