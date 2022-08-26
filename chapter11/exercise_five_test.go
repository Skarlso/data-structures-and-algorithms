package chapter11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePath(t *testing.T) {
	n := UniquePath(3, 7)
	assert.Equal(t, 28, n)
}
