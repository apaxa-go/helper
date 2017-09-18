package emojih

// IsEmojiModifierYes reports whether the rune has property "Emoji_Modifier"="Yes".
// Value "Yes" known as "Yes", "Y", "True", "T".
func IsEmojiModifierYes(r rune) bool { return (r >= 0x1f3fb && r <= 0x1f3ff) }

// IsEmojiModifierNo reports whether the rune has property "Emoji_Modifier"="No".
// Value "No" known as "No", "N", "False", "F".
func IsEmojiModifierNo(r rune) bool {
	return (r >= 0x0 && r <= 0xffff) || (r >= 0x10000 && r <= 0x1f3fa) || (r >= 0x1f400 && r <= 0x10ffff)
}
