package main

import (
	"fmt"
)

/*
	square root
		input: integer n
		output: smallest value such that x^2 >= v

	example:
		25 -> 5
		26 -> 6
		36 -> 6
	 	100 -> 10
	 	99 -> 10

	 don't focus so much on the _exact_ boundary, just focus on the question
*/

func sqRoot(n int) int {
	low := 0
	high := n
	mid := low + (high - low) / 2
	x := low * low
	// want low^2 <= n
	// 
	
	for x <= n  {
		x := low * low
		if x >= n {
			return mid
		}
		// mid := low + (high - low) / 2
	}
	return -1
}

func ex1(arr []int, left int, right int) {
	pivot = arr[left]
	i := left+1
	j := left + 1
	for j < right {
		if arr[j] <= pivot {
			// swap
			swap(arr, i, j)
			i += 1
			j += 1
		} else {
			j += 1
		}
		swap(arr, i-1, left)
	}
	return i-1
}

func main() {
	fmt.Println(sqRoot(25))
}