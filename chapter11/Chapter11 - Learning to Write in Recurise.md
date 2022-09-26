# Learning to write in recursive

Knowing how to read, or even knowing how to understand recursion is not enough to be able to write recursive functions.
It takes practice and time to achieve such skill as it requires a different way of thinking about a problem.
Disregarding the performance of recursion for now, let's talk about some ways to help you better navigate this new land.

## Recursive Category: Repeatedly Execute
The book discusses several categories for recursions that make up most cases. The first one, and perhaps the
easiest of them all, is repeatedly executing something.

The countdown example from the previous chapter fits into this category perfectly. It repeatedly executes printing a
number.

Usually, in these categories, there is a single call to the function at the end, with some value being decreased
or changed until finished. The folder searcher algorithm is also one of these types.

_Note_: If you see a function that has a function call at the end to itself, that's called a _tails recursion_.

### Recursive Trick: Passing Extra Parameters
We'll take an array and double each value for the following example. We aren't going to return a new array; we'll be
modifying the existing one.

That is also a repeated task. We start with the first number, then repeat that, then the next, until we reach the end.

Let's try writing this. We know our last call will be the function itself, so we can add that.

```go
func double(list []int) {
    double(list)
}
```

Next, add the doubling code. But which number? Let's try with the first:

```go
func double(list []int) {
    list[0] *= 2
    double(list)
}
```

But here comes the problem. We don't have a way to keep track of the `index`. Or the next number to modify. We can add
that in as a parameter to the recursive function!

```go
func double(list []int, index int) {
    list[index] *= 2
    double(list, index+1)
}
```

Now, we just add the _base case_ on how to stop or when to stop. We have our index, so it's easy to check whether we are
at the end of the array.

```go
func double(list []int, index int) {
    if index >= len(list) {
        return
    }
    list[index] *= 2
    double(list, index+1)
}
```

And we are done. You would call this by starting at whatever index you wish. For example: `double([]int{1,2,3,4}, 0)`.

## Recursive Category: Calculations

This category contains problems that require some calculation, like summing an array. Or find the greatest
element in an array.

What all of these have in common is that _it can make a calculation based on a subproblem of the problem at hand_.

For example, let's take the factorial problem.

```go
func factorial(n int) int {
    product := 1
    for i := 1; i <= n; i++ {
        product *= i
    }
    return product
}
```

Straightforward loop based implementation. We could approach this differently by considering a subproblem of the main
problem. A _subproblem_ is the same as the problem but applied to a smaller input.

Let's see.

Factorial is: 6 * 5 * 4 * 3 * 2 * 1. But it really uses the previous calculations result as the next calculation so in
essence it's _6 * factorial(5)_. _factorial(5)_ is a subset of the main problem which is _factorial(6)_.

In code, from the previous chapter this looks like this:

```go
func factorial(n int) int {
    if n == 1 {
        return 1
    }

    return n * factorial(n - 1)
}
```

In the last line, we return our calculation for the subset of the problem.

## Two approaches for Calculations

- _bottom up_
- _top down_

The bottom-up approach was the for loop for factorial. We can try implementing that using recursion. But it requires
several more parameters to keep track of where we are.

```go
func factorial(n, i, product int) int {
    if i > n {
        return product
    }
    return factorial(n, i + 1, product * i)
}
```

- `n`: is the number
- `i`: index to keep track of the current number and helps us to know when to stop
- `product`: is the end result

This approach is not very elegant though.

Now for the meat of the matter.

## Top-Down Recursion: A New Way of Thinking

The top-down approach is a mindset change. It allows a different way of thinking about a problem. Specifically, we want
to "kick the problem down the road".

The factorial problem's last line returns` n * factorial(n - 1)`. Technically, you are deferring the main
problem down the road to the next function call. You expect the function to know how to calculate it, so you pass it
along. It's a bit weird because we are writing the function here, but let's not focus on that for now.

But this is a great approach. We can solve the problem without knowing how to solve the problem. You can say something
like I'm deferring this problem to the subproblem.

## The Top-Down Thought Process

The book outlines the following steps that make it easier to think about these kinds of problems:

1. Imagine the function you're writing has already been implemented
2. Identify the subproblem of the problem
3. See what happens when you call the function on the subproblem and go from there

