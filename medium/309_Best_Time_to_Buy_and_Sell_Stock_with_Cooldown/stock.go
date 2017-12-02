package main

import "fmt"

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
	buy := make([]int, len(prices))
	sell := make([]int, len(prices))
	buy[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		if i == 1 {
			buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		} else {
			buy[i] = max(buy[i-1], sell[i-2]-prices[i])
		}
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[len(prices)-1]
}

func testMaxProfit(prices []int) {
	fmt.Printf("%v, get %d\n", prices, maxProfit(prices))
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	testMaxProfit(prices)
	prices = []int{2, 1, 4}
	testMaxProfit(prices)
}
