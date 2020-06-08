package main

import (
	"examples/stack"
	"fmt"
)

type PAIRS map[rune]rune

var PAIRINGS = PAIRS{
	'(': ')',
	'{': '}',
	'[': ']',
}

func main() {
	in := "((()))"
	out := isBalanced(in)
	fmt.Println(in, " => ", out)
}

func isBalanced(input string) bool {
	s := stack.New()
	for _, r := range input {
		if PAIRINGS.InKeys(r) {
			s.Push(r)
			continue
		}
		expectedOpeningSymbol := s.Pop()
		if expectedOpeningSymbol == nil {
			return false
		}

		openingSymbol := *expectedOpeningSymbol
		if r != PAIRINGS[openingSymbol.(rune)] {
			return false
		}

	}
	return s.Size() == 0
}

func (p *PAIRS) InKeys(in rune) bool {
	for key := range *p {
		if key == in {
			return true
		}
	}
	return false
}
