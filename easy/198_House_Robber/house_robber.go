package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prev := 0
		if i > 1 {
			prev = dp[i-2]
		}
		dp[i] = max(dp[i-1], prev+nums[i])
	}
	return dp[len(nums)-1]
}

func main() {
	fmt.Println("vim-go")
}
