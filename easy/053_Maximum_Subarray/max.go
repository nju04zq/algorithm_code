package main

import "fmt"
import "math/rand"
import "time"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSub := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		maxSub = max(dp[i], maxSub)
	}
	return maxSub
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int()%maxLen + 1
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int()%maxElement - 8
	}
	return a
}

func maxSubBF(nums []int) int {
	maxSub := nums[0]
	for i := 0; i < len(nums); i++ {
		total := 0
		for j := i; j < len(nums); j++ {
			total += nums[j]
			maxSub = max(total, maxSub)
		}
	}
	return maxSub
}

func testMaxSubArray(nums []int) {
	res := maxSubArray(nums)
	ans := maxSubBF(nums)
	if res != ans {
		panic(fmt.Errorf("nums %v, get %d, expect %d", nums, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testMaxSubArray(MakeRandArray())
	}
}
