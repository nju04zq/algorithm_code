package main

import "fmt"

func insert(matrix [][]int, res []int, idx, i, j, k, di, dj int) int {
	for m := 0; m < k-1; m++ {
		res[idx] = matrix[i][j]
		i += di
		j += dj
		idx++
	}
	return idx
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	oi, oj, idx := 0, 0, 0
	for m > 1 && n > 1 {
		idx = insert(matrix, res, idx, oi, oj, n, 0, 1)
		idx = insert(matrix, res, idx, oi, oj+n-1, m, 1, 0)
		idx = insert(matrix, res, idx, oi+m-1, oj+n-1, n, 0, -1)
		idx = insert(matrix, res, idx, oi+m-1, oj, m, -1, 0)
		oi++
		oj++
		m -= 2
		n -= 2
	}
	if m > 1 && n > 0 {
		idx = insert(matrix, res, idx, oi, oj, m+1, 1, 0)
	} else if n > 1 && m > 0 {
		idx = insert(matrix, res, idx, oi, oj, n+1, 0, 1)
	} else if m == 1 && n == 1 {
		res[idx] = matrix[oi][oj]
	}
	return res
}

func dumpMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		fmt.Println("Null")
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func testSpiral(matrix [][]int) {
	dumpMatrix(matrix)
	res := spiralOrder(matrix)
	fmt.Println(res)
}

func main() {
	a := [][]int{}
	testSpiral(a)
	a = [][]int{
		[]int{1},
	}
	testSpiral(a)
	a = [][]int{
		[]int{1, 2, 3, 4},
	}
	testSpiral(a)
	a = [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}
	testSpiral(a)
	a = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}
	testSpiral(a)
	a = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
		[]int{10, 11, 12},
	}
	testSpiral(a)
	a = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	testSpiral(a)
	a = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
		[]int{13, 14, 15, 16},
	}
	testSpiral(a)
	a = [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
	}
	testSpiral(a)
}
