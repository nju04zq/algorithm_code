package main

import "fmt"
import "math"

const (
	INT_MAX_SQRT = 46340
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func sqrt(y int) int {
	x := min(y, INT_MAX_SQRT)
	for x*x > y {
		x = (x + y/x) / 2
	}
	return x
}

func dfs(squares []int, n, total int, res int) int {
	if n == 0 {
		return total
	} else if len(squares) == 0 {
		return -1
	}
	maxCur := n / squares[0]
	m := maxCur * squares[0]
	total += maxCur
	for {
		if total >= res {
			break
		}
		t := dfs(squares[1:], n-m, total, res)
		if t != -1 {
			res = min(res, t)
		}
		m -= squares[0]
		if m < 0 {
			break
		}
		total--
	}
	return res
}

func numSquaresDFS(n int) int {
	m := sqrt(n)
	squares := make([]int, m)
	for i := 0; i < m; i++ {
		squares[m-i-1] = (i + 1) * (i + 1)
	}
	return dfs(squares, n, 0, math.MaxInt32)
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func testSquares(n int) {
	res := numSquares(n)
	ans := numSquaresDFS(n)
	if ans != res {
		panic(fmt.Errorf("%d, get %d, expect %d\n", n, res, ans))
	}
}

func main() {
	for i := 1; i < 1000; i++ {
		testSquares(i)
	}
}
