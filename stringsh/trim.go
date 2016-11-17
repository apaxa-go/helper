package stringsh

func TrimBytes(s string, cut byte) string {
	from := 0
	for ; from < len(s); from++ {
		if s[from] != cut {
			break
		}
	}

	to := len(s) - 1
	for ; to > from; to-- {
		if s[to] != cut {
			break
		}
	}
	return s[from : to+1]
}

func TrimLeftBytes(s string, cut byte) string {
	from := 0
	for ; from < len(s); from++ {
		if s[from] != cut {
			break
		}
	}
	return s[from:]
}

func TrimRightBytes(s string, cut byte) string {
	to := len(s) - 1
	for ; to >= 0; to-- {
		if s[to] != cut {
			break
		}
	}
	return s[:to+1]
}
