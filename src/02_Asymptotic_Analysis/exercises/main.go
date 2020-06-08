package main

import (
	"fmt"
)

func fac(n int) int {
	if n == 1 {
		return n
	}
	return n * fac(n*n-1)
}

func main() {
	output := fac(10)
	fmt.Println(output)
}
