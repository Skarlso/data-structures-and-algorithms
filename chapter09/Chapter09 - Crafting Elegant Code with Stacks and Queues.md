# Crafting Elegant Code with Stacks and Queues

Until now, we mostly focused on performance and not the data structure. However, having a variety of data structures
allows you to create code that is easier and simpler to read.

Stacks and queues are really just arrays with added restrictions. But these restrictions allow for some very elegant
coding.

They both are used mainly for temporary data such as, mail queue, food order, print jobs etc.

## Stacks

Three constraints:
- Data can be inserted only at the end
- Data can be deleted only from the end
- Only the last element of a stack can be read

This is the definition of LIFO. Last In, First Out.

To make it visually more appealing, let's turn them up-side-down:

1 -> inserted last.
3
4
8 -> inserted first.

You can only remove or read `1` since that was inserted last and it's at the top of the stack. Like a tall building, you
can only remove something safely from the top.

### Abstract Data Types

Most languages don't offer a stack implementation. Lucky for us, it's super simple to write one. You either create a nice
API or you just stick to using a slice as a stack.

Example API:

```go
package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func NewStack[T any](data []T) *Stack[T] {
	return &Stack[T]{
		data: data,
	}
}

// Push pushes a value to the end of the stack.
func (s *Stack[T]) Push(i T) {
	s.data = append(s.data, i)
}

// Pop pops an item from the stack.
func (s *Stack[T]) Pop() T {
	var i T
	i, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return i
}

// Read reads the last item in the stack.
func (s *Stack[T]) Read() T {
	return s.data[len(s.data)-1]
}

func main() {
	stack := NewStack([]int{1, 2, 3, 4})
	fmt.Println("read: ", stack.Read())
	i := stack.Pop()
	fmt.Println(i, stack.Read()) // 4 3
}
```

Example Slice:

```go
func main() {
    stack := []int{}
    stack = append(stack, 1, 2, 3)
    var i int
    // Pop first
    i, stack = stack[len(stack)-1], stack[:len(stack)-1]
    fmt.Println(i, stack) // 3 [1 2]
}
```

Go does have a stack implementation: https://pkg.go.dev/github.com/golang-collections/collections/stack

The implementation uses an interface and `interface{}` as a type so it's sort of a generic stack but not a generic type
stack. That said, it doesn't use a slice in the background, but rather a linked list. Which makes sense for a Stack.

### Stacks in Action

Let's implement a language linter. Implementing a linter is complicated, but let's focus on a single aspect which is
detecting and validating opening and closing braces.

There are four types of syntax errors:

- `(var x int`
- `var x int)`
- `((var x int)`
- `(var x int}`

The logic is as follows:

- we encounter an opening character we push it into the stack
- if the next closing isn't an equal to the last opened one, error
- if the next one is an opening then closing and then the string ends, there is still an item left in the stack, error
- we encounter a closing character but there is nothing on the queue or the last opening is not of the same type, error

The whole code can be found here: [Linter](https://github.com/Skarlso/data-structures-and-algoritms/blob/main/chapter09/linter.go).

## Why the abstraction?

Since an array can do what this constrained API can do, why bother with using the abstraction?

- Prevent potential bugs in mis-indexing and mis-using the array
- A new mental model to work with
- The code become more elegant

## Queues

Queue is similar to Stack but has a different order. Queues are FIFO based. First In, First Out.

- Data can be inserted only at the end
- Data can be deleted only from the front
- Only the element at the front can be read

### Queue Implementation

```go
package main

import "fmt"

type Queue[T any] struct {
	data []T
}

func NewQueue[T any](data []T) *Queue[T] {
	return &Queue[T]{
		data: data,
	}
}

// Enqueue pushes a value to the end of the Queue.
func (s *Queue[T]) Enqueue(i T) {
	s.data = append(s.data, i)
}

// Dequeue pops an item from the front of the Queue.
func (s *Queue[T]) Dequeue() (T, error) {
	var i T
	if len(s.data) == 0 {
		return i, fmt.Errorf("empty")
	}
	i, s.data = s.data[0], s.data[1:]
	return i, nil
}

// Read reads the last item in the Queue.
func (s *Queue[T]) Read() (T, error) {
	if len(s.data) == 0 {
		var i T
		return i, fmt.Errorf("empty")
	}
	return s.data[0], nil
}

func main() {
	queue := NewQueue([]int{1, 2, 3, 4})
	v, _ := queue.Dequeue()
	fmt.Println("read: ", v) // 1
	queue.Enqueue(5)
	k, _ := queue.Read()
	fmt.Println(i, k) // 2
}
```

### Queue In Action

Printing jobs from various computers:

```go
package chapter09

import "fmt"

type PrintManager[T any] struct {
	queue *Queue[T]
}

func (p *PrintManager[T]) QueuePrintJob(job T) {
	p.queue.Enqueue(job)
}

func (p *PrintManager[T]) Run() {
	for !p.queue.Empty() {
		v, _ := p.queue.Dequeue()
		fmt.Println("now printing: ", v)
	}
}
```

Queues are great for processing async requests. It will ensure that all requests are processed in the order as they were
received.

## Examples

1. Call center which places calls on hold and then assigns them to the next available representative. What would you use?
Stack or Queue?

Answer: Queue, we have to process people in the order in which they made the call.

2. If you pushed numbers onto a stack in the following order: 1, 2, 3, 4, 5, 6, and then popped two items, which number
would you be able to read from the stack?

Answer: 4. You popped 5 and 6 from the end of the stack.

3. If you inserted numbers onto a queue in the following order: 1, 2, 3, 4, 5, 6, and then dequeued two items, which number
would you be able to read from the stack?

Answer: 3. We removed 1 and 2 from the front of the queue.

4. Write a function that uses a stack to reverse a string.

```go
// ReverseString simplified reverse not dealing with any unicode magic. Just
// plain reverse using a stack.
func ReverseString(s string) string {
	stack := []string{}
	for _, c := range s {
		stack = append(stack, string(c))
	}

	var (
		result string
		elem string
	)
	for len(stack) > 0 {
		elem, stack = stack[len(stack)-1], stack[:len(stack)-1]
		result += elem
	}
	return result
}
```
