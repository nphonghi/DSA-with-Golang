package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
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

		// duyệt danh sách
		cur := head
		for cur != nil {
			fmt.Print(cur.data, " ")
			cur = cur.next
		}
	} else {
		fmt.Println("Chạy lại chương trình!")
	}
}
