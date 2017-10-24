package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := high - (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[low] < nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func searchBF(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

func testSearch(nums []int, target int) {
	ans := searchBF(nums, target)
	res := search(nums, target)
	if ans != res {
		panic(fmt.Errorf("Fail on %v, target %d, get %d, should %d\n",
			nums, target, res, ans))
	}
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func rotate(nums []int, k int) []int {
	nums1 := make([]int, len(nums))
	for i, num := range nums {
		nums1[i] = num
	}
	n := len(nums1)
	reverse(nums1[:n-k])
	reverse(nums1[n-k:])
	reverse(nums1)
	return nums1
}

func main() {
	testSearch([]int{1}, 1)
	nums := []int{2, 4, 6, 8, 10}
	targets := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := 0; i < len(nums); i++ {
		for _, target := range targets {
			nums1 := rotate(nums, i)
			testSearch(nums1, target)
		}
	}
}
