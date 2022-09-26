package chapter20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSumHash(t *testing.T) {
	numbers := []int{2, 0, 4, 1, 7, 9}
	got := TwoSumWithHash(numbers)
	assert.True(t, got)
}
