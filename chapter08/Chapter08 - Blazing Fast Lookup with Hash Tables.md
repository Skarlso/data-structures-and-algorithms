# Blazing fast Lookup with Hash Tables

Consider the following scenario. An ordering program that gets prices based on names. If we stored the items in a
list, we would have difficulty gathering the needed costs.

For example: `[["french fries", 0.75], ["hamburger", 2.5], ["hot dog", 1.5]]`

However, a hash table makes the lookup an O(1) complexity.

```go
var menu = map[string]float64{
	"french fries": 0.75,
	"hamburger":    2.5,
	"hot dog":      1.5,
}
```

A lot easier.

## Hashing with Hash Functions

Consider the following mapping of characters to numbers:

A = 1
B = 2
C = 3
D = 4
E = 5

According to this:

ACE = 135
CAB = 312
DAB = 412
BAD = 214

A hash function is something, for example, that takes the sum of each number. Another option for hashing is using
multiplication. So BAD would become `8`. In reality, hashmap hash functions are much more complex, but multiplication is
enough for our purposes.

Note that with this hash function, DAB will also convert to 8. This will cause some issues, which we will discuss later.

## Building a Thesaurus for Fun and Profit, but mainly Profit

Write a thesaurus that returns a single synonym instead of every possible synonym.

Let's pair them up in a hash map.

```go
thesaurus := make(map[string]string)
```

In memory a hash looks something like this:
| 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  | 11  | 12  | 13  | 14  | 15  | 16  |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |

We left off `0` since our hash function won't have `0`.

Let's store some values into a hash. BAD => EVIL, CAB => TAXI, ACE => STAR.

BAD = 8
CAB = 6
ACE = 15

| 1   | 2   | 3   | 4   | 5   | 6    | 7   | 8    | 9   | 10  | 11  | 12  | 13  | 14  | 15   | 16  |
| --- | --- | --- | --- | --- | ---- | --- | ---- | --- | --- | --- | --- | --- | --- | ---- | --- |
|     |     |     |     |     | TAXI |     | EVIL |     |     |     |     |     |     | STAR |     |

## Hash Table Lookups

A lookup is simply using a key. `thesaurus["BAD"]` -> would return `EVIL`. How? The computer hashes the key the same way
it hashed it when it stored the value. Thus, it returns to the same place and returns the value. That is why it's so
fast and can be used to speed up various algorithms that require looking up something.

## One-Directional Lookups

Fast lookup only works when we know the key. If we only know the value, a hashmap will not help. We'll have to loop
through all the values to find the right one.


## Dealing with Collisions

What happens when we save DAB? DAB => PAT would calculate to 8 and would try to overwrite the existing value. That is
called a _collision_. There are ways around this problem.

One such approach is _separate chaining_. If multiple values exist for a single cell, it will put a list of values
into that place instead of a single one. Our value would become something like this:
| 1   | 2   | 3   | 4   | 5   | 6    | 7   | 8                         | 9   | 10  | 11  | 12  | 13  | 14  | 15   | 16  |
| --- | --- | --- | --- | --- | ---- | --- | ------------------------- | --- | --- | --- | --- | --- | --- | ---- | --- |
|     |     |     |     |     | TAXI |     | [[BAD, EVIL], [DAB, PAT]] |     |     |     |     |     |     | STAR |     |

It contains sub-arrays where the first value is the key; the second is the value.
A lookup will check if the result is an array; if yes, it will search through the items until it finds the correct value.

Since looping through an array requires O(N) complexity, a good hashing function must be used that minimizes the need to
create array values.

## Making an Efficient Hash Table

There are three critical properties of an efficient hash map:

1. How much data are we storing
2. How many cells are available in the hash map
3. Which hash function is used

A good hash function will distribute the values into each cell instead of partially filling it up.

### The Great Balancing Act

We could allocate a LOT of cells to accommodate each item in a list. However, that would not be very memory friendly.
So, we have to balance how much memory we allocate. For example, consider allocating a 1000 cells for a data of size 5.
That would be a lot of wasted space.

A good hash table will avoid most collisions and minimize the amount of memory used.

Computer scientists developed a rule of thumb: For every 7 data elements stored in a hash table, it should have
10 cells. If you are planning on 14 items, it should have 20 cells.

