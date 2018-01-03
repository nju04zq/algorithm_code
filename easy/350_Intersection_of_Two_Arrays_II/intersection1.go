package main

import "fmt"
import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	sort.Ints(nums1)
	sort.Ints(nums2)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			res = append(res, nums1[i])
			i++
			j++
		}
	}
	return res
}

func testIntersect(nums1, nums2 []int) {
	fmt.Printf("%v, %v, get %v\n", nums1, nums2, intersect(nums1, nums2))
}

func main() {
	testIntersect([]int{4, 7, 9, 7, 6, 7}, []int{5, 0, 0, 6, 1, 6, 2, 2, 4})
}
