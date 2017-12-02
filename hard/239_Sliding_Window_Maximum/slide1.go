package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	res := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		if i%k == 0 {
			left[i] = nums[i]
		} else {
			left[i] = max(left[i-1], nums[i])
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if (i+1)%k == 0 || i == len(nums)-1 {
			right[i] = nums[i]
		} else {
			right[i] = max(right[i+1], nums[i])
		}
	}
	for i := 0; i < len(nums)-k+1; i++ {
		res[i] = max(right[i], left[i+k-1])
	}
	return res
}

func testSliding(nums []int, k int) {
	res := maxSlidingWindow(nums, k)
	fmt.Printf("%v, k %d, get %v\n", nums, k, res)
}

func main() {
	testSliding([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}
