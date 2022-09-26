# Optimizing for Optimistic Scenarios

The worst-case scenario is definitely not the only thing to consider when designing and using algorithms.

For example, as we saw previously with Bubble Sort and Selection Sort, Selection Sort is twice as fast.

## Insertion Sort

- starting from index 1, take the element and remove it from the list (not literally, just save it's value in a temp val)
- we begin a shifting phase where we take each value to the left and compare it to the current value
- if the value is is greater we shift that value to the right
- with this the gap moves left
- we stop if end is found or a smaller value than our temp
- insert the temp value into the _current gap_

Example:

4 2 7 1 3
  ^

  2
4   7 1 3
->

<-2
  4 7 1 3

2 4 7 1 3

2 4 7 1 3
    ^

    7
2 4 v 1 3


2 4 7 1 3
      ^

      1
2 4 7   3
    ->

      1
2 4   7 3
  ->

      1
2   4 7 3
->

<-----1
v 2 4 7 3

1 2 4 7 3

- and then finish with the last one

For 1 we've seen the special case where it shifts all the numbers to the left until there are no more numbers that are
smaller than the currently selected number. Once that happens, it puts the currently selected number in the remaining
gap.

Code:

```go
func insertionSort(list []int) {
	for i := 1; i < len(list); i++ {
    	tmp := list[i]
    	position := i - 1

    	for position >= 0 {
      		if list[position] <= tmp {
        		break
      		}
      		list[position + 1] = list[position]
      		position--
    	}
    	list[position + 1] = tmp
  	}
}
```

## Efficiency of Insertion Sort

Four operations:
- remove
- compare
- shift
- insert

### Compare

Compare happens each time we compare a value to the left of the gap with the `tmp` value. Worst case it's all of them.

1 + 2 + 3 + .. + (N - 1) compares

If we examine this, it's apparent that we have ~O(n^2/2) comparisons.

### Shift

Occurs each time we move a value on cell to the right. Worst case there are as many shifts as there are compares.

N^2 / 2 compare
+ N^2 / 2 shifts

### Remove & Insert

Remove & Insert happens once / pass. There are always N - 1 pass-throughs, therefore, there are N - 1 removes.

So far, we have N^2 comparisons and shifts combined + N - 1 removals + N - 1 insertions.

N^2 + 2N - 2 steps.

We remove the constants. This leaves us with O(N^2 + N).

Another aspect of Big O is the _it only cares about the highest order of N_. If we would have something like:
O(N^3 + N^2 + N) we would simply choose O(N^3) because that is the highest order.

## The Average Case

We can see that Selection Sort _is_ faster than Insertion Sort. Why would you choose Insertion Sort then? Turns out that
there is a thing called the _average-case scenario_.

Best- and worst-case scenarios happen relatively infrequently. In the real world _average scenarios_ happen more often!

It rarely happens that something is sorted completely backwards or not sorted at all. Most likely, the values are all
over the place.

Let's have a look. Worst-case we compare and shift all values. Best-case we don't shift and make one comparison per
pass-through.

For the average scenario, compare and shift about _half_ the data. If it takes O(n^2) in worst case, it takes about
~O(n^2/2) for the average case.

So which is better? It depends on your use-case. If your data is all over the place, insertion sort will be faster than
selection sort. If you have no idea, on the average scale, both will equal.

## Example

When you need to get the intersection between two arrays.

A naive implementation would be something like this:

```go
func intersection(a, b []int) []int {
    var result []int
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(b); j++ {
            if a[i] == b[j] {
                result = append(result, a[i])
            }
        }
    }
    return result
}
```

If the arrays aren't equal in size this would be an O(n * m).

Is there any way to improve this?

Exit early.

_Side note_: I have no idea what all of this has to do with insertion sort or sorting in general. The provided solution
to "improve" this has nothing to do with sorting either.

```go
func intersection(a, b []int) []int {
    var result []int
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(b); j++ {
            if a[i] == b[j] {
                result = append(result, a[i])
                break
            }
        }
    }
    return result
}
```

Breaking early, because we don't have to check the other values once we determined that 9 is part of list 2. It also adds
a significant boost to the algorithm's runtime.

## Exercises

1. Use Big O Notation to describe the efficiency of an algorithm that takes 3N^2 + 2N + 1 steps
Answer: O(n^2)

2. Use Big O Notation to describe the efficiency of an algorithm that takes N + logN steps
Answer: O(n) we always take the higher N.

3. Check if array contains a pair of two numbers that add up to 10.

```go
func twoSum(list []int) bool {
    for i := 0; i < len(list); i++ {
        for j := 0; j < len(list); j++ {
            if i != j && list[i] + list[j] == 10 {
                return true
            }
        }
    }

    return false
}
```

What are the best-, average- and worst-case scenarios?
Worst is O(n^2) when there are no matches. Best case is: if the first two pairs return 10. Average is O(n^2/2).

4. Return true if `X` is present in the String.

```go
func containsX(s string) {
    foundX := false
    for _, c := range s {
        if c == 'X' {
            foundX = true
        }
    }
    return foundX
}
```

It's important to note that we don't return immediately once we find X but we keep on looping. Which means, this will
always be O(n). A simple optimization would be to return immediately.
