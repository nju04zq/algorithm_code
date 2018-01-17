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

func mergeWithFind(sum, temp []int, low, high, target int) int {
	if (low + 1) >= high {
		return math.MinInt32
	}
	//fmt.Printf("low %d, high %d\n", low, high)
	mid := low + (high-low)/2
	m1 := mergeWithFind(sum, temp, low, mid, target)
	m2 := mergeWithFind(sum, temp, mid, high, target)
	i, j, k := low, mid, 0
	m3 := math.MinInt32
	for i < mid && j < high {
		if sum[i]+target < sum[j] {
			i++
		} else {
			m3 = max(m3, sum[j]-sum[i])
			j++
		}
	}
	i, j, k = low, mid, 0
	for i < mid || j < high {
		if j >= high || (i < mid && sum[i] <= sum[j]) {
			temp[k] = sum[i]
			i++
		} else {
			temp[k] = sum[j]
			j++
		}
		k++
	}
	i, k = low, 0
	for i < high {
		sum[i] = temp[k]
		i++
		k++
	}
	//fmt.Printf("low %d, high %d, mid %d, m3 %d\n", low, high, mid, m3)
	m3 = max(m3, m2)
	m3 = max(m3, m1)
	return m3
}

func findMax(a []int, sum, temp []int, k int) int {
	m := len(a)
	for i := 1; i <= m; i++ {
		sum[i] = sum[i-1] + a[i-1]
	}
	maxSum := mergeWithFind(sum, temp, 0, len(sum), k)
	return maxSum
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	maxSum := math.MinInt32
	m, n := len(matrix), len(matrix[0])
	col := make([]int, m+1)
	temp := make([]int, m+1)
	for i := 0; i < n; i++ {
		jSum := make([]int, m)
		for j := i; j < n; j++ {
			for k := 0; k < m; k++ {
				jSum[k] += matrix[k][j]
			}
			//fmt.Println("jSum ", i, j, jSum)
			tempSum := findMax(jSum, temp, col, k)
			maxSum = max(maxSum, tempSum)
		}
	}
	return maxSum
}

func testMax(matrix [][]int, k int) {
	fmt.Printf("%v, k %d, get %d\n", matrix, k, maxSumSubmatrix(matrix, k))
}

func testFindMax() {
	k := 2
	a := []int{1, 1}
	sum := make([]int, len(a)+1)
	temp := make([]int, len(a)+1)
	res := findMax(a, sum, temp, k)
	fmt.Printf("find max %d\n", res)
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
