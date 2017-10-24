package main

import "fmt"

func copySlice(a []int) []int {
	b := make([]int, len(a))
	for i, num := range a {
		b[i] = num
	}
	return b
}

func combineInternal(nums []int, start, k int, path []int, res [][]int) [][]int {
	if k == 0 {
		res = append(res, copySlice(path))
		return res
	} else if (start + k) > len(nums) {
		return res
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		res = combineInternal(nums, i+1, k-1, path, res)
		path = path[:len(path)-1]
	}
	return res
}

func combine(n int, k int) [][]int {
	path := make([]int, 0, k)
	res := make([][]int, 0)
	nums := make([]int, n)
	for i, _ := range nums {
		nums[i] = i + 1
	}
	res = combineInternal(nums, 0, k, path, res)
	return res
}

func testCombine(n, k int) {
	res := combine(n, k)
	fmt.Printf("n %d, k %d, get %v\n", n, k, res)
}

func main() {
	testCombine(4, 1)
	testCombine(4, 2)
	testCombine(4, 3)
	testCombine(4, 4)
}
