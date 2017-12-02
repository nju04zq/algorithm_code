package main

import "fmt"
import "math"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestConsecutive(nums []int) int {
	var longest, left, right int
	tbl := make(map[int]bool, len(nums))
	for _, num := range nums {
		tbl[num] = true
	}
	for num, free := range tbl {
		if !free {
			continue
		}
		right = num
		for i := num + 1; i <= math.MaxInt32; i++ {
			free, ok := tbl[i]
			if !ok || !free {
				break
			}
			tbl[i] = false
			right = i
		}
		left = num
		for i := num - 1; i >= math.MinInt32; i-- {
			free, ok := tbl[i]
			if !ok || !free {
				break
			}
			tbl[i] = false
			left = i
		}
		longest = max(longest, right-left+1)
	}
	return longest
}

func testLongest(nums []int) {
	longest := longestConsecutive(nums)
	fmt.Printf("%v, get %d\n", nums, longest)
}

func main() {
	testLongest([]int{100, 4, 200, 1, 3, 2})
	testLongest([]int{100, 4, 200, 1})
}
