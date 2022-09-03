package chapter14

import "fmt"

// Number covers all numeric types.
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Leaf is a node in a binary search tree.
type Leaf[T Number] struct {
	Left  *Leaf[T]
	Right *Leaf[T]
	Data  T
}

func (l *Leaf[T]) Insert(val T) {
	if val < l.Data {
		if l.Left == nil {
			l.Left = &Leaf[T]{Data: val}
		} else {
			l.Left.Insert(val)
		}
	} else if val > l.Data {
		if l.Right == nil {
			l.Right = &Leaf[T]{Data: val}
		} else {
			l.Right.Insert(val)
		}
	}
}

func Delete[T Number](node *Leaf[T], val T) *Leaf[T] {
	if node == nil {
		return nil
	}

	if val < node.Data {
		node.Left = Delete(node.Left, val)
		return node
	} else if val > node.Data {
		node.Right = Delete(node.Right, val)
		return node
	} else if val == node.Data {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		node.Right = lift(node.Right, node)
		return node
	}

	return nil
}

func lift[T Number](node *Leaf[T], toDelete *Leaf[T]) *Leaf[T] {
	if node.Left != nil {
		node.Left = lift(node.Left, toDelete)
		return node
	}
	toDelete.Data = node.Data
	return node.Right
}

func Search[T Number](l *Leaf[T], val T) *Leaf[T] {
	if l == nil || l.Data == val {
		return l
	}

	if val < l.Data {
		return Search(l.Left, val)
	}
	if val > l.Data {
		return Search(l.Right, val)
	}
	return nil
}

func Display[T Number](l *Leaf[T]) {
	if l == nil {
		return
	}
	if l.Left != nil {
		fmt.Printf("%v-->%v\n", l.Data, l.Left.Data)
		Display(l.Left)
	}
	if l.Right != nil {
		fmt.Printf("%v-->%v\n", l.Data, l.Right.Data)
		Display(l.Right)
	}
}

func NewBinarySearchTree[T Number](nums []T) *Leaf[T] {
	if len(nums) == 0 {
		return nil
	}
	first := &Leaf[T]{
		Data: nums[0],
	}

	for _, n := range nums {
		first.Insert(n)
	}

	return first
}
