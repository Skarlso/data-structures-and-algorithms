package chapter08

func Intersection[T comparable](list1, list2 []T) []T {
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
	var result []T
	for _, v := range smallerList {
		if _, ok := biggerListHash[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
