package main

import "fmt"

func swap(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func InsertSort(arr []int, n int) {
	for i := 1; i < n; i++ {
		key := arr[i] 
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j + 1] = arr[j]
			j--
		}
		arr[j + 1] = key
	}
}

func main() {
	var n int
	fmt.Printf("Nhập số phần tử của mảng: ")
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("Nhập giá trị cho phần tử ", i + 1, ": ")
		fmt.Scanln(&arr[i])
	}
	
	InsertSort(arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}