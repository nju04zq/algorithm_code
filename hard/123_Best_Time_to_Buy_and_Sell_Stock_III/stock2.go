package main

import (
	"fmt"
	"math"
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
	release := make([]int, 3)
	hold := make([]int, 3)
	hold[1], hold[2] = math.MinInt32, math.MinInt32
	for i := 0; i < len(prices); i++ {
		release[2] = max(hold[2]+prices[i], release[2])
		hold[2] = max(release[1]-prices[i], hold[2])
		release[1] = max(hold[1]+prices[i], release[1])
		hold[1] = max(-prices[i], hold[1])
	}
	return release[2]
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
