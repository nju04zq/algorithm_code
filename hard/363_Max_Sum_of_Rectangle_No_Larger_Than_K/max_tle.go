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

func buildMatrixSum(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	sum := make([][]int, m)
	jSum := make([]int, n)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
		for j := 0; j < n; j++ {
			jSum[j] += matrix[i][j]
			if j == 0 {
				sum[i][j] = jSum[j]
			} else {
				sum[i][j] = jSum[j] + sum[i][j-1]
			}
		}
	}
	return sum
}

func getMatrixSum(sum [][]int, ui, uj, bi, bj int) int {
	m, n := len(sum), len(sum[0])
	get := func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		} else {
			return sum[i][j]
		}
	}
	return get(bi, bj) - get(ui-1, bj) - get(bi, uj-1) + get(ui-1, uj-1)
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	sum := buildMatrixSum(matrix)
	m, n := len(matrix), len(matrix[0])
	maxSum := math.MinInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			bi, bj := i, j
			for ui := i; ui >= 0; ui-- {
				for uj := j; uj >= 0; uj-- {
					tempSum := getMatrixSum(sum, ui, uj, bi, bj)
					if tempSum <= k {
						maxSum = max(maxSum, tempSum)
						//if maxSum == tempSum {
						//	fmt.Println("max: ", maxSum, ui, uj, bi, bj)
						//}
					}
				}
			}
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
