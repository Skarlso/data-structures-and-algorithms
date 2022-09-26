# Recursively Recurse with Recursion

Recursion can be used to solve certain types of tricky problems in surprisingly nice and tidy ways.

_Recursion_ is the term for a function that calls itself during its runtime.

Infinite recursions are possible if there is no termination logic. Just like for loops.

## Recurse instead of loop

We don't _have_ to use a loop. There is another option in the form of recursion. Let's transform a regular loop operation
into a recursive call:

```go
func countdown(from int) {
    for i := from; i > 0; i-- {
        fmt.Println(i)
    }
}
```

This becomes:

```go
func countdown(from int) {
    fmt.Println(from)

    countdown(from-1)
}
```

It will now display the same thing, but it has a flaw. There is no stop. There is no _Base Case_.

## The Base Case

What is the call stack for this countdown function? Let's explore step-by-step.

- countdown(10)
    - display the number
    - call function with current number-1; which will be 9.
        - countdown(9)
            - display the number
            - call function with current number-1; which will be 9.
                - countdown(8)
                ...
                - countdown(0)
                - countdown(-1)

It doesn't stop. Ever. This is an _infinite recursion_. This will eat up the call-stack and eventually error with a
stack overflow. Recursions have two parts to them. The _base case_ and the _recursive step_. The base case is what
happens when you want the recursion to stop. The single task you want to perform on sub-elements of the same nature.
It's also the mechanism to prevent the function from calling itself forever.

In this case, when do we want to stop? When the count is `0`. So let's add that:

```go
func countdown(from int) {
    fmt.Println(from)
    if from == 0 {
        return
    }
    countdown(from-1)
}
```

Done. This will now print until `0` and then stop the recursive process.

## Reading Recursive Code

You need to sets of skills. Reading and Writing recursive code. This might seem trivial, but reading recursive code
requires a mind-set switch from the normal, sequential way of reading code. First, identify the base-case. Then you can
move on to identifying the recurrence.

Let's use the ever popular _factorial_ to read some recursion.

As a reminder, _factorial_ is 3 * 2 * 1 = 6.

```go
func factorial(n int) int {
    if n == 1 {
        return 1
    }
    return n * factorial(n - 1)
}
```

What is happening here? Let's walk through it step-by-step:

- factorial(3)
    - is n == 1?
    - no: call 3 * factorial(2)
        - factorial(2)
            - is n == 1?
            - no: 2 * factorial(1)
                - factorial(1)
                    - is n == 1?
                    - yes: return 1
            - return 2 * 1 (2)
    - return 3 * 2 (6)

Did you catch that? It calls, calls, calls until n is 1 and then unravels the whole call-chain with the returned values.
It's pretty elegant. And guess what? The function calls are all in a _stack_. Which we learned about previously. That's
how it keeps track of the calls. And that's why it can, eventually, overflow. It runs out of memory if there is an
infinite number of calls.

## Recursion in the Eyes of the Computer

This might seem trivial to us, but for the computer it's a bit more complicated. The function call doesn't end when it
encounters a recursion. It will put that call into a stack than call it again and put that call on top of the stack and
perform the function. Basically the working of the previous function is suspended until the next call returns something
that isn't a function call to itself.

## The Call Stack

I touched on this previously but how does it look like? Well, just how you'd imagine it.

| Stack        |
| ------------ |
| factorial(3) |

First, it starts off with the first call. Then the next one gets on top.

| Stack        |
| ------------ |
| factorial(2) |
| factorial(3) |

And finally, the last one.

| Stack        |
| ------------ |
| factorial(1) |
| factorial(2) |
| factorial(3) |

Once the last one returns, it unravels the call stack by popping off calls from it.

Some refer to this idea as _passing a value up through the call stack_.

## Stack Overflow

This is an allocation. The call stack is a stack in memory. Which means it can hold a limited number of items. During
normal operations, the call stack is just fine. It doesn't get too many items. But once recursion comes into play, oh boy.

That means it's **critically** important that a base case exists; otherwise, you might end up crashing your application
spectacularly.

## Filesystem Traversal

Let's look at a practical example now. Imagine a filesystem with multiple sub-folders and files. You would like to list
all file names starting from a selected location.

Doing this in a loop would be tricky or even ugly. It could look something like this:

```go
func findFiles(location string)
    files, err := ioutil.ReadDir(location)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
        if file.IsDir() {
            files, err := ioutil.ReadDir(location)
            if err != nil {
                log.Fatal(err)
            }
            for _, file := range files {
                fmt.Println(file.Name())
                if file.IsDir() {
                    files, err := ioutil.ReadDir(location)
                    if err != nil {
                        log.Fatal(err)
                    }
                    for _, file := range files {
                        fmt.Println(file.Name())
                        if file.IsDir() {
                            files, err := ioutil.ReadDir(location)
                            if err != nil {
                                log.Fatal(err)
                            }
                        }
                    }
                }
            }
        }
    }
}
```

And so on and so forth. And, besides this being insanely complex it will only work with folders three level deep. But
folders can be N number of deeply nested. You can also see that there is a code part where we just keep repeating the
same thing in every step. It seems we have our _recursive step_. So, what's our base case? Listing the filenames.

Let's try this again in a recursive way:

```go
func findFiles(location string) {
	files, err := os.ReadDir(location)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			findFiles(filepath.Join(location, file.Name()))
		}
		fmt.Println(file.Name())
	}
}
```

Pretty neat, right? It's small, compact and works for any number of sub-folders. It's a lot simpler to read and reason
about too.

## Exercises

1. The following function displays ever other number starting from `low` to `high`. For example, low is 0 high is 10 it
would print: 0, 2, 4, 6, 8, 10.

```go
func printEveryOther(low, high int) {
    if low > high {
        return
    }
    fmt.Println(low)
    printEveryOther(low + 2, high)
}
```

Identify the base-case of this recursion.

Answer: When low gets larger than high (12 in this case) it should stop the process. If it's equal to or lower it prints
the given number.

2. This is a changed factorial recursion. Can you guess what happens when it's called with factorial(10)?

```go
func factorial(n int) int {
    if n == 1 {
        return 1
    }
    return n * factorial(n - 2)
}
```

Answer: It's an infinite recursion because n will not get to be `1` it will go down as 10, 8, 6, 4, 2, 0, -2, -4...

3. Following function sums all numbers from `low` to `high`. Can you write the correct base case?

```go
func sum(low, high int) int {
    return high + sum(low, high - 1)
}
```

Answer:

```go
func sum(low, high int) int {
    // Basically, we want it to stop as soon as we have equal numbers. That's when we return a concrete number
    // so the call chain can add them together while it unravels.
    if low == high {
        return low
    }
    return high + sum(low, high - 1)
}
```

4. Here is an array containing arrays and number. Only print the numbers!

```
[
    1,
    2,
    3,
    [4, 5, 6],
    7,
    [8, 9, 10, [
            11, 12, 13,[
                14, 15, 16
            ]
        ],
    ],
    17,
    18,
    19, [
        20, 21, 22
    ]
]
```

This is a bit tricky in Go as Go is a typed language.
```go
func PrintNumber(slice []any) {
	for _, item := range slice {
		if s, ok := item.([]any); ok {
			PrintNumber(s)
		} else {
            fmt.Println(item)
        }
	}
}
```

This prints:

```
=== RUN   TestPrintNumber
1
2
3
4
5
6
7
8
9
```

For a slice definition as such:

```go
	slice := []any{
		1,
		2,
		3,
		[]any{
			4, 5, 6,
		},
		7, 8, 9,
	}
```
