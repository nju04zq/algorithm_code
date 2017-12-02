package main

import "fmt"

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= m; i++ {
		for j := n; j >= 1; j-- {
			if t[j-1] == s[i-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n]
}

func testNumDistinct(s, t string) {
	fmt.Printf("t %q, s %q, get %d\n", t, s, numDistinct(s, t))
}

func main() {
	testNumDistinct("rabbbit", "rabbit")
	testNumDistinct("ccc", "c")
}
