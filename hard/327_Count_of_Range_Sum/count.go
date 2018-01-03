package main

import "fmt"

func countRange(sums, counts, temp []int, upper, left, mid, right, sign int) {
	i, j := left, mid
	//fmt.Printf("sums %v, left %d, mid %d, right %d\n", sums, left, mid, right)
	for i < mid && j < right {
		for j < right && sums[j]-sums[i] < upper {
			j++
		}
		counts[i] += (sign * (right - j))
		i++
	}
}

func merge(sums, temp []int, left, mid, right int) {
	i, j, k := left, mid, 0
	for i < mid || j < right {
		//fmt.Println("merge in", i, j, left, mid, right)
		if j >= right || (i < mid && sums[i] < sums[j]) {
			temp[k] = sums[i]
			i++
		} else if i >= mid || sums[j] <= sums[i] {
			temp[k] = sums[j]
			j++
		}
		k++
	}
	for i, k := left, 0; i < right; i++ {
		sums[i] = temp[k]
		k++
	}
}

func mergeCount(sums, counts, temp []int, lower, upper, left, right int) {
	if left >= right-1 {
		return
	}
	mid := left + (right-left)/2
	mergeCount(sums, counts, temp, lower, upper, left, mid)
	mergeCount(sums, counts, temp, lower, upper, mid, right)
	countRange(sums, counts, temp, lower, left, mid, right, 1)
	countRange(sums, counts, temp, upper+1, left, mid, right, -1)
	merge(sums, temp, left, mid, right)
}

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = nums[i-1] + sums[i-1]
	}
	temp := make([]int, n+1)
	counts := make([]int, n+1)
	mergeCount(sums, counts, temp, lower, upper, 0, n+1)
	total := 0
	for i := 0; i <= n; i++ {
		total += counts[i]
	}
	return total
}

func testCount(nums []int, lower, upper int) {
	fmt.Printf("%v, lower %d, upper %d, get %d\n", nums, lower, upper, countRangeSum(nums, lower, upper))
}

func main() {
	testCount([]int{0}, 0, 0)
	testCount([]int{0, 0, 0}, 0, 0)
	testCount([]int{-2, 5, -1}, -2, 2)
}
