package chapter20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMissingLink(t *testing.T) {
	list := []int{2, 3, 0, 6, 1, 5}
	result := FindMissingLink(list)
	assert.Equal(t, 4, result)
	result = FindMissingLinkUsingSums(list)
	assert.Equal(t, 4, result)
	list = []int{8, 2, 3, 9, 4, 7, 5, 0, 6}
	result = FindMissingLink(list)
	assert.Equal(t, 1, result)

}

func BenchmarkFindMissingLink(b *testing.B) {
	list := []int{8, 2, 3, 9, 4, 7, 5, 0, 6}
	for i := 0; i < b.N; i++ {
		result := FindMissingLink(list)
		if result != 1 {
			b.Fail()
		}
	}
}

func BenchmarkFindMissingLinkWithSums(b *testing.B) {
	list := []int{8, 2, 3, 9, 4, 7, 5, 0, 6}
	for i := 0; i < b.N; i++ {
		result := FindMissingLinkUsingSums(list)
		if result != 1 {
			b.Fail()
		}
	}
}
