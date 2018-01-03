package main

import "fmt"

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func main() {
	nums := []int{1, 2, 2, 4}
	fmt.Println(nums, findDuplicate(nums))
	nums = []int{1, 3, 4, 2, 2}
	fmt.Println(nums, findDuplicate(nums))
}
