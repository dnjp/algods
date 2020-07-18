package main

import (
	"fmt"
)

/*
 	example:

 	arr =  [2, 4, 6, 8]
	temp:                              4 => 2 => 1
	i:                                     0 => 0 => 0 => 1      <= ANSEWR
	half:                                2 => 1 => 0
	result (arr[half] >= 4):     t =>  t =>  f

 	arr =  [2, 4, 6, 8]
	temp:                              4 => 4
	i:                                     0 => 3 => 4   <= ANSEWR
	half:                                2 => 2
	result (arr[half] > 8):       f =>  f

*/
func binarySearch(arrLen int, f func(int) bool) int {
	low := 0
	high := arrLen
	for low < high {
		mid := low + (high-low)/2 // avoid overflow. 
		if !f(mid) {
			low = mid + 1 // preserves f(low-1) == false
		} else {
			high = mid // preserves f(high) == true
		}
	}
	return low
}

func main() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	greaterThanEq28 := func(i int) bool {
		return a[i] >= 28
	}
	i := binarySearch(len(a), greaterThan36)
	fmt.Println(a, i)
}
