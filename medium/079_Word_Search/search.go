package main

import "fmt"

func existInternal(board [][]byte, visited [][]bool, word string, i, j, start int) bool {
	if start == len(word) {
		return true
	} else if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	} else if visited[i][j] == true {
		return false
	} else if board[i][j] != word[start] {
		return false
	}
	visited[i][j] = true
	if existInternal(board, visited, word, i-1, j, start+1) {
		return true
	} else if existInternal(board, visited, word, i+1, j, start+1) {
		return true
	} else if existInternal(board, visited, word, i, j-1, start+1) {
		return true
	} else if existInternal(board, visited, word, i, j+1, start+1) {
		return true
	}
	visited[i][j] = false
	return false
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	if n == 0 {
		return false
	}
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if existInternal(board, visited, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func testExist(board [][]byte, word string, ans bool) {
	res := exist(board, word)
	if res != ans {
		panic(fmt.Errorf("%v, %s, get %t, should %t", board, word, res, ans))
	}
}

func main() {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	testExist(board, "SEE", true)
	testExist(board, "ABCCED", true)
	testExist(board, "SEE", true)
	testExist(board, "ABCB", false)
}
