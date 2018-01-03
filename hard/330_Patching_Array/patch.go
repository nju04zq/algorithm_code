package main

import "fmt"

func minPatches(nums []int, n int) int {
	miss, i, cnt := 1, 0, 0
	for miss <= n {
		if i < len(nums) && miss >= nums[i] {
			miss += nums[i]
			i++
		} else {
			miss += miss
			cnt++
		}
	}
	return cnt
}

func testPatch(nums []int, n int) {
	fmt.Printf("%v, n %d, get %d\n", nums, n, minPatches(nums, n))
}

func main() {
	testPatch([]int{1, 3}, 6)
	testPatch([]int{1, 5, 10}, 20)
	testPatch([]int{1, 2, 2}, 5)
}
