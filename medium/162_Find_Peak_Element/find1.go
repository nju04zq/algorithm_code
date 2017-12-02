package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low < high {
		mid1 := low + (high-low)/2
		mid2 := mid1 + 1
		if nums[mid1] > nums[mid2] {
			high = mid1
		} else {
			low = mid2
		}
	}
	return low
}

func testFind(nums []int) {
	fmt.Printf("%v, get %d\n", nums, findPeakElement(nums))
}

func main() {
	testFind([]int{1})
	testFind([]int{1, 2, 3, 1})
}
