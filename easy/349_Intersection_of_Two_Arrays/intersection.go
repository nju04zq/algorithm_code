package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	tbl1 := make(map[int]bool)
	res := make([]int, 0)
	for _, num := range nums1 {
		tbl1[num] = true
	}
	for _, num := range nums2 {
		if _, ok := tbl1[num]; ok && tbl1[num] {
			res = append(res, num)
			tbl1[num] = false
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
