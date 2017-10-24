package main

import "fmt"
import "sort"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < nums[high] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return nums[low]
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
}
