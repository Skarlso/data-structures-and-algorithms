package chapter11

var startingNumbers = []int{1, 3, 6, 10, 15, 21}

// TriangularNumbers returns the correct number in the sequence of Triangular Numbers.
func TriangularNumbers(n int) int {
	if n-1 < 6 {
		return startingNumbers[n-1]
	}

	return n + TriangularNumbers(n-1)
}
