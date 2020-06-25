package main

import (
	"strings"
)

type Operator func(int, int) int

var OPERATORS = map[rune]Operator{
	'+': func(a int, b int) int {
		return a + b
	},
	'-': func(a int, b int) int {
		return a - b
	},	
	'*': func(a int, b int) int {
		return a * b
	},
	'/': func(a int, b int) int {
		return a / b
	},		
}

var LeftParen rune = '('
var RightParen rune = ')'

type Node struct {
	Token rune
	Left *Node
	Right *Node
}

func NewNode(token rune) *Node {
	return &Node{Token: token}
}

func buildParseTree(expression string) {
	exp := strings.ReplaceAll(expression, " ", "")
	for _, t := range exp {
		if token == LeftParen {
			
		}
	}
}

func main() {
	buildParseTree("(3 + (4 * 5))") // 23
}