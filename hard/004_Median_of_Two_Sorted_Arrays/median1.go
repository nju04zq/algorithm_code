package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// k starts from 1
func findKth(nums1 []int, nums2 []int, k int) int {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findKth(nums2, nums1, k)
	}
	if m == 0 {
		return nums2[k-1]
	} else if n == 0 {
		return nums1[k-1]
	} else if k == 1 {
		return min(nums1[0], nums2[0])
	}
	pa := min(k/2, m)
	pb := k - pa
	if nums1[pa-1] < nums2[pb-1] {
		return findKth(nums1[pa:], nums2[:pb], k-pa)
	} else {
		return findKth(nums1[:pa], nums2[pb:], k-pb)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var a, b int
	m, n := len(nums1), len(nums2)
	k := (m + n) / 2
	if (m+n)%2 == 0 {
		a = findKth(nums1, nums2, k)
		b = findKth(nums1, nums2, k+1)
	} else {
		a = findKth(nums1, nums2, k+1)
		b = a
	}
	return float64(a+b) / 2

}

func makeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func makeRandSortedArray() []int {
	a := makeRandArray()
	sort.Ints(a)
	return a
}

func medianBF(nums1 []int, nums2 []int) float64 {
	var k1, k2, a1, a2 int
	m, n := len(nums1), len(nums2)
	if (m+n)%2 == 0 {
		k1 = (m+n)/2 - 1
		k2 = (m + n) / 2
	} else {
		k1 = (m + n) / 2
		k2 = -1
	}
	var i, j, cur, k int
	for i < m || j < n {
		if i >= m {
			cur = nums2[j]
			j++
		} else if j >= n {
			cur = nums1[i]
			i++
		} else if nums1[i] <= nums2[j] {
			cur = nums1[i]
			i++
		} else {
			cur = nums2[j]
			j++
		}
		if k == k1 {
			a1 = cur
		} else if k == k2 {
			a2 = cur
		} else if k > max(k1, k2) {
			break
		}
		k++
	}
	if k2 == -1 {
		return float64(a1)
	} else {
		return float64(a1+a2) / 2
	}

}

func testMedian() {
	for i := 0; i < 10000; i++ {
		a := makeRandSortedArray()
		b := makeRandSortedArray()
		if len(a) == 0 && len(b) == 0 {
			continue
		}
		mid1 := medianBF(a, b)
		mid2 := findMedianSortedArrays(a, b)
		if mid1 != mid2 {
			panic(fmt.Sprintf("a %v, b %v, get %d, should %d",
				a, b, mid2, mid1))
		}
	}
}

func main() {
	testMedian()
	//res := findMedianSortedArrays([]int{15, 16, 17, 18}, []int{1, 1, 5, 8, 10, 11, 11, 16})
	//fmt.Printf("res %.2f\n", res)
}
