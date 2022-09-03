package chapter14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTreeDisplay(t *testing.T) {
	tree := NewBinarySearchTree([]int{50, 25, 75, 10, 33, 4, 11, 30, 40, 56, 89, 52, 61, 82, 95})
	Display(tree)
}

func TestBinarySearchTreeSearch(t *testing.T) {
	tree := NewBinarySearchTree([]int{50, 25, 75, 10, 33, 4, 11, 30, 40, 56, 89, 52, 61, 82, 95})
	n := Search(tree, 10)
	assert.Equal(t, &Leaf[int]{
		Data:  10,
		Left:  &Leaf[int]{Data: 4},
		Right: &Leaf[int]{Data: 11},
	}, n)
}

func TestBinarySearchTreeDelete(t *testing.T) {
	tree := NewBinarySearchTree([]int{50, 25, 75, 10, 33, 4, 11, 30, 40, 56, 89, 52, 61, 82, 95})
	Delete(tree, 50)
	n := Search(tree, 56)
	assert.Equal(t, &Leaf[int]{
		Data:  56,
		Right: &Leaf[int]{Data: 61},
	}, n)
}
