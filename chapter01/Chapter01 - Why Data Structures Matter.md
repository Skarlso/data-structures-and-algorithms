# Chapter 01 - Why Data Structures Matter

Data structures define how data is organized.

Code example on string.

```go
s := "asdf" + "asdf2" + "asdf3"
s2 := []string{"asdf", "asdf2", "asdf3"}
s3 := "asdfasdf2asdf3"
```

Example on efficiency.

```go
count := 0
for i := 0; i < 100; i++ {
    if i % 2 == 0 {
        count++
    }
}
fmt.Println("count: ", count)
```

vs

```go
count := 0
i := 2
for i < 100 {
    count++
    // implicitly dividable by two.
    i += 2
}
fmt.Println("count: ", count)
```

Same result but the second one will only loop 50 times.

## The Array

Array is the most basic structure.

Basic array concepts: `Size`, `Index`.

Index starts from 0.

```go
s := []string{"1", "2", "3"}
fmt.Println(s[0]) // prints 1
fmt.Println(s[1]) // prints 2
fmt.Println(s[2]) // prints 3
```

Four basic operations: _Read_, _Search_, _Insert_, _Delete_

Speed is measured based on how many **steps** and operation took and not how much **time** it took. Time is relative
to the hardware. Steps are constant.

### Reading
It takes _one_ step because of how memory works and arrays work. Arrays are stored in memory as a sequentially increasing
addressed space. The first element is at address 1011, the next is 1012, next is 1013. The _begin_ address is stored
with an array, so once you as for the 3rd element it says: _begin address_ + 3.

### Searching

Searching is tedious as there is no way to jump. Has to check all cells to find a particular item. This means it takes
as many steps to search an array as there are elements. N elements means N steps.

```go
n := []int{1,2,3,4,5}
for i, v := range n {
    if n == 4 {
        fmt.Println("found at index: ", i)
        break
    }
}
```

### Insertion

Insertion's efficiency depends on where the item needs to be inserted. To the _end_ just **one** step. In the middle,
the array's existing elements have to be shifted around. Worst-case is the beginning, where it takes N+1 steps for N
elements because we have to shift all elements plus one for the actual insertion operation.

### Deletion

Deletion is almost the same as insertion, but we only need to close the gap that is produced. Again, deletion from the
end is the most efficient, and deletion from the beginning is the least efficient, taking N steps in an array containing
N elements.

_Note_ it's only N and not N + 1 because we just removed one element and shifted N-1 elements.

## The Set

Have a single rule: No duplicate items. This changes one of the operations. Can you guess which? It's _insertion_.
Because before we insert, we have to check if the array already contains the item or not. This requires a search in case
of a hash map that would be a single operation. But in the case of an array, this is, as we saw previously, N steps
(worst case).

And thus, insertion for a set requires N+1 for insertion and N for search. Which is 2N+1.

## Exercises

1. For an array containing a 100 elements, provide the number of steps for the following:
    a. Reading -> 1
    b. Search for a value not containing within the array -> N
    c. Insertion at the begin -> 1 + N shifts
    d. Insertion at the end -> 1 + 0 shifts
    e. Deletion at the begin -> 1 + N - 1 shifts
    f. Deletion at the end -> 1 + 0 shifts
2. For an array-based set containing a 100 elements, provide the number of steps for the following:
    a. Reading -> 1
    b. Search for a value not containing within the array -> N
    c. Insertion at the begin -> 1 + N for search + N for shifting -> 2N + 1
    d. Insertion at the end -> 1 + N for searching + 0 for shifting -> N + 1
    e. Deletion at the begin -> 1 + N - 1 shifts
    f. Deletion at the end -> 1 + 0 shifts
3. How many steps does a search for all occurrences of an item X in an array with N elements take?
    Answer: N steps if we are looking for N occurrences. If we look for 1 occurrence we might be able to cut it short
    after we find it.