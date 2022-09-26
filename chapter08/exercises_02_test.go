package chapter08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDuplicate(t *testing.T) {
	list := []string{"a", "b", "c", "d", "c", "e", "f"}
	result := Duplicate(list)
	assert.Equal(t, "c", result)
}
