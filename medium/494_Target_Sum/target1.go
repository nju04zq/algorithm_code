package main

import "fmt"

func dump(dp [][]int, m, n int) {
	for i := 0; i <= m; i++ {
		fmt.Printf("%02d ", i)
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		for j := 0; j <= m; j++ {
			fmt.Printf("%02d ", dp[i][j])
		}
		fmt.Println()
	}
}

func findTargetSumWays(nums []int, sum int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
		total += nums[i]
	}
	if sum > total || sum < -total || (sum+total)%2 != 0 {
		return 0
	}
	m := (sum + total) / 2
	dp := make([]int, m+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := m; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[m]
}

func testFind(nums []int, sum int) {
	fmt.Printf("nums %v, sum %d, get %d\n", nums, sum, findTargetSumWays(nums, sum))
}

func main() {
	testFind([]int{1, 1, 1, 1, 1}, 3)
	testFind([]int{1, 1, 1, 1, 1}, -3)
	testFind([]int{1}, 2)
	testFind([]int{1, 0}, 1)
}
