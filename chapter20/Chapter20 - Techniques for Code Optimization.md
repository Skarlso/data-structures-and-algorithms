# Techniques for Code Optimization

These are some notes about how to optimize code, or better said, how to spot where, when, how and what to optimize.

## Prerequisite

Before we start Optimization, there is always a prerequisite that needs to be taken. That is, determine the
current Big O of your algorithm. You should always know what you're trying to optimize.

## Start here

Next, think about which could be the _best imaginable Big O_ for your current algorithm. Now, this is the hard part and
has some art to it.

The steps are as follows:

1. Determine the Big O category
2. Determine the best-imaginable Big O category
3. If these two don't align, optimize

## Stretch your imagination

Now, take the previous steps a notch further and think about what would be an _amazing_ Big O for your algorithm.
If someone came along and told you that it could run in `O(1)`, would you believe them?

## Magical Lookup

In many cases, what takes the longest time is some data lookup. If we could speed up those lookups and make the O(1)
constant time, most of our algorithms would become significantly faster.

Most of the time, this is done by using HashMaps. But it's not always trivial what the key could be! Sometimes it's a
complex data structure.

Consider the following data:

```go
var (
	authors = []map[string]any{
		{
			"id":   1,
			"name": "Virginia Woolf",
		},
		{
			"id":   2,
			"name": "Leo Tolstoy",
		},
		{
			"id":   3,
			"name": "Dr. Seuss",
		},
		{
			"id":   4,
			"name": "J. R. R. Tolkien",
		},
		{
			"id":   5,
			"name": "Mark Twain",
		},
	}
	books = []map[string]any{
		{
			"authorID": 3,
			"name":     "Hop on Pop",
		},
		{
			"authorID": 1,
			"name":     "Mrs. Dalloway",
		},
		{
			"authorID": 4,
			"name":     "The Fellowship of the Ring",
		},
		{
			"authorID": 1,
			"name":     "To the Lighthouse",
		},
		{
			"authorID": 2,
			"name":     "Anna Karenina",
		},
		{
			"authorID": 5,
			"name":     "The Adventures of Tom Sawyer",
		},
		// you get the idea...
	}
)
```

Let's disregard how ineffective it is to store something like this.

A naive way of connecting these books with their respective authors is using nested loops. Now, nested loops have their
time and place, but usually, they are a good indicator that the logic could be optimized.

```go
func ConnectBooksWithAuthors() []map[string]string {
	booksWithAuthors := make([]map[string]string, 0)
	for _, book := range books {
		for _, author := range authors {
			if book["authorID"] == author["id"] {
				booksWithAuthors = append(booksWithAuthors, map[string]string{
					"title":  book["title"].(string),
					"author": author["name"].(string),
				})
			}
		}
	}
	return booksWithAuthors
}
```

This algorithm has O(N * M) complexity. Let's see if we can make it better. Ask yourself, "if I could find the desired
information in O(1) time, could I make my algorithm run faster?".

## Bringing in the Extra Data Structure

Here's one possibility to bring in a HashMap as an extra data structure to help speed up the algorithm.

```go
func ConnectBooksWithAuthorsWithMaps() []map[string]string {
	booksWithAuthors := make([]map[string]string, 0)
	authorHash := make(map[string]string)
	for _, author := range authors {
		authorHash[author["id"].(string)] = author["name"].(string)
	}

	for _, book := range books {
		booksWithAuthors = append(booksWithAuthors, map[string]string{
			"title":  book["title"].(string),
			"author": authorHash[book["authorID"].(string)],
		})
	}

	return booksWithAuthors
}
```

This new function has a total time complexity of `O(N + M)`. However, there is a caveat to it. It does now take up
`O(M)` space because we store all authors again in a different structure.

The book tries to convey the point of constantly imagining that you can do O(1) lookups.

## The Two-Sum problem

The task is to write a function that accepts an array of numbers and returns true or false depending on whether there
are any two numbers in the array that add up t 10.

For example, consider the following list of numbers: `[2, 0 4, 1, 7, 9]`.

Again, the naive approach compares all numbers with all other numbers.

```go
// TwoSumNaive defines an approach which uses nested loops.
func TwoSumNaive(list []int) bool {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if i != j && list[i]+list[j] == 10 {
				return false
			}
		}
	}

	return false
}
```

It looks like we have no choice but to compare all numbers with all other numbers. But what if you imagine IS an O(n)
approach for this.

Let's ask ourselves, _is there an O(1) lookup of desired information that could help me make this_
_algorithm faster_?

