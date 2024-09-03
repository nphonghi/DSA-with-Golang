package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Front() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Front()) // 1
	fmt.Println(queue.Dequeue()) // 1
	fmt.Println(queue.Front()) // 2
	fmt.Println(queue.Dequeue()) // 2
	fmt.Println(queue.Front()) // 3
	fmt.Println(queue.Dequeue()) // 3
	fmt.Println(queue.IsEmpty()) // true
}
