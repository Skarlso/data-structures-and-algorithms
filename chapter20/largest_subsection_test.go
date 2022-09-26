package chapter20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestSubsection(t *testing.T) {
	list := []int{1, 1, 0, -3, 5}
	sum := LargestSubsection(list)
	assert.Equal(t, 5, sum)
}
