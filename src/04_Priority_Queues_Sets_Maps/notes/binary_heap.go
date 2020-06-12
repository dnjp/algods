package main

import "fmt"

type BinaryHeap struct {
	items []int
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{
		items: []int{0},
	}
}

func (b *BinaryHeap) Len() int {
	return len(b.items) - 1
}

func (b *BinaryHeap) percolateUp() {
	i := b.Len()
	for i/2 > 0 {
		if b.items[i] < b.items[i/2] {
			b.items[i/2], b.items[i] = b.items[i], b.items[i/2]
		}

		i = i / 2
	}
}

func (b *BinaryHeap) percolateDown(i int) {
	for i*2 <= b.Len() {
		mc := b.minChild(i)
		if b.items[i] > b.items[mc] {
			b.items[i], b.items[mc] = b.items[mc], b.items[i]
		}
		i = mc
	}
}

// - Because the tree is comlete, the left child of a parent (p) is the node that
//   is found in position 2p in the list. The right child of the parent is at
//   position 2p + 1 in the list.
// - To find the parent of any node in the tree, we can use integer division
//   (discarding remainder). Given that a node is at position n in the list, the
//   parent is at position n/2.

// Note the 2p and 2p+1 relationship between parent and children. The list
// representation of the tree, along with the full structure property allows us to
// efficiently traverse a complete binary tree using only a few mathematical
// operations.

// parent = p
//
func (b *BinaryHeap) minChild(i int) int {
	if i*2+1 > b.Len() {
		return i * 2
	}

	if b.items[i*2] < b.items[i*2+1] {
		return i * 2
	}

	return i*2 + 1
}

func (b *BinaryHeap) pop() int {
	length := b.Len()
	item := b.items[length]
	b.items = b.items[:length-1]
	return item
}

func (b *BinaryHeap) DeleteMin() int {
	length := b.Len()
	retVal := b.items[1]
	b.items[1] = b.items[length]
	b.pop()
	b.percolateDown(1)
	return retVal
}

func (b *BinaryHeap) Insert(elem int) {
	b.items = append(b.items, elem)
	b.percolateUp()
}

func (b *BinaryHeap) BuildHeap(elems []int) {
	i := b.Len() / 2
	items := []int{0}
	items = append(items, elems...)
	b.items = items

	// starting at the last element in the array and moving down
	for i > 0 {
		b.percolateDown(i)
		i = i - 1
	}
}

func (b *BinaryHeap) Print() {
	fmt.Printf("%+v\n", b.items)
}

func main() {
	b := NewBinaryHeap()
	nums := []int{9, 6, 5, 2, 3}
	b.BuildHeap(nums)
	b.Print()
}
