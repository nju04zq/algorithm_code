package main

import "fmt"
import "sort"

func twoSum(nums []int, target int) [][]int {
	results := make([][]int, 0)
	i, j := 0, len(nums)-1
	for i < j {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}
		if j < (len(nums)-1) && nums[j] == nums[j+1] {
			j--
			continue
		}
		sum := nums[i] + nums[j]
		if sum < target {
			i++
		} else if sum > target {
			j--
		} else {
			results = append(results, []int{nums[i], nums[j]})
			i++
			j--
		}
	}
	return results
}

func threeSum(nums []int, target int) [][]int {
	results := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		res := twoSum(nums[i+1:], target-nums[i])
		if len(res) == 0 {
			continue
		}
		for _, a := range res {
			b := []int{nums[i]}
			b = append(b, a...)
			results = append(results, b)
		}
	}
	return results
}

func fourSum(nums []int, target int) [][]int {
	results := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		res := threeSum(nums[i+1:], target-nums[i])
		if len(res) == 0 {
			continue
		}
		for _, a := range res {
			b := []int{nums[i]}
			b = append(b, a...)
			results = append(results, b)
		}
	}
	return results
}

func testFourSum(nums []int, target int) {
	res := fourSum(nums, target)
	fmt.Printf("target %d, %v, get %v\n", target, nums, res)
}

func main() {
	testFourSum([]int{}, 0)
	testFourSum([]int{1}, 0)
	testFourSum([]int{0, 0, 0, 0, 0, 0}, 0)
	testFourSum([]int{1, 0, -1, 0, -2, 2}, 0)
}
