package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var j, cur int
	for i, num := range nums {
		if i == 0 {
			cur = num
			continue
		}
		if num != cur {
			j++
			nums[j] = num
			cur = num
		}
	}
	return j + 1
}

func testRemove(nums []int) {
	fmt.Printf("Before remove, %v\n", nums)
	i := removeDuplicates(nums)
	fmt.Printf("After remove, %v\n", nums[:i])
}

func main() {
	testRemove([]int{})
	testRemove([]int{1})
	testRemove([]int{1, 2})
	testRemove([]int{1, 2, 3})
	testRemove([]int{1, 1})
	testRemove([]int{1, 2, 2})
	testRemove([]int{1, 2, 2, 3})
}
