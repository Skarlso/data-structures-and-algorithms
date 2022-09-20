package chapter19

import "sort"

func Reverse(list []int) {
	for i := len(list)/2 - 1; i >= 0; i-- {
		opp := len(list) - 1 - i
		list[i], list[opp] = list[opp], list[i]
	}
}

func HasDuplicateValue(list []int) bool {
	sort.Ints(list)

	for i := 0; i < len(list)-1; i++ {
		if list[i] == list[i+1] {
			return true
		}
	}
	return false
}

func Double1(list []int) []int {
	var newList []int

	for _, v := range list {
		newList = append(newList, v*2)
	}

	return newList
}

func Double2(list []int) []int {
	for i := 0; i < len(list); i++ {
		list[i] *= 2
	}
	return list
}

func Double3(list []int, index int) []int {
	if index >= len(list) {
		return list
	}

	list[index] *= 2

	return Double3(list, index+1)
}
