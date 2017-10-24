package main

import "fmt"

func addNumStats(stats []int, num int) {
	for i := 0; i < 32; i++ {
		stats[i] += (num & 0x1)
		num >>= 1
	}
}

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	stats := make([]int, 32)
	for _, num := range nums {
		addNumStats(stats, num)
	}
	var num uint32
	var weight uint32 = 1
	for _, cnt := range stats {
		num += uint32(cnt%3) * weight
		weight *= 2
	}
	return int(int32(num))
}

func testSingleNumber(nums []int, ans int) {
	res := singleNumber(nums)
	if res != ans {
		panic(fmt.Errorf("Fail on %v, should %d, get %d", nums, ans, res))
	}
}

func main() {
	testSingleNumber([]int{-1}, -1)
	testSingleNumber([]int{1}, 1)
	testSingleNumber([]int{2}, 2)
	testSingleNumber([]int{1, 1, 1, 2}, 2)
}
