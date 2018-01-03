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

func numSquares(n int) int {
	m := sqrt(n)
	squares := make([]int, m)
	for i := 0; i < m; i++ {
		squares[m-i-1] = (i + 1) * (i + 1)
	}
	return dfs(squares, n, 0, math.MaxInt32)
}

func testSquares(n int) {
	fmt.Printf("%d, get %d\n", n, numSquares(n))
}

func main() {
	testSquares(12)
	testSquares(13)
}
