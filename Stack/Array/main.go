package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item
}

func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.items[len(s.items) - 1]
}

func (s *Stack) Top() int {
	return s.Peek()
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
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
