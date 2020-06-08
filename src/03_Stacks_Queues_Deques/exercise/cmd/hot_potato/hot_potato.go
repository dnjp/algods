package main

import (
	"exercise/queue"
	"fmt"
)

func main() {
	out := hotPotato([]string{"Bill", "David", "Susan", "Jane", "Kent", "Brad"}, 9)
	fmt.Println(out)
}

func hotPotato(names []string, num int) string {
	q := queue.New()
	for _, name := range names {
		q.Enqueue(name)
	}

	for q.Size() > 1 {
		for i := 0; i < num; i++ {
			val := q.Dequeue()
			if val == nil {
				panic("nothing left in queue")
			}
			q.Enqueue(*val)
		}
		q.Dequeue()
	}

	out := q.Dequeue()
	if out == nil {
		panic("nothing left in queue")
	}
	val := *out
	valStr := val.(string)
	return valStr
}
