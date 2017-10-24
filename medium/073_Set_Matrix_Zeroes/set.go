package main

import "fmt"

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	rows := make([]int, len(matrix))
	cols := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows[i], cols[j] = 1, 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if rows[i] == 1 || cols[j] == 1 {
				matrix[i][j] = 0
			}
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
		[]int{1, 1, 1},
	}
	testSet(m)
}
