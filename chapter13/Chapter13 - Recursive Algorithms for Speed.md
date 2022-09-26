# Recursive Algorithms for Speed

This chapter is all about `Quicksort` and `Quickselect`. It details them greatly through examples which I'll attempt to
summarize.
Quicksort is a high-speed sorting algorithm for the average scenario. For worst-case, it behaves just as Insertion
Sort and Selection Sort. But in the wild, the average case reigns. So we need something that will behave best in those
cases.
Quicksort has a complexity of O(nLogN), a new category. It's just after O(n) but below O(n^2). The book goes
into great lengths with step-by-step sentences and calculations as to why it's O(nLogN).

It's logN because it halves the array. And it's N times because it does that N times. That's it.

Quicksort depends on the method called partitioning. It's a prime example of divide and conquer-based recursions. It
merges two concepts. One is partitioning. And two, recursion. It does the partitioning recursively.

Let's take a closer look.

## Partitioning

The process goes as follows: Take a random value ( there are actually researches based on how to select the most optimal
pivot, but that's out of scope in here ) from the array called _pivot_ and then, everything greater than the pivot ends
up on the right side and anything smaller than the pivot ends up on the left side.

To see this in action, consider the following array:

0 5 2 1 6 3

We select a pivot and set the right and left pointers to the two sides of the pivot.

0 5 2 1 6 [3]
^       ^

Then we move each pointer towards each other.
On the right side: We stop if we find anything higher than the pivot.
On the left side: We stop if we find anything lower than the pivot.

After just two steps, the two pointers will point to 5 and 1.

0 5 2 1 6 [3]
  ^   ^

Now, we swap them with each other.

0 1 2 5 6 [3]

Once the swap is done, we continue this process until the two meet.

0 1 2 5 6 [3]
    ^ ^

5 is greater so left stops. 2 is smaller so right continues.

0 1 2 5 6 [3]
      ^
      ^

Both end up on 5. Now, lastly, we swap the pivot with the last value.

0 1 2 [3] 6 5

With this, the partitioning ends. Let's see an implementation of that.


The book offers this implementation:
```go
func partition(list []int, left, right int) int {

    pIndex := right
    // Set up the pivot
    pivot = list[pIndex]

    // Start immediately to the left of the pivot.
    right--

    for {
        // Move the left pointer to the right as long as it points to a value
        // that is less than the pivot:
        for list[left] < pivot {
            left++
        }

        // Move the right pointer to the left as long as it points to a value
        // that is greater than the pivot:
        for list[right] > pivot {
            right--
        }

        // We reached a point where we stopped moving both pointers.

        // Break if they met at a single number so we can swap the pivot.
        if left >= right {
            break
        }

        // Otherwise, we swap the two numbers.
        list[right], list[left] = list[left], list[right]

        // Move the left pointer off from the changed item so the next round can begin.
        left--
    }

    // Finally, swap the pivot with the selected single value.
    // At this point, right and left should be the same, so it doesn't matter
    // which we choose.
    list[left], list[pIndex] = list[pIndex], list[left]

    // Return the left pointer, this we'll use later in quicksort.
    return left
}
```

The Wikipedia implementation tells this:
```go
// Wikipedia implementation of quicksort.
func partition[T constraints.Ordered](list []T, low, high int) int {
	i := low - 1
	pivot := list[high]

	for j := low; j < high; j++ {
		if list[j] <= pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	i++
	list[i], list[high] = list[high], list[i]
	return i
}
```

I find that the Wikipedia implementation is a lot more effective.

Now, let's get to the recursion part.

Quicksort is a repeat of the following steps:

1. Partition
2. Subdivide the list to the right and left of the pivot and call step 1, 2 on each of the sides
3. If the list is empty or has one item, we stop

These steps are illustrated with an array in the book, which I'm going to leave out, because for me, it is
understandable and makes sense.

The book offers this implementation:

```go
func quicksort(list []int, left, right int) {
    if right - left <= 0 {
        return
    }

    pIndex := partition(list, left, right)

    // Call sort on the left side.
    quicksort(list, left, pIndex - 1)

    // Call sort on the right side.
    quicksort(list, pIndex + 1, right)
}
```

Wikipedia:
```go
func quicksort[T constraints.Ordered](list []T, low, high int) {
	if len(list) == 1 {
		return
	}
	if low >= high || low < 0 {
		return
	}

	p := partition(list, low, high)
	quicksort(list, low, p-1)
	quicksort(list, p+1, high)
}
```

These two are similar in nature and offer the same algorithm.

The book here dives deep into the Big O and efficiency of this algorithm. Worst case is O(n^2) average, as we've seen is
O(nlogn).

|                | Base Case  | Average Case | Worst Case |
| -------------- | ---------- | ------------ | ---------- |
| Insertion Sort | O(N)       | O(N^2)       | O(N^2)     |
| Quicksort      | O(N log N) | O( N log N)  | O(N^2)     |

This makes Quicksort an ideal choice for sorting.

## Quickselect

There is an interesting algorithm called Quickselect. It uses Quicksort and returns the Nth sorted item. So let's say
that you have a list like [1, 5, 6, 9, 2, 3]. And you want the 2nd sorted number, which in this case would be 2.

To get it, normally, you would sort the entire list. But quickselect ALSO uses partitioning to select one of the sides
of an array. Here, we have the opportunity to throw away the side that we no longer need.

For example, you know that the 2nd number is below the pivot, you can throw away everything that is above and not sort
it. That's a huge save on a large list.

Implementation:

```go
func quickselect(list []int, nthLowest, left, right) int {
    // base case, we found our number
    if right - left == 0 {
        return list[0]
    }

    // partition and grab the index of the pivot
    pIndex := partition(list, left, right)

    if nthLowest < pIndex {
        // To the left
        quickselect(list, nthLowest, left, pIndex - 1)
    } else if nthLowest > pIndex {
        // To the right
        quickselect(list, nthLowest, pIndex + 1, right)
    } else if nthLowest == pIndex {
        // We found our value
        return list[pIndex]
    }
}
```

## Sorting as a Key to Other Algorithms

Sorting can also be used in other algorithms. For example, consider the find duplicate value algorithm from the previous
chapter. It required O(n^2) to compare all values. But we can make it actually a lot better if we pre-sort the list with
quicksort. Chances are that we get an average case and then we get O(nlogn) speed. Since we just have to then go through
the list and stop when the next value is the same as the previous value.

## Exercises

1. Given an array of positive numbers, write a function that returns the greatest product of any three numbers. Use
sorting in a way that will make it O(nlogN). Otherwise, if you try looping it would be O(n^3).

Answer: any three number is a bit misleading here, I think. You either have the option to select any number or you
return the greatest product. I find this mutually exclusive. I'll assume this means just to return the greatest product.
In which case, we just sort the array, then return the product of the last three elements.
```go
func GreatestThree(list []int) int {
    sort.Ints(list)

    return list[len(list)-1] * list[len(list)-2] * list[len(list)-3]
}
```
And, indeed, that was the solution.

2. The following function finds the "missing number" from an array integers. That is, the array is expected to have all
integers from 0 up to the array's length, but one is missing. Here is an implementation that is O(N^2).

```go
func findMissingNumber(list []int) int {
    for i := 0; i < len(list); i++ {
        if !contains(list, i) {
            return i
        }
    }
    // All numbers are present
    return -1
}
```

Use sorting to write a new implementation of this function.

Answer:

```go
func findMissingNumber(list []int) int {
    sort.Ints(list)

    for i := 0; i < len(list); i++ {
        if list[i] != i {
            return i
        }
    }

    return -1
}
```

3. Write three different implementations of a function that finds the greatest number within an array. Write one
function that is O(N^2), one that is O(N log N), and one that is O(N).

Answer:
```go
func FindMax1(list []int) int {
	max := list[0]
	for i := 0; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}

// FindMax2 is O(n^2)
func FindMax2(list []int) int {
	for i := 0; i < len(list); i++ {
        greatest := false
        for j := 0; j < len(list); j++ {
            if list[i] > list[j] {
                greatest = true
            }
        }
        if greatest {
            return list[i]
        }
	}

	return -1
}

// FindMax3 is O(N log N)
func FindMax3(list []int) int {
	sort.Ints(list)

	return list[len(list)-1]
}
```