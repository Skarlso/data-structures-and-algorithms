# Dealing with Space Constraints

We focused more or less on running speed these past few chapters, but there is another variant you need to consider when
dealing with data. It's _space complexity_. How much space your algorithm uses? How much memory does
it allocate when running?

## Big O of Space Complexity

With memory, we need to ask _if there_ are N data elements; how memory units will the algorithm consume?_

Let's see an example. Consider the following code:

```go
func makeUppercase(list []string) []string {
    var result []string
    for _, w := range result {
        result = append(result, strings.ToUpper(w))
    }
    return result
}
```

This code allocates a new slice. And adds each element into it. Meaning the new slice grows with the data. It creates as
many new items as there are data items. This is O(N) space complexity. Let's make this O(1).

```go
func MakeUppercaseO1(list []string) {
	for i, w := range list {
		list[i] = strings.ToUpper(w)
	}
}
```

This will not allocate any **new** space. The **new** is an important distinction. Big O notation doesn't consider
counting existing space. It counts newly created space. The extra space is called _auxiliary_ space.

## Trade-Off between time and space

( This sounds like a Marvel Movie idea... )

Now, the question for you to decide is when you are writing your algorithm, what are your constraints? Do you have
limited capacity for RAM or CPU? You can't optimize for both in most cases. ( sometimes you can ).

For example, let's consider the following method that returns whether there are duplicate values in an list or not:

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

This is obviously largely inefficient and we can do better in terms of speed by using a hashmap.

```go
func hasDuplicateValue(list []int) bool {
    items := map[int]struct{}{}

    for _, v := range list {
        if _, ok := items[v]; ok {
            return true
        }
        items[v] = struct{}{}
    }

    return false
}
```

Now, this is a lot faster. It's O(n) speed is vastly superior to the O(n^2) speed of the previous solution. However,
there is a tradeoff. This uses an extra hashmap and potentially stores all N values again if there are no duplicates.
If your space requirement can't manage that extra memory because you are running in an embedded system, for example, you
need an algorithm that is faster, but space efficient. You can use sorting first ( as we learned, sorting
takes O(nlogN) which is still a magnitude better than O(n^2) ).

```go
func hasDuplicateValue(list []int) bool {
	sort.Ints(list)

	for i := 0; i < len(list)-1; i++ {
		if list[i] == list[i+1] {
			return true
		}
	}
	return false
}
```

Depending on the sorting implementation, this might still allocate some data. But the top space efficiency will max out
at O(logN) which is not O(1) but not that bad either. Heapsort, for example, is O(1) space efficiency and O(nlogN) speed
efficiency. Meanwhile, Quicksort that is used by most standard libraries ( or a variation of it ) uses O(logN) space.

## Hidden cost of recursion

Now, we talked about recursion and time complexity, but we have yet to discuss recursion and space complexity. It turns
out that recursion does result in none constant space efficiency. Because it allocates a new value each time it's called
on the stack. The call stack also uses space. A 100 recursive calls means a 100 items in the call stack. Eventually, the
call stack might fill up. Which will result in a call stack size exceeded error.

And that's why Quicksort uses O(logN) space. It uses recursive calls and it does that exactly O(logN) times.

## Exercises

1. Describe the space complexity of the word builder algorithm.
```go
func wordBuilder(list []string) []string {
    var collections []string
    for i := 0; i < len(list); i++ {
        for j := 0; j < len(list); j++ {
            if i != j {
                collections = append(collections, string(list[i]) + string(list[j]))
            }
        }
    }
    return collections
}
```

Answer: O(n^2) because it creates an array which will end up containing N^2 items.

2. Following is a function that reverses an array. Describe its space complexity.

```go
func reverse(list []int) []int {
	var list []int
	for i := len(list) - 1; i >= 0; i-- {
		list = append(list, list[i])
	}

	return list
}
```

Answer: O(n) because it allocates as many items as there are items in the list.

3. Create a new function that reverses an array with O(1) space.

```go
func reverse(list []int) {
	for i := len(list)/2 - 1; i >= 0; i-- {
		opp := len(list) - 1 - i
		list[i], list[opp] = list[opp], list[i]
	}
}
```

4. Define the space complexity of the following three functions:

```go
func Double1(list []int) []int {
	var newList []int

	for _, v := range list {
		newList = append(newList, v*2)
	}

	return newList
}

func Double2(list []int) []int {
	for i := 0; i < len(list); i++ {
		list[i] *= 2
	}
	return list
}

func Double3(list []int, index int) []int {
	if index >= len(list) {
		return list
	}

	list[index] *= 2

	return Double3(list, index+1)
}
```

Answer:

| Version  | Time Complexity | Space Complexity |
| -------- | --------------- | ---------------- |
| Double 1 | O(n)            | O(n)             |
| Double 2 | O(n)            | O(1)             |
| Double 3 | O(n)            | O(n)             |