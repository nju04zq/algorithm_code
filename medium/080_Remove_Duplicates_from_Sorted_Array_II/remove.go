package main

import "fmt"

func removeDuplicates(nums []int) int {
	var j, cnt, cur int
	for i, num := range nums {
		if i == 0 {
			cur = num
		}
		if num == cur {
			cnt++
		} else {
			cur = num
			cnt = 1
		}
		if cnt <= 2 {
			if i != j {
				nums[j] = num
			}
			j++
		}
	}
	return j
}

func testRemove(nums []int) {
	fmt.Printf("Before remove: %v\n", nums)
	ans := removeDuplicates(nums)
	fmt.Printf("After remove: %d, %v\n", ans, nums[:ans])
}

func main() {
	testRemove([]int{})
	testRemove([]int{1})
	testRemove([]int{1, 1})
	testRemove([]int{1, 1, 1})
	testRemove([]int{1, 1, 1, 2, 2, 3})
	testRemove([]int{1, 1, 1, 2, 2, 3, 3, 3})
}
