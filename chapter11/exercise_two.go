package chapter11

// FilterEven takes a list of numbers and returns only the ones that are Even.
func FilterEven(n []int) []int {
	if len(n) == 0 {
		return []int{}
	}

	if n[0]%2 == 0 {
		return append([]int{n[0]}, FilterEven(n[1:])...)
	}
	return FilterEven(n[1:])
}
