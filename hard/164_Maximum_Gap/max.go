package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findMinMax(nums []int) (int, int) {
	minNum, maxNum := -1, -1
	for _, num := range nums {
		if minNum == -1 {
			minNum = num
		} else {
			minNum = min(minNum, num)
		}
		maxNum = max(maxNum, num)
	}
	return minNum, maxNum
}

func ceil(a, b int) int {
	m := a / b
	n := a % b
	if n == 0 {
		return m
	} else {
		return m + 1
	}
}

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	minNum, maxNum := findMinMax(nums)
	if minNum == maxNum {
		return 0
	}
	gap := ceil(maxNum-minNum, n-1)
	bucketN := (maxNum-minNum)/gap + 1
	minBucket := make([]int, bucketN)
	maxBucket := make([]int, bucketN)
	for i := 0; i < bucketN; i++ {
		minBucket[i], maxBucket[i] = -1, -1
	}
	for i := 0; i < n; i++ {
		k := (nums[i] - minNum) / gap
		if minBucket[k] == -1 {
			minBucket[k] = nums[i]
		} else {
			minBucket[k] = min(minBucket[k], nums[i])
		}
		maxBucket[k] = max(maxBucket[k], nums[i])
	}
	prev, maxGap := -1, 0
	for i := 0; i < bucketN; i++ {
		if minBucket[i] == -1 {
			continue
		}
		if prev != -1 {
			maxGap = max(maxGap, minBucket[i]-maxBucket[prev])
		}
		prev = i
	}
	return maxGap
}

func testMaxGap(nums []int) {
	fmt.Printf("%v, get %d\n", nums, maximumGap(nums))
}

func main() {
	testMaxGap([]int{2, 2, 2, 2})
	testMaxGap([]int{1, 2, 8, 10})
	testMaxGap([]int{1, 2, 8, 9})
	testMaxGap([]int{1, 2, 8, 11})
}
