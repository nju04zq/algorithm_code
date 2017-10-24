package main

import (
	"fmt"
	"math/rand"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getCap(heights []int, left int, right int) int {
	return min(heights[left], heights[right]) * (right - left)
}

func maxArea(heights []int) int {
	left, right := 0, len(heights)-1
	capMax := 0
	for left < right {
		cap := getCap(heights, left, right)
		if cap > capMax {
			capMax = cap
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return capMax
}

func MakeRandArray() []int {
	maxLen, maxElement := 100, 100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func maxAreaBf(heights []int) int {
	capMax := 0
	for i := 0; i < len(heights); i++ {
		for j := i; j < len(heights); j++ {
			cap := getCap(heights, i, j)
			if cap > capMax {
				capMax = cap
			}
		}
	}
	return capMax
}

func testMaxArea() bool {
	a := MakeRandArray()
	ans := maxAreaBf(a)
	res := maxArea(a)
	if ans != res {
		fmt.Printf("For %v, expect %d, get %d\n", a, ans, res)
		return false
	}
	return true
}

func main() {
	for i := 0; i < 1000; i++ {
		res := testMaxArea()
		if !res {
			break
		}
	}
}
