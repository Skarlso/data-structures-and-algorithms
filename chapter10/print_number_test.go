package chapter10

import "testing"

func TestPrintNumber(t *testing.T) {
	slice := []any{
		1,
		2,
		3,
		[]any{
			4, 5, 6,
		},
		7, 8, 9,
	}
	PrintNumber(slice)
}
