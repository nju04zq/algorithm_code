package main

import "fmt"

type trie struct {
	root trieNode
}

type trieNode struct {
	exist    bool
	children [26]*trieNode
}

func (t *trie) add(s string) {
	node := &t.root
	for i := 0; i < len(s); i++ {
		j := s[i] - 'a'
		if node.children[j] == nil {
			node.children[j] = &trieNode{exist: true}
		}
		node = node.children[j]
	}
}

func (t *trie) search(s string) (bool, string) {
	node := &t.root
	for i := 0; i < len(s); i++ {
		j := s[i] - 'a'
		node = node.children[j]
		if node == nil {
			return true, s[:i]
		} else if node.exist == false {
			return false, s[:i]
		}
	}
	return true, s
}

func findInternal(board [][]byte, mask [][]bool, word string, t *trie,
	i, j, start int) bool {
	if start == len(word) {
		return true
	} else if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	} else if board[i][j] != word[start] {
		return false
	} else if mask[i][j] == true {
		return false
	}
	mask[i][j] = true
	if findInternal(board, mask, word, t, i-1, j, start+1) {
		return true
	}
	if findInternal(board, mask, word, t, i+1, j, start+1) {
		return true
	}
	if findInternal(board, mask, word, t, i, j-1, start+1) {
		return true
	}
	if findInternal(board, mask, word, t, i, j+1, start+1) {
		return true
	}
	mask[i][j] = false
	t.add(word[:start+1])
	return false
}

func findWord(board [][]byte, word string, t *trie) bool {
	valid, prefix := t.search(word)
	if !valid {
		return false
	} else if prefix == word {
		return true
	}
	mask := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		mask[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if findInternal(board, mask, word, t, i, j, 0) {
				t.add(word)
				return true
			}
		}
	}
	return false
}

func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	visited := make(map[string]bool)
	t := new(trie)
	for _, word := range words {
		if _, ok := visited[word]; ok {
			continue
		}
		if findWord(board, word, t) {
			res = append(res, word)
		}
		visited[word] = true
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
