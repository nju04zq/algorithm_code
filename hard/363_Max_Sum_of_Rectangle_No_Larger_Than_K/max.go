// Time Limit Exceeded

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

func findMax(a []int, sum []int, k int) int {
	m := len(a)
	for i := 0; i < m; i++ {
		if i == 0 {
			sum[i] = a[i]
		} else {
			sum[i] = sum[i-1] + a[i]
		}
	}
	maxSum := math.MinInt32
	for i := 0; i < m; i++ {
		for j := 0; j <= i; j++ {
			x := 0
			if j == 0 {
				x = sum[i]
			} else {
				x = sum[i] - sum[j-1]
			}
			if x <= k {
				maxSum = max(maxSum, x)
			}
		}
	}
	return maxSum
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	maxSum := math.MinInt32
	m, n := len(matrix), len(matrix[0])
	col := make([]int, m)
	for i := 0; i < n; i++ {
		jSum := make([]int, m)
		for j := i; j < n; j++ {
			for k := 0; k < m; k++ {
				jSum[k] += matrix[k][j]
			}
			//fmt.Println("jSum ", i, j, jSum)
			tempSum := findMax(jSum, col, k)
			maxSum = max(maxSum, tempSum)
		}
	}
	return maxSum
}

func testMax(matrix [][]int, k int) {
	fmt.Printf("%v, k %d, get %d\n", matrix, k, maxSumSubmatrix(matrix, k))
}

func main() {
	matrix := [][]int{
		[]int{1, 0, 1},
		[]int{0, -2, 3},
	}
	testMax(matrix, 2)
	testMax(matrix, 3)
	matrix = [][]int{
		[]int{5, -4, -3, 4},
		[]int{-3, -4, 4, 5},
		[]int{5, 1, 5, -4},
	}
	testMax(matrix, 8)
}
