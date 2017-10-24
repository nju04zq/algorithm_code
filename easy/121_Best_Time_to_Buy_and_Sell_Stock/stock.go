package main

import (
	"fmt"
	"math/rand"
	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	minPrice, maxProfit := -1, 0
	for _, price := range prices {
		if minPrice == -1 {
			minPrice = price
			continue
		}
		if price >= minPrice {
			maxProfit = max(price-minPrice, maxProfit)
		} else {
			minPrice = price
		}
	}
	return maxProfit
}

func maxProfitBf(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			maxProfit = max(maxProfit, prices[j]-prices[i])
		}
	}
	return maxProfit
}

func MakeRandArray() []int {
	maxLen, maxElement := 100, 100
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
	ans := maxProfitBf(a)
	res := maxProfit(a)
	if ans != res {
		fmt.Printf("Fail on %v, get %d, expect %d\n", a, ans, res)
		return false
	}
	return true
}

func main() {
	for i := 0; i < 1000; i++ {
		res := testMaxProfit()
		if !res {
			break
		}
	}
}
