package main

import "fmt"

func insert(m [][]int, idx, i, j, n, di, dj int) int {
	for k := 0; k < n-1; k++ {
		m[i][j] = idx
		i += di
		j += dj
		idx++
	}
	return idx
}

func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i, _ := range m {
		m[i] = make([]int, n)
	}
	oi, oj, idx := 0, 0, 1
	for ; n > 1; n -= 2 {
		idx = insert(m, idx, oi, oj, n, 0, 1)
		idx = insert(m, idx, oi, oj+n-1, n, 1, 0)
		idx = insert(m, idx, oi+n-1, oj+n-1, n, 0, -1)
		idx = insert(m, idx, oi+n-1, oj, n, -1, 0)
		oi++
		oj++
	}
	if n == 1 {
		m[oi][oj] = idx
	}
	return m
}

func testSpiral(n int) {
	fmt.Printf("n: %d\n", n)
	res := generateMatrix(n)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%2d ", res[i][j])
		}
		fmt.Println()
	}
}

func main() {
	testSpiral(0)
	testSpiral(1)
	testSpiral(2)
	testSpiral(3)
	testSpiral(4)
}