Let's do a couple of examples to solidify these ideas.

## Array Sum

Let's say we want to sum all numbers in an array. `[1, 2, 3, 4, 5] == 15`.

First, imagine the `sum` function has already been implemented. Next, identify the subproblem. That isn't easy. It takes
a lot of practice and sometimes hits and misses. In this case, the subproblem is the rest of the array. So
`[2, 3, 4, 5]`. Without the first element. Next, what happens when we call our function on the subproblem?
`sum([2, 3, 4, 5])`. The answer is we get `14`. That means finding the solution to the original problem is adding the
first item to the result of the sum of the subproblem.

Something like `return list[0] * sum(the rest of the array)`. And that's it! We found our solution except for the base
case.

```go
func sum(list []int) int {
    return list[0] + sum(list[1:])
}
```

This thinking helped us defer the actual problem-solving to a later stage. We assumed that `sum` just worked already.
Now, to identify the _base_ case_. What will be the end of this? Once there is only a single number left, we'll
return.

```go
func sum(list []int) int {
    if len(list) == 1 {
        return list[0]
    }
    return list[0] + sum(list[1:])
}
```

## String reversal

Let's try another example. The string reverse from the stacks problem.

First, identify the subproblem. Before reading on, try it yourself.

Again, the subproblem is a subset of the main problem. In this case, let's say we have a string like `abcde`. A subset
of this would be `bcde`. From here, we just throw the `a` at the end of the string if someone already wrote `reverse`
for us. Let's say `reverse` already exists than, the solution would be something like
`return reverse(rest of the string) + s[0]`. And again, as before, we stop when there is a single character left.

```go
func reverse(s string) string {
    if len(s) == 1 {
        return string(s[0])
    }
    return reverse(string(s[1:])) + string(s[0])
}
```

Done.

## Counting X

Let's do another one. Count the number of `X`s in a string.

Given a string like `axbcxdex` it would return 3.

What is the subproblem? The subset of the main problem `xbcxdex`. Let's assume counting has been done already.
Then we can call that function with `return countX(remainder of string) + 1 ( if the first character is an x )`.

If it's not `x` we don't add anything to it.

```go
func countX(s string) int {
    n := 0
    if s[0] == 'x' {
        n++
    }
    return n + countX(string(s[1:]))
}
```

But we aren't done yet. We need the _base case_ to stop the function. Which is, when there is only one character left.

```go
func countX(s string) int {
    if len(s) == 1 {
        if s[0] == 'x' {
            return 1
        }
        return 0
    }
    n := 0
    if s[0] == 'x' {
        n++
    }
    return n + countX(string(s[1:]))
}
```

We can simplify this a bit knowing that in Go s[1:] returns an empty string once there are no more characters left. And
we already perform the X check on the last character.

```go
func countX(s string) int {
    if len(s) == 0 {
        return 0
    }
    n := 0
    if s[0] == 'x' {
        n++
    }
    return n + countX(string(s[1:]))
}
```

## THe Staircase Problem

Finally, to prove the point why you need a mental shift in thinking, here is the Staircase Problem which goes like this:

"Let's say we have a staircase of N steps, and a person has the ability to climb one, two, or three steps at a time. How
many different possible "paths" can someone take to reach the top? Write a function that will calculate this for `N`
steps."

For example, for 5 staircases, there are many possible ways, but let's take three.

```
1, 1, 1, 1, 1
2, 1, 2
3, 2
```

With the top-down approach, thinking about this problem is surprisingly easy. For 11 steps, the subproblem becomes 10.
Let's assume we know how many combinations there are for 10, we just add 1 to it for the last step. This is flawed since
we know people can jump more than one step at a time. When you think about it a bit more, you'll notice, that when you
are going from 8 to 11 in one jump, you won't include all the ones. You only include 1 which is the 3 step approach.
Which means we need the sum of all the path to stairs 10, 9 and 8. Because there is no other combination, since we can't
jump from 7 to 11, we get that for N steps `numberOfPaths(n - 1) + numberOfPaths(n - 2) + numberOfPaths(n - 3)`.

Save for the _base case_ that is our function:

