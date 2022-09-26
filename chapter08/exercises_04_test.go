package chapter08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonDuplicated(t *testing.T) {
	s := "minimum"
	result := NonDuplicated(s)
	assert.Equal(t, "n", result)

}
