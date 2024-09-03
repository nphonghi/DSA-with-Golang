package main

import "fmt"

func binarySearch(arr []int, left, right, x int) int {
	if right >= left {
		mid := left + (right - left) / 2

		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			return binarySearch(arr, left, mid - 1, x)
		}
		return binarySearch(arr, mid + 1, right, x)
	}
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	x := 4
	result := binarySearch(arr, 0, len(arr) - 1, x)
	if result != -1 {
		fmt.Printf("%d nằm tại index %d\n trong mảng", x, result)
	} else {
		fmt.Printf("%d không tồn tại trong mảng", x)
	}
}
