package main

import "fmt"

func bubleSort(arr []int, n int) {
	
	for i := 0; i < n - 1; i++ {
		swapped := false
		for j := 0; j < n - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
}

func main() {
	var n int
	fmt.Print("Nhập số lượng phần tử muốn sắp xếp: ")
	fmt.Scanln(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Print("Nhập phần tử thứ ", i + 1, ": ")
		fmt.Scanln(&arr[i])
	}

	bubleSort(arr, n)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}