package stack

// IStack is the interface for a stack
type IStack interface {
	// New creates a new, empty stack
	New() *IStack

	// Push adds given item to top of stack
	Push(interface{})

	// Pop removes and returns the top item from the stack
	Pop() *interface{}

	// Peek returns the top item from the stack but doesn't remove it
	Peek() *interface{}

	// IsEmpty returns whether the stack is empty
	IsEmpty() bool

	// Size returns the number of items on the stack
	Size() int
}

// Stack is a LIFO data structure where additions and removals occur at the top of the stack
type Stack struct {
	items []interface{}
}

// New creates a new, empty stack
func New() *Stack {
	return &Stack{}
}

// Push adds given item to top of stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item from the stack
func (s *Stack) Pop() *interface{} {
	length := len(s.items)
	if length > 0 {
		item := s.items[length-1]
		s.items = s.items[:length-1]
		return &item
	}
	return nil
}

// Peek returns the top item from the stack but doesn't remove it
func (s *Stack) Peek() *interface{} {
	length := len(s.items)
	if length > 0 {
		item := s.items[len(s.items)-1]
		return &item
	}
	return nil
}

// IsEmpty returns whether the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items on the stack
func (s *Stack) Size() int {
	return len(s.items)
}
