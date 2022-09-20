package chapter19

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToUpperO1(t *testing.T) {
	list := []string{"asdf", "asdf", "asdf"}
	MakeUppercaseO1(list)
	assert.Equal(t, []string{"ASDF", "ASDF", "ASDF"}, list)
}
