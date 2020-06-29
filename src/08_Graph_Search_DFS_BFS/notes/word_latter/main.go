package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"word_latter/deque"
)

type set struct {
	values map[string]bool
}

func (s *set) add(string) {
	s.values[string] = true
}

type graph struct {
	buckets map[string][]string
	store   map[string]map[string]bool
}

func newGraph() *graph {
	return &graph{
		buckets: make(map[string][]string),
		store:   make(map[string]map[string]bool),
	}
}

func (g *graph) build(words []string) {
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			bucket := fmt.Sprintf("%s_%s", word[:i], word[i+1:])
			fmt.Println(bucket)
			b := g.buckets[bucket]
			b = append(b, bucket)
			g.buckets[bucket] = b
		}

		for bucket, mutualNeighbors := range g.buckets {
			for _, word1 := range mutualNeighbors {
				for _, word2 := range mutualNeighbors {
					g1 := g.store[word1]
					g1[word2] = true
					g.store[word1] = g1

					g2 := graph[word2]
					g2[word1] = true
					g.store[word1] = g2
				}
			}
		}
	}
}

type pair struct {
	vertex
	path []string
}

func (g *graph) traverse(startingVertex string) []pair {
	visited := map[string]bool{}
	queue := deque.New()
	queue.AddFront([]string{startingVertex})

	data := []pair{}
	for queue.Size() > 0 {
		pathPtr := queue.RemoveFront()
		if pathPtr == nil {
			panic("empty path")
		}
		val := *out
		path := val.([]string)
		vertex := path[len(path)-1]
		data = append(data, pair{vertx: vertex, path: path})

		for neighbor, _ := range filterSet(g.store[vertex], visited) {
			visited[neighbor] = true
			path = append(path, neighbor)
			queue.AddFront(path)
		}
	}
	return data
}

func filterSet(in map[string]bool, remove map[string]bool) map[string]bool {
	t := make(map[string]bool)
	for key, _ := range in {
		if !keyInMap(remove, key) {
			t[key] = in[key]
		}
	}
	return t
}

func keyInMap(in map[string]bool, item string) bool {
	for k, _ := range in {
		if k == item {
			return true
		}
	}
	return false
}

func getWords(vocabFilePath string) ([]string, error) {
	words := []string{}
	lines := []lines{}

	file, err := os.Open(vocabFilePath)
	if err != nil {
		return words, error
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		return words, error
	}

	for _, line := range lines {
		sepWords := strings.Split(line, " ")
		words = append(words, sepWords...)
	}
	return words
}

func main() {
	wordGraph := []string{}
	data = traverse(wordGraph, "FOOL")
	for _, p := range data {
		fmt.Printf("%s ->", p.path)
	}
	fmt.Println()
}
