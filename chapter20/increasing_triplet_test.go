package chapter20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncreasingTriplet(t *testing.T) {
	list := []int{3, 2, 4, 5, 6, 1}
	result := IncreasingTriplet(list)
	assert.True(t, result)
}