Sometimes, it helps to walk through the code and ask the question stated above for each line. As we go through the list,
let's first take the number `2`. What information could we look up in O(1) time that could help determine if the sum is
`10` or not? We could look for the number `8`. If we could get `8` in O(1) time and see if `8` is in the list, we could
significantly improve the algorithm.

```go
func TwoSumWithHash(list []int) bool {
	numberHash := make(map[int]struct{})
	for i := 0; i < len(list); i++ {
		if _, ok := numberHash[10-list[i]]; ok {
			return true
		}
		numberHash[i] = struct{}{}
	}

	return false
}
```

This approach works splendidly. It does take O(n) extra space, but the runtime improvement is significant enough to
ignore the additional memory.

## Recognizing Patterns

One of the most helpful strategies in solving these problems is trying to spot patterns. Often the discovery
of a pattern will lead to the solution.

### The Coin Game

The coin game consists of two players playing with some coins. The rule is simple. The player who removes the last coin
from the pile _loses_. During each turn, you can take one or two coins.

Let's start by analyzing a few games. If there is only a single coin, the one who starts loses. If there are two, the
other player will lose. If there are three, you could take two, so the other player loses. The interesting scenario is
four coins. If you take one, the other player takes two, and you lose. If you take two, the other player will take one
and you lose.

There are a few ways to solve this. A naive solution is using top-down recursion.

```go
func GameWinner(numberOfCoins int, currentPlayer string) string {
	var nextPlayer string
	if numberOfCoins <= 0 {
		return currentPlayer
	}
	if currentPlayer == "you" {
		nextPlayer = "them"
	} else {
		nextPlayer = "you"
	}

	if GameWinner(numberOfCoins-1, nextPlayer) == currentPlayer || GameWinner(numberOfCoins-2, nextPlayer) == currentPlayer {
		return currentPlayer
	}

	return nextPlayer
}
```

This solution has the complexity of O(2^N), which is really bad. Now, imagine someone telling you there is a better
way and that this can be solved in O(1).

### Generating Examples

The best way to spot patterns is through creating samples and analyzing the results. Let's look at a test case:

```go
func TestCoinGame(t *testing.T) {
	winner := GameWinner(1, "you")
	assert.Equal(t, "them", winner)
	winner = GameWinner(2, "you")
	assert.Equal(t, "you", winner)
	winner = GameWinner(3, "you")
	assert.Equal(t, "you", winner)
	winner = GameWinner(4, "you")
	assert.Equal(t, "them", winner)
	winner = GameWinner(5, "you")
	assert.Equal(t, "you", winner)
	winner = GameWinner(6, "you")
	assert.Equal(t, "you", winner)
	winner = GameWinner(7, "you")
	assert.Equal(t, "them", winner)
	winner = GameWinner(8, "you")
	assert.Equal(t, "you", winner)
}
```

Can you spot the pattern? It's `you, them, you, you, them, you, you, them, you`.
We see that each "them" ends up on a number divisible by three. This makes it easy to calculate the solution immediately.

```go
func GameWinnerWithMath(numberOfCoins int) string {
	if numberOfCoins%3 == 0 {
		return "them"
	}
	return "you"
}
```

Boom! Done.

## The Sum Swap Problem

Here is an example of using pattern recognition and fast lookups together.

Here is the problem: Take two lists. Sum them up. And see if you can exchange a single item from each list to have the
sum be equal in both lists.

For example:

```go
list1 := []int{5, 3, 2, 9, 1} // sum: 20
list2 := []int{1, 12, 5} // sum: 18
```

You could exchange the `1` from `list2` with the `2` from `list1`, and the sum would come together being `19`.

Again, it seems that we have no choice but to check all numbers against each number, which would be an O(M * N) approach
( because the lists have different sizes ).

To keep things simple, we won't perform the actual swap. We'll only check if it is possible.

Can we do better than O(M * N)? Let's see if we can spot a pattern with a few more examples.

| Before Swap                                                       | After Swap                                         |
| ----------------------------------------------------------------- | -------------------------------------------------- |
| list1 = [5, 3, 3, 7] sum = 18<br /> list2 = [4, 1, 1, 6] sum = 12 | [5, 3, 3, 4] sum = 15 <br /> [7, 1, 1, 6] sum = 15 |
| list1 = [1, 2, 3, 4, 5] sum = 15<br /> list2 = [6, 7, 8] sum = 21 | [1, 2, 6, 4, 5] sum = 18 <br /> [3, 7, 8] sum = 18 |
| list1 = [10, 15, 20] sum = 45<br /> list2 = [5, 30] sum = 35      | [5, 15, 20] sum = 40 <br /> [10, 30] sum = 40      |

