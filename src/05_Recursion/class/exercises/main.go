package main

import (
	"fmt"
)

/*
- given a paper strip and it is n squares wide
- want to tile it
  - red = 1 square wide
  - blue = 2 squares wide

example (n = 7): [ red, blue, blue, blue, blue, red, red ]
                            [ blue, blue, blue, blue, red, red, red ] is different

given n, how many different patterns can you make?

n = 1:
  - only 1 pattern - red

n = 2:
  - 2 patterns - [red, red], [blue, blue]

n = 1, answer = 1
	[ red ]
n = 2, answer = 2
	[ red, red ]
	[  { blue } ]

n = 3, answer = 3
	[ red,  { blue } ]
	[  { blue }, red ]
	[ red, red, red ]

n = 4, answer = 5
	[ red, red, red, red ]
	[  { blue }, red, red ]
	[  red, { blue }, red ]
	[  red, red, { blue } ]
	[  { blue },  { blue } ]

n = 5, answer = 8
	[  { blue }, red, red ]
	[  red, { blue }, red ]
	[  red, red, { blue } ]
	[  { blue },  { blue } ]


KEY IDEA: answer = sum of previous two results
f(n) = f(n - 2) + f(n - 1) is valid for n > 2

Why does this pattern come up?

because of the size of the blocks?
if blue were equal to 3 spaces instead of 2, i'd suspect that the answer would still be the sum of the previous two.


*/

var BLUEWIDTH int = 2
var REDWIDTH int = 1


func tile(n int) int {
	if n <= 1 {
		return 1
	}
	tiles := make([]int, n)

	// blue := len(tiles) / BLUEWIDTH
	red := len(tiles) / REDWIDTH

	for i := 1; i < n; i+= 2 {
		tiles[i] = 1
		tiles[i-1] = 1
	}

	for i := 0; i < red; i++ {
		if tiles[i] != 1 {
			tiles[i] = 2
		}
	}

	return tiles[0]
}

func main() {
	fmt.Println("1", tile(1))
	fmt.Println("2", tile(2))
// 	fmt.Println("3", tile(3))
}