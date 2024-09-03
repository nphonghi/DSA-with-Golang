package main

import (
	"fmt"
)

func selectionSort(arr []int, n int) {
	for index := range arr {
		minIndex := index
		for j := index + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[index], arr[minIndex] = arr[minIndex], arr[index]
	}
}

func main() {
	var n int
	fmt.Print("Nhập số phần tử muốn sắp xếp: ")
	fmt.Scanln(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Nhập vào phần tử thứ ", i + 1, ": ")
		fmt.Scanln(&arr[i])
	}

	selectionSort(arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}