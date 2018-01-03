package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

var dirs = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}

func dfs(matrix [][]int, cache [][]int, x, y int) int {
	if cache[x][y] != 0 {
		return cache[x][y]
	}
	maxLen := 0
	for _, dir := range dirs {
		dx, dy := dir[0], dir[1]
		i, j := x+dx, y+dy
		if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) {
			continue
		}
		if matrix[x][y] >= matrix[i][j] {
			continue
		}
		curLen := dfs(matrix, cache, i, j)
		maxLen = max(maxLen, curLen)
	}
	maxLen++
	cache[x][y] = maxLen
	return maxLen
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m)
	for i, _ := range cache {
		cache[i] = make([]int, n)
	}
	maxLen := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pathLen := dfs(matrix, cache, i, j)
			maxLen = max(maxLen, pathLen)
		}
	}
	fmt.Println(cache)
	return maxLen
}

func testLongest(matrix [][]int) {
	fmt.Printf("%v, get %d\n", matrix, longestIncreasingPath(matrix))
}

func main() {
	matrix := [][]int{
		[]int{9, 4, 3},
		[]int{8, 5, 2},
		[]int{7, 6, 1},
	}
	testLongest(matrix)
}
