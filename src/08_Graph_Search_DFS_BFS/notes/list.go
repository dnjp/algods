package main

import (
	"fmt"
)

type Vertex struct {
	Key       int
	Neighbors map[int]Vertex
	Weight    float64
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key: key,
		Neighbors: make(map[int]Vertex),
	}
}

func (v *Vertex) AddNeighbor(neighbor Vertex, weight float64) {
	v.Neighbors[neighbor.Key] = neighbor
}

func (v *Vertex) ToString() string {
	return fmt.Sprintf("%d neighbors: %v", v.Key, v.getKeys())
}

func (v *Vertex) getKeys() []int {
	keys := []int{}
	for key, _ := range v.Neighbors {
		keys = append(keys, key)
	}
	return keys
}

func (v *Vertex) GetConnections() []int {
	return v.getKeys()
}

func (v *Vertex) GetWeight(neighbor Vertex) float64 {
	n := v.Neighbors[neighbor.Key]
	return n.Weight
}

type Graph struct {
	Verticies map[int]Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Verticies: make(map[int]Vertex),
	}
}

func (g *Graph) getKeys() []int {
	keys := []int{}
	for key, _ := range g.Verticies {
		keys = append(keys, key)
	}
	return keys
}

func (g *Graph) AddVertex(v Vertex) {
	g.Verticies[v.Key] = v
}

func (g *Graph) GetVertex(key int) Vertex {
	return g.Verticies[key]
}

func (g *Graph) Contains(key int) bool {
	for k, _ := range g.Verticies {
		if k == key {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(fromKey int, toKey int, weight float64) {
	if !g.Contains(fromKey) {
		v := NewVertex(fromKey)
		g.AddVertex(*v)
	}
	if !g.Contains(toKey) {
		v := NewVertex(toKey)
		g.AddVertex(*v)
	}
	n := g.Verticies[fromKey]
	n.AddNeighbor(g.Verticies[toKey], weight)
}

func (g *Graph) GetVertices() []int {
	return g.getKeys()
}

func main() {
	g := NewGraph()
	for i := 1; i <= 6; i++ {
		v := NewVertex(i)
		g.AddVertex(*v)	
	}
	for _, vert := range g.Verticies {
		fmt.Printf("%+v\n", vert)
	}
	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 5, 2)
	g.AddEdge(1, 2, 4)
	g.AddEdge(2, 3, 9)
	g.AddEdge(3, 4, 7)
	g.AddEdge(3, 5, 3)
	g.AddEdge(4, 0, 1)
	g.AddEdge(5, 4, 8)
	g.AddEdge(5, 2, 1)
	
	for _, v := range g.Verticies {
		for _, key := range v.GetConnections() {
			fmt.Printf("%d -> %d\n", v.Key, key)	
		}
	}
}