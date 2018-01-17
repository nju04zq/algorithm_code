package main

import "fmt"

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
	}
	iIdx, jIdx := make([]int, 0), make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				iIdx = append(iIdx, i)
				jIdx = append(jIdx, j)
			}
		}
	}
	vecs := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	depth := 0
	for len(iIdx) > 0 {
		depth++
		cnt := len(iIdx)
		for k := 0; k < cnt; k++ {
			i, j := iIdx[k], jIdx[k]
			for _, vec := range vecs {
				ii, jj := i+vec[0], j+vec[1]
				if ii < 0 || ii >= m || jj < 0 || jj >= n {
					continue
				}
				if matrix[ii][jj] == 0 || dist[ii][jj] > 0 {
					continue
				}
				dist[ii][jj] = depth
				iIdx = append(iIdx, ii)
				jIdx = append(jIdx, jj)
			}
		}
		iIdx = iIdx[cnt:]
		jIdx = jIdx[cnt:]
	}
	return dist
}

func testMatrix(matrix [][]int) {
	fmt.Printf("Matrix:\n%v\n", matrix)
	fmt.Printf("Get:\n%v\n", updateMatrix(matrix))
}

func main() {
	matrix := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	testMatrix(matrix)
	matrix = [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{1, 1, 1},
	}
	testMatrix(matrix)
}
