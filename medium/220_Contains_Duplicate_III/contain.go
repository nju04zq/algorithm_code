package main

import "fmt"
import "math"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k <= 0 || t < 0 {
		return false
	}
	tbl := make(map[int]int)
	for i, num := range nums {
		remapped := num + math.MaxInt32
		j := remapped / max(1, t)
		if _, ok := tbl[j]; ok {
			return true
		}
		if _, ok := tbl[j-1]; ok && num-tbl[j-1] <= t {
			return true
		}
		if _, ok := tbl[j+1]; ok && tbl[j+1]-num <= t {
			return true
		}
		tbl[j] = num
		if i >= k {
			j = (nums[i-k] + math.MaxInt32) / max(1, t)
			delete(tbl, j)
		}
	}
	return false
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func bf(nums []int, k, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if abs(nums[i]-nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
