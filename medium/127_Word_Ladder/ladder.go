package main

import "fmt"

func makeGraph(words []string) map[string][]string {
	graph := make(map[string][]string)
	for _, word := range words {
		graph[word] = make([]string, 0)
	}
	for _, word := range words {
		neighbors := make([]string, 0)
		wordBuf := []byte(word)
		for i, _ := range wordBuf {
			temp := wordBuf[i]
			for ch := byte('a'); ch <= byte('z'); ch++ {
				if ch == temp {
					continue
				}
				wordBuf[i] = ch
				target := string(wordBuf)
				if _, ok := graph[target]; ok {
					neighbors = append(neighbors, target)
				}
			}
			wordBuf[i] = temp
		}
		graph[word] = neighbors
	}
	return graph
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	graph := makeGraph(wordList)
	visited := make(map[string]bool)
	words := []string{beginWord}
	visited[beginWord] = true
	length := 1
	for len(words) > 0 {
		//fmt.Println("==========")
		length++
		cnt := len(words)
		for i := 0; i < cnt; i++ {
			word := words[i]
			//fmt.Printf("%q: ", word)
			for _, neighbor := range graph[word] {
				//fmt.Printf("%q ", neighbor)
				if neighbor == endWord {
					return length
				}
				if !visited[neighbor] {
					words = append(words, neighbor)
					visited[neighbor] = true
				}
			}
		}
		words = words[cnt:]
		fmt.Println()
	}
	return 0
}

func testLadderLength(beginWord, endWord string, wordList []string) {
	fmt.Printf("\nFrom %q, to %q, in %v, get %d\n",
		beginWord, endWord, wordList,
		ladderLength(beginWord, endWord, wordList))
}

func main() {
	words := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	testLadderLength("hit", "cog", words)
}
