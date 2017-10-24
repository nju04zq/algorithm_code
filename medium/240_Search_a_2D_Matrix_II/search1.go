package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	x, y := m-1, 0
	for x >= 0 && y < n {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			y++
		} else {
			x--
		}
	}
	return false
}

func dump(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func testSearch(matrix [][]int, target int, ans bool) {
	res := searchMatrix(matrix, target)
	if res != ans {
		dump(matrix)
		fmt.Printf("search %d, get %t, should %t\n", target, res, ans)
	}
}

func main() {
	m := [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	testSearch(m, 0, false)
	testSearch(m, 1, true)
	testSearch(m, 7, true)
	testSearch(m, 3, true)
	testSearch(m, 9, true)
	testSearch(m, 24, true)
	testSearch(m, 30, true)
	testSearch(m, 20, false)
	testSearch(m, 40, false)
}
