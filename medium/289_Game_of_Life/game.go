package main

import "fmt"

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	prev := make([]int, n)
	cur := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cur[j] = board[i][j]
			cnt := 0
			if prev[j] == 1 {
				cnt++
			}
			if j > 0 {
				if cur[j-1] == 1 {
					cnt++
				}
				if prev[j-1] == 1 {
					cnt++
				}
			}
			if j < n-1 {
				if prev[j+1] == 1 {
					cnt++
				}
				if board[i][j+1] == 1 {
					cnt++
				}
			}
			if i < m-1 {
				if j > 0 && board[i+1][j-1] == 1 {
					cnt++
				}
				if j < n-1 && board[i+1][j+1] == 1 {
					cnt++
				}
				if board[i+1][j] == 1 {
					cnt++
				}
			}
			if cnt < 2 {
				board[i][j] = 0
			} else if cnt == 3 {
				board[i][j] = 1
			} else if cnt > 3 {
				board[i][j] = 0
			}
		}
		prev, cur = cur, prev
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
