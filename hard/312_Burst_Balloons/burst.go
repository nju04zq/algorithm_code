package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxCoins(nums []int) int {
	nums1 := make([]int, len(nums)+2)
	nums1[0] = 1
	var i, j int
	for i, j = 0, 1; i < len(nums); i++ {
		if nums[i] != 0 {
			nums1[j] = nums[i]
			j++
		}
	}
	nums1[j] = 1
	n := j + 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	fmt.Println(nums1)
	for k := 2; k < n; k++ {
		for left := 0; left+k < n; left++ {
			right := left + k
			for i := left + 1; i < right; i++ {
				c := nums1[left] * nums1[i] * nums1[right]
				c += (dp[left][i] + dp[i][right])
				dp[left][right] = max(dp[left][right], c)
			}
		}
	}
	return dp[0][n-1]
}

func dfs(nums []int, mask []bool, total int, maxCoin *int) {
	n := len(mask)
	prev, cur, next := -1, 0, n
	for ; cur < n; cur++ {
		if !mask[cur] {
			break
		}
	}
	if cur == n {
		*maxCoin = max(*maxCoin, total)
		return
	}
	for cur < n {
		for next = cur + 1; next < n; next++ {
			if !mask[next] {
				break
			}
		}
		coin := nums[cur]
		if prev != -1 {
			coin *= nums[prev]
		}
		if next != n {
			coin *= nums[next]
		}
		mask[cur] = true
		dfs(nums, mask, total+coin, maxCoin)
		mask[cur] = false
		prev, cur = cur, next
	}
}

func bf(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	mask := make([]bool, n)
	total, maxCoin := 0, 0
	dfs(nums, mask, total, &maxCoin)
	return maxCoin
}

func testMaxCoin() {
	nums := []int{3, 1, 5, 8}
	fmt.Printf("%v, get %d\n", nums, maxCoins(nums))
}

func main() {
	testMaxCoin()
}
