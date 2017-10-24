package main

import "fmt"

func removeElement(nums []int, val int) int {
	var j int
	for i, num := range nums {
		if num == val {
			continue
		}
		if i != j {
			nums[j] = num
		}
		j++
	}
	return j
}

func testRemove(nums []int, val int) {
	fmt.Printf("Before remove %d: %v\n", val, nums)
	ans := removeElement(nums, val)
	fmt.Printf("After remove: %v\n", nums[:ans])
}

func main() {
	testRemove([]int{}, 0)
	testRemove([]int{1}, 0)
	testRemove([]int{1}, 1)
	testRemove([]int{1, 2, 2, 3}, 0)
	testRemove([]int{1, 2, 2, 3}, 1)
	testRemove([]int{1, 2, 2, 3}, 2)
	testRemove([]int{1, 2, 2, 3}, 3)
}
