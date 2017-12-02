package main

import "fmt"

func initList(board [][]byte) ([]int, []int) {
	xList, yList := []int{}, []int{}
	if len(board) == 1 {
		for i := 0; i < len(board[0]); i++ {
			if board[0][i] == 'O' {
				xList = append(xList, 0)
				yList = append(yList, i)
			}
		}
		return xList, yList
	} else if len(board[0]) == 1 {
		for i := 0; i < len(board); i++ {
			if board[i][0] == 'O' {
				xList = append(xList, i)
				yList = append(yList, 0)
			}
		}
		return xList, yList
	}
	x, y := 0, 0
	for i := 0; i < len(board[0])-1; i++ {
		if board[x][y] == 'O' {
			xList = append(xList, x)
			yList = append(yList, y)
		}
		y++
	}
	for i := 0; i < len(board)-1; i++ {
		if board[x][y] == 'O' {
			xList = append(xList, x)
			yList = append(yList, y)
		}
		x++
	}
	for i := 0; i < len(board[0])-1; i++ {
		if board[x][y] == 'O' {
			xList = append(xList, x)
			yList = append(yList, y)
		}
		y--
	}
	for i := 0; i < len(board)-1; i++ {
		if board[x][y] == 'O' {
			xList = append(xList, x)
			yList = append(yList, y)
		}
		x--
	}
	fmt.Println(xList, yList)
	return xList, yList
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	xList, yList := initList(board)
	for len(xList) > 0 {
		cnt := len(xList)
		for i := 0; i < cnt; i++ {
			x, y := xList[i], yList[i]
			board[x][y] = 'Y'
			if x-1 >= 0 && board[x-1][y] == 'O' {
				xList = append(xList, x-1)
				yList = append(yList, y)
			}
			if x+1 < m && board[x+1][y] == 'O' {
				xList = append(xList, x+1)
				yList = append(yList, y)
			}
			if y-1 >= 0 && board[x][y-1] == 'O' {
				xList = append(xList, x)
				yList = append(yList, y-1)
			}
			if y+1 < n && board[x][y+1] == 'O' {
				xList = append(xList, x)
				yList = append(yList, y+1)
			}
		}
		xList = xList[cnt:]
		yList = yList[cnt:]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}

func dump(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func testSolve(board [][]byte) {
	fmt.Println("Before:")
	dump(board)
	solve(board)
	fmt.Println("After:")
	dump(board)
}

func main() {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X', 'X'},
		[]byte{'X', 'O', 'X', 'X', 'X'},
		[]byte{'X', 'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X', 'X'},
	}
	testSolve(board)
	board = [][]byte{[]byte{'O'}}
	testSolve(board)
}
