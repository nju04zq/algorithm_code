package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		val := nums[i]
		if val < 0 {
			val = -val
		}
		if nums[val-1] > 0 {
			nums[val-1] = -nums[val-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func testMissing(nums []int) {
	fmt.Printf("Missing in %v, ", nums)
	fmt.Printf("get %v\n", findDisappearedNumbers(nums))
}

func main() {
	testMissing([]int{4, 3, 2, 7, 8, 2, 3, 1})
}
