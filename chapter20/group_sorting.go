package chapter20

func GroupSort[T comparable](list []T) []T {
	hash := make(map[T]int)
	for _, v := range list {
		hash[v]++
	}

	var result []T
	for k, v := range hash {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}

	return result
}
