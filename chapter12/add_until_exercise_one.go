package chapter12

func AddUntil100(list []int) int {
	if len(list) == 0 {
		return 0
	}

	n := AddUntil100(list[1:])
	if list[0]+n > 100 {
		return n
	}
	return list[0] + n
}
