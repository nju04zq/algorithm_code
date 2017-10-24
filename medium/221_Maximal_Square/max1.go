package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var maxLen int
	dp0 := make([]int, len(matrix[0])+1)
	dp1 := make([]int, len(matrix[0])+1)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				dp1[j+1] = 0
			} else {
				dp1[j+1] = min(min(dp1[j], dp0[j+1]), dp0[j]) + 1
			}
			maxLen = max(maxLen, dp1[j+1])
		}
		dp0, dp1 = dp1, dp0
	}
	return maxLen * maxLen
}

func main() {
	matrix := [][]byte{
		[]byte{'1', '0', '1', '0', '0'},
		[]byte{'1', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1'},
		[]byte{'1', '0', '0', '1', '0'},
	}
	fmt.Printf("Get %d\n", maximalSquare(matrix))
}
