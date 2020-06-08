package main

import (
	"examples/stack"
	"fmt"
)

func main() {
	runStackTest()
}

func runStackTest() {
	var t = struct {
		Name string
	}{
		Name: "daniel",
	}
	s := stack.New()
	s.Push(1)
	s.Print()

	item1 := s.Peek()
	fmt.Println("item1:", *item1)
	s.Print()

	s.Push("hello")
	s.Print()

	fmt.Println("is empty:", s.IsEmpty())
	fmt.Println("size:", s.Size())

	s.Push(t)
	s.Print()

	someItem1 := s.Pop()
	fmt.Printf("someItem1: %+v\n", *someItem1)
	s.Print()
}
