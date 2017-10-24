package main

import "fmt"

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if target <= nums[low] {
			return low
		} else if nums[low] < target && target < nums[mid] {
			high = mid - 1
		} else if nums[mid] < target && target < nums[high] {
			low = mid + 1
		} else if nums[high] == target {
			return high
		} else {
			return high + 1
		}
	}
	return 0
}

func searchBF(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

func testSearch(nums []int, target int) {
	res := searchInsert(nums, target)
	ans := searchBF(nums, target)
	if res != ans {
		panic(fmt.Errorf("nums %v, target %d, res %d, ans %d", nums, target, res, ans))
	}
}

func main() {
	testSearch([]int{1, 3, 5, 6}, 5)
	testSearch([]int{1, 3, 5, 6}, 2)
	testSearch([]int{1, 3, 5, 6}, 7)
	testSearch([]int{1, 3, 5, 6}, 0)
}
