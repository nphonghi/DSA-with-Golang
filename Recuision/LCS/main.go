package main

import "fmt"

func lcs(X string, Y string, m int, n int) int {
    if m == 0 || n == 0 {
        return 0
    }
    if X[m-1] == Y[n-1] {
        return 1 + lcs(X, Y, m-1, n-1)
    } else {
        return max(lcs(X, Y, m, n-1), lcs(X, Y, m-1, n))
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    X := "AGGTAB"
    Y := "GXTXAYB"
    m := len(X)
    n := len(Y)
    fmt.Printf("Length of LCS is %d\n", lcs(X, Y, m, n))
}