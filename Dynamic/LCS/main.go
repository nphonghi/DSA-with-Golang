package main

import "fmt"

func lcs(X string, Y string, m int, n int) int {
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for i := 0; i <= m; i++ {
        for j := 0; j <= n; j++ {
            if i == 0 || j == 0 {
                dp[i][j] = 0
            } else if X[i-1] == Y[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    return dp[m][n]
}

func main() {
    X := "AGGTAB"
    Y := "GXTXAYB"
    m := len(X)
    n := len(Y)
    fmt.Printf("Length of LCS is %d\n", lcs(X, Y, m, n))
}