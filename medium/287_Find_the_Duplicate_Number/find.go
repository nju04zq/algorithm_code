package main

import "fmt"

func findDuplicate(nums []int) int {
	n := len(nums)
	low, high := 1, n-1
	for low < high {
		mid := low + (high-low)/2
		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	nums := []int{1, 2, 2, 4}
	fmt.Println(nums, findDuplicate(nums))
}
