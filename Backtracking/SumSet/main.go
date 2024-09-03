package main

import "fmt"

func subsetSum(set []int, n, sum int) bool {
	if sum == 0 {
		return true
	}
	if n == 0 && sum != 0 {
		return false
	}
	if set[n-1] > sum {
		return subsetSum(set, n-1, sum)
	}
	return subsetSum(set, n-1, sum) || subsetSum(set, n-1, sum-set[n-1])
}

func main() {
	set := []int{3, 34, 4, 12, 5, 2}
	sum := 10
	n := len(set)
	if subsetSum(set, n, sum) {
		fmt.Println("Found a subset with given sum")
	} else {
		fmt.Println("No subset with given sum")
	}
}
