package chapter20

func FindMissingLink(list []int) int {
	// first thought, find min, max, and re-create the list. convert list ot map and look if each number is in there.
	hashNumbers := make(map[int]struct{})
	min, max := list[0], list[0]
	for _, v := range list {
		hashNumbers[v] = struct{}{}
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}

	for i := min; i <= max; i++ {
		if _, ok := hashNumbers[i]; !ok {
			return i
		}
	}

	return -1
}

func FindMissingLinkUsingSums(list []int) int {
	var sum int
	for _, v := range list {
		sum += v
	}

	var fullSum int
	for i := 1; i < len(list)+1; i++ {
		fullSum += i
	}
	return fullSum - sum
}
