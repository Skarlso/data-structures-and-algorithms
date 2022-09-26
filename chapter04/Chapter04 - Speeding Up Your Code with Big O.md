# Speeding Up You Code with Big O

Use Big O to compare your solutions to other algorithms in the wild.

## Bubble Sort

This is a sorting algorithm. At its core, it performs swaps of values until no swaps are performed anymore.
On every pass-through it "Bubbles up" values to the right until every value is at the correct position.

Let's take a look:

1. Select two consecutive numbers
2 1 3 5
^ ^
2. Swap them if necessary
1 2 3 5
3. Move one tile
1 2 3 5
  ^ ^
4. Repeat step 1,2.
5. End if no swap was performed during a pass-through

Example:

4 2 7 1 3
^ ^
- swap
2 4 7 1 3
  ^ ^
- nothing
2 4 7 1 3
    ^ ^
- swap
2 4 1 7 3
      ^ ^
- swap
2 4 1 3 7
- we reached the end, start again
2 4 1 3 7
^ ^
- nothing
2 4 1 3 7
  ^ ^
- swap
2 1 4 3 7
    ^ ^
- swap
2 1 4 3 7
    ^ ^
- swap
2 1 3 4 7
      ^ ^
- nothing
- start again
2 1 3 4 7
^ ^
- swap
1 2 3 4 7
  ^ ^
- nothing
1 2 3 4 7
    ^ ^
- nothing
1 2 3 4 7
      ^ ^
- nothing
- we start again because there was a swap from 2->1
- there are no swaps in the last pass-through so we stop

Code:

```go
func BubbleSort(list []int) {
    unsortedUntilIndex := len(list) - 1
    sorted := false
    for !sorted {
        sorted = true
        // Because we know that at each pass-through the highest number gets
        // put into its rightmost place, we only check until the last index we
        // sorted.
        for i := 0; i < unsortedUntilIndex; i++ {
            if list[i] > list[i+1] {
                list[i], list[i+1] = list[i+1], list[i]
                sorted = false
            }
        }
        unsortedUntilIndex++
    }
}
```

## Efficiency of Bubble Sort

Operations: _Comparisons_, _Swaps_

First, calculate how many _comparisons_ there are.

For the above example:

4 + 3 + 2 + 1 = 10

That is:

(N - 1) + (N - 2) + (N - 3) + ... + 1 comparisons

_Swaps_:

Worst case we swap all. So 10 swaps. 10 comparisons + 10 swaps = 20 steps.

For a list with 10 values you get 45 swaps and 45 comparisons. That's 90 steps. For 20 values you get 380 steps.
The growth is exponential. (20^2 == 400 ~ 380).

Therefor, the Big O complexity of Bubble Sort is O(n^2). This is called a _quadratic time_. Most often, it's considered
an inferior solution, albeit, sometimes, unavoidable. But it's a good indicator that you probably need to re-think your
algorithm's design.

## A Quadratic Problem

Let's try and replace an O(n^2) with an O(n). Consider the following: You have to analyze ratings from people to products.
You are checking if an array contains a duplicate value or not. A trivial approach would be something like this:

```go
func hasDuplicateValue(list []int) bool {
    for i := 0; i < len(list); i++ {
        for j := 0; j < len(list); j++ {
            if i != j && list[i] == list[j] {
                return true
            }
        }
    }
    return false
}
```

This goes through all items and compares them one by one. This is clearly an O(n^2) solution. A good indicator is always
a nested loop. That should ring some alarm bells in your head.

How could we solve this in one pass? I give you a bit of time to think about it... done.

You can either use a `map` to store each elements and see if the value is greater than 1 or you can use another list and
store the values as indexes and if there is another such index, we increase that location. I don't like the second approach
because it potentially creates a VERY LARGE slice if we would want to do that. For example, if one of the values is a
large int. Also, it restricts the algorithm to only work with `int`s.

The map is a more elegant solution although it still will increase memory consumption the calculation boost is significant.

```go
// needs a comparable constraint because only values that are comparable can be map keys.
func hasDuplicateValue[T comparable](list []T) bool {
    m := make(map[T]int)
    for _, v := range list {
        m[v]++ // in Go, map values are initialized to the given types 0 value
        if m[v] > 1 {
            return true
        }
    }
    return false
}
```

This is much more elegant, and requires only an O(n) approach as we only check all values once. And the size of the map
is negligible.

## Exercises

1. Replace the question marks denoting how many steps the various things will take
| N Elements | O(n) | O(log N) | O(n^2) |
| ---------- | ---- | -------- | ------ |
| 100        | 100  | ?        | ?      |
| 2000       | ?    | ?        | ?      |

Solutions:
| N Elements | O(n) | O(log N) | O(n^2)    |
| ---------- | ---- | -------- | --------- |
| 100        | 100  | ~7       | 10000     |
| 2000       | 2000 | ~11      | 4.000.000 |

You literally just have to perform the function O(n) like f(x).

2. If we have an O(N^2) algorithm that processes an array and find that it takes 256 steps, what is the size of the array?

√n = 16.

3. Use Big O Notation to describe the time complexity of the following function. It find the greatest product of any
pair of two numbers within a given array.

```go
func greatestProduct(list []int) int {
    // ignoring all concerns of size... :D
    greatestSoFar := list[0] + list[1]
    for i, iv := range list {
        for j, jv := range list {
            if i != j && iv * jv > greatestSoFar {
                greatestSoFar = iv * jv
            }
        }
    }
    return greatestSoFar
}
```

Solution: O(n^2)

4. The following function finds the greatest single number but has an efficiency of O(n^2). Rewrite it so that it becomes
O(n).

```go
func greatestNumber(list []int) int {
    for _, v := range list {
        isVValTheGreatest := true
        for _, k := range list {
            if k > v {
                isVValTheGreatest = false
            }
        }
        if isVValTheGreatest {
            return v
        }
    }
}
```

Solution:
```go
func max(list []int) int {
    m := list[0]
    for _, v := range list {
        if v > m {
            m = v
        }
    }
    return m
}
```
