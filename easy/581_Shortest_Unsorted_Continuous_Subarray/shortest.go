package main

import "fmt"
import "sort"

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

func testFind(a []int) {
	fmt.Printf("%v, get %d\n", a, findUnsortedSubarray(a))
}

func main() {
	testFind([]int{2, 6, 4, 8, 10, 9, 15})
	testFind([]int{2, 4, 8, 10, 15})
	testFind([]int{2, 4, 8, 6, 10, 9, 12, 15})
	testFind([]int{5, 4, 3, 2, 1})
	testFind([]int{2, 1})
	testFind([]int{1, 3, 2, 2, 2})
	testFind([]int{1, 2, 4, 5, 3})
}
