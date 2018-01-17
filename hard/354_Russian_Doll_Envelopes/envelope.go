package main

import "fmt"
import "sort"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func isFit(a, b []int) bool {
	if a[0] > b[0] && a[1] > b[1] {
		return true
	} else {
		return false
	}
}

func findSmaller(a [][]int, end, w int) int {
	low, high := 0, end
	for low < high {
		mid := high - (high-low)/2
		if a[mid][0] >= w {
			high = mid - 1
		} else {
			low = mid
		}
	}
	if low == high && a[low][0] < w {
		return low
	} else {
		return -1
	}
}

func maxEnvelopes(a [][]int) int {
	if len(a) == 0 || len(a[0]) == 0 {
		return 0
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] > a[j][0] {
			return false
		} else if a[i][0] < a[j][0] {
			return true
		} else if a[i][1] > a[j][1] {
			return false
		} else {
			return true
		}
	})
	n := len(a)
	dp := make([]int, n)
	maxN := 0
	for i := 0; i < n; i++ {
		maxJ := findSmaller(a, i-1, a[i][0])
		for j := 0; j <= maxJ; j++ {
			if isFit(a[i], a[j]) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] == 0 {
			dp[i] = 1
		}
		maxN = max(maxN, dp[i])
	}
	return maxN
}

func testEnvelope(a [][]int) {
	fmt.Printf("%v, get %d\n", a, maxEnvelopes(a))
}

func main() {
	a := [][]int{[]int{5, 4}, []int{6, 4}, []int{6, 7}, []int{2, 3}}
	testEnvelope(a)
	a = [][]int{[]int{1, 3}, []int{3, 5}, []int{6, 7}, []int{6, 8}, []int{8, 4}, []int{9, 5}}
	testEnvelope(a)
	a = [][]int{[]int{1, 1}}
	testEnvelope(a)
}
