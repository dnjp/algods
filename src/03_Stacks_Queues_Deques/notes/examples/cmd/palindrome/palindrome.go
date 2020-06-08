package main

import (
	"examples/deque"
)

func isPalindrome(input string) bool {
	d := deque.New()
	for _, r := range input {
		d.AddFront(r)
	}

	for d.Size() > 1 {
		first := d.RemoveFront()
		last := d.RemoveRear()
		if first == nil || last == nil {
			panic("first or last element is nil")
		}
		if *first != *last {
			return false
		}
	}
	return true
}
