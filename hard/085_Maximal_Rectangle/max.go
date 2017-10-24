package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func largest(heights []int) int {
	var top, width, j, maxArea int
	stack := make([]int, 0)
	for i := 0; i < len(heights); {
		top = len(stack) - 1
		if top < 0 || heights[i] > heights[stack[top]] {
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

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxArea := 0
	heights := make([]int, len(matrix[0])+1)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j] += 1
			}
		}
		maxArea = max(maxArea, largest(heights))
	}
	return maxArea
}

func main() {
	matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
	matrix = [][]byte{
		[]byte{'0', '0', '0'},
		[]byte{'0', '0', '0'},
		[]byte{'1', '1', '1'},
	}
	fmt.Println(maximalRectangle(matrix))
	matrix = [][]byte{
		[]byte{'0', '0', '0', '0', '1', '1', '1', '0', '1'},
		[]byte{'0', '0', '1', '1', '1', '1', '1', '0', '1'},
		[]byte{'0', '0', '0', '1', '1', '1', '1', '1', '0'},
	}
	fmt.Println(maximalRectangle(matrix))
}
