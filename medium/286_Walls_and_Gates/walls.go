package main

import "fmt"

// You are given a m x n 2D grid initialized with these three possible values.
//
// 1. -1 - A wall or an obstacle.
// 2. 0 - A gate.
// 3. INF - Infinity means an empty room. We use the value 2**31 - 1 = 2147483647 to represent INF as you may assume that the distance to a gate is less than 2147483647.
// Fill each empty room with the distance to its nearest gate. If it is impossible to reach a gate, it should be filled with INF.
//
// For example, given the 2D grid:
// INF  -1  0  INF
// INF INF INF  -1
// INF  -1 INF  -1
//   0  -1 INF INF
// After running your function, the 2D grid should be:
//   3  -1   0   1
//   2   2   1  -1
//   1  -1   2  -1
//   0  -1   3   4

func nearest(grid [][]int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	m, n := len(grid), len(grid[0])
	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				dist[i][j] = -1
			}
		}
	}
	iIdx, jIdx := make([]int, 0), make([]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
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
				if grid[ii][jj] == 0 || grid[ii][jj] == -1 || dist[ii][jj] > 0 {
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

func main() {
	grid := [][]int{
		[]int{-2, -1, 0, -2},
		[]int{-2, -2, -2, -1},
		[]int{-2, -1, -2, -1},
		[]int{0, -1, -2, -2},
	}
	fmt.Printf("grid:\n%v\n", grid)
	fmt.Printf("get:\n%v\n", nearest(grid))
}
