package main

/*
 * if smaller number comes before larger number, treat it as subtraction
 * otherwise, numbers go bigger to smaller
 *
 * not all pairse are allowed:
 *
 * I can precede V or X
 * X can precede L or C
 * C can precede D or M
 *
 * What is the unknown?
 * The number that the roman romanNumerals total to.
 * What are the data?
 * The given roman romanNumerals in all upper case.
 * What is the condition?
 *
 *
 * Challenge: Given a roman numeral, convert it into the shortest possible roman numeral that communicates that value
 *
 */

import (
	"fmt"
)

type Test struct {
	phrase string
	value  int
}

type Num struct {
	value    int
	subChars map[string]bool
}

func getNum(phrase string) int {
	romanNumerals := map[string]Num{
		"I": Num{value: 1, subChars: map[string]bool{"V": true, "X": true}},
		"V": Num{value: 5, subChars: map[string]bool{}},
		"X": Num{value: 10, subChars: map[string]bool{"L": true, "C": true}},
		"L": Num{value: 50, subChars: map[string]bool{}},
		"C": Num{value: 100, subChars: map[string]bool{"D": true, "M": true}},
		"D": Num{value: 500, subChars: map[string]bool{}},
		"M": Num{value: 1000, subChars: map[string]bool{}},
	}

	total := 0
	i := 0
	for i < len(phrase) {
		curChar := phrase[i]
		curNum := romanNumerals[string(curChar)]
		if curNum.value == 0 {
			// illegal entry
			return -1
		}

		if i+1 > len(phrase)-1 {
			total += curNum.value
			i++
			continue
		}
		nextChar := phrase[i+1]
		if curNum.subChars[string(nextChar)] {
			nextNum := romanNumerals[string(nextChar)]
			curVal := nextNum.value - curNum.value
			total += curVal
			i += 2
		} else {
			total += curNum.value
			i++
		}
	}
	return total
}

func main() {
	tests := []Test{
		{"LXIII", 63},
		{"LXIV", 64},
		{"LXXXIX", 89},
		{"XCIV", 94},
		{"C", 100},
	}
	for _, t := range tests {
		output := getNum(t.phrase)
		if output != t.value {
			fmt.Printf("failed: phrase: %s, value: %d, received: %d\n", t.phrase, t.value, output)
		} else {
			fmt.Printf("passed: phrase: %s, value: %d, received: %d\n", t.phrase, t.value, output)
		}
	}
}
