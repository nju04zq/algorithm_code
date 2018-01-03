package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := -board[i][j]
			for k := max(i-1, 0); k <= min(i+1, m-1); k++ {
				for l := max(j-1, 0); l <= min(j+1, n-1); l++ {
					cnt += (board[k][l] & 0x1)
				}
			}
			if (cnt | board[i][j]) == 3 {
				board[i][j] |= 0x2
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}
}

func dumpBoard(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", board[i][j])
		}
		fmt.Println()
	}
}

func testGame(board [][]int) {
	fmt.Println("Before:")
	dumpBoard(board)
	gameOfLife(board)
	fmt.Println("After:")
	dumpBoard(board)
}

func main() {
	board := [][]int{
		[]int{1, 0, 1, 0},
		[]int{0, 1, 0, 1},
		[]int{1, 0, 1, 0},
		[]int{0, 1, 0, 1},
	}
	testGame(board)
}
