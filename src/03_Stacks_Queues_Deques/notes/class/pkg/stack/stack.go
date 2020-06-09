package stack

import (
	"fmt"
)

type IStack interface {
	// Creates a new, empty stack
	New() *IStack

	// Adds given item to top of stack
	Push(interface{})

	// Removes and returns the top item from the stack
	Pop() *interface{}

	// Returns the top item from the stack but doesn't remove it
	Peek() *interface{}

	// Returns whether the stack is empty
	IsEmpty() bool

	// Returns the number of items on the stack
	Size() int
}

type Stack struct {
	items []interface{}
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() *interface{} {
	length := len(s.items)
	if length > 0 {
		item := s.items[length-1]
		s.items = s.items[:length-1]
		return &item
	}
	return nil
}

func (s *Stack) Peek() *interface{} {
	length := len(s.items)
	if length > 0 {
		item := s.items[len(s.items)-1]
		return &item
	}
	return nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Print() {
	fmt.Printf("contents: %+v\n", s.items)
}
