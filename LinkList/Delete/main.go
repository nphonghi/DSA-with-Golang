package main

import (
	"fmt"
)

// Định nghĩa cấu trúc Node
type Node struct {
	data int
	next *Node
}

// Xóa nút khỏi danh sách
func deleteNode(head *Node, del *Node) *Node {
	if del == head {
		head = del.next
	} else {
		prev := head
		for prev.next != del {
			prev = prev.next
		}
		prev.next = del.next
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

	head = deleteNode(head, head)

	fmt.Print("Danh sách sau khi xóa là: ")
	for cur := head; cur != nil; cur = cur.next {
		fmt.Print(cur.data, " ")
	}
}
