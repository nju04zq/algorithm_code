package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dpLen := make([]int, n)
	dpCnt := make([]int, n)
	var maxLen, maxLenCnt int
	for i := 0; i < n; i++ {
		dpLen[i], dpCnt[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if dpLen[i] == dpLen[j]+1 {
				dpCnt[i] += dpCnt[j]
			} else if dpLen[i] < dpLen[j]+1 {
				dpLen[i] = dpLen[j] + 1
				dpCnt[i] = dpCnt[j]
			}
		}
		if dpLen[i] > maxLen {
			maxLen = dpLen[i]
			maxLenCnt = dpCnt[i]
		} else if dpLen[i] == maxLen {
			maxLenCnt += dpCnt[i]
		}
	}
	return maxLenCnt
}

func testFind(nums []int) {
	fmt.Printf("%v, get %d\n", nums, findNumberOfLIS(nums))
}

func main() {
	testFind([]int{1})
	testFind([]int{3, 2, 1})
	testFind([]int{1, 3, 5, 4, 7})
	testFind([]int{2, 2, 2, 2, 2})
	testFind([]int{1, 2, 4, 3, 5, 4, 7, 2})
}
