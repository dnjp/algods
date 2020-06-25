package main

import (
	"fmt"
)

/*
exponentiation:
	power(a, b int) -> a^b
*/

// O(n)
func power(a, b int) int {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	isEven := b % 2 == 0
	x := power(a, b / 2)
	if isEven {
		return x * x
	} 
	return a * x * x
}

/*
Thanos is attacking an avengers base. The base is n units wide where n is a power of 2. At each step, thanos has 2 options:
	1. If a region is empty, use E energy to wipe out the base
	2. If a region is not empty, then it costs `(a * Z) * sizeOfRegion` to take out the region:
		a is the number of avengers
		Z is the cost (given) per avenger
	3. Alternatively, you can split them up into 2 regions and attack them separately
	
	func getNumAvengers(left, right) is given to you (runs in constant time)
	
	given:
		E - cost to attack empty region
		Z - cost per avenger for populated region
		getNumAvengers(left, right)
		n - size of the base
		
	Return the minimum total amount of energy to take out the base
*/

func attackBase(base []int, costForEmpty, costPer, sizeOfBase int) int {
	numAvengers := getNumAvengers(base, 0, len(base))
	if numAvengers == 0 {
		return costForEmpty
	}
	
	mid := len(base) / 2
	
}

func getCost(numAvengers, costPer, size int) int {
	return (numAvengers * costPer) * size
}

func getNumAvengers(arr []int, left, right int) int {
	count := 0
	for idx, val := range arr {
		if idx >= left && idx <= right {
			count += val
		}
	}
	return count
}

func swap(data []int, m, n int) []int {
	temp := data[m]
	data[m] = data[n]
	data[n] = temp
	return data
}

func main() {
	fmt.Println(power(2, 4))
}
