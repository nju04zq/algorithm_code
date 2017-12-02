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
	dp := alloc2dArray(k+1, len(prices))
	for i := 1; i <= k; i++ {
		tmpMax := -prices[0]
		for j := 1; j < len(prices); j++ {
			dp[i][j] = max(dp[i][j-1], tmpMax+prices[j])
			tmpMax = max(tmpMax, dp[i-1][j-1]-prices[j])
		}
	}
	dump(dp, prices)
	return dp[k][len(prices)-1]
}

func dump(dp [][]int, prices []int) {
	fmt.Printf("  ")
	for i := 0; i < len(prices); i++ {
		fmt.Printf("%3d ", prices[i])
	}
	fmt.Println()
	fmt.Printf("  ")
	for i := 1; i <= len(dp[0]); i++ {
		fmt.Printf("%3d ", i)
	}
	fmt.Println()
	for i := 0; i < len(dp); i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < len(dp[i]); j++ {
			fmt.Printf("%3d ", dp[i][j])
		}
		fmt.Println()
	}
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
	maxLen, maxElement := 20, 20
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
	if len(a) != 10 {
		return true
	}
	//a := []int{80, 60, 11, 87, 7, 92, 60}
	a = []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 10}
	ans := maxProfitBf(a)
	res := maxProfit(4, a)
	return false
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
