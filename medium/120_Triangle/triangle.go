package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}
	n := len(triangle[m-1])
	if n != m {
		return 0
	}
	dp := make([]int, m)
	dp[0] = triangle[0][0]
	ans := dp[0]
	for i := 1; i < m; i++ {
		var prev, temp int
		for j := 0; j <= i; j++ {
			temp = dp[j]
			if j == i {
				dp[j] = prev
			} else if j > 0 {
				dp[j] = min(prev, dp[j])
			}
			dp[j] += triangle[i][j]
			prev = temp
			if j == 0 {
				ans = dp[j]
			} else {
				ans = min(ans, dp[j])
			}
		}
	}
	return ans
}

func testMinimumTotal(a [][]int) {
	fmt.Printf("%v, get %d\n", a, minimumTotal(a))
}

func main() {
	a := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	testMinimumTotal(a)
}
