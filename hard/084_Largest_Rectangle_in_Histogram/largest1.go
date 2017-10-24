package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	var maxArea, p int
	lessLeft := make([]int, len(heights))
	lessRight := make([]int, len(heights))
	lessLeft[0], lessRight[len(heights)-1] = -1, len(heights)
	for i := 1; i < len(heights); i++ {
		p = i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = lessLeft[p]
		}
		lessLeft[i] = p
	}
	for i := len(heights) - 1; i >= 0; i-- {
		p = i + 1
		for p < len(heights) && heights[p] >= heights[i] {
			p = lessRight[p]
		}
		lessRight[i] = p
	}
	for i, height := range heights {
		maxArea = max(maxArea, height*(lessRight[i]-lessLeft[i]-1))
	}
	return maxArea
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Printf("heights %v, get %d\n", heights, largestRectangleArea(heights))
}
