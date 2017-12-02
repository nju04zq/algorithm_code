package main

import "fmt"

func make2dArray(m, n int) [][]int {
	a := make([][]int, m)
	for i, _ := range a {
		a[i] = make([]int, n)
	}
	return a
}

func numDistinct(s string, t string) int {
	m, n := len(t), len(s)
	dp := make2dArray(m+1, n+1)
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	//dump(t, s, dp)
	return dp[m][n]
}

func dump(t, s string, dp [][]int) {
	fmt.Printf("  ")
	fmt.Printf("  ")
	for i, _ := range s {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
	for i := 0; i <= len(t); i++ {
		if i == 0 {
			fmt.Printf("  ")
		} else {
			fmt.Printf("%c ", t[i-1])
		}
		for j, _ := range dp[i] {
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Println()
	}
}

func testNumDistinct(s, t string) {
	fmt.Printf("t %q, s %q, get %d\n", t, s, numDistinct(s, t))
}

func main() {
	testNumDistinct("rabbbit", "rabbit")
	testNumDistinct("ccc", "c")
}
