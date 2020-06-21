package main

import (
	"fmt"
	"time"
)

func binarySearch(list []int, item int) bool {
	if len(list) == 0 {
		return false
	}

	mid := len(list) / 2

	if list[mid] == item {
		return true
	}

	if item < list[mid] {
		return binarySearch(list[:mid], item)
	}

	return binarySearch(list[mid+1:], item)
}

func binarySearchIter(list []int, item int) bool {
	for len(list) > 0 {
		mid := len(list) / 2
		if list[mid] == item {
			return true
		}
		if item < list[mid] {
			list = list[:mid]
		} else {
			list = list[mid+1:]
		}
	}

	return false
}

// func binarySearch2(list []int, item int) bool {
//     low := 0
//     high := len(list) - 1
//  
//     for low <= high{
//         median := (low + high) / 2
//  
//         if list[median] < needle {
//             low = median + 1
//         }else{
//             high = median - 1
//         }
//     }
//  
//     if low == len(list) || list[low] != needle {
//         return false
//     }
//  
//     return true
// }


func main() {
	testList := []int{
		121,
		134,
		176,
		301,
		680,
		695,
		787,
		805,
		938,
		1029,
		1178,
		1303,
		1338,
		1402,
		1438,
		1538,
		1773,
		1909,
		2040,
		2119,
		2198,
		2229,
		2491,
		2507,
		2529,
		2560,
		2864,
		2942,
		2993,
		2996,
		3185,
		3227,
		3359,
		3513,
		3780,
		3907,
		4150,
		4162,
		4248,
		4250,
		4273,
		4532,
		4586,
		4589,
		4776,
		4784,
		4828,
		4832,
		5037,
		5139,
		5341,
		5391,
		5403,
		5458,
		5575,
		5728,
		5784,
		5875,
		5892,
		5925,
		5983,
		5984,
		6005,
		6006,
		6034,
		6086,
		6148,
		6266,
		6356,
		6380,
		6399,
		6439,
		6525,
		6655,
		6735,
		7110,
		7165,
		7550,
		7720,
		8507,
		8641,
		8666,
		8682,
		8822,
		8924,
		8992,
		9067,
		9164,
		9253,
		9265,
		9297,
		9377,
		9583,
		9601,
		9786,
		9889,
		9892,
		9897,
		9930,
		9981,
	}
	start := time.Now()
	a1 := binarySearch(testList, 9265)
	elapsed := time.Since(start)
	fmt.Println("recur 1: ", a1, elapsed)

	start = time.Now()
	a2 := binarySearchIter(testList, 9265)
	elapsed = time.Since(start)
	fmt.Println("iter 1: ", a2, elapsed)

	start = time.Now()
	b1 := binarySearch(testList, 10000)
	elapsed = time.Since(start)
	fmt.Println("recur 2: ", b1, elapsed)

	start = time.Now()
	b2 := binarySearchIter(testList, 10000)
	elapsed = time.Since(start)
	fmt.Println("iter 2: ", b2, elapsed)
}
