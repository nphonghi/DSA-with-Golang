package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents an edge with a destination and weight
type Edge struct {
	dest, weight int
}

// Graph represents a graph using an adjacency list
type Graph struct {
	adjList map[int][]Edge
}

// NewGraph creates a new Graph instance
func NewGraph() *Graph {
	return &Graph{adjList: make(map[int][]Edge)}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(src, dest, weight int) {
	g.adjList[src] = append(g.adjList[src], Edge{dest, weight})
}

// PriorityQueue implements a priority queue for Dijkstra's algorithm
type PriorityQueue struct {
	items []PQItem
}

// PQItem represents an item in the priority queue
type PQItem struct {
	vertex, distance int
}

// Len returns the length of the priority queue
func (pq PriorityQueue) Len() int {
	return len(pq.items)
}

// Less compares two items in the priority queue
func (pq PriorityQueue) Less(i, j int) bool {
	return pq.items[i].distance < pq.items[j].distance
}

// Swap swaps two items in the priority queue
func (pq PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// Push adds an item to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	pq.items = append(pq.items, x.(PQItem))
}

// Pop removes and returns the item with the highest priority
func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.items = old[0 : n-1]
	return item
}

// Dijkstra finds the shortest paths from the source vertex to all other vertices
func Dijkstra(g *Graph, source int) map[int]int {
	dist := make(map[int]int)
	for v := range g.adjList {
		dist[v] = math.MaxInt64
	}
	dist[source] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, PQItem{vertex: source, distance: 0})

	for pq.Len() > 0 {
		u := heap.Pop(pq).(PQItem).vertex
		for _, edge := range g.adjList[u] {
			v := edge.dest
			weight := edge.weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, PQItem{vertex: v, distance: dist[v]})
			}
		}
	}

	return dist
}

func main() {
	// Create a new graph
	g := NewGraph()

	// Add edges to the graph
	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 1)
	g.AddEdge(2, 1, 2)
	g.AddEdge(1, 3, 1)
	g.AddEdge(2, 3, 5)
	g.AddEdge(3, 4, 3)

	// Source vertex
	source := 0

	// Get the shortest paths using Dijkstra's algorithm
	distances := Dijkstra(g, source)

	// Print the shortest paths
	fmt.Println("Shortest distances from source vertex:", source)
	for vertex, distance := range distances {
		fmt.Printf("Vertex %d: %d\n", vertex, distance)
	}
}
