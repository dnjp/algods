# Asymptotic Analysis

Algorithm analysis is a way to compare the time and space efficiency of programs
with respect to their possible inputs, but irrespective of other context.

[Example](./sum_of_nums.go)

# Big O Notation

An algorithm is a series of steps required to perform a task. If we treat each
step as a basic unit of computation, then an algorithm's execution time can be
expressed as the number of steps required to solve the problem.

In the example above, `sumOfNums` requires more steps to complete than
`arithmeticSum`. Given that steps are repeated in `sumOfNums`, the program takes
longer to complete if we increase the value of n.

The most expensive unit of computation in `sumOfNums` is variable assignment.
The initial assignment `total := 0` is performed once, followed by a loop that
executes `(total += i)n` times.

This can be denoted as T(n) = 1 + n.

n is often used to refer to the "size of the problem". The above can be read as
"T(n) is the time it takes to solve a problem of size n, namely 1 + n steps."

In order to be more precise, we're not going to worry about the exact number of
operations an algorithm performs and determine the dominant part of the T(n)
function. As the problem gets larger, some portion of the T(n) function tends to
overpower the rest - this is what we care about.

The _order of magnitude_ function (or Big O) describes the part of T(n) that increases
fastest as the value of n increases. We write it as O(f(n)) where f(n) is the
dominant part of the original T(n).

In the above example, we saw that T(n) = 1 + n. As n gets larger, the constant 1
will become less significant to the result. If we are looking for an
approximation of T(n), then we can drop the 1 and say that the running time is
O(n).

Suppose we have an algorithm whose exact number of steps can be expressed as
T(n) = 5n^2 + 27n + 1005. When n is small, the constant 1005 seems to be the
dominant part of the function. As n gets larger 5n^2 dwarfs the other terms in
the final result.

For an approximation of T(n) at large values of n, we can focus on 5n^2 and
ignore the other terms. As n gets even larger, the coefficient 5 becomes
insignificant. We can then say T(n) has an order of magnitude f(n) = n^2, or
more simply O(n^2).

The common order of magnitude functions are listed below in order from lowest to
highest:

```
f(n)        |   	Name
----------------------------
1                constant
log n            logarithmic
n                linear
n log n          log linear
n^2              quadratic
n^3              cubic
2^n              exponential
```

~[Big O Graph](https://bradfieldcs.com/algos/analysis/big-o-notation/figures/big-o-plot.png)

Go example:

```go

func bigFunc(n int) {
  a := 5
  b := 6                          // 3 assignments                   => 3
  c := 10

  for i,_ := range n {
    for j, _ := range n {
      x := i * i
      y := j * j                 // 3 assignments in nested iteration => 3n^2
      z := i * j
    }
  }

  for k,_ := range n {
    w := a * k + 45              // 2 assignments iterated n times    => 2n
    v := b * b
  }
  d := 33                        // 1 assignment                      => 1
}
```

Putting these all together: T(n) = 3 + 3n^2 + 2n + 1 = 3n^2 + 2n + 4. By looking
at the exponents, we can see that the n^2 term will be dominant, so this
function is O(n^2).
