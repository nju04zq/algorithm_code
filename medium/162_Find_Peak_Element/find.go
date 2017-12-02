package main

import "fmt"

func isPeak(nums []int, i int) bool {
	if len(nums) == 1 {
		return true
	} else if i == 0 {
		if nums[i] > nums[i+1] {
			return true
		} else {
			return false
		}
	} else if i == len(nums)-1 {
		if nums[i] > nums[i-1] {
			return true
		} else {
			return false
		}
	} else {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return true
		} else {
			return false
		}
	}
	return false
}

func findInternal(nums []int, low, high int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	if isPeak(nums, mid) {
		return mid
	}
	left := findInternal(nums, low, mid-1)
	if left != -1 {
		return left
	}
	right := findInternal(nums, mid+1, high)
	if right != -1 {
		return right
	}
	return -1
}

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	return findInternal(nums, 0, len(nums)-1)
}

func testFind(nums []int) {
	fmt.Printf("%v, get %d\n", nums, findPeakElement(nums))
}

func main() {
	testFind([]int{1})
	testFind([]int{1, 2, 3, 1})
}
