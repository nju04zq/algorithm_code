package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}
	buy := make([]int, len(prices))
	sell := make([]int, len(prices))
	buy[0] = -prices[0] - fee
	for i := 1; i < len(prices); i++ {
		buy[i] = max(buy[i-1], sell[i-1]-prices[i]-fee)
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[len(prices)-1]
}

func testMaxProfit(prices []int, fee int) {
	fmt.Printf("prices %v, fee %v, get %d\n",
		prices, fee, maxProfit(prices, fee))
}

func main() {
	testMaxProfit([]int{1, 3, 2, 8, 4, 9}, 2)
}
