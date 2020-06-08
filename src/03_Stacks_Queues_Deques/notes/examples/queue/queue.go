package queue

type IQueue interface {
	// constructs a new empty queue
	New() *IQueue

	// Adds a new item to the rear of the queue
	Enqueue(interface{})

	// Removes the front item from the queue
	Dequeue() *interface{}

	// Returns whether or not the queue is empty
	IsEmpty() bool

	// Returns the number of items in the queue
	Size() int
}

type Queue struct {
	items []interface{}
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, 0)
	copy(q.items[1:], q.items[0:])
	q.items[0] = item
}

func (q *Queue) Dequeue() *interface{} {
	length := len(q.items)
	if length > 0 {
		item := q.items[length-1]
		q.items = q.items[:length-1]
		return &item
	}
	return nil
}

func (q *Queue) Size() int {
	return len(q.items)
}
