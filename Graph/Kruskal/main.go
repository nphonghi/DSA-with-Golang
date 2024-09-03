package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	src, dest, weight int
}

// UnionFind is a data structure to keep track of disjoint sets
type UnionFind struct {
	parent, rank []int
}

// NewUnionFind creates a new UnionFind instance with 'n' elements
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

// Find finds the root of the set containing 'x'
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

// Union joins the sets containing 'x' and 'y'
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

// Kruskal finds the Minimum Spanning Tree (MST) using Kruskal's algorithm
func Kruskal(n int, edges []Edge) []Edge {
	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	uf := NewUnionFind(n)
	mst := []Edge{}

	for _, edge := range edges {
		if uf.Find(edge.src) != uf.Find(edge.dest) {
			mst = append(mst, edge)
			uf.Union(edge.src, edge.dest)
		}
	}

	return mst
}

func main() {
	// Define the graph edges with weights
	edges := []Edge{
		{0, 1, 4},
		{0, 2, 4},
		{1, 2, 2},
		{1, 3, 5},
		{2, 3, 5},
		{2, 4, 7},
		{3, 4, 6},
	}

	// Number of vertices in the graph
	n := 5

	// Get the Minimum Spanning Tree using Kruskal's algorithm
	mst := Kruskal(n, edges)

	// Print the edges in the MST
	fmt.Println("Edges in the Minimum Spanning Tree:")
	for _, edge := range mst {
		fmt.Printf("%d -- %d == %d\n", edge.src, edge.dest, edge.weight)
	}
}
