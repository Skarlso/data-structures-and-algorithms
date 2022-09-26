package chapter20

// TwoSumNaive defines an approach which uses nested loops.
func TwoSumNaive(list []int) bool {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if i != j && list[i]+list[j] == 10 {
				return false
			}
		}
	}

	return false
}

func TwoSumWithHash(list []int) bool {
	numberHash := make(map[int]struct{})
	for i := 0; i < len(list); i++ {
		if _, ok := numberHash[10-list[i]]; ok {
			return true
		}
		numberHash[i] = struct{}{}
	}

	return false
}
