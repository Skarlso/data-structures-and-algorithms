package chapter12

func Golomb(n int, m map[int]int) int {
	if n == 1 {
		return 1
	}

	if v, ok := m[n]; ok {
		return v
	}

	m[n] = 1 + Golomb(n-Golomb(Golomb(n-1, m), m), m)

	return m[n]
}