We can spot three patterns in these examples:

- to achieve equality, the larger sum array needs to trade a larger number with the smaller sum array's smaller number
- the numbers change by the same amount; while list one's sum changes from 18 to 15 `-3`, the other is `+3`
- and lastly, the swap causes the two arrays to meet in the middle

We are looking for counterparts again. We know what the sum will/should be, so we have to check if the
other list contains our counterpart or not. And there is, where we can use fast lookups. How do we express this in code?

```go
shiftAmount := (sum1 - sum2) / 2
```

`shiftAmount` is the amount that we will have to look for. Just like in [Two Sum](#the-two-sum-problem) problem.

```go
// SumSwap determines if two lists can swap a single item to achieve equality in their sums. It returns the indexes of
// the two items to swap.
func SumSwap(list1, list2 []int) []int {
	// hashList stores the values of list1 with their indexes for a later swap
	hashList1 := make(map[int]int)
	var (
		sum1 int
		sum2 int
	)

	for i, v := range list1 {
		hashList1[v] = i
		sum1 += v
	}

	for _, v := range list2 {
		sum2 += v
	}

	shiftAmount := (sum1 - sum2) / 2

	for i, num := range list2 {
		if v, ok := hashList1[num+shiftAmount]; ok {
			return []int{v, i}
		}
	}
	return nil
}
```

This approach changes our complexity from O(M * N) to O(M + N) but also includes an O(M) additional memory space now.

## Greedy Algorithm

A greedy algorithm is an algorithm that, for each step, chooses the best possible option at that moment in time.

### Array Max

I bet you didn't guess that maxing a list of numbers is a greedy algorithm.

```go
func Max(list []int) int {
    max := list[0] // greedy choice!
    for _, v := range list {
        if v > max {
            max = v
        }
    }

    return max
}
```

It assumes that the first number is the maximum number. That's a greedy assumption.

### Largest Subsection Sum

The problem is as follows: Write a function that accepts a list of numbers and returns the largest sum that can be
computed from any "subsection" of the list. The "subsection" MUST be continuous.

For example:

```go
int{3, -4, 4, -3, 5, -9}
```

A naive approach would generate a check for each subsection. But there are N^2 / 2 subsections, so this would be an
O(N^2) approach. Let's see if we can imagine doing better than that. We could try to use a greedy approach.

Taking the previous list, a greedy approach would get the `3` and declare it as max. The current sum would be `3` as
well. Continuing, it would encounter `-4`. `3 - 4 = -1`, which is not greater than our max sum, so we continue. Next, it
gets to `4`. Our current sum becomes `3` and our greatest sum remains `3`. Next, `-3`. Current sum is `0`, max is `3`.
Next, `5`, current sum is `5`, which is greater than `3`, so we set it as max. `-9` is disregarded again since that drops
current max below the greatest max.

Now, our answer is `5`. But that's wrong because the max is `6`, using `4, -3, 5`. Our greedy approach didn't work out
ultimately this bit. But we _almost_ got it. Let's try tweaking it a bit and maybe find a pattern that would help us
with a couple of examples.


```go
int{1, 1, 0, -3, 5} // 5, [5]
int{5, -2, 3, -8, 4} // 6, [5, -2, 3]
int{2, -3, 1, 2, -1} // 3 [1, 2]
int{5, -8, 2, 1, 0} // 5 [5]
```

We can see the following pattern. Every time we encounter a negative value, our streak is broken. But that's not
completely true. Since there are sums with negative values in them. Further checking this fact, we detect that it breaks
the streak only if it falls below `0`. Now, all we have to do is write that down in code.

```go
// LargestSubsection returns the largest sum a continuous subsection of a list could produce.
func LargestSubsection(list []int) int {
	var (
		currentSum  int
		greatestSum int
	)

	for _, v := range list {
		if currentSum+v < 0 {
			currentSum = 0
		} else {
			currentSum += v

			if currentSum > greatestSum {
				greatestSum = currentSum
			}
		}
	}

	return greatestSum
}
```

While discovering the pattern helped us in solving this one, we only came by it by using the greedy approach first.

### Greedy Stock Prediction

One last greedy example. Let's say we are in the business of predicting stock value. The algorithm should look for a
positive trend in a given stock.

Given a list of stock prices, let's see if ANY three numbers would generate an upward trend.

[22, 25, 21, (18), (19.6), 17, 16, (20.5)]

To clarify, going from left to right, there are three prices for which the "right hand" price is greater than the
"middle" price, which in turn is greater than the "left hand" price.

A naive approach would be to use three nested loops. That is not the right approach. Thinking about what would be the
best is using the greedy approach. Just grab whatever is the best at any given time. It would be nice to find the lowest
of our trio and then keep collecting the followers ( the middle and highest item ) in that loop.

First, we'll assume that the first price on the list is the lowest. For the middle, we'll designate a price guaranteed
to be higher than anything in the list. That would be a positive max `int` value.

Then, our single pass would look like this:

1. If the current price is lower than the lowest price we've encountered so far, this price becomes the new lowest
2. If the current price is higher than the lowest price but lower than the middle price, this price becomes the middle
3. If the current price is higher than the lowest and the middle price, we found our upward trend!

Let's look at some code to solidify this idea:

```go
// IncreasingTriplet finds three items in the list can result in an upwards trend.
func IncreasingTriplet(list []int) bool {
	if len(list) == 0 {
		return false
	}
	lowest := list[0]
	middle := math.MaxInt
	for _, v := range list {
		if v < lowest {
			lowest = v
		} else if v <= middle {
			middle = v
		} else {
			return true
		}
	}

	return false
}
```

Super easy. Barely an inconvenience. Again, our greed resulted in the right approach. ( I wonder if there is a lesson
to be learned hidden in this... ).


## Change the Data Structure

If we can't find a better way to solve a situation, try looking at how the data is stored. Changing the data
structure in which the sample data is stored can help figure out a better solution.

### The Anagram Checker

Let's say we want to know if a string is an anagram of another string. We could generate all anagrams and compare, but
all we need here is if all the characters of string 1 are part of string 2. If we delete characters from string 2 that
we encounter in string 1, we should end with an empty string 2.

Sorting and comparing them would take O(nLogN) time, but we can do better with the second approach, which would only be
O(n)!

```go
func AreAnagrams(first, second string) bool {
	secondList := make([]byte, len(second))
	for i := range second {
		secondList[i] = second[i]
	}

	for i := range first {
		// our second list ran out of letters.
		if len(secondList) == 0 {
			return false
		}

		for j := 0; j < len(secondList); j++ {
			if first[i] == secondList[j] {
                // Delete the item from the list and stop there, since we don't want to delete
                // any other characters that might equal.
				secondList = append(secondList[:j], secondList[j+1:]...)
				break
			}
		}
	}

	return len(secondList) == 0
}
```

We could sort the two strings and compare them. This would take O(nLogM + nLogN) time. Which is better than O(M * N).
But we can go even further. We are aiming for O(M + N). This is where using an alternative data structure can help.

We can use our good friend the hash map. Store the characters as keys and store how many of these characters there are
for the value.

So `balloon` would something like this: `{"b": 1, "a": 1, "l": 2, "o": 2, "n": 1}`. There is a bit of data loss here
because we don't know the order of the characters, but that is unimportant for the result!

Once we convert both strings into hash tables, we can compare them, and if they are equal, we have a winner.

```go
func AreAnagramsOnSteroid(first, second string) bool {
	firstHash := make(map[byte]int)
	for i := range first {
		firstHash[first[i]]++
	}
	secondHash := make(map[byte]int)
	for i := range second {
		secondHash[second[i]]++
	}

	for k, v1 := range firstHash {
		v2, ok := secondHash[k]
		if !ok || v1 != v2 {
			return false
		}
	}

	return true
}
```

This works out rather nicely.

### Group Sorting

For our last example, consider the following: We have an array that contains different values, and we want to reorder
the data so that the same values are grouped. The order of the group doesn't matter.

For example:

```
["a", "c", "d", "b", "b", "c", "a", "d", "c", "b", "a", "d"]
```

The goal is to have this:

```
["c", "c", "c", "a", "a", "a", "d", "d", "d", "b", "b", "b"]
```

The order doesn't matter, so any group can be first or last.

We know that sorting takes O(nLogN) time, but can we imagine something better? We don't do a complete sort, we just do a
partial one. Can we exploit that?

_Note_: My initial thought was to count the letters and repeat them as we loop through a hash as many times as the value.
And indeed, that's what the book does as well.

```go
func GroupSort[T comparable](list []T) []T {
	hash := make(map[T]int)
	for _, v := range list {
		hash[v]++
	}

	var result []T
	for k, v := range hash {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}

	return result
}
```

Nice, and generic too.

## Exercises

1. You're working on software that analyzes sports players. Following are two arrays of players of different sports:

```go
basketballPlayers := []map[string]string{
    {
        "firstName": "Jill",
        "lastName": "Huang",
        "team": "Gators",
    },
    {
        "firstName": "Janko",
        "lastName": "Barton",
        "team": "Sharks",
    },
    {
        "firstName": "Wanda",
        "lastName": "Vakulsak",
        "team": "Sharks",
    },
    {
        "firstName": "Jill",
        "lastName": "Moloney",
        "team": "Sharks",
    },
    {
        "firstName": "Luuk",
        "lastName": "Watkins",
        "team": "Sharks",
    },
}

footballPlayers := []map[string]string{
    {
        "firstName": "Hanzla",
        "lastName": "Radosti",
        "team": "32ers",
    },
    {
        "firstName": "Tine",
        "lastName": "Watkins",
        "team": "Barleycorns",
    },
    {
        "firstName": "Alex",
        "lastName": "Patel",
        "team": "32ers",
    },
    {
        "firstName": "Jill",
        "lastName": "Huang",
        "team": "Barleycorns",
    },
    {
        "firstName": "Wanda",
        "lastName": "Vakulsak",
        "team": "Barleycorns",
    },
}
```

There are some players who play in both. Gather them and list their names like `["Jill Huang", "Wanda Vakulskas"]`.

Answer:

```go

func FindMultiSportsPlayers(sportOne, sportTwo []map[string]string) []string {
	// we just need the name
	namesHash := make(map[string]struct{})
	keyF := func(first, last string) string {
		return fmt.Sprintf("%s %s", first, last)
	}
	for _, n := range sportOne {
		namesHash[keyF(n["firstName"], n["lastName"])] = struct{}{}
	}

	var result []string
	for _, v := range sportTwo {
		key := keyF(v["firstName"], v["lastName"])
		if _, ok := namesHash[key]; ok {
			result = append(result, key)
		}
	}
	return result
}
```

2. Return a missing number from a continuous list of numbers.

```go
[]int{2, 3, 0, 6, 1, 5} // missing 4
[]int{8, 2, 3, 9, 4, 7, 5, 0, 6} // missing 1
```

Your task: find an O(N) algorithm. First thought, find min, max, and re-create the list. convert list ot map and look if
each number is in there.

Answer:

```go
func FindMissingLink(list []int) int {
	// first thought, find min, max, and re-create the list. convert list ot map and look if each number is in there.
	hashNumbers := make(map[int]struct{})
	min, max := list[0], list[0]
	for _, v := range list {
		hashNumbers[v] = struct{}{}
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}

	for i := min; i <= max; i++ {
		if _, ok := hashNumbers[i]; !ok {
			return i
		}
	}

	return -1
}
```

What does the book have to say? The book went a different approach. It calculates the sum of a list which would include
all numbers and the sum of the given list. Then subtract the results and the end is the result. Interestingly enough, this
would always return 0 even if the missing number IS 0. This doesn't cover the case where no number is missing. But I
guess that wasn't part of the problem.

```go
func FindMissingLinkUsingSums(list []int) int {
	var sum int
	for _, v := range list {
		sum += v
	}

	var fullSum int
	for i := 1; i < len(list)+1; i++ {
		fullSum += i
	}
	return fullSum - sum
}
```

Benchmarks:

My solution:
```
goos: darwin
goarch: amd64
pkg: github.com/Skarlso/data-structures-and-algoritms/chapter20
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkFindMissingLink
BenchmarkFindMissingLink-12    	 3553544	       319.9 ns/op	     160 B/op	       1 allocs/op
PASS
ok  	github.com/Skarlso/data-structures-and-algoritms/chapter20	1.775s
```

Book solution:
```
goos: darwin
goarch: amd64
pkg: github.com/Skarlso/data-structures-and-algoritms/chapter20
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkFindMissingLinkWithSums
BenchmarkFindMissingLinkWithSums-12    	100000000	        11.16 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/Skarlso/data-structures-and-algoritms/chapter20	1.242s
```

To no surprise, the book's solution runs faster; thus, it wins.

3. You're working on some more stock-prediction software. The function you're writing accepts an array of predicted
prices for a particular stock over the course of time. For example:

```go
[]int{10, 7, 5, 8, 11, 2, 6}
```

This predicts the following: On day 1 the stock price will be 10, on day 2, 7. On day 3, 5... And so on.
Calculate the best price possible with a single buy and a single sell transaction.

## Conclusion and parting thoughts

And with that, we finished the book. Thank you for reading this far. This was quite the journey. I found the book to be
too easy in the beginning. But in retrospect, this is an excellent primer for beginners. The later chapters also
contained a lot of food for thought items which I enjoyed reading. All in all, thank you, Jay Wengrow, for a time well spent.
