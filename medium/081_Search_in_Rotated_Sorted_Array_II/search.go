package main

import "fmt"

func searchInternal(nums []int, low, high, target int) bool {
	if low > high {
		return false
	}
	mid := high - (high-low)/2
	if nums[mid] == target {
		return true
	}
	if nums[low] < nums[mid] {
		if nums[low] <= target && target < nums[mid] {
			return searchInternal(nums, low, mid-1, target)
		} else {
			return searchInternal(nums, mid+1, high, target)
		}
	} else if nums[mid] < nums[high] {
		if nums[mid] < target && target <= nums[high] {
			return searchInternal(nums, mid+1, high, target)
		} else {
			return searchInternal(nums, low, mid-1, target)
		}
	} else {
		res := searchInternal(nums, low, mid-1, target)
		if res {
			return true
		}
		res = searchInternal(nums, mid+1, high, target)
		if res {
			return true
		}
	}
	return false
}

func search(nums []int, target int) bool {
	return searchInternal(nums, 0, len(nums)-1, target)
}

func searchBF(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func testSearch(nums []int, target int) {
	ans := searchBF(nums, target)
	res := search(nums, target)
	if ans != res {
		panic(fmt.Errorf("Fail on %v, target %d, get %t, should %t\n",
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
	testSearch([]int{1, 1, 1, 3, 1}, 3)
	return
	testSearch([]int{1}, 1)
	testSearch([]int{2, 2, 2, 2, 2}, 1)
	testSearch([]int{2, 2, 2, 2, 2}, 2)
	testSearch([]int{1, 1, 1, 3, 1}, 3)
	nums := []int{2, 4, 6, 8, 10}
	targets := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := 0; i < len(nums); i++ {
		for _, target := range targets {
			nums1 := rotate(nums, i)
			testSearch(nums1, target)
		}
	}
}
