package chapter19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDouble3(t *testing.T) {
	list := []int{1, 2, 3, 4}
	list = Double3(list, 0)
	assert.Equal(t, []int{2, 4, 6, 8}, list)
}
