package main

import "fmt"

func dfs(grid [][]byte, i, j int, ch byte) {
	if grid[i][j] != '1' {
		return
	}
	m, n := len(grid), len(grid[0])
	grid[i][j] = ch
	if i-1 >= 0 {
		dfs(grid, i-1, j, ch)
	}
	if i+1 < m {
		dfs(grid, i+1, j, ch)
	}
	if j-1 >= 0 {
		dfs(grid, i, j-1, ch)
	}
	if j+1 < n {
		dfs(grid, i, j+1, ch)
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j, '2')
				cnt++
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '2' {
				grid[i][j] = '1'
			}
		}
	}
	return cnt
}

func testIslands(grid [][]byte) {
	fmt.Printf("grid %v, get %d\n", grid, numIslands(grid))
}

func main() {
	grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}
	testIslands(grid)
	grid = [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}
	testIslands(grid)
}
