package chapter08

func NonDuplicated(value string) string {
	hash := make(map[rune]int)
	for _, v := range value {
		hash[v]++
	}

	for _, v := range value {
		if hash[v] == 1 {
			return string(v)
		}
	}

	return ""
}
