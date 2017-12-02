package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestConsecutive(nums []int) int {
	var longest int
	tbl := make(map[int]bool, len(nums))
	for _, num := range nums {
		tbl[num] = true
	}
	for num, _ := range tbl {
		if _, ok := tbl[num-1]; ok {
			continue
		}
		start := num
		end := num + 1
		for {
			if _, ok := tbl[end]; !ok {
				break
			}
			end++
		}
		longest = max(longest, end-start)
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
