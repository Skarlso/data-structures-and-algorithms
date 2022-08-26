package chapter11

// Anagram takes a string and returns how many combinations are possible based on the characters.
func Anagram(s string) []string {
	// If we have a single letter left, return that
	if len(s) == 1 {
		return []string{string(s[0])}
	}

	collection := []string{}

	anagrams := Anagram(string(s[1:]))
	// Do the iteration over all the returned items
	for _, a := range anagrams {
		for i := 0; i < len(a)+1; i++ {
			collection = append(collection, a[:i]+string(s[0])+a[i:])
		}
	}

	return collection
}
