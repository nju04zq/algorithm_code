package main

import "fmt"
import "sort"

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func twoSumClosest(nums []int, target int) int {
	i, j := 0, len(nums)-1
	closeSum := nums[i] + nums[j]
	for i < j {
		sum := nums[i] + nums[j]
		if abs(target-sum) < abs(target-closeSum) {
			closeSum = sum
		}
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			break
		}
	}
	return closeSum
}

func threeSumClosest(nums []int, target int) int {
	var closeSum int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		res := twoSumClosest(nums[i+1:], target-nums[i])
		if i == 0 {
			closeSum = nums[i] + res
		} else {
			sum := nums[i] + res
			if abs(target-sum) < abs(target-closeSum) {
				closeSum = sum
			}
		}
		if closeSum == target {
			break
		}
	}
	return closeSum
}

func testThreeSumClosest(nums []int, target int) {
	res := threeSumClosest(nums, target)
	fmt.Printf("target %d, %v, get %v\n", target, nums, res)
}

func main() {
	testThreeSumClosest([]int{}, 0)
	testThreeSumClosest([]int{1}, 0)
	testThreeSumClosest([]int{-1, 2, 1, -4}, 1)
}
