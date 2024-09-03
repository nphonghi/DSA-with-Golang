package main

import (
	"container/heap"
	"fmt"
	"math"
)

// Point represents a point in the grid
type Point struct {
	x, y int
}

// Node represents a node in the pathfinding
type Node struct {
	point    Point
	cost     int
	priority int
	parent   *Node
	index    int
}

// PriorityQueue implements a priority queue
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil // avoid memory leak
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

func (pq *PriorityQueue) update(node *Node, point Point, cost, priority int) {
	node.point = point
	node.cost = cost
	node.priority = priority
	heap.Fix(pq, node.index)
}

// heuristic estimates the cost from a point to the goal
func heuristic(a, b Point) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

// aStar finds the shortest path using the A* algorithm
func aStar(start, goal Point, grid [][]int) []Point {
	directions := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)
	startNode := &Node{point: start, cost: 0, priority: 0}
	heap.Push(&openSet, startNode)

	cameFrom := make(map[Point]*Node)
	costSoFar := make(map[Point]int)
	cameFrom[start] = nil
	costSoFar[start] = 0

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*Node)

		if current.point == goal {
			path := []Point{}
			for current != nil {
				path = append([]Point{current.point}, path...)
				current = current.parent
			}
			return path
		}

		for _, direction := range directions {
			next := Point{current.point.x + direction.x, current.point.y + direction.y}

			if next.x >= 0 && next.x < len(grid) && next.y >= 0 && next.y < len(grid[0]) && grid[next.x][next.y] == 0 {
				newCost := costSoFar[current.point] + 1
				if cost, ok := costSoFar[next]; !ok || newCost < cost {
					costSoFar[next] = newCost
					priority := newCost + heuristic(next, goal)
					nextNode := &Node{point: next, cost: newCost, priority: priority, parent: current}
					heap.Push(&openSet, nextNode)
					cameFrom[next] = current
				}
			}
		}
	}
	return nil
}

func main() {
	grid := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 0, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 1, 1, 1, 1, 0, 1, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 1, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	start := Point{0, 0}
	goal := Point{9, 9}

	path := aStar(start, goal, grid)
	if path != nil {
		fmt.Println("Path found: ")
		for _, p := range path {
			fmt.Printf("(%d, %d) ", p.x, p.y)
		}
		fmt.Println()
	} else {
		fmt.Println("No path found")
	}
}
