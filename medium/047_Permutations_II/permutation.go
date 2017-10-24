package main

import "fmt"

func makeCopy(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func permuteInternal(nums []int, start int, ans [][]int) [][]int {
	if start == len(nums) {
		ans = append(ans, makeCopy(nums))
	}
	tbl := make(map[int]bool, len(nums))
	for i := start; i < len(nums); i++ {
		if _, ok := tbl[nums[i]]; ok {
			continue
		}
		tbl[nums[i]] = true
		nums[i], nums[start] = nums[start], nums[i]
		ans = permuteInternal(nums, start+1, ans)
		nums[i], nums[start] = nums[start], nums[i]
	}
	return ans
}

func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = permuteInternal(nums, 0, ans)
	return ans
}

func testPermute(nums []int) {
	ans := permuteUnique(nums)
	fmt.Printf("nums %v, get %v\n", nums, ans)
}

func main() {
	testPermute([]int{1})
	testPermute([]int{1, 1})
	testPermute([]int{1, 1, 2})
}
