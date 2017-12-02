package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	for i := 0; i < len(nums); i++ {
		start := i
		for i+1 < len(nums) && nums[i+1] == nums[i]+1 {
			i++
		}
		if i == start {
			res = append(res, fmt.Sprintf("%d", nums[start]))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i]))
		}
	}
	return res
}

func testSummary(nums []int) {
	res := summaryRanges(nums)
	fmt.Printf("%v, summary %q\n", nums, res)
}

func main() {
	testSummary([]int{0})
	testSummary([]int{0, 1, 2, 4, 5, 7})
	testSummary([]int{0, 2, 3, 4, 6, 8, 9})
}
