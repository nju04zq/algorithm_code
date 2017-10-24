// https://segmentfault.com/a/1190000003483697

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func alloc2dArray(m, n int) [][]int {
	a := make([][]int, m)
	for i, _ := range a {
		a[i] = make([]int, n)
	}
	return a
}

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	if k > len(prices)/2 {
		maxProfit := 0
		for i := 1; i < len(prices); i++ {
			maxProfit += max(0, prices[i]-prices[i-1])
		}
		return maxProfit
	}
	global := alloc2dArray(len(prices), k+1)
	local := alloc2dArray(len(prices), k+1)
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		for j := 1; j <= k; j++ {
			local[i][j] = max(global[i-1][j-1]+max(0, diff), local[i-1][j]+diff)
			global[i][j] = max(global[i-1][j], local[i][j])
		}
	}
	return global[len(prices)-1][k]
}

func maxProfitBf(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := (prices[j] - prices[i])
			maxProfit = max(maxProfit, profit)
		}
	}
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			for k := j + 1; k < len(prices); k++ {
				for l := k + 1; l < len(prices); l++ {
					profit := (prices[j] - prices[i])
					profit += (prices[l] - prices[k])
					maxProfit = max(maxProfit, profit)
				}
			}
		}
	}
	return maxProfit
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func testMaxProfit() bool {
	a := MakeRandArray()
	//a := []int{80, 60, 11, 87, 7, 92, 60}
	ans := maxProfitBf(a)
	res := maxProfit(2, a)
	if ans != res {
		fmt.Printf("Fail on %v, get %d, expect %d\n", a, res, ans)
		return false
	}
	return true
}

func main() {
	for i := 0; i < 100; i++ {
		res := testMaxProfit()
		if !res {
			break
		}
	}
}
