package main

import "fmt"

func compare(a, b []int) int {
	var i int
	for i = 0; i < len(a) && i < len(b); i++ {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	if i == len(a) {
		return -1
	} else {
		return 1
	}
}

func replace(a, b []int) {
	for i := 0; i < len(a); i++ {
		a[i] = b[i]
	}
}

func copyInts(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}

func remove(a []int, j int) []int {
	for i := j; i < len(a)-1; i++ {
		a[i] = a[i+1]
	}
	a = a[:len(a)-1]
	return a
}

func getDp(nums1 []int, k int) [][]int {
	n := len(nums1)
	dp := make([][]int, k+1)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = nums1[i]
	}
	var i int
	for start := 0; len(temp) > 0; {
		if len(temp) <= k {
			dp[len(temp)] = copyInts(temp)
		}
		for i = start; i < len(temp)-1; i++ {
			if temp[i] < temp[i+1] {
				break
			}
		}
		temp = remove(temp, i)
		if i == 0 {
			start = 0
		} else {
			start = i - 1
		}
	}
	return dp
}

func merge(a []int, b []int) []int {
	i, j, k := 0, 0, 0
	m, n := len(a), len(b)
	c := make([]int, m+n)
	for i < m || j < n {
		if j >= n || (i < m && compare(a[i:], b[j:]) >= 0) {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
		k++
	}
	return c
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := make([]int, k)
	dp1 := getDp(nums1, k)
	dp2 := getDp(nums2, k)
	for i := 0; i <= k; i++ {
		if len(dp1[i])+len(dp2[k-i]) < k {
			continue
		}
		knums := merge(dp1[i], dp2[k-i])
		if compare(knums, res) > 0 {
			replace(res, knums)
		}
	}
	return res
}

func dfs(nums1 []int, nums2 []int, k int, path []int, res []int) {
	if k == 0 {
		if compare(path, res) > 0 {
			replace(res, path)
		}
		return
	}
	for i := 0; i < len(nums1); i++ {
		path = append(path, nums1[i])
		dfs(nums1[i+1:], nums2, k-1, path, res)
		path = path[:len(path)-1]
	}
	for i := 0; i < len(nums2); i++ {
		path = append(path, nums2[i])
		dfs(nums1, nums2[i+1:], k-1, path, res)
		path = path[:len(path)-1]
	}
}

func bf(nums1 []int, nums2 []int, k int) []int {
	path := make([]int, 0)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = -1
	}
	dfs(nums1, nums2, k, path, res)
	return res
}

func testCreate() {
	nums1 := []int{8, 6, 9}
	nums2 := []int{1, 7, 5}
	k := 3
	res := maxNumber(nums1, nums2, k)
	fmt.Printf("%v, %v, k %d, get %v\n", nums1, nums2, k, res)
}

func main() {
	testCreate()
}
