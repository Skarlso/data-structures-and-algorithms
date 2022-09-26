package chapter08

// Duplicate returns the first duplicate value.
func Duplicate[T comparable](list []T) T {
	hash := make(map[T]int)

	var result T
	for _, v := range list {
		hash[v]++
		if hash[v] > 1 {
			result = v
			break
		}
	}

	return result
}
