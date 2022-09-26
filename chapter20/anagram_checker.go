package chapter20

func AreAnagrams(first, second string) bool {
	secondList := make([]byte, len(second))
	for i := range second {
		secondList[i] = second[i]
	}

	for i := range first {
		// our second list ran out of letters.
		if len(secondList) == 0 {
			return false
		}

		for j := 0; j < len(secondList); j++ {
			if first[i] == secondList[j] {
				secondList = append(secondList[:j], secondList[j+1:]...)
				break
			}
		}
	}

	return len(secondList) == 0
}

func AreAnagramsOnSteroid(first, second string) bool {
	firstHash := make(map[byte]int)
	for i := range first {
		firstHash[first[i]]++
	}
	secondHash := make(map[byte]int)
	for i := range second {
		secondHash[second[i]]++
	}

	for k, v1 := range firstHash {
		v2, ok := secondHash[k]
		if !ok || v1 != v2 {
			return false
		}
	}

	return true
}
