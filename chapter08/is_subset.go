package chapter08

func IsSubset[T comparable](list1, list2 []T) bool {
	biggerListHash := make(map[T]struct{})
	var (
		biggerList  []T
		smallerList []T
	)
	if len(list1) > len(list2) {
		biggerList = list1
		smallerList = list2
	} else {
		biggerList = list2
		smallerList = list1
	}
	for _, v := range biggerList {
		biggerListHash[v] = struct{}{}
	}
	for _, v := range smallerList {
		if _, ok := biggerListHash[v]; !ok {
			return false
		}
	}
	return true
}
