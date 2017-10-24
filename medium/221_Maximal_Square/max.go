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
	var maxArea, n int
	dp := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == '1' {
			dp[i] = 1
			maxArea = 1
		}
	}
	//fmt.Println(dp)
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == '1' {
			dp[0] = 1
			maxArea = max(maxArea, 1)
		} else {
			dp[0] = 0
		}
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				dp[j] = 0
				continue
			}
			if dp[j] == dp[j-1] {
				n = dp[j]
				if matrix[i-n][j-n] == '1' {
					n++
				}
			} else {
				n = min(dp[j], dp[j-1]) + 1
			}
			dp[j] = n
			maxArea = max(maxArea, n*n)
		}
		//fmt.Println(dp)
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
	fmt.Printf("Get %d\n", maximalSquare(matrix))
}
