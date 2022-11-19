package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Adjacency List
// 1 - 2, 4
// 2 - 1
// 3 - 1, 4, 5

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      string
	weight   int
	adjacent []*Vertex
}

func NewGraphFromRules(rules []string) *Graph {
	graph := &Graph{}
	for _, rule := range rules {
		luggages := strings.Split(rule, " contain ")
		mainVertex := luggages[0]
		graph.AddVertex(mainVertex, 1)
		for _, bags := range strings.Split(luggages[1], ", ") {
			if bags == "no other" {
				continue
			} else if val, err := strconv.Atoi(string(bags[0])); err == nil {
				graph.AddEdge(mainVertex, strings.TrimSpace(bags[1:]), val)
			}
		}
	}

	return graph
}

// AddVertex
func (g *Graph) AddVertex(key string, weight int) {
	if g.contains(key) {
		log.Fatalf("The key: %s already exist in the Graph", key)
	}
	g.vertices = append(g.vertices, &Vertex{key: key, weight: weight})
}

// AddEdge
func (g *Graph) AddEdge(fromKey, toKey string, weight int) {
	for _, v := range g.vertices {
		if v.key == fromKey {
			v.adjacent = append(v.adjacent, &Vertex{key: toKey, weight: weight})
		}
	}
}

// Get All Vertices Key
func (g *Graph) GetAllVerticesKeys() []string {
	vertices := []string{}
	for _, v := range g.vertices {
		vertices = append(vertices, v.key)
	}

	return vertices
}

func (g *Graph) GetPathWeight(start, end string) int {
	return -1
}

func (g *Graph) IsReachable(start, end string) bool {
	if start == end {
		return true
	}
	visited := map[string]bool{start: true}

	queue := []string{start}
	for len(queue) != 0 {
		s := queue[0]
		queue = queue[1:]
		i := findIndex(g.vertices, s)
		if i == -1 {
			queue = queue[1:]
			continue
		}
		for _, v := range g.vertices[i].adjacent {
			if v.key == end {
				return true
			}

			if !visited[v.key] {
				visited[v.key] = true
				queue = append(queue, v.key)
			}
		}
	}

	return false
}

func findIndex(vertices []*Vertex, key string) int {
	for i, v := range vertices {
		if v.key == key {
			return i
		}
	}

	return -1
}

func (g *Graph) contains(key string) bool {
	for _, v := range g.vertices {
		if v.key == key {
			return true
		}
	}

	return false
}

// Print
func (g *Graph) Print() {
	for i, v := range g.vertices {
		fmt.Printf("Vertex[%d,%s]: ", i, v.key)
		for _, av := range v.adjacent {
			fmt.Printf("%s=%d ", av.key, av.weight)
		}
		fmt.Println()
	}
}

func (g *Graph) GoString() string {
	return fmt.Sprintf("Graph{vertices: %#v}", g.vertices)
}

func (v *Vertex) GoString() string {
	return fmt.Sprintf("Vertex{key: %s, weight: %d, adjacent: %#v}", v.key, v.weight, v.adjacent)
}
