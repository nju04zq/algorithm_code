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
	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		ans = permuteInternal(nums, start+1, ans)
		nums[i], nums[start] = nums[start], nums[i]
	}
	return ans
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = permuteInternal(nums, 0, ans)
	return ans
}

func testPermute(nums []int) {
	ans := permute(nums)
	fmt.Printf("nums %v, ans %v\n", nums, ans)
}

func main() {
	testPermute([]int{1})
	testPermute([]int{1, 2})
	testPermute([]int{1, 2, 3})
}
