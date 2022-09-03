package chapter12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGolomb(t *testing.T) {
	n := Golomb(6, map[int]int{})
	assert.Equal(t, 4, n)
}
