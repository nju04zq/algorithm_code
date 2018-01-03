package main

import "fmt"
import "math"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n <= 2 {
		return false
	}
	temp := make([]int, n)
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			temp[i] = nums[0]
		} else {
			temp[i] = min(nums[i-1], temp[i-1])
		}
	}
	//fmt.Println(temp)
	maxTill := math.MinInt32
	for i := n - 2; i >= 1; i-- {
		maxTill = max(maxTill, nums[i+1])
		if nums[i] > temp[i] && nums[i] < maxTill {
			//fmt.Println(i)
			return true
		}
	}
	return false
}

func testIncreasing(nums []int) {
	fmt.Printf("%v, get %t\n", nums, increasingTriplet(nums))
}

func main() {
	testIncreasing([]int{1, 2, 3, 4, 5})
	testIncreasing([]int{5, 4, 3, 2, 1})
	testIncreasing([]int{2, 4, -2, -3})
	testIncreasing([]int{5, 1, 5, 5, 2, 5, 4})
	testIncreasing([]int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2})
}
