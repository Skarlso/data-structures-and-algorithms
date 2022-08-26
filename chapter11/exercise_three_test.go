package chapter11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTriangularNumbers(t *testing.T) {
	n := TriangularNumbers(7)
	assert.Equal(t, 28, n)
	n = TriangularNumbers(8)
	assert.Equal(t, 36, n)
	n = TriangularNumbers(6)
	assert.Equal(t, 21, n)
}
