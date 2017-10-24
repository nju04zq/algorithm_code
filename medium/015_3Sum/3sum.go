package main

import "fmt"
import "sort"

func twoSum(nums []int, target int) [][]int {
	i, j := 0, len(nums)-1
	results := make([][]int, 0)
	for i < j {
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		} else if j < (len(nums)-1) && nums[j] == nums[j+1] {
			j--
			continue
		}
		sum := nums[i] + nums[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			a := []int{nums[i], nums[j]}
			results = append(results, a)
			i++
			j--
		}
	}
	return results
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		resTwo := twoSum(nums[i+1:], -nums[i])
		if len(resTwo) == 0 {
			continue
		}
		for _, res := range resTwo {
			a := []int{nums[i]}
			a = append(a, res...)
			results = append(results, a)
		}
	}
	return results
}

func testThreeSum(nums []int) {
	results := threeSum(nums)
	fmt.Printf("%v, get %v\n", nums, results)
}

func main() {
	testThreeSum([]int{})
	testThreeSum([]int{1})
	testThreeSum([]int{-1, 0, 1, 2, -1, -4})
	testThreeSum([]int{0, 0, 0, 0, 0, 0})
}
