package main

import (
	"class/pkg/queue"
	"fmt"
)

/*
cs courses:
    0       1       2        3         4         5          6
["algos", "os", "graphics", "ml", "compilers", "arch", "dist sys"]

Pre reqs (pre-req, course)
[(5, 6), (0, 4), (1, 2), (1, 6), (0, 1), (4, 1), (3, 2)]

what order to take the courses?

write function that takes in these two lists and returns how many prereqs each course has

for each course: what depends on it?
*/

type Course struct {
	id     int
	prereq int
}

func getPreReqs(courseIDs []int, courses []Course) []int {

	coursesToTake := queue.New()
	prereqCount := make(map[int]int)
	depCourses := make(map[int][]int)

	for _, v := range courseIDs {
		depCourses[v] = []int{}
		prereqCount[v] = 0
	}

	for _, course := range courses {
		id := course.id
		pre := course.prereq

		prereqCount[id]++

		dependencies := depCourses[pre]
		if !includes(dependencies, id) {
			dependencies = append(dependencies, id)
			depCourses[pre] = dependencies
		}
	}

	// add courses to the queue that do not have any prerequisites
	for course, prereqs := range prereqCount {
		if prereqs == 0 {
			coursesToTake.Enqueue(course)
		}
	}

	courseOrder := []int{}
	for coursesToTake.Size() > 0 {
		// take the course
		c := coursesToTake.Dequeue()
		cp := *c
		course := cp.(int)
		courseOrder = append(courseOrder, course)

		// take the courses that now have all prereqs met
		for _, dependent := range depCourses[course] {
			// decrement prereqs for the course, because we have taken one of the prereqs
			prereqCount[dependent] -= 1
			if prereqCount[dependent] == 0 {
				coursesToTake.Enqueue(dependent)
			}
		}
	}
	return courseOrder
}

func includes(in []int, e int) bool {
	for _, v := range in {
		if v == e {
			return true
		}
	}
	return false
}

func main() {
	courses := []int{0, 1, 2, 3, 4, 5, 6}

	// Pre reqs (pre-req, course)
	// [(5, 6), (0, 4), (1, 2), (1, 6), (0, 1), (4, 1), (3, 2)]
	reqs := []Course{
		{
			prereq: 5,
			id:     6,
		},
		{
			prereq: 0,
			id:     4,
		},
		{
			prereq: 1,
			id:     2,
		},
		{
			prereq: 1,
			id:     6,
		},
		{
			prereq: 0,
			id:     1,
		},
		{
			prereq: 4,
			id:     1,
		},
		{
			prereq: 3,
			id:     2,
		},
	}
	fmt.Println(getPreReqs(courses, reqs))
}
