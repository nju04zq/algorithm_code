package main

import "fmt"

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	row0, col0 := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			row0 = true
			break
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if row0 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	if col0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
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
