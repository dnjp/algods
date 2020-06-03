package main

import "fmt"

type test struct {
	given    int
	expected int
}

/*
sumToNum returns the sum of the first n numbers

Runs in O(n) or linear time
*/
func sumToNum(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

/*
arithmeticSum returns the sum of the first n numbers

Runs in O(1) or constant time
*/
func arithmeticSum(n int) int {
	total := (n * (n + 1)) / 2
	return total
}

func main() {
	tests := []test{
		{given: 1, expected: 1},

		// 5 + 1, 4 + 2, 3
		{given: 5, expected: 15},
		{given: 10, expected: 55},
	}

	for _, t := range tests {
		output := sumToNum(t.given)
		if output != t.expected {
			fmt.Println("sumToNum: expected", t.expected, "but received", output)
		}
		output = arithmeticSum(t.given)
		if output != t.expected {
			fmt.Println("arithmeticSum: expected", t.expected, "but received", output)
		}
	}
}
