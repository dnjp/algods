package main

import (
	"fmt"
	"math"
	"time"
)

func bubbleSort(data []int) []int {
	N := len(data)

	// start at the end and work your way down until you are 1 above the first element
	for j := N - 1; j > 0; j-- {
		// sort starting from beginning of dataay up to j
		for i := 0; i < j; i++ {
			if data[i] > data[i+1] {
				// if the rightmost item in the pair is higher than the left, swap them
				data = swap(data, i, i+1)
			}
		}
	}
	return data
}

func selectionSort(data []int) []int {
	N := len(data)
	// iterate through the entire length of the array
	for i := 0; i < N-1; i++ {
		// find the index of the minimum value starting with index i
		k := minValPos(data, i)

		// if the minimum value is greater than that at i, swap it with i if i
		if i != k {
			data = swap(data, i, k)
		}
	}
	return data
}

func insertionSort(data []int) []int {
	N := len(data)
	// iterate through the whole array starting at the second element
	for j := 1; j < N; j++ {
		/*
			starting with idx j, reverse iterate through the array until you reach an element
			that is greater than the one before it
		*/
		for i := j; i > 0 && data[i] < data[i-1]; i-- {
			// you've found an out of order element, swap them
			data = swap(data, i, i-1)
		}
	}
	return data
}

func shellSort(data []int, intervals []int) []int {
	N := len(data)
	for k := len(intervals) - 1; k >= 0; k-- {
		interval := intervals[k]
		for m := 0; m < interval; m++ {
			for j := m + interval; j < N; j += interval {
				for i := j; i >= interval && data[i] < data[i-interval]; i -= interval {
					data = swap(data, i, i-interval)
				}
			}
		}
	}
	return data
}

func quicksort(data []int, lowI, highI int) []int {
	afterSmall := lowI // first idx: starts at 0
	beforeBig := highI // last elem: starts at len(data)-1

	midpoint := (lowI+highI)/2
	// find the middle value using integer division
	
	/*
	pivot must meet 3 criteria:
		1. Correct position in final array
		2. Items to the left are smaller
		3. Items to the right are larger
	*/
	pivot := data[midpoint]

	/*
	In loop data[i] <= pivot if i < afterSmall
	            data[i] >= pivot if i > beforeBig
	            region with aftersmall <= i <= beforeBig shrinks to nothing
	*/

	for afterSmall <= beforeBig {
		// start from beginning of array, working up until the element is less than the pivot
		for data[afterSmall] < pivot {
			afterSmall++
		}

		// start from end of array, working down until the element is greater than the pivot
		for data[beforeBig] > pivot {
			beforeBig--
		}

		if afterSmall <= beforeBig {
			data = swap(data, afterSmall, beforeBig)
			afterSmall++
			beforeBig--
		}
	}

	/*
	   After loop: beforeBig < afterSmall, and
	     data[i] <= pivot for i <= beforeBig,
	     data[i] == pivot for i if beforeBig < i < afterSmall,
	     data[i] >= pivot for i >= afterSmall
	*/

	// at least two elements
	if lowI < beforeBig {
		return quicksort(data, lowI, beforeBig)
	}

	if highI > afterSmall {
		return quicksort(data, afterSmall, highI)
	}
	
	return data
}

func minValPos(data []int, start int) int {
	N := len(data)
	minPos := start
	for pos := start + 1; pos < N; pos++ {
		if data[pos] < data[minPos] {
			minPos = pos
		}
	}
	return minPos
}

func swap(data []int, m, n int) []int {
	temp := data[m]
	data[m] = data[n]
	data[n] = temp
	return data
}

func generateIntervals(n int) []int {
	if n < 2 {
		return []int{0}
	}

	t := int(math.Max(1.0, log(float64(n), 3.0)-1))
	intervals := make([]int, t)
	for i := 1; i < t; i++ {
		intervals[i] = 3*intervals[i-1] + 1
	}
	return intervals
}

func log(x, base float64) float64 {
	return math.Log(x) / math.Log(base)
}

func main() {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{12, 8, -5, 22, 9, 2},
			expected: []int{-5, 2, 8, 9, 12, 22},
		},
	}

	for _, t := range tests {
		// bubble sort
		bubble := bubbleSort(createInput(t.input))
		passes := testArr(bubble, t.expected)
		print("Bubble Sort", passes, bubble, t.expected)

		// selection sort
		selection := selectionSort(createInput(t.input))
		passes = testArr(selection, t.expected)
		print("Selection Sort", passes, selection, t.expected)

		// insertion sort
		insertion := insertionSort(createInput(t.input))
		passes = testArr(insertion, t.expected)
		print("Insertion Sort", passes, insertion, t.expected)

		// shell sort
		shell := shellSort(createInput(t.input), generateIntervals(int(time.Now().Unix())))
		passes = testArr(shell, t.expected)
		print("Shell Sort", passes, shell, t.expected)

		// quicksort
		in := createInput(t.input)
		quick := quicksort(in, 0, len(in)-1)
		passes = testArr(quick, t.expected)
		print("Quicksort", passes, quick, t.expected)
	}
}

func print(name string, passes bool, in, exp []int) {
	if !passes {
		fmt.Printf("%s failed: expected %+v but received %+v\n", name, exp, in)
	} else {
		fmt.Printf("%s passed\n", name)
	}
}

func createInput(in []int) []int {
	testInput := []int{}
	for _, val := range in {
		testInput = append(testInput, val)
	}
	return testInput
}

func testArr(in, exp []int) bool {
	for idx, val := range in {
		if val != exp[idx] {
			return false
		}
	}
	return true
}
