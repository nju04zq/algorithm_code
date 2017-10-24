package main

import (
	"fmt"
	"math/rand"
	"time"
)

func maxProfit(prices []int) int {
	minPrice, maxPrice, profit := -1, -1, 0
	for _, price := range prices {
		if minPrice == -1 {
			minPrice, maxPrice = price, price
			continue
		}
		if price > maxPrice {
			maxPrice = price
		} else {
			if maxPrice != -1 {
				profit += (maxPrice - minPrice)
			}
			minPrice, maxPrice = price, price
		}
	}
	if maxPrice != -1 {
		profit += (maxPrice - minPrice)
	}
	return profit
}

func maxProfitDp(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		for j := i + 1; j < len(prices); j++ {
			var profit int
			if prices[j] > prices[i] {
				profit = prices[j] - prices[i] + dp[j]
			} else {
				profit = dp[j]
			}
			dp[i] = max(dp[i], profit)
		}
	}
	return dp[0]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfitBf(prices []int) int {
	maxProfit, profit := 0, 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit = prices[j] - prices[i]
			maxProfit = max(maxProfit, profit+maxProfitBf(prices[j:]))
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
	res := maxProfit(a)
	dp := maxProfitDp(a)
	if ans != res || dp != ans {
		fmt.Printf("Fail on %v, get %d, dp %d, expect %d\n", a, res, dp, ans)
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
