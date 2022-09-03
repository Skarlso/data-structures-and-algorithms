package chapter12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUntil(t *testing.T) {
	n := AddUntil100([]int{1, 2, 3, 4, 100, 5, 6})
	assert.Equal(t, 21, n)

}
