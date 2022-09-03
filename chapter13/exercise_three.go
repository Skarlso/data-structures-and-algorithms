package chapter13

import "sort"

// FindMax1 is O(n)
func FindMax1(list []int) int {
	max := list[0]
	for i := 0; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}

// FindMax2 is O(n^2)
func FindMax2(list []int) int {
	max := list[0]
	for i := 0; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}

// FindMax3 is O(N log N)
func FindMax3(list []int) int {
	sort.Ints(list)
	return list[len(list)-1]
}
