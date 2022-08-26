package chapter11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterEven(t *testing.T) {
	n := FilterEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, []int{2, 4, 6, 8, 10}, n)
}
