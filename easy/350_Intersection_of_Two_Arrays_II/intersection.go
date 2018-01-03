package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	tbl1 := make(map[int]int)
	for _, num := range nums1 {
		if _, ok := tbl1[num]; ok {
			tbl1[num]++
		} else {
			tbl1[num] = 1
		}
	}
	res := make([]int, 0)
	for _, num := range nums2 {
		if _, ok := tbl1[num]; ok && tbl1[num] > 0 {
			res = append(res, num)
			tbl1[num]--
		}
	}
	return res
}

func main() {
	fmt.Println("vim-go")
}
