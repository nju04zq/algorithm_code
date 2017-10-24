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

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	left := make([]int, len(prices))
	minPrice, maxProfit := prices[0], 0
	for i := 1; i < len(prices)-1; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-minPrice)
		}
		left[i] = maxProfit
	}
	maxPrice, maxProfit := prices[len(prices)-1], left[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		} else {
			maxProfit = max(maxProfit, maxPrice-prices[i]+left[i])
		}
	}
	return maxProfit
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
	res := maxProfit(a)
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
