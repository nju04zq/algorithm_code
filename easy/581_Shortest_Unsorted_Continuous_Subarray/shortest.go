package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	a := make([]int, len(nums))
	for i := 0; i < len(a); i++ {
		a[i] = nums[i]
	}
	sort.Ints(a)
	var i int
	for i = 0; i < len(a); i++ {
		if a[i] != nums[i] {
			break
		}
	}
	start := i
	for i = len(a) - 1; i >= 0; i-- {
		if a[i] != nums[i] {
			break
		}
	}
	end := i
	if start >= end {
		return 0
	} else {
		return end - start + 1
	}
}

func main() {
	fmt.Println("vim-go")
}
