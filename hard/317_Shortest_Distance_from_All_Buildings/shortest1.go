package main

import "fmt"
import "math"

// You want to build a house on an empty land which reaches all buildings in the shortest amount of distance. You can only move up, down, left and right. You are given a 2D grid of values 0, 1 or 2, where:
//
// Each 0 marks an empty land which you can pass by freely.
// Each 1 marks a building which you cannot pass through.
// Each 2 marks an obstacle which you cannot pass through.
// For example, given three buildings at (0,0), (0,4), (2,2), and an obstacle at (0,2):
//
// 1 - 0 - 2 - 0 - 1
// |   |   |   |   |
// 0 - 0 - 0 - 0 - 0
// |   |   |   |   |
// 0 - 0 - 1 - 0 - 0
// The point (1,2) is an ideal empty land to build a house, as the total travel distance of 3+3+1=7 is minimal. So return 7.
//
// Note:
// There will be at least one building. If it is not possible to build such house according to the above rules, return -1.
//

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

var vecs = [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}

func bfs(grid [][]int, i, j int, total [][]int, land int) int {
	m, n := len(grid), len(grid[0])
	idxi, idxj := []int{i}, []int{j}
	d, minDist := 0, -1
	for len(idxi) > 0 {
		d++
		cnt := len(idxi)
		for i := 0; i < cnt; i++ {
			for _, vec := range vecs {
				x, y := idxi[i]+vec[0], idxj[i]+vec[1]
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != land {
					continue
				}
				total[x][y] += d
				if minDist == -1 || minDist > total[x][y] {
					minDist = total[x][y]
				}
				grid[x][y]--
				idxi = append(idxi, x)
				idxj = append(idxj, y)
			}
		}
		idxi = idxi[cnt:]
		idxj = idxj[cnt:]
	}
	return minDist
}

func copyArray(a [][]int) [][]int {
	m, n := len(a), len(a[0])
	b := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			b[i][j] = a[i][j]
		}
	}
	return b
}

func shortest(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	d := math.MaxInt32
	grid = copyArray(grid)
	total := copyArray(grid)
	land := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			d = bfs(grid, i, j, total, land)
			if d == -1 {
				return -1
			}
			land--
		}
	}
	return d
}

func testShortest(grid [][]int) {
	fmt.Printf("grid: %v\nget %d\n", grid, shortest(grid))
}

func main() {
	grid := [][]int{
		[]int{1, 0, 2, 0, 1},
		[]int{0, 0, 0, 0, 0},
		[]int{0, 0, 1, 0, 0},
	}
	testShortest(grid)
	grid = [][]int{
		[]int{1, 2, 2},
		[]int{1, 2, 0},
		[]int{1, 1, 2},
	}
	testShortest(grid)
}
