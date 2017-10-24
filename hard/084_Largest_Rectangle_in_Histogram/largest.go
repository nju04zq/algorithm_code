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
	var height, top, width, maxArea, j int
	heights = append(heights, 0)
	stack := make([]int, 0)
	for i := 0; i < len(heights); {
		height = heights[i]
		top = len(stack) - 1
		if top < 0 || height > heights[stack[top]] {
			stack = append(stack, i)
			i++
		} else {
			j, stack = stack[top], stack[:top]
			top = len(stack) - 1
			if top < 0 {
				width = i
			} else {
				width = i - stack[top] - 1
			}
			maxArea = max(maxArea, width*heights[j])
		}
	}
	return maxArea
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	fmt.Printf("heights %v, get %d\n", heights, largestRectangleArea(heights))
}
