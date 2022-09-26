# Chapter 02 - Why Algorithms Matter

Definition _algorithm_: Set of instructions for completing a specific task.

Creating a bowl of cereals is also an algorithm of specific steps.

## Ordered Arrays

All items are ordered such as:

```go
n := []int{1, 2, 3, 4, 5}
```

How they are ordered is of no importance to this chapter.

Insertion requires a Search first to see where the insertion needs to happen. This takes N + 2 steps.

## Searching an Ordered Array

We must look at each element, but we have an advantage. We can stop early if we encounter a number that is higher than
our current number.

```go
func linearSearch(list []int, value int) int {
    for i, v := range list {
        if v == value {
            return i
        }
        if v > value {
            break
        }
    }
    return -1
}
```

The binary search is much more efficient. We can only use binary search with an ordered array. It halves the array as it
searches. It performs the following steps:
- Choose the middle of the array
- Check the value if it's greater or smaller
- If smaller, halve the right side, if larger, half the left side and repeat

```go
func BinarySearch[T constraints.Ordered](list []T, item T) (T, error) {
	var result T
	head := 0
	tail := len(list) - 1
	for head <= tail {
		middle := (head + tail)/2
		guess := list[middle]
		if guess == item {
			return guess, nil
		}
		if guess > item {
			tail = middle - 1
		} else {
			head = middle + 1
		}
	}
	return result, errors.New("not found")
}
```

## Binary search vs Linear Search

Linear Search: 100 steps
Binary Search: 7 steps

Each time you double the size, binary search increases by 1.

Ordered lists are better than none-ordered lists in some regards. Searching is more efficient, but insertion takes longer.
You have to decide which one you need more based on your application. If you do a lot of inserts, don't use an ordered
list. If you have more reads, an ordered list will improve your performance for such cases.

## Exercises

1. How many steps would it take to perform a linear search for the number 8 in the ordered list [2,4,6,8,10,12,13]?
    4 because we look at each element.
2. How many steps would Binary Search take?
    1 because we halve the array and choose the left side element first for comparison.
3. What is the maximum number of steps it would take to perform a binary search on an array of size 100.000?
    log2(100000) - 16.61 -> 16 times. The book says how many times does it take to halve 100000 until we get to 1. That's
    log2(X).