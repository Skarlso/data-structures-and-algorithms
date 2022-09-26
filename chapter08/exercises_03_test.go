package chapter08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMissingLetter(t *testing.T) {
	s := "the quick brown box jumps over the lazy dog"
	result := FindMissingLetter(s)
	assert.Equal(t, "f", result)

}
