package chapter20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagramChecker(t *testing.T) {
	first := "secure"
	second := "rescue"
	result := AreAnagrams(first, second)
	assert.True(t, result)
	result = AreAnagramsOnSteroid(first, second)
	assert.True(t, result)
}
