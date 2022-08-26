package chapter11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindX(t *testing.T) {
	n := FindX("abcdefghijklmnopqrstuvwxyz")
	assert.Equal(t, 23, n)
}
