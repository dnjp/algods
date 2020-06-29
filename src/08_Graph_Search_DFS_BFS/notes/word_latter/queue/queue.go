package queue

import (
	"word_latter/stack"
)

// IQueue is the interface for a queue
type IQueue interface {

	// Enqueue adds a new item to the rear of the queue
	Enqueue(interface{})

	// Dequeue removes an element from the front of the queue
	Dequeue() *interface{}

	// IsEmpty returns whether or not the queue is empty
	IsEmpty() bool

	// Size returns the number of elements in queue
	Size() int
}

/*
Queue uses two underlying stacks in order to implement the queue interface. This implementation
moves items from the primary stack to the cache during insertion, and then re-adds old items after
inserting the new item in order to place the new item at the rear of the stack. This makes the cost
of additions O(n), while removals are O(1) because the item can simply be popped off the front of
the primary stack.
*/
type Queue struct {
	s     *stack.Stack
	cache *stack.Stack
}

// New constructs a new Queue
func New() *Queue {
	return &Queue{
		s:     stack.New(),
		cache: stack.New(),
	}
}

// IsEmpty returns whether or not the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.s.IsEmpty()
}

// Enqueue adds a new item to the rear of the queue
func (q *Queue) Enqueue(item interface{}) {

	// move items from primary stack to cache
	for q.s.Size() > 0 {
		item := q.s.Pop()
		if item != nil {
			q.cache.Push(*item)
		}
	}

	// add the new item to front of primary stack
	q.s.Push(item)

	// move items back from cache to primary stack
	for q.cache.Size() > 0 {
		item := q.cache.Pop()
		if item != nil {
			q.s.Push(*item)
		}
	}
}

// Dequeue removes an element from the front of the queue
func (q *Queue) Dequeue() *interface{} {
	item := q.s.Pop()
	return item
}

// Size returns the number of elements in queue
func (q *Queue) Size() int {
	return q.s.Size()
}
