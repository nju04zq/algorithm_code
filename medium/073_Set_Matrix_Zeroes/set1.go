package main

import "fmt"

func findFirstZero(matrix [][]int) (int, int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	row, col := findFirstZero(matrix)
	if row == -1 {
		return
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][col] == 0 {
			matrix[i][col] = 1
		} else {
			matrix[i][col] = 0
		}
	}
	for i := 0; i < len(matrix[row]); i++ {
		if matrix[row][i] == 0 {
			matrix[row][i] = 1
		} else {
			matrix[row][i] = 0
		}
	}
	for i := 0; i < len(matrix); i++ {
		if i == row {
			continue
		}
		for j := 0; j < len(matrix[i]); j++ {
			if j == col {
				continue
			}
			if matrix[i][j] == 0 {
				matrix[i][col] = 1
				matrix[row][j] = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		if i == row {
			continue
		}
		for j := 0; j < len(matrix[i]); j++ {
			if j == col {
				continue
			}
			if matrix[i][col] == 1 || matrix[row][j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
	for i := 0; i < len(matrix[row]); i++ {
		matrix[row][i] = 0
	}
}

func dump(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func testSet(matrix [][]int) {
	fmt.Println("Before:")
	dump(matrix)
	fmt.Println("After:")
	setZeroes(matrix)
	dump(matrix)
}

func main() {
	m := [][]int{
		[]int{0, 1, 1},
		[]int{1, 1, 1},
		[]int{1, 1, 0},
	}
	testSet(m)
	m = [][]int{
		[]int{0, 0, 0, 5},
		[]int{4, 3, 1, 4},
		[]int{0, 1, 1, 4},
		[]int{1, 2, 1, 3},
		[]int{0, 0, 1, 1},
	}
	testSet(m)
}
