package chapter20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSwap(t *testing.T) {
	list1 := []int{5, 3, 3, 7}
	list2 := []int{4, 1, 1, 6}
	indexes := SumSwap(list1, list2)
	assert.Equal(t, []int{3, 0}, indexes)
}
