package chapter11

// CalculateChars takes a list of strings and calculates each's length and adds them up.
// It uses recursion to do so.
func CalculateChars(list []string) int {
	if len(list) == 0 {
		return 0
	}
	return len(list[0]) + CalculateChars(list[1:])
}
