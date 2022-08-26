package chapter11

// FindX finds an X in a string and returns its index.
func FindX(s string) int {
	if s[0] == 'x' {
		return 0
	}
	return FindX(s[1:]) + 1
}
