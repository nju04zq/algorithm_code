package main

import "fmt"

func searchInRow(matrix [][]int, x, y, cols, target int) int {
	row := matrix[x]
	low, high := y, cols-1
	for low < high {
		mid := high - (high-low)/2
		if row[mid] == target {
			return mid
		} else if row[mid] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}
	if row[low] <= target {
		return low
	} else {
		return -1
	}
}

func searchInternal(matrix [][]int, x, y, rows, cols int, target int) bool {
	if rows == 0 || cols == 0 {
		return false
	}
	k := searchInRow(matrix, x, y, cols, target)
	if k == -1 {
		return false
	} else if matrix[x][k] == target {
		return true
	}
	return searchInternal(matrix, x+1, y, rows-1, k+1, target)
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	return searchInternal(matrix, 0, 0, m, n, target)
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
