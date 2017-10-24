package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func canJump(nums []int) bool {
	var maxForward int
	for i := 0; i < len(nums); i++ {
		if maxForward < i || maxForward >= len(nums)-1 {
			break
		}
		maxForward = max(maxForward, i+nums[i])
	}
	if maxForward >= len(nums)-1 {
		return true
	} else {
		return false
	}
}

func canJumpBF(nums []int) bool {
	for i := 1; i <= nums[0] && i < len(nums); i++ {
		if canJumpBF(nums[i:]) {
			return true
		}
	}
	return false
}

func testCanJump(nums []int) {
	ans1 := canJump(nums)
	ans2 := canJump(nums)
	fmt.Printf("nums %v, get %t, get %t\n", nums, ans1, ans2)
}

func main() {
	testCanJump([]int{2, 3, 1, 1, 4})
	testCanJump([]int{3, 2, 1, 0, 4})
}
