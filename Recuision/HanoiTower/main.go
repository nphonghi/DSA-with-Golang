package main

import "fmt"

func towerOfHanoi(n int, fromRod string, toRod string, auxRod string) {
    if n == 1 {
        fmt.Printf("Move disk 1 from %s to %s\n", fromRod, toRod)
        return
    }
    towerOfHanoi(n-1, fromRod, auxRod, toRod)
    fmt.Printf("Move disk %d from %s to %s\n", n, fromRod, toRod)
    towerOfHanoi(n-1, auxRod, toRod, fromRod)
}

func main() {
    n := 3
    towerOfHanoi(n, "A", "C", "B")
}