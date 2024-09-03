package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{value: value}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}

func (q *Queue) Dequeue() int {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return -1
	}
	value := q.front.value
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return value
}

func (q *Queue) Front() int {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.front.value
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
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
