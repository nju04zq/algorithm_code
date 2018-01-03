package main

import "fmt"
import "math"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := 1; i < len(f); i++ {
		f[i] = math.MinInt32
	}
	for i := 0; i < len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j-coins[i] < 0 {
				continue
			}
			tmp := f[j-coins[i]]
			if tmp != math.MinInt32 {
				f[j] = max(f[j], tmp-1)
			}
		}
	}
	if f[amount] == math.MinInt32 {
		return -1
	} else {
		return -f[amount]
	}
}

func testCoin(coins []int, amount int) {
	fmt.Printf("%v, amount %d, get %d\n",
		coins, amount, coinChange(coins, amount))
}

func main() {
	testCoin([]int{1, 2, 5}, 11)
	testCoin([]int{2}, 3)
}
