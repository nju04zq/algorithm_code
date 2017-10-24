package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		k--
		i--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
}

func testMerge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println("Before merge:")
	fmt.Println(nums1[:m])
	fmt.Println(nums2[:n])
	merge(nums1, m, nums2, n)
	fmt.Println("After merge:")
	fmt.Println(nums1[:m+n])
}

func main() {
	nums1 := []int{1, 3, 5, 0, 0}
	nums2 := []int{2, 4}
	testMerge(nums1, 3, nums2, 2)
}
