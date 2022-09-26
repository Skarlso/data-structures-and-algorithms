package chapter08

func FindMissingLetter(list string) string {
	hash := make(map[rune]struct{})
	for _, v := range list {
		hash[v] = struct{}{}
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, a := range alphabet {
		if _, ok := hash[a]; !ok {
			return string(a)
		}
	}

	return ""
}