```go
func numberOfPaths(n int) int {
    return numberOfPaths(n - 1) + numberOfPaths(n - 2) + numberOfPaths(n - 3)
}
```

The base case is a bit tricky. We could hardcode the 3 base steps:

```go
func numberOfPaths(n int) int {
    if n >= 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    if n == 4 {
        return 3
    }
    return numberOfPaths(n - 1) + numberOfPaths(n - 2) + numberOfPaths(n - 3)
}
```

But this can be improved with a neat little trick. We'll rig the system. We know that we definitely want
`numberOfPaths(1)` to be 1. So we can start with that.

```go
    if n == 1 {
        return 1
    }
```

Now we cheat a bit with 2. `numberOfPaths(2)` should return 2. But we know that `numberOfPaths(1)` returns 1, and
`numberOfPaths(0)` returns 0 and `numberOfPaths(-1)` returns 0 as well. If we made `numberOfPaths(0)` return 1 instead
of 0, we happened to correctly calculate the right paths for number 2.

```go
    if n < 0 {
        return 0
    }
    if n == 1 || n == 0 {
        return 1
    }
```

Let's look at 3. It returns the sum of 2, 1, 0. We know that we want the result to be 4. That actually already works.
`numberOfPaths(2)` returns 2 and `numberOfPaths(1)` returns 1 AND `numberOfPaths(0)` also returns 1. That's
`2 + 1 + 1 = 4`!

And we're done:

```go
func numberOfPaths(n int) int {
    if n < 0 {
        return 0
    }
    if n == 1 || n == 0 {
        return 1
    }
    return numberOfPaths(n - 1) + numberOfPaths(n - 2) + numberOfPaths(n - 3)
}
```

This approach is elegant and a nice way of thinking about a problem.

## Anagram Generation

Now comes the big gun. A harder problem yet to figure out to practice thinking this way.

As a refresher. An anagram is reordering of the characters within a string.

Anagrams of `abc` are, `["abc", "acb", "bac", "bca", "cab", "cba"]`.

Let's say we want to display the anagrams for `abcd`. We could say the subproblem is `abc`. If we had a working anagram
function that returned all anagrams how could we use that to return all anagrams for `abcd`? Try thinking about this for
a moment.

( I would cycle through and put d in all places since we have all previous combinations. YES! This is what the book also
suggested. Nice. :) ).

The book suggests ( although there are many more options ) that, since we already have the permutations of the previous
round, we just go and put our remaining character into all possible places.

So you have `["abc", "acb", "bac", "bca", "cab", "cba"]`. You cycle through all of them and put `d` into every possible
location. So `abc` would become `["dabc", "adbc", "abdc", "abcd"]`.

Let's code that up.

```go
func Anagram(s string) []string {
	// If we have a single letter left, return that
	if len(s) == 1 {
		return []string{string(s[0])}
	}

	collection := []string{}

	anagrams := Anagram(string(s[1:]))
	// Do the iteration over all the returned items
	for _, a := range anagrams {
        // Without the len(a)+1 ( the + 1 ) we would miss inserting our character at the end of the string.
        // And thus break the whole loop and end up a couple anagrams short.
		for i := 0; i < len(a)+1; i++ {
			collection = append(collection, a[:i]+string(s[0])+a[i:])
		}
	}

	return collection
}
```

## The efficiency of Anagram Generation

As a last, let's look at what our efficiency looks like. Since we create permutations for all characters and each
character, this is a factorial pattern.

4 characters = 4 * 3 * 2 * 1 anagrams
5 characters = 5 * 4 * 3 * 2 * 1 anagrams
6 characters = 6 * 5 * 4 * 3 * 2 * 1 anagrams

This is a factorial pattern. It's depicted as `O(n!)`. This is the lowest possible efficiency. The slowest algorithm.
But we have no choice because we have to create these permutations.

## Exercises

1. Use recursion to write a function that accepts an array of strings and returns the total number of characters across
all strings. For example: `["ab", "c", "def", "ghij"]` would return 10.

Answer: This is a calculation. If there are no items in the list, we return 0; we return the length of the first item
`+` the recursion for the rest of the items.

