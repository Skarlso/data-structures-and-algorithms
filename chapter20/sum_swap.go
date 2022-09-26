package chapter20

// SumSwap determines if two lists can swap a single item to achieve equality in their sums. It returns the indexes of
// the two items to swap.
func SumSwap(list1, list2 []int) []int {
	// hashList stores the values of list1 with their indexes for a later swap
	hashList1 := make(map[int]int)
	var (
		sum1 int
		sum2 int
	)

	for i, v := range list1 {
		hashList1[v] = i
		sum1 += v
	}

	for _, v := range list2 {
		sum2 += v
	}

	shiftAmount := (sum1 - sum2) / 2

	for i, num := range list2 {
		if v, ok := hashList1[num+shiftAmount]; ok {
			return []int{v, i}
		}
	}
	return nil
}
