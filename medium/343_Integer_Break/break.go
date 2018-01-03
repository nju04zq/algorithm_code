package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		maxVal := 0
		for j := 1; j <= i-1; j++ {
			maxVal = max(maxVal, j*max(i-j, dp[i-j]))
		}
		dp[i] = maxVal
	}
	return dp[n]
}

func testBreak(n int) {
	fmt.Printf("%d, get %d\n", n, integerBreak(n))
}

func main() {
	for i := 2; i <= 20; i++ {
		testBreak(i)
	}
}
