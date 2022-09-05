package chapter14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}
	tree := NewLinkedListFromSlice(items)
	d := tree.FirstNode
	assert.Equal(t, 1, d.Data)
	node := tree.Read(5)
	assert.NotNil(t, node)
	assert.Equal(t, 6, node.Data)
	trav := tree.Traverse()
	assert.Equal(t, items, trav)
	i := tree.Search(3)
	assert.Equal(t, 2, i)
}

func TestLinkedListInsert(t *testing.T) {
	tree := NewLinkedList(&Node[int]{
		Data: 1,
		Next: &Node[int]{
			Data: 2,
			Next: &Node[int]{
				Data: 3,
			},
		},
	})
	tree.Insert(3, 4)
	assert.Equal(t, &Node[int]{
		Data: 1,
		Next: &Node[int]{
			Data: 2,
			Next: &Node[int]{
				Data: 3,
				Next: &Node[int]{
					Data: 4,
				},
			},
		},
	}, tree.FirstNode)

	tree.Insert(0, 5)
	assert.Equal(t, &Node[int]{
		Data: 5,
		Next: &Node[int]{
			Data: 1,
			Next: &Node[int]{
				Data: 2,
				Next: &Node[int]{
					Data: 3,
					Next: &Node[int]{
						Data: 4,
					},
				},
			},
		},
	}, tree.FirstNode)

	tree.Insert(1, 6)
	assert.Equal(t, &Node[int]{
		Data: 5,
		Next: &Node[int]{
			Data: 6,
			Next: &Node[int]{
				Data: 1,
				Next: &Node[int]{
					Data: 2,
					Next: &Node[int]{
						Data: 3,
						Next: &Node[int]{
							Data: 4,
						},
					},
				},
			},
		},
	}, tree.FirstNode)
}
