package chapter08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{2, 4, 9, 8}
	result := Intersection(list1, list2)
	assert.Equal(t, []int{2, 4}, result)
	list3 := []int{10, 11, 12}
	result = Intersection(list1, list3)
	assert.Nil(t, result)
}
