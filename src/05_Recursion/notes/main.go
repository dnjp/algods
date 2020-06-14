package main

import "fmt"

var CharForInt string = "0123456789abcdef"

func toString(n int, base int) string {
	if n < base {
		return string(CharForInt[n])
	}

	return toString(n/base, base) + string(CharForInt[n%base])
}

func numPathsDP(height int, width int) int {
	memo := make([][]int, height)
	for i := 0; i < height; i++ {
		memo[i] = make([]int, width)
		for j := 0; j < width; j++ {
			memo[i][j] = 1
		}
	}

	for i := 1; i < len(memo); i++ {
		row := memo[i]
		for j := 1; j < len(row); j++ {
			memo[i][j] = memo[i - 1][j] + memo[i][j - 1]
		}
	}
	fmt.Printf("memo: %+v\n", memo)
	return memo[height-1][width-1]
}

func main() {
	// fmt.Println(toString(1453, 10))
	fmt.Println(numPathsDP(3, 3))
}
