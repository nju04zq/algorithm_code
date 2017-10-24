package main

import "fmt"

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		total := nums[i] + nums[j]
		if total == target {
			return []int{i + 1, j + 1}
		} else if total > target {
			j--
		} else {
			i++
		}
	}
	return nil
}

func main() {
	nums := []int{1, 5, 7, 9}
	target := 12
	fmt.Printf("twoSum %v\n", twoSum(nums, target))
}
