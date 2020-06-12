package main

import (
	"fmt"
)

/*
- going to be given 1000 sorted iterators. An iterator looks like this:
iterator {
	next() int
	hasNext() bool
}
- you are guaranteed that they come back in order
- goal: return the lowest 100 globally


fn GetLowest100(iters) {
	// put pair of (value, iterator) in queue
	pq = PriorityQueue()

	minIters = []Iterator{}

	// inserts the lowest value for each iterator that has a next value
	// what if the first iterator has the lowest value?
	for iter in iters {
		if pq.Size() == 100 {
			break;
		}
		if !iter.hasNext() {
			continue
		}

		num = iter.next()
		if num < pq.FindMin() {
			pq.Enque(num)
		} else {
			continue
		}
		// what if the next iterators value is
	}

	arr := []int{}
	for i < 100 {
		num = pq.DeleteMin()
		arr = append(arr, num)
		// call next on iterators here and add to list
	}
	return arr
}

*/

func main() {
	fmt.Println("hello world")
}
