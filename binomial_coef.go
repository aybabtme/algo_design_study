package main

import (
	"fmt"
)

func Choose(n, k int) int {
	memo := make([][]int, n)

	for i := 0; i < n; i++ {
		memo[i] = make([]int, k)

		for j := 0; j < min(i, k); j++ {
			if j == 0 || j == i {
				memo[i][j] = 1
			} else {
				memo[i][j] = memo[i-1][j-1] + memo[i-1][j]
			}
		}
	}
	return memo[n-1][k-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	for n := 1; n < 20; n++ {
		for k := 1; k < n; k++ {
			fmt.Printf("%d choose %d = %d\n", n, k, Choose(n, k))
		}
	}
}
