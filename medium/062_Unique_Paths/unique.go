package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}
	return dp[n-1]
}

func main() {
	m, n := 3, 7
	fmt.Printf("m %d, n %d, get %d\n", m, n, uniquePaths(m, n))
}
