package main

import "fmt"

func reverse(nums []int, start, end int) {
	i, j := start, end-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n)
	reverse(nums, 0, k)
	reverse(nums, k, n)
}

func testRotate(nums []int, k int) {
	fmt.Printf("Before rotate: %v\n", nums)
	rotate(nums, k)
	fmt.Printf("After rotate: %v\n", nums)
}

func main() {
	testRotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	testRotate([]int{1, 2, 3, 4, 5, 6, 7}, 10)
}
