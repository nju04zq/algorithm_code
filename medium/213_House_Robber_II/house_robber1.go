package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func robInternal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prev := 0
		if i >= 2 {
			prev = dp[i-2]
		}
		dp[i] = max(dp[i-1], prev+nums[i])
	}
	return dp[len(nums)-1]
}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	return max(robInternal(nums[1:n]), robInternal(nums[0:n-1]))
}

func main() {
	a := []int{1}
	fmt.Printf("%v, get %d\n", a, rob(a))
	a = []int{1, 2}
	fmt.Printf("%v, get %d\n", a, rob(a))
	a = []int{1, 2, 3}
	fmt.Printf("%v, get %d\n", a, rob(a))
	a = []int{1, 2, 3, 4}
	fmt.Printf("%v, get %d\n", a, rob(a))
}
