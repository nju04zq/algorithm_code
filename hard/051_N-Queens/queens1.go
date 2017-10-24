package main

import "fmt"

type matrix struct {
	buf  []byte
	n    int
	cols []bool
}

func (m *matrix) init(n int) *matrix {
	m.buf = make([]byte, n*n)
	for i := 0; i < len(m.buf); i++ {
		m.buf[i] = '.'
	}
	m.cols = make([]bool, n)
	m.n = n
	return m
}

func (m *matrix) index(i, j int) int {
	return i*m.n + j
}

func (m *matrix) getVal(i, j int) byte {
	return m.buf[m.index(i, j)]
}

func (m *matrix) setVal(i, j int, c byte) {
	m.buf[m.index(i, j)] = c
}

func (m *matrix) conflict(i, j, di, dj int) bool {
	getCnt := func(i, j int) int {
		cnt := 0
		for i >= 0 && i < m.n && j >= 0 && j < m.n {
			if m.getVal(i, j) == 'Q' {
				cnt++
				break
			}
			i += di
			j += dj
		}
		return cnt
	}
	cnt := getCnt(i, j)
	if cnt > 0 {
		return true
	}
	di, dj = -di, -dj
	cnt += getCnt(i, j)
	if cnt > 0 {
		return true
	}
	return false
}

func (m *matrix) set(i, j int) bool {
	if m.cols[j] {
		return false
	}
	if m.conflict(i, j, 1, 1) {
		return false
	}
	if m.conflict(i, j, -1, 1) {
		return false
	}
	m.setVal(i, j, 'Q')
	m.cols[j] = true
	return true
}

func (m *matrix) unset(i, j int) {
	m.setVal(i, j, '.')
	m.cols[j] = false
}

func (m *matrix) format() []string {
	res := make([]string, m.n)
	for i := 0; i < m.n*m.n; i += m.n {
		s := string(m.buf[i : i+m.n])
		res[i/m.n] = s
	}
	return res
}

func solveInternal(i, n int, m *matrix, res [][]string) [][]string {
	if i == n {
		res = append(res, m.format())
		return res
	}
	for j := 0; j < n; j++ {
		if !m.set(i, j) {
			continue
		}
		res = solveInternal(i+1, n, m, res)
		m.unset(i, j)
	}
	return res
}

func solveNQueens(n int) [][]string {
	m := new(matrix).init(n)
	res := make([][]string, 0)
	res = solveInternal(0, n, m, res)
	return res
}

func dumpRes(res [][]string) {
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			for k := 0; k < len(res[i][j]); k++ {
				fmt.Printf("%c ", res[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func testQueen(n int) {
	res := solveNQueens(n)
	fmt.Printf("n: %d\n", n)
	dumpRes(res)
}

func main() {
	testQueen(1)
	//testQueen(2)
	testQueen(4)
	testQueen(8)
}
