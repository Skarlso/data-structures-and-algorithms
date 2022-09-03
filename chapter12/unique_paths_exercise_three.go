package chapter12

type pair struct {
	row int
	col int
}

func UniquePath(rows, columns int, m map[pair]int) int {
	p := pair{row: rows, col: columns}
	if rows == 1 || columns == 1 {
		return 1
	}

	if v, ok := m[p]; ok {
		return v
	}

	m[p] = UniquePath(rows-1, columns, m) + UniquePath(rows, columns-1, m)

	return m[p]
}
