package main

import "fmt"

func dump(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}
}

type line struct {
	x, y int
}

func (l *line) getij(offset, k, i int) (int, int) {
	if l.x == 1 {
		return offset, offset + i
	} else if l.x == -1 {
		return offset + k - 1, offset + k - 1 - i
	} else if l.y == 1 {
		return offset + i, offset + k - 1
	} else {
		return offset + k - 1 - i, offset
	}
}

func swap(matrix [][]int, offset, k int, l1, l2 *line) {
	for i := 0; i < k-1; i++ {
		i1, j1 := l1.getij(offset, k, i)
		i2, j2 := l2.getij(offset, k, i)
		matrix[i1][j1], matrix[i2][j2] = matrix[i2][j2], matrix[i1][j1]
	}
}

func rotate(matrix [][]int) {
	l1, l2 := &line{}, &line{}
	offset, k := 0, len(matrix[0])
	for k > 1 {
		l1.x, l1.y = 1, 0
		l2.x, l2.y = 0, 1
		swap(matrix, offset, k, l1, l2)
		l2.x, l2.y = -1, 0
		swap(matrix, offset, k, l1, l2)
		l2.x, l2.y = 0, -1
		swap(matrix, offset, k, l1, l2)
		offset++
		k -= 2
	}
}

func testRotate(matrix [][]int) {
	fmt.Println("Before rotate.")
	dump(matrix)
	rotate(matrix)
	fmt.Println("After rotate.")
	dump(matrix)
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	testRotate(matrix)
	matrix = [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	testRotate(matrix)
}