```go
func CalculateChars(list []string) int {
	if len(list) == 0 {
		return 0
	}
	return len(list[0]) + CalculateChars(list[1:])
}
```

2. Use recursion to write a function that accepts an array of numbers and returns a new array containing just the even
numbers.

Answer: we divide here. Base case is if our list is empty. We return an empty list in that case. The recursion is that
if our first element is even, we create a list out of it and append the recursive call's result to that list. If our
first number is not even, we just return the recursion for the rest of the items.

```go
// FilterEven takes a list of numbers and returns only the ones that are Even.
func FilterEven(n []int) []int {
	if len(n) == 0 {
		return []int{}
	}

	if n[0]%2 == 0 {
		return append([]int{n[0]}, FilterEven(n[1:])...)
	}
	return FilterEven(n[1:])
}
```

3. There is a numerical sequence known as Triangular Numbers. The pattern begins as 1, 3, 6, 10, 15, 21 and continues
onward with the Nth number in the pattern, which is N plus the previous number. For example, the 7th number in the
sequence is 28, since it's 7 (which is N) plus 21 (the previous number). Write a function that returns the correct
number in the sequence.

Answer: This is much simpler than at first glance. We have to call the same thing with a lower number
and add it to the result of the last call. And for the base numbers, we just create a constant list with the numbers and
return them if n is below 7.

```go
var startingNumbers = []int{1, 3, 6, 10, 15, 21}

// TriangularNumbers returns the correct number in the sequence of Triangular Numbers.
func TriangularNumbers(n int) int {
	if n-1 < 6 {
		return startingNumbers[n-1]
	}

	return n + TriangularNumbers(n-1)
}
```

4. Use recursion to write a function that accepts a string and returns the first index that contains the character "x".
For example, "abcdefghijklmnopqrstuvwxyz" has an "x" at 23. To keep it simple; assume there is always an x.

Answer: keep track of the index; return it if the item at index equals "x". ( This feels a bit cheap since it feels like
it's a bottom-up approach. But I'll check what the book has to say. We need to keep track of the index so we can't just
divide the list. Because then we lose at what place we were before. )

```go
// FindX finds an X in a string and returns its index.
func FindX(s string, index int) int {
	if index >= len(s) {
		return -1
	}

	if s[index] == 'x' {
		return index
	}
	return FindX(s, index+1)
}
```

There is a lot better solution since we always expect an 'x' to be present:

```go
// FindX finds an X in a string and returns its index.
func FindX(s string) int {
	if s[0] == 'x' {
		return 0
	}
	return FindX(s[1:]) + 1
}
```

And the list not to be empty...

5. Unique Path problem. This is a well-known exercise. Find all unique paths in a matrix from start to finish.

Given a matrix such as:

| S   |     |     |     |     |     |
| --- | --- | --- | --- | --- | --- |
|     |     |     |     |     |     |
|     |     |     |     |     |     |
|     |     |     |     |     |     |
|     |     |     |     |     |     |
|     |     |     |     |     | F   |

Find the paths that can lead from S to F, only going down and to the right. It should calculate the _number_ of
shortest paths. How many are there? The function should accept several rows and columns. So NxM.

( I would use a queue for this and dfs. But since we write recursion... )

Answer: What is our subproblem? We have to reduce it to the very base case. If we have no columns or rows, the only
possible route is the number of possible steps to the right and to down.

This looks something like this:

| S   | .   | .   | .   | .   | .   |
| --- | --- | --- | --- | --- | --- |
| .   |     |     |     |     |     |
| .   |     |     |     |     |     |
| .   |     |     |     |     |     |
| .   |     |     |     |     |     |
| .   |     |     |     |     | F   |

And we do this with an ever-decreasing number of rows and columns.
It is expressed like this:

```go
return UniquePaths(rows - 1, cols) + UniquePaths(rows, cols - 1)
```

All we need in the base case is to add up the numbers.

```go
if rows == 1 || cols == 1 {
    return 1
}
return UniquePaths(rows - 1, cols) + UniquePaths(rows, cols - 1)
```

Put together:

```go
func UniquePath(rows, cols int) int {
	if rows == 1 || cols == 1 {
		return 1
	}
	return UniquePath(rows-1, cols) + UniquePath(rows, cols-1)
}
```
