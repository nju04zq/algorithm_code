package main

import "fmt"
import "sort"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minVal := nums[0]
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[low] < nums[mid] {
			minVal = min(minVal, nums[low])
			low = mid + 1
		} else if nums[mid] < nums[high] {
			minVal = min(minVal, nums[mid])
			high = mid - 1
		} else {
			minVal = min(minVal, nums[high])
			high = high - 1
		}
	}
	return minVal
}

func findBF(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func testFindMin(nums []int) {
	res := findMin(nums)
	ans := findBF(nums)
	if res != ans {
		panic(fmt.Errorf("nums %v, get %d, expect %d", nums, res, ans))
	}
}

func main() {
	testFindMin([]int{1, 2, 3, 4, 5})
	testFindMin([]int{5, 1, 2, 3, 4})
	testFindMin([]int{4, 5, 1, 2, 3})
	testFindMin([]int{3, 4, 5, 1, 2})
	testFindMin([]int{2, 3, 4, 5, 1})
	testFindMin([]int{1, 2, 2, 2, 2})
	testFindMin([]int{2, 1, 2, 2, 2})
	testFindMin([]int{2, 2, 1, 2, 2})
	testFindMin([]int{2, 2, 2, 1, 2})
	testFindMin([]int{2, 2, 2, 2, 1})
	testFindMin([]int{2, 2, 2, 2, 2})
}
