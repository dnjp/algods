package reverselist

import (
	"exercise/stack"
)

// ReverseList takes a list of strings as input and returns the reverse of that list
func ReverseList(input []string) []string {
	s := stack.New()
	for _, val := range input {
		s.Push(val)
	}

	newList := []string{}
	for s.Size() > 0 {
		out := s.Pop()
		if out == nil {
			panic("items have not been properly added to stack")
		}
		val := *out
		newList = append(newList, val.(string))
	}
	return newList
}