## Hash Tables for Organization

Hash tables can be used for various applications where the data is key/value. Or if we can construct a clear, unique key
for our data elements. Examples include shipping systems, ordering systems, keeping track of political party scores, etc.
In some cases, it can even simplify conditional logic.

For example, consider the following:
```go
func StatusCodeMeaning(code int) string {
    if code == 200 {
        return "OK"
    }
    if code == 301 {
        return "Moved Permanently"
    }
    if code == 401 {
        return "Unauthorized"
    }
    if code == 404 {
        return "Not Found"
    }
    if code == 500 {
        return "Internal Server Error"
    }
}
```

This could be a lot better if we had data like this:
```go
codes := map[int]string{
    200: "OK",
    ...
}
```

And then used this map to return the right value:

```go
func StatusCodeMeaning(code int) string {
    v, ok := StatusCodes["Unknown"]
    if !ok {
        return "Unknown"
    }

    return v
}
```

We can even have a list of map values that could group various attributes of an object.
```
[
    {"Name": "Fido", "Breed": "Pug"},
    {"Name": "Lady", "Breed": "Poodle"},
    {"Name": "Spot", "Breed": "Dalmatian"},
]
```

## Hash Tables for Speed

If your data isn't paired, you can construct your key. For example, consider the following list: [61, 30, 123, 34, 5, 9].
If you searched this array to locate 9, you would have to use an O(N) approach.

_Note_: this next section is weird. The book suggests converting the array to a map and doing a lookup which would then
be O(1). But to convert the list to a map, you'll have to go through all the items anyways.

## Array Subset

Determine if a given array is a subset of another array.

```
["a", "b", "c", "d", "e", "f"]
["b", "d", "f"]
```

We could use nested loops to determine the outcome, but this is a chapter about hash maps. :) So let's use one.

```go
func IsSubset[T comparable](list1, list2 []T) bool {
	biggerListHash := make(map[T]struct{})
	var (
		biggerList  []T
		smallerList []T
	)
	if len(list1) > len(list2) {
		biggerList = list1
		smallerList = list2
	} else {
		biggerList = list2
		smallerList = list1
	}
	for _, v := range biggerList {
		biggerListHash[v] = struct{}{}
	}
	for _, v := range smallerList {
		if _, ok := biggerListHash[v]; !ok {
			return false
		}
	}
	return true
}
```

This approach is an O(n) time complexity. The nested loop approach would be O(N * M).

## Exercises

1. Write a function that returns the intersection of two arrays.

Answer:
```go
func Intersection[T comparable](list1, list2 []T) []T {
	biggerListHash := make(map[T]struct{})
	var (
		biggerList  []T
		smallerList []T
	)
	if len(list1) > len(list2) {
		biggerList = list1
		smallerList = list2
	} else {
		biggerList = list2
		smallerList = list1
	}
	for _, v := range biggerList {
		biggerListHash[v] = struct{}{}
	}
	var result []T
	for _, v := range smallerList {
		if _, ok := biggerListHash[v]; ok {
			result = append(result, v)
		}
	}
	return result
}
```

2. Write a function that accepts an array of strings and returns the first duplicate value it finds.

Answer:

```go
// Duplicate returns the first duplicate value.
func Duplicate[T comparable](list []T) T {
	hash := make(map[T]int)

	var result T
	for _, v := range list {
		hash[v]++
		if hash[v] > 1 {
			result = v
			break
		}
	}

	return result
}
```

3. Write a function that accepts a string that contains all the letters of the alphabet except one and returns the
missing letter. For example, "the quick brown box jumps over the lazy dog" contains all letters except `f`.

Answer:

```go
func FindMissingLetter(list string) string {
	hash := make(map[rune]struct{})
	for _, v := range list {
		hash[v] = struct{}{}
	}

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, a := range alphabet {
		if _, ok := hash[a]; !ok {
			return string(a)
		}
	}

	return ""
}
```

4. Write a function that returns the first non-duplicated character in a string.

Answer:

```go
func NonDuplicated(value string) string {
	hash := make(map[rune]int)
	for _, v := range value {
		hash[v]++
	}

	for _, v := range value {
		if hash[v] == 1 {
			return string(v)
		}
	}

	return ""
}
```