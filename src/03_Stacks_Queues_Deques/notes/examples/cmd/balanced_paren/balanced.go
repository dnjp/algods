package main

import (
	"examples/stack"
	"fmt"
)

func main() {
	in := "((()))"
	out := isBalanced(in)
	fmt.Println(in, " => ", out)
}

func isBalanced(input string) bool {
	opening := '('
	closing := ')'

	s := stack.New()
	for _, r := range input {
		if r == opening {
			s.Push(r)
		}
		if r == closing {
			val := s.Pop()
			if val == nil {
				return false
			}
		}
	}
	return s.Size() == 0
}
