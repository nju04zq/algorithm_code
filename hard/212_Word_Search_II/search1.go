package main

import "fmt"

type trieNode struct {
	word     string
	children [26]*trieNode
}

func buildTrie(words []string) *trieNode {
	root := new(trieNode)
	for _, word := range words {
		p := root
		for i := 0; i < len(word); i++ {
			j := word[i] - 'a'
			if p.children[j] == nil {
				p.children[j] = new(trieNode)
			}
			p = p.children[j]
		}
		p.word = word
	}
	return root
}

func dfs(board [][]byte, p *trieNode, i, j int, res []string) []string {
	k := board[i][j] - 'a'
	if board[i][j] == '#' || p.children[k] == nil {
		return res
	}
	p = p.children[k]
	if p.word != "" {
		res = append(res, p.word)
		p.word = ""
	}
	c := board[i][j]
	board[i][j] = '#'
	if i-1 >= 0 {
		res = dfs(board, p, i-1, j, res)
	}
	if i+1 < len(board) {
		res = dfs(board, p, i+1, j, res)
	}
	if j-1 >= 0 {
		res = dfs(board, p, i, j-1, res)
	}
	if j+1 < len(board[0]) {
		res = dfs(board, p, i, j+1, res)
	}
	board[i][j] = c
	return res
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 || len(words) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	t := buildTrie(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			res = dfs(board, t, i, j, res)
		}
	}
	return res
}

func testFind(board [][]byte, words []string) {
	fmt.Printf("words %v, find %v\n", words, findWords(board, words))
}

func main() {
	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	testFind(board, words)
	board = [][]byte{
		[]byte{'a', 'b'},
		[]byte{'a', 'a'},
	}
	words = []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}
	testFind(board, words)
}
