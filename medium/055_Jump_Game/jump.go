package main

import "fmt"

func canJump(nums []int) bool {
	j := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= j {
			j = i
		}
	}
	if j == 0 {
		return true
	} else {
		return false
	}
}

func canJumpDp(nums []int) bool {
	f := make([]bool, len(nums))
	f[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j <= nums[i]+i && j < len(nums); j++ {
			if f[j] == true {
				f[i] = true
				break
			}
		}
	}
	return f[0]
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
