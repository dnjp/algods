package main

import (
	"examples/stack"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	in := 42
	out := convertToBinary(in)
	fmt.Println(in, " => ", out)
}

func convertToBinary(decNum int) string {
	sremainder := stack.New()
	for decNum > 0 {
		remainder := decNum % 2
		sremainder.Push(remainder)
		fNum := math.Floor(float64(decNum) / 2)
		decNum = int(fNum)
	}

	binary := []string{}
	for sremainder.Size() > 0 {
		out := sremainder.Pop()
		val := *out
		valStr := strconv.Itoa(val.(int))
		binary = append(binary, valStr)
	}
	return strings.Join(binary, "")
}
