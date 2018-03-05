package main

import "fmt"

func subarraySum(nums []int, k int) int {
	tbl := make(map[int]int, len(nums))
	sum, cnt := 0, 0
	tbl[sum] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		x := sum - k
		y, ok := tbl[x]
		if ok {
			cnt += y
		}
		if _, ok := tbl[sum]; ok {
			tbl[sum]++
		} else {
			tbl[sum] = 1
		}
	}
	return cnt
}

func testSubarray(nums []int, k int) {
	fmt.Printf("nums %v, k %d get %d\n", nums, k, subarraySum(nums, k))
}

func main() {
	testSubarray([]int{1, 1, 1}, 2)
}
