package main

import "fmt"

func knapsack(W int, weight []int, value []int, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				dp[i][w] = 0
			} else if weight[i-1] <= w {
				dp[i][w] = max(value[i-1] + dp[i-1][w-weight[i-1]], dp[i-1][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	weight := []int{1, 3, 4, 5}
	value := []int{1, 4, 5, 7}
	W := 7
	n := len(weight)
	fmt.Printf("Maximum value in Knapsack = %d\n", knapsack(W, weight, value, n))
}
