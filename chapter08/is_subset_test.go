package chapter08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubset(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{1, 2, 3}
	result := IsSubset(list1, list2)
	assert.True(t, result)
	list3 := []int{10, 11, 12}
	result = IsSubset(list1, list3)
	assert.False(t, result)
}
