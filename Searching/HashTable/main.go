package main

import (
	"fmt"
)

type Node struct {
	Key, Value string
	Next       *Node
}

type HashMap struct {
	buckets []*Node
	size    int
}

func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func hashFunction(key string, size int) uint {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return uint(hash % size)
}

func (hm *HashMap) Insert(key string, value string) {
	index := hashFunction(key, hm.size)
	node := &Node{
		Key:   key,
		Value: value,
	}

	if hm.buckets[index] == nil {
		hm.buckets[index] = node
	} else {
		current := hm.buckets[index]
		for current != nil {
			if current.Key == key {
				current.Value = value // Cập nhật giá trị tồn tại
				return
			}
			if current.Next == nil {
				current.Next = node
				return
			}
			current = current.Next
		}
	}
}

func (hm *HashMap) Search(key string) (string, bool) {
	index := hashFunction(key, hm.size)
	current := hm.buckets[index]
	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}
	return "", false
}

func (hm *HashMap) Delete(key string) {
	index := hashFunction(key, hm.size)
	current := hm.buckets[index]
	var prev *Node
	for current != nil {
		if current.Key == key {
			if prev == nil {
				hm.buckets[index] = current.Next
			} else {
				prev.Next = current.Next
			}
			return
		}
		prev = current
		current = current.Next
	}
}

func main() {
	// Tạo một bảng băm mới với kích thước 10
	hashMap := NewHashMap(10)

	// Chèn các cặp khóa tương ứng với giá trị vào bảng băm
	hashMap.Insert("Họ", "Nguyễn")
	hashMap.Insert("Tên", "Nhật Phong")

	// Tìm và in ra giá trị của khóa "Họ"
	value1, found1 := hashMap.Search("Họ")
	if found1 {
		fmt.Println("Tìm Họ:", value1)
	} else {
		fmt.Println("Họ không có trong hash map")
	}

	// Xóa khóa "Tên" khỏi bảng băm
	hashMap.Delete("Tên")

	// Tìm và in ra giá trị của khóa "Tên" sau khi xóa
	value2, found2 := hashMap.Search("Tên")
	if found2 {
		fmt.Println("Tìm Tên:", value2)
	} else {
		fmt.Println("Tên không có trong hash map")
	}
}