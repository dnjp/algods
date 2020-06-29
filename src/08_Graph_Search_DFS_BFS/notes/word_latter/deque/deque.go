package deque

type IDeque interface {
	// Adds a new item to the front of the deque
	AddFront(interface{})

	// Adds a new item to the rear of the deque
	AddRear(interface{})

	// Removes the front item from the deque
	RemoveFront() *interface{}

	// Removes the rear item from the deque
	RemoveRear() *interface{}

	// Returns whether the deque is empty
	IsEmpty() bool

	// Returns the number of items in the deque
	Size() int
}

type Deque struct {
	items []interface{}
}

func New() *Deque {
	return &Deque{}
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) Size() int {
	return len(d.items)
}

func (d *Deque) AddFront(item interface{}) {
	d.items = append(d.items, item)
}

func (d *Deque) RemoveFront() *interface{} {
	length := len(d.items)
	if length > 0 {
		item := d.items[length-1]
		d.items = d.items[:length-1]
		return &item
	}
	return nil
}

func (d *Deque) AddRear(item interface{}) {
	d.items = append(d.items, 0)
	copy(d.items[1:], d.items[0:])
	d.items[0] = item
}

func (d *Deque) RemoveRear() *interface{} {
	item := d.items[0]
	d.items = d.items[1:]
	return &item
}
