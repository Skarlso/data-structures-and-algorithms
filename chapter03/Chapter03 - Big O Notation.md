# Big O Notation

We can't simply label algorithms as something like "22-steps algorithm". Because the number of steps changes.
In the case of linear search, a better approach would be to say N steps for N elements.

Big O notation allows us to categorize efficiency easily. There are tools for analyzing and determining Big O notation.

Big O simplified explanation: How many steps relative to N elements?

Key Question: _If there are N data elements how many steps will the algorithm take?_

The answer lies in the `()` section of the notation. For linear, that would be O(N). The algorithm will take N steps.
O(n) is called _Linear Time_.

Reading a single value takes O(1) steps. This is called _constant time_.

Big O defines the _upper bound_ of the growth rate.

But there is a deeper meaning still. The soul of Big O is to define how an algorithm will behave when the data _grows_.

For O(1) there is no change if data grows. Or O(3) or O(4). It's **constant**. But for linear, for example, the steps
change. Linearly.

![linear](./linear-function.png)

O(n) is the _worst case_ scenario. Linear isn't always O(n) but the worst case _is_.

So what does that make Binary Search? For that we have to analyze how many steps binary search takes compared to growth.
It increased once for every size doubling. We talked about it being a logarithmic scale. We can easily describe that
using Big O by writing O(logN). So the three now in sequence are: O(1), O(logN), O(n).

On a graph that looks something like this:

![3](./3-together.png)

You can see in red O(n), blue is O(logN) and black is O(1).

_Simply said, O(logN) means the algorithm takes as many steps as it takes to keep halving the data elements until we remain with 1._

## Exercises

1. Q: Use Big O Notation to describe the time complexity of the following function that determines whether a given year is
a leap year.
```go
func isLeapYear(year int) {
    if year % 100 == 0 {
        return year % 400 == 0
    }
    return year % 4 == 0
}
```

A: O(1) since there are no loops and it's basically just a read operation and a calculation.

2. Q: Use Big O Notation to describe the time complexity of the following function that determines whether a given year is
a leap year.

```go
func sum(list []int) int {
    sum := 0
    for _, v := range list {
        sum += v
    }
    return sum
}
```

A: O(n) since it loops through all elements and performs an operation.

3. Q: The following describes the power of compounding interest:
    Imagine you have a chessboard and put a single grain of rice on one square. On the second square you put 2 grains.
    On 3rd you put 4 ( double the previous square ). The 5th will be 16.
    The following function calculates which square you would put a given number of grains on. For 16 it returns 5.

```go
func chessboardSpace(n int) int {
    space := 1
    placedGrains := 1

    for placedGrains < n {
        placedGrains *= 2
        space++
    }
    return space
}
```

This is rather convoluted example. Simply return log2(16) + 1.
Each time we _double_ the placedGrains. Which means this is O(logN).


4. Function which filters out all strings from a list of strings that start with the character `a`.
> This would be entirely different if it would be a `contains`.

```go
func selectAString(s []string{}) []string {
    var ret []string
    for _, v := range s {
        if len(s) > 0 && s[0] == 'a' {
            ret = append(ret, v)
        }
    }
    return ret
}
```

Since it's only a single loop checking every item, it's O(n). We don't take `append` into account.

5. Calculate the median from an _ordered array_.

```go
func median(list []int) int {
    l := len(list)
    middle := l / 2
    if l % 2 == 0 {
        return (list[middle-1] + list[middle]) / 2
    }
    return list[middle]
}
```

Remember, the list is ordered, so we don't have to loop around at all. The `len` calculation is also not a loop
so this is O(1).