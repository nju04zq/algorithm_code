package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	res := make([]int, 0, n)
	for i := 0; i < n; {
		if nums[i] == i+1 || nums[nums[i]-1] == nums[i] {
			i++
		} else {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

func testMissing(nums []int) {
	fmt.Printf("Missing in %v, %v\n", nums, findDisappearedNumbers(nums))
}

func main() {
	testMissing([]int{4, 3, 2, 7, 8, 2, 3, 1})
}
