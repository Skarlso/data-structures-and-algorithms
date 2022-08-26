package chapter11

// UniquePath calculates the number of unique paths from start to finish if
// start is located at the upper left and finish is located at the bottom right cell.
func UniquePath(rows, cols int) int {
	if rows == 1 || cols == 1 {
		return 1
	}
	return UniquePath(rows-1, cols) + UniquePath(rows, cols-1)
}
