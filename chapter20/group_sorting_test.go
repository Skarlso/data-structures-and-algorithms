package chapter20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupSorting(t *testing.T) {
	list := []string{"a", "c", "d", "b", "b", "c", "a", "d", "c", "b", "a", "d"}
	result := GroupSort(list)
	// This ends up being this same list all the time. I would have thought that it wouldn't be since I'm iterating over
	// a map.
	assert.Equal(t, []string{"a", "a", "a", "c", "c", "c", "d", "d", "d", "b", "b", "b"}, result)
}
