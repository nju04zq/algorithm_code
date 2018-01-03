package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// There are a row of n houses, each house can be painted with one of the k colors. The cost of painting each house with a certain color is different. You have to paint all the houses such that no two adjacent houses have the same color.
//
// The cost of painting each house with a certain color is represented by a n x k cost matrix. For example, costs[0][0] is the cost of painting house 0 with color 0; costs[1][2]is the cost of painting house 1 with color 2, and so on... Find the minimum cost to paint all houses.
//
// Note:
// All costs are positive integers.
//
// Follow up:
// Could you solve it in O(nk) runtime?

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func calc(a, res []int) {
	for i := 0; i < len(a); i++ {
		if i == 0 {
			res[i] = math.MaxInt32
		} else {
			res[i] = min(res[i-1], a[i-1])
		}
	}
	x := math.MaxInt32
	for i := len(a) - 1; i >= 0; i-- {
		if i == len(a)-1 {
			continue
		} else {
			x = min(x, a[i+1])
			res[i] = min(x, res[i])
		}
	}
}

func paint(cost [][]int) int {
	if len(cost) == 0 || len(cost[0]) == 0 {
		return 0
	}
	m, n := len(cost), len(cost[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	temp := make([]int, n)
	for i := 0; i < m; i++ {
		if i == 0 {
			for j := 0; j < n; j++ {
				temp[j] = cost[i][j]
			}
		} else {
			for j := 0; j < n; j++ {
				temp[j] = cost[i][j] + dp[i-1][j]
			}
		}
		calc(temp, dp[i])
	}
	total := math.MaxInt32
	for i := 0; i < n; i++ {
		total = min(total, dp[m-1][i])
	}
	return total
}

func dfs(cost [][]int, start int, prev int) int {
	m, n := len(cost), len(cost[0])
	if start == m {
		return 0
	}
	total := math.MaxInt32
	for i := 0; i < n; i++ {
		if i == prev {
			continue
		}
		subcost := dfs(cost, start+1, i)
		total = min(total, cost[start][i]+subcost)
	}
	return total
}

func bf(cost [][]int) int {
	if len(cost) == 0 || len(cost[0]) == 0 {
		return 0
	}
	return dfs(cost, 0, -1)
}

func MakeRandInt() int {
	maxNum := 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()%maxNum + 1
}

func MakeRandArray(len int) []int {
	maxElement := 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int()%maxElement + 1
	}
	return a
}

func testCost() {
	m, n := MakeRandInt(), 3
	cost := make([][]int, m)
	for i := 0; i < m; i++ {
		cost[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		a := MakeRandArray(m)
		for i := 0; i < m; i++ {
			cost[i][j] = a[i]
		}
	}
	res := paint(cost)
	ans := bf(cost)
	if res != ans {
		panic(fmt.Errorf("%v, get %d, ans %d\n", cost, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		fmt.Printf("\r%d", i)
		testCost()
	}
	fmt.Println()
}
