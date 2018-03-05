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
	if sum > total || sum < -total {
		return 0
	}
	dp := make([][]int, n)
	m := 2 * total
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m+1)
	}
	j := total + nums[0]
	if j <= m {
		dp[0][j] += 1
	}
	j = total - nums[0]
	if j <= m {
		dp[0][j] += 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= m; j++ {
			k := j + nums[i]
			if k >= 0 && k <= m {
				dp[i][j] += dp[i-1][k]
			}
			k = j - nums[i]
			if k >= 0 && k <= m {
				dp[i][j] += dp[i-1][k]
			}
		}
	}
	//dump(dp, m, n)
	return dp[n-1][sum+total]
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
