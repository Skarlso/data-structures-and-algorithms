package chapter11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	result := Anagram("abc")
	assert.Equal(t, []string{"abc", "bac", "bca", "acb", "cab", "cba"}, result)
}
