package main

import (
	"fmt"
)

// Định nghĩa cấu trúc Node
type Node struct {
	data int
	next *Node
}

// Chèn nút vào đầu danh sách
func insertToHead(head *Node, data int) *Node {
	newNode := &Node{data: data}
	newNode.next = head
	head = newNode
	return head
}

// Chèn nút sau nút hiện tại
func insertAfter(cur *Node, data int) *Node {
	newNode := &Node{data: data}
	newNode.next = cur.next
	cur.next = newNode
	return newNode
}

// Chèn nút trước nút hiện tại
func insertBefore(head *Node, cur *Node, data int) *Node {
	newNode := &Node{data: data}
	if head == nil {
		head = newNode
	} else if cur == head {
		newNode.next = cur
		head = newNode
	} else {
		prev := head
		for prev.next != cur {
			prev = prev.next
		}
		prev.next = newNode
		newNode.next = cur
	}
	return newNode
}

// Chèn nút vào cuối danh sách
func insertToLast(head *Node, data int) *Node {
	newNode := &Node{data: data}
	if head == nil {
		head = newNode
	} else {
		last := head
		for last.next != nil {
			last = last.next
		}
		last.next = newNode
	}
	return head
}

func main() {
	var n, m int // số nút và dữ liệu trong 1 nút

	fmt.Print("Nhập vào số lượng nút trong danh sách liên kết đơn: ")
	fmt.Scan(&n)

	var head, nodeNew, tail *Node

	if n > 0 {
		for i := 1; i <= n; i++ {
			fmt.Printf("Nhập dữ liệu cho nút thứ %d: ", i)
			fmt.Scan(&m)
			nodeNew = &Node{data: m}
			if i == 1 {
				head = nodeNew
				tail = nodeNew
			} else {
				tail.next = nodeNew
				tail = nodeNew
			}
		}
		tail.next = nil
	}

	// head = insertToHead(head, 10)
	// insertAfter(head.next, 20)
	// insertBefore(head, head.next, 5)
	insertToLast(head, 30)

	fmt.Print("Danh sách sau khi thêm là: ")
	for cur := head; cur != nil; cur = cur.next {
		fmt.Print(cur.data, " ")
	}
}
