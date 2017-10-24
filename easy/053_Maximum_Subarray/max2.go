package main

import "fmt"
import "math/rand"
import "time"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minSub := nums[0]
	total := nums[0]
	for i := 1; i < len(nums); i++ {
		if total > 0 {
			total = nums[i]
		} else {
			total = total + nums[i]
		}
		minSub = min(total, minSub)
	}
	return minSub
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

func minSubBF(nums []int) int {
	minSub := nums[0]
	for i := 0; i < len(nums); i++ {
		total := 0
		for j := i; j < len(nums); j++ {
			total += nums[j]
			minSub = min(total, minSub)
		}
	}
	return minSub
}

func testMinSubArray(nums []int) {
	res := minSubArray(nums)
	ans := minSubBF(nums)
	if res != ans {
		panic(fmt.Errorf("nums %v, get %d, expect %d", nums, res, ans))
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		testMinSubArray(MakeRandArray())
	}
}
