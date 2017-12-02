package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	if len(nums) == 0 {
		return res
	}
	nums = append(nums, nums[len(nums)-1])
	n := len(nums)
	start, end := nums[0], 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1]+1 {
			continue
		}
		end = nums[i-1]
		if start == end {
			res = append(res, fmt.Sprintf("%d", start))
		} else {
			res = append(res, fmt.Sprintf("%d->%d", start, end))
		}
		start = nums[i]
	}
	nums = nums[:n-1]
	return res
}

func testSummary(nums []int) {
	res := summaryRanges(nums)
	fmt.Printf("%v, summary %q\n", nums, res)
}

func main() {
	testSummary([]int{0, 1, 2, 4, 5, 7})
	testSummary([]int{0, 2, 3, 4, 6, 8, 9})
}
