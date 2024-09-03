package main

import "fmt"

type TP interface {
	int64 | float64
}

func partitionLomuto[T TP](arr []T, left, right int) int {
	pivot := arr[right]
	i := left - 1
	for k := left; k < right; k++ {
		if arr[k] <= pivot {
			i++
			arr[i], arr[k] = arr[k], arr[i]
		}
	}
	i++
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

func QuickSort[T TP](arr []T, left, right int) {
	if left >= right {
		return
	}
	pivot := partitionLomuto(arr, left, right)
	QuickSort(arr, left, pivot - 1)
	QuickSort(arr, pivot + 1, right)
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

	QuickSort(arr, 0, n - 1)

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}