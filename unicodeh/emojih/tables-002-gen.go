package emojih

import "unicode"

// Emoji property "Emoji_Modifier".
// Kind of property: "Binary".
// Based on file "emoji-data.txt".
var (
	EmojiModifierYes = emojiModifierYes // Value "Yes" (known as "Yes", "Y", "True", "T").
	EmojiModifierNo  = emojiModifierNo  // Value "No" (known as "No", "N", "False", "F").
)

var (
	emojiModifierYes = &unicode.RangeTable{nil, []unicode.Range32{{0x1f3fb, 0x1f3ff, 0x1}}, 0}
	emojiModifierNo  = &unicode.RangeTable{[]unicode.Range16{{0x0, 0xffff, 0x1}}, []unicode.Range32{{0x10000, 0x1f3fa, 0x1}, {0x1f400, 0x10ffff, 0x1}}, 0}
)
