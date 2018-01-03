package main

import "fmt"

type NumMatrix struct {
	buf [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := new(NumMatrix)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return *m
	}
	m.buf = make([][]int, len(matrix))
	temp := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		m.buf[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			temp[j] += matrix[i][j]
			if j > 0 {
				m.buf[i][j] = m.buf[i][j-1] + temp[j]
			} else {
				m.buf[i][j] = temp[j]
			}
		}
	}
	return *m
}

func (m *NumMatrix) sum(row, col int) int {
	if row < 0 || col < 0 {
		return 0
	} else {
		return m.buf[row][col]
	}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if len(m.buf) == 0 || len(m.buf[0]) == 0 {
		return 0
	}
	return m.sum(row2, col2) - m.sum(row1-1, col2) - m.sum(row2, col1-1) + m.sum(row1-1, col1-1)
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func testSum(matrix [][]int, row1, col1, row2, col2 int) {
	numMatrix := Constructor(matrix)
	m := &numMatrix
	fmt.Printf("matrix:\n%v\n", matrix)
	fmt.Printf("%d, %d, %d, %d, get %d\n", row1, col1, row2, col2, m.SumRegion(row1, col1, row2, col2))
}

func main() {
	m := [][]int{
		[]int{3, 0, 1, 4, 2},
		[]int{5, 6, 3, 2, 1},
		[]int{1, 2, 0, 1, 5},
		[]int{4, 1, 0, 1, 7},
		[]int{1, 0, 3, 0, 5},
	}
	testSum(m, 2, 1, 4, 3)
	testSum(m, 1, 1, 2, 2)
	testSum(m, 1, 2, 2, 4)
}
