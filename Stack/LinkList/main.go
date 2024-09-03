package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(value int) {
	newNode := &Node{
		value: value,
		next:  s.top,
	}
	s.top = newNode
}

func (s *Stack) Pop() int {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return -1
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

func (s *Stack) Peek() int {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.top.value
}

func (s *Stack) Top() int {
	return s.Peek()
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Top()) // 3
	fmt.Println(stack.Pop()) // 3
	fmt.Println(stack.Top()) // 2
	fmt.Println(stack.Pop()) // 2
	fmt.Println(stack.Top()) // 1
	fmt.Println(stack.Pop()) // 1
	fmt.Println(stack.IsEmpty()) // true
}
