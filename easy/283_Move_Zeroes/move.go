package main

import "fmt"

func moveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			if i != j {
				nums[j] = nums[i]
			}
			j++
		}
		i++
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}

func testMove(nums []int) {
	fmt.Printf("Before move, %v\n", nums)
	moveZeroes(nums)
	fmt.Printf("After move, %v\n", nums)
}

func main() {
	testMove([]int{0, 1, 0, 3, 12})
	testMove([]int{1})
}
