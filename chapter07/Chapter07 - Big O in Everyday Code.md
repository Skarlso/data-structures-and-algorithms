# Big O in Everyday Code

Determining the efficiency of our code is the first step towards optimizing it.

Once we know the category, we can determine if we can optimize it or not.

## Mean Average of Even Numbers

The following code calculates the mean average of all its _even_ numbers.

```go
func averageOfEvenNumbers(list []int) int {
    // It will be calculated determining sum of the even numbers dividing the number of even numbers.
    sum := 0
    count := 0

    for _, v := range list {
        if v % 2 == 0 {
            sum += v
            count++
        }
    }

    return sum / count
}
```

- Determine what is our dataset. The N number of items.
- How many steps does it take to process N values?
- Since we loop through all items, we know that the algorithm takes at least N steps.
- _inside_ the loop the steps wary based on the number and type of items.
- Big O focuses primarily on the worst-case scenario.
- The worst case is if every item is even.
- In that case, we perform 3 steps each round -> O(3n).
- The variable init is another two steps O(3n + 2).
- And at the end, we have the division which is another step. O(3n + 3).
- But, since we don't care about constants... O(N).

## Word Builder

The next example collects every combination of two-character strings built from an array of single characters.

For example:

["a", "b", "c", "d"]:

[
    'ab', 'ac', 'ad', 'ba', 'bc', 'bd',
    'ca', 'cb', 'cd', 'da', 'db', 'dc'
]

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

The first thing to notice is the nested looping. That's O(n^2) on the same list.

If we would try to do this with three characters...

```go
func wordBuilder(list []string) []string {
    var collections []string
    for i := 0; i < len(list); i++ {
        for j := 0; j < len(list); j++ {
            for k := 0; k < len(list); k++ {
                if i != j && j != k && i != k {
                    collections = append(collections, string(list[i]) + string(list[j]) + string(list[k]))
                }
            }
        }
    }
    return collections
}
```

We have two nested loops now. This results in O(n^3).


## Array Sample

Takes a small sample of an array.

```go
func sample[T Numbers](list []T) []T {
    first := list[0]
    middle := list[len(list)/2]
    last := list[len(list)-1]
    return []T{first, middle, last}
}
```

No loops, and the number of steps is constant. Therefor, O(1).

## Average Celsius Reading

