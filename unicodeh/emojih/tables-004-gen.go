package emojih

import "unicode"

// Emoji property "Emoji_Component".
// Kind of property: "Binary".
// Based on file "emoji-data.txt".
var (
	EmojiComponentYes = emojiComponentYes // Value "Yes" (known as "Yes", "Y", "True", "T").
	EmojiComponentNo  = emojiComponentNo  // Value "No" (known as "No", "N", "False", "F").
)

var (
	emojiComponentYes = &unicode.RangeTable{[]unicode.Range16{{0x23, 0x2a, 0x7}, {0x30, 0x39, 0x1}, {0x200d, 0x20e3, 0xd6}, {0xfe0f, 0xfe0f, 0x1}}, []unicode.Range32{{0x1f1e6, 0x1f1ff, 0x1}, {0x1f3fb, 0x1f3ff, 0x1}, {0xe0020, 0xe007f, 0x1}}, 2}
	emojiComponentNo  = &unicode.RangeTable{[]unicode.Range16{{0x0, 0x22, 0x1}, {0x24, 0x29, 0x1}, {0x2b, 0x2f, 0x1}, {0x3a, 0x200c, 0x1}, {0x200e, 0x20e2, 0x1}, {0x20e4, 0xfe0e, 0x1}, {0xfe10, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x1f1e5, 0x1}, {0x1f200, 0x1f3fa, 0x1}, {0x1f400, 0xe001f, 0x1}, {0xe0080, 0x10ffff, 0x1}}, 3}
)