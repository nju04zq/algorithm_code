package main

import "fmt"
import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) <= 2 {
		return false
	}
	c1, c2 := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num <= c1 {
			c1 = num
		} else if num <= c2 {
			c2 = num
		} else {
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
}
