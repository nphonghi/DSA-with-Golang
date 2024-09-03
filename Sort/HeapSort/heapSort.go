package main

import (
	"fmt"
)

type TP interface {
	int64 | float64
}

func buildMaxHeap[T TP](arr []T) {
	n := len(arr)
	for i := n / 2 - 1; i >= 0; i-- {
		maxHeappify(arr, i, n)
	}
}

func maxHeappify[T TP](arr []T, i, n int) {
	subLeft := 2*i + 1
	subRight := 2*i + 2
	var largest int
	if subLeft <= n && arr[subLeft] > arr[i] {
		largest = subLeft
	} else {
		largest = i
	}
	if subRight <= n && arr[subRight] > arr[largest] {
		largest = subRight
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeappify(arr, largest, n)
	}
}

func heapSort[T TP](arr []T) {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeappify(arr, 0, i - 1)
	}
}

func main() {
	var n int
	fmt.Print("Nhập số phần tử muốn sắp xếp: ")
	fmt.Scanln(&n)
	arr := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Print("Nhập vào phần tử thứ ", i + 1, ": ")
		fmt.Scanln(&arr[i])
	}

	heapSort(arr)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}