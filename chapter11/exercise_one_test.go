package chapter11

import "testing"

func TestCalculateChars(t *testing.T) {
	count := CalculateChars([]string{"ab", "c", "def", "ghij"})
	if count != 10 {
		t.Fatalf("want: %d; got: %d", 10, count)
	}
}
