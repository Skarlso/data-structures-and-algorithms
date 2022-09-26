# Optimizing code with and without Big O

Big O is not the only tool we should rely on. There are cases when Big O describes two algorithms as being similar, yet
one of them is faster than the other.

## Selection Sort

### How does it work?

- we select the index of the lowest value in the list
- once done with the pass-through we do a swap with the value from where we began

### Example

4 2 7 1 3
^

- select the lowest element

4 2 7 1 3
      ^

- swap with the original

1 2 7 4 3

- repeat until sorted

### Code

```go
func selectionSort(list []int) {
    for i := 0; i < len(list); i++ {
        lowest := i
        for j := i + 1; j < len(list); j++ {
            if list[j] < list[lowest] {
                lowest = j
            }
        }

        if lowest != i {
            list[i], list[lowest] = list[lowest], list[i]
        }
    }
}
```

### The efficiency of Selection Sort

Two types of steps: _comparison_ and _swap_.

(N - 1) + (N - 2) + (N - 3) + ... + 1 comparison

Because the further we get the less we have to travel.

For swaps, we either swap or not per pass.

| N Elements | Max # of Steps in Bubble Sort | Max # of Steps in Selection Sort |
| ---------- | ----------------------------- | -------------------------------- |
| 5          | 20                            | 14 (10 comp + 4 swaps)           |
| 10         | 90                            | 54 (45 comp + 9 swaps)           |
| 20         | 380                           | 199 (180 comp + 19 swaps)        |
| 40         | 1560                          | 819 (780 comp + 39 swaps)        |
| 80         | 6329                          | 3239 (3160 comp + 79 swaps)      |

This makes Selection Sort _twice_ as fast.

## Ignoring Constants

Labeling this as O(n^2 / 2) would seem reasonable. But Big O ignores constants as considering the big picture, they
don't really matter when comparing such values like 2N vs N^2. The difference is so vast that not even 100N would make
a dent into the N^2's curve.

This is called _general categories_ for Big O Notation.

Which leaves us with a **"warning"**. The same category doesn't mean the same speed. It means further analysis is required!

## Significant Steps

There is another benefit to generalizing. By focusing on the significant steps in the loop, we can ignore the fluff
around and inside it mostly. We can concentrate on how many times the loop runs instead of what is the detail inside it.

## Exercise

1. Use Big O Notation to describe the time complexity of an algorithm that takes 4N + 16 steps.
Answer: O(n)
2. Use Big O Notation to describe the time complexity of an algorithm that takes 2N^2 steps.
Answer: O(n^2)

3. Use Big O Notation to describe the time complexity of the following function:

Sum of all numbers of an array after they have been doubled.

```go
import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func doubleThenSum[T Number](list []T) T {
	var doubles []T
	for _, v := range list {
		doubles = append(doubles, v*2)
	}
	var sum T
	for _, v := range doubles {
		sum += v
	}
	return sum
}
```

Answer: O(n) because although it loops twice it is in sequence. So it remains O(n).

4. Describe the following:

```go
func print(list []string) {
    for _, v := range list {
        fmt.Println("upcase: ", strings.ToUpper(v))
        fmt.Println("lower: ", strings.ToLower(v))
        fmt.Println("capitalized: ", strings.Title(v)) // I know, strings.Title is deprecated...
    }
}
```

This is tricky because toUpper and toLower both loop through each character. But that doesn't add to the overall
complexity of the main runner which is a simple loop through all the items. You don't loop through all items again.
So it's considered a simple step. Complexity is O(n).

5. Iterate over a number and for each number whose _index_ is even it prints the sum of that number plus every number in
the array.

```go
func everyOther(list []int) {
    for i, v := range list {
        if i % 2 == 0 {
            for _, k := range list {
                fmt.Println(v + k)
            }
        }
    }
}
```

This is O(n^2) because it loops twice over the list of items. Even though it's technically O(n^2/2) because it does it
only for halve of the values.