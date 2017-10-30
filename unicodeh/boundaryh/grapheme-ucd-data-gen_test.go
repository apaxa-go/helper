package boundaryh

var ucdGraphemeClusterTests = []ucdTest{
	{[]rune{0x20, 0x20}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x20, 0xd}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x20, 0xa}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x20, 0x1}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x20, 0x300}, []int{0, 2}},
	{[]rune{0x20, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x20, 0x600}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x20, 0x903}, []int{0, 2}},
	{[]rune{0x20, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x20, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x20, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x20, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x20, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x20, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x20, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x20, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x20, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x20, 0x200d}, []int{0, 2}},
	{[]rune{0x20, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x20, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x20, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x20, 0x378}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x20, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x20, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0xd, 0x20}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x20}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0xd}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0xd}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0xa}, []int{0, 2}},
	{[]rune{0xd, 0x308, 0xa}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x1}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x1}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x300}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x300}, []int{0, 1, 3}},
	{[]rune{0xd, 0x600}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x600}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x903}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x903}, []int{0, 1, 3}},
	{[]rune{0xd, 0x1100}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x1100}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x1160}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x1160}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x11a8}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0xac00}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0xac00}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0xac01}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0xac01}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x1f1e6}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x261d}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x261d}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x1f3fb}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x200d}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x200d}, []int{0, 1, 3}},
	{[]rune{0xd, 0x2640}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x2640}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x1f466}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0x378}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0x378}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0xd800}, []int{0, 1, 2}},
	{[]rune{0xd, 0x308, 0xd800}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x20}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x20}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0xd}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0xd}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0xa}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0xa}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x1}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x1}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x300}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x300}, []int{0, 1, 3}},
	{[]rune{0xa, 0x600}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x600}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x903}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x903}, []int{0, 1, 3}},
	{[]rune{0xa, 0x1100}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x1100}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x1160}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x1160}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x11a8}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0xac00}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0xac00}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0xac01}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0xac01}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x1f1e6}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x261d}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x261d}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x1f3fb}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x200d}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x200d}, []int{0, 1, 3}},
	{[]rune{0xa, 0x2640}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x2640}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x1f466}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0x378}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0x378}, []int{0, 1, 2, 3}},
	{[]rune{0xa, 0xd800}, []int{0, 1, 2}},
	{[]rune{0xa, 0x308, 0xd800}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x20}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x20}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0xd}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0xd}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0xa}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0xa}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x1}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x1}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x300}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x300}, []int{0, 1, 3}},
	{[]rune{0x1, 0x600}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x600}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x903}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x903}, []int{0, 1, 3}},
	{[]rune{0x1, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x1100}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x1160}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x11a8}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0xac00}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0xac01}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x1f1e6}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x261d}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x1f3fb}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x200d}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x200d}, []int{0, 1, 3}},
	{[]rune{0x1, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x2640}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x1f466}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0x378}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0x378}, []int{0, 1, 2, 3}},
	{[]rune{0x1, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x1, 0x308, 0xd800}, []int{0, 1, 2, 3}},
	{[]rune{0x300, 0x20}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x300, 0xd}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x300, 0xa}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x300, 0x1}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x300, 0x300}, []int{0, 2}},
	{[]rune{0x300, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x300, 0x600}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x300, 0x903}, []int{0, 2}},
	{[]rune{0x300, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x300, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x300, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x300, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x300, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x300, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x300, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x300, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x300, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x300, 0x200d}, []int{0, 2}},
	{[]rune{0x300, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x300, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x300, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x300, 0x378}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x300, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x300, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x600, 0x20}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x600, 0xd}, []int{0, 1, 2}},
	{[]rune{0x600, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x600, 0xa}, []int{0, 1, 2}},
	{[]rune{0x600, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x600, 0x1}, []int{0, 1, 2}},
	{[]rune{0x600, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x600, 0x300}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x600, 0x600}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x600, 0x903}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x600, 0x1100}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x600, 0x1160}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x600, 0x11a8}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x600, 0xac00}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x600, 0xac01}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x600, 0x1f1e6}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x600, 0x261d}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x600, 0x1f3fb}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x600, 0x200d}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x600, 0x2640}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x600, 0x1f466}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x600, 0x378}, []int{0, 2}},
	{[]rune{0x600, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x600, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x600, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x903, 0x20}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x903, 0xd}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x903, 0xa}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x903, 0x1}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x903, 0x300}, []int{0, 2}},
	{[]rune{0x903, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x903, 0x600}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x903, 0x903}, []int{0, 2}},
	{[]rune{0x903, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x903, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x903, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x903, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x903, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x903, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x903, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x903, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x903, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x903, 0x200d}, []int{0, 2}},
	{[]rune{0x903, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x903, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x903, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x903, 0x378}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x903, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x903, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x20}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x1100, 0xd}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x1100, 0xa}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x1}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x300}, []int{0, 2}},
	{[]rune{0x1100, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x1100, 0x600}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x903}, []int{0, 2}},
	{[]rune{0x1100, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x1100, 0x1100}, []int{0, 2}},
	{[]rune{0x1100, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x1160}, []int{0, 2}},
	{[]rune{0x1100, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x1100, 0xac00}, []int{0, 2}},
	{[]rune{0x1100, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x1100, 0xac01}, []int{0, 2}},
	{[]rune{0x1100, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x200d}, []int{0, 2}},
	{[]rune{0x1100, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x1100, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x378}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x1100, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x1100, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x20}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x1160, 0xd}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x1160, 0xa}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x1}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x300}, []int{0, 2}},
	{[]rune{0x1160, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x1160, 0x600}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x903}, []int{0, 2}},
	{[]rune{0x1160, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x1160, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x1160}, []int{0, 2}},
	{[]rune{0x1160, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x11a8}, []int{0, 2}},
	{[]rune{0x1160, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x1160, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x1160, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x200d}, []int{0, 2}},
	{[]rune{0x1160, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x1160, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x1160, 0x378}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x1160, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x1160, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x20}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0xd}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0xa}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x1}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x300}, []int{0, 2}},
	{[]rune{0x11a8, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x11a8, 0x600}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x903}, []int{0, 2}},
	{[]rune{0x11a8, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x11a8, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x11a8}, []int{0, 2}},
	{[]rune{0x11a8, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x200d}, []int{0, 2}},
	{[]rune{0x11a8, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x11a8, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0x378}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x11a8, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x11a8, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x20}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0xac00, 0xd}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0xac00, 0xa}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x1}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x300}, []int{0, 2}},
	{[]rune{0xac00, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0xac00, 0x600}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x903}, []int{0, 2}},
	{[]rune{0xac00, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0xac00, 0x1100}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x1160}, []int{0, 2}},
	{[]rune{0xac00, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x11a8}, []int{0, 2}},
	{[]rune{0xac00, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0xac00, 0xac00}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0xac00, 0xac01}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x261d}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x200d}, []int{0, 2}},
	{[]rune{0xac00, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0xac00, 0x2640}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0xac00, 0x378}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0xac00, 0xd800}, []int{0, 1, 2}},
	{[]rune{0xac00, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x20}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0xac01, 0xd}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0xac01, 0xa}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x1}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x300}, []int{0, 2}},
	{[]rune{0xac01, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0xac01, 0x600}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x903}, []int{0, 2}},
	{[]rune{0xac01, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0xac01, 0x1100}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x1160}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x11a8}, []int{0, 2}},
	{[]rune{0xac01, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0xac01, 0xac00}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0xac01, 0xac01}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x261d}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x200d}, []int{0, 2}},
	{[]rune{0xac01, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0xac01, 0x2640}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x378}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0xac01, 0xd800}, []int{0, 1, 2}},
	{[]rune{0xac01, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x20}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0xd}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0xa}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x1}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x300}, []int{0, 2}},
	{[]rune{0x1f1e6, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x1f1e6, 0x600}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x903}, []int{0, 2}},
	{[]rune{0x1f1e6, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x1f1e6, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x1f1e6}, []int{0, 2}},
	{[]rune{0x1f1e6, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x200d}, []int{0, 2}},
	{[]rune{0x1f1e6, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x1f1e6, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x378}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x1f1e6, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x20}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x261d, 0xd}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x261d, 0xa}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x1}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x300}, []int{0, 2}},
	{[]rune{0x261d, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x261d, 0x600}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x903}, []int{0, 2}},
	{[]rune{0x261d, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x261d, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x261d, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x261d, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x1f3fb}, []int{0, 2}},
	{[]rune{0x261d, 0x308, 0x1f3fb}, []int{0, 3}},
	{[]rune{0x261d, 0x200d}, []int{0, 2}},
	{[]rune{0x261d, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x261d, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x261d, 0x378}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x261d, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x261d, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x20}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0xd}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0xa}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x1}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x300}, []int{0, 2}},
	{[]rune{0x1f3fb, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x1f3fb, 0x600}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x903}, []int{0, 2}},
	{[]rune{0x1f3fb, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x1f3fb, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x200d}, []int{0, 2}},
	{[]rune{0x1f3fb, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x1f3fb, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0x378}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x1f3fb, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x1f3fb, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x20}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x200d, 0xd}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x200d, 0xa}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x1}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x300}, []int{0, 2}},
	{[]rune{0x200d, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x200d, 0x600}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x903}, []int{0, 2}},
	{[]rune{0x200d, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x200d, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x200d, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x200d, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x200d}, []int{0, 2}},
	{[]rune{0x200d, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x200d, 0x2640}, []int{0, 2}},
	{[]rune{0x200d, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x1f466}, []int{0, 2}},
	{[]rune{0x200d, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x200d, 0x378}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x200d, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x200d, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x20}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x2640, 0xd}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x2640, 0xa}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x1}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x300}, []int{0, 2}},
	{[]rune{0x2640, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x2640, 0x600}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x903}, []int{0, 2}},
	{[]rune{0x2640, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x2640, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x2640, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x2640, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x200d}, []int{0, 2}},
	{[]rune{0x2640, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x2640, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x2640, 0x378}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x2640, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x2640, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x20}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0xd}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0xa}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x1}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x300}, []int{0, 2}},
	{[]rune{0x1f466, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x1f466, 0x600}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x903}, []int{0, 2}},
	{[]rune{0x1f466, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x1f466, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x1f3fb}, []int{0, 2}},
	{[]rune{0x1f466, 0x308, 0x1f3fb}, []int{0, 3}},
	{[]rune{0x1f466, 0x200d}, []int{0, 2}},
	{[]rune{0x1f466, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x1f466, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x378}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x1f466, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0x378, 0x20}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x20}, []int{0, 2, 3}},
	{[]rune{0x378, 0xd}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0xd}, []int{0, 2, 3}},
	{[]rune{0x378, 0xa}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0xa}, []int{0, 2, 3}},
	{[]rune{0x378, 0x1}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x1}, []int{0, 2, 3}},
	{[]rune{0x378, 0x300}, []int{0, 2}},
	{[]rune{0x378, 0x308, 0x300}, []int{0, 3}},
	{[]rune{0x378, 0x600}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x600}, []int{0, 2, 3}},
	{[]rune{0x378, 0x903}, []int{0, 2}},
	{[]rune{0x378, 0x308, 0x903}, []int{0, 3}},
	{[]rune{0x378, 0x1100}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x378, 0x1160}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x1160}, []int{0, 2, 3}},
	{[]rune{0x378, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x11a8}, []int{0, 2, 3}},
	{[]rune{0x378, 0xac00}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0xac00}, []int{0, 2, 3}},
	{[]rune{0x378, 0xac01}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0xac01}, []int{0, 2, 3}},
	{[]rune{0x378, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x1f1e6}, []int{0, 2, 3}},
	{[]rune{0x378, 0x261d}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x378, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x1f3fb}, []int{0, 2, 3}},
	{[]rune{0x378, 0x200d}, []int{0, 2}},
	{[]rune{0x378, 0x308, 0x200d}, []int{0, 3}},
	{[]rune{0x378, 0x2640}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x2640}, []int{0, 2, 3}},
	{[]rune{0x378, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x1f466}, []int{0, 2, 3}},
	{[]rune{0x378, 0x378}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0x378}, []int{0, 2, 3}},
	{[]rune{0x378, 0xd800}, []int{0, 1, 2}},
	{[]rune{0x378, 0x308, 0xd800}, []int{0, 2, 3}},
	{[]rune{0xd800, 0x20}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x20}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0xd}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0xd}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0xa}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0xa}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x1}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x1}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x300}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x300}, []int{0, 1, 3}},
	{[]rune{0xd800, 0x600}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x600}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x903}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x903}, []int{0, 1, 3}},
	{[]rune{0xd800, 0x1100}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x1100}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x1160}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x1160}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x11a8}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x11a8}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0xac00}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0xac00}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0xac01}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0xac01}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x1f1e6}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x1f1e6}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x261d}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x261d}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x1f3fb}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x1f3fb}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x200d}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x200d}, []int{0, 1, 3}},
	{[]rune{0xd800, 0x2640}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x2640}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x1f466}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x1f466}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0x378}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0x378}, []int{0, 1, 2, 3}},
	{[]rune{0xd800, 0xd800}, []int{0, 1, 2}},
	{[]rune{0xd800, 0x308, 0xd800}, []int{0, 1, 2, 3}},
	{[]rune{0xd, 0xa, 0x61, 0xa, 0x308}, []int{0, 2, 3, 4, 5}},
	{[]rune{0x61, 0x308}, []int{0, 2}},
	{[]rune{0x20, 0x200d, 0x646}, []int{0, 2, 3}},
	{[]rune{0x646, 0x200d, 0x20}, []int{0, 2, 3}},
	{[]rune{0x1100, 0x1100}, []int{0, 2}},
	{[]rune{0xac00, 0x11a8, 0x1100}, []int{0, 2, 3}},
	{[]rune{0xac01, 0x11a8, 0x1100}, []int{0, 2, 3}},
	{[]rune{0x1f1e6, 0x1f1e7, 0x1f1e8, 0x62}, []int{0, 2, 3, 4}},
	{[]rune{0x61, 0x1f1e6, 0x1f1e7, 0x1f1e8, 0x62}, []int{0, 1, 3, 4, 5}},
	{[]rune{0x61, 0x1f1e6, 0x1f1e7, 0x200d, 0x1f1e8, 0x62}, []int{0, 1, 4, 5, 6}},
	{[]rune{0x61, 0x1f1e6, 0x200d, 0x1f1e7, 0x1f1e8, 0x62}, []int{0, 1, 3, 5, 6}},
	{[]rune{0x61, 0x1f1e6, 0x1f1e7, 0x1f1e8, 0x1f1e9, 0x62}, []int{0, 1, 3, 5, 6}},
	{[]rune{0x61, 0x200d}, []int{0, 2}},
	{[]rune{0x61, 0x308, 0x62}, []int{0, 2, 3}},
	{[]rune{0x61, 0x903, 0x62}, []int{0, 2, 3}},
	{[]rune{0x61, 0x600, 0x62}, []int{0, 1, 3}},
	{[]rune{0x261d, 0x1f3fb, 0x261d}, []int{0, 2, 3}},
	{[]rune{0x1f466, 0x1f3fb}, []int{0, 2}},
	{[]rune{0x200d, 0x1f466, 0x1f3fb}, []int{0, 3}},
	{[]rune{0x200d, 0x2640}, []int{0, 2}},
	{[]rune{0x200d, 0x1f466}, []int{0, 2}},
	{[]rune{0x1f466, 0x1f466}, []int{0, 1, 2}},
}
