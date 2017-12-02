package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minCut(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = n - i - 1
	}
	p := make([][]bool, n)
	for i, _ := range p {
		p[i] = make([]bool, n)
	}
	for i := n - 2; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i < 2 || p[i+1][j-1]) {
				p[i][j] = true
				if j == n-1 {
					dp[i] = 0
				} else {
					dp[i] = min(dp[i], dp[j+1]+1)
				}
			}
		}
	}
	return dp[0]
}

func testMinCut(s string) {
	fmt.Printf("On %q, get %d\n", s, minCut(s))
}

func main() {
	testMinCut("")
	testMinCut("aaa")
	testMinCut("aab")
}
