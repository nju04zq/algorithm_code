package main

import "fmt"

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tbl := make(map[int]int)
	for _, num := range nums {
		tbl[num]++
	}
	for _, num := range nums {
		if tbl[num] == 1 {
			return num
		}
	}
	return 0
}

func testSingleNumber(nums []int, ans int) {
	res := singleNumber(nums)
	if res != ans {
		panic(fmt.Errorf("Fail on %v, should %d, get %d", nums, ans, res))
	}
}

func main() {
	testSingleNumber([]int{1, 1, 2}, 2)
}
