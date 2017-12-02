package main

import "fmt"

type graphNode struct {
	word      string
	visited   bool
	neighbors []*graphNode
	parents   []*graphNode
}

func makeGraph(words []string) map[string]*graphNode {
	graph := make(map[string]*graphNode)
	for _, word := range words {
		graph[word] = &graphNode{word: word}
	}
	for word, _ := range graph {
		wordBuf := []byte(word)
		neighbors := make([]*graphNode, 0)
		for i, _ := range wordBuf {
			temp := wordBuf[i]
			for ch := byte('a'); ch <= byte('z'); ch++ {
				if ch == temp {
					continue
				}
				wordBuf[i] = ch
				target := string(wordBuf)
				if _, ok := graph[target]; ok {
					neighbors = append(neighbors, graph[target])
				}
			}
			wordBuf[i] = temp
		}
		graph[word].neighbors = neighbors
		graph[word].parents = make([]*graphNode, 0)
	}
	return graph
}

func ladderLength(beginWord, endWord string, graph map[string]*graphNode) int {
	cur_level := map[*graphNode]bool{graph[beginWord]: true}
	next_level := map[*graphNode]bool{}
	length := 1
	found := false
	for len(cur_level) > 0 {
		length++
		for node, _ := range cur_level {
			node.visited = true
		}
		for node, _ := range cur_level {
			for _, neighbor := range node.neighbors {
				if neighbor.visited {
					continue
				}
				neighbor.parents = append(neighbor.parents, node)
				if neighbor.word == endWord {
					found = true
				}
				if _, ok := next_level[neighbor]; !ok {
					next_level[neighbor] = true
				}
			}
		}
		if found {
			return length
		}
		cur_level, next_level = next_level, map[*graphNode]bool{}
	}
	return 0
}

func reverseCopy(path []string) []string {
	n := len(path)
	res := make([]string, n)
	for i, word := range path {
		res[n-i-1] = word
	}
	return res
}

func dfs(beginWord string, graph map[string]*graphNode, node *graphNode, path []string, res [][]string) [][]string {
	if node.word == beginWord {
		return append(res, reverseCopy(path))
	}
	for _, parent := range node.parents {
		path = append(path, parent.word)
		res = dfs(beginWord, graph, parent, path, res)
		path = path[:len(path)-1]
	}
	return res
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := make([][]string, 0)
	wordList = append(wordList, beginWord)
	graph := makeGraph(wordList)
	length := ladderLength(beginWord, endWord, graph)
	if length == 0 {
		return res
	}
	//dumpParents(graph)
	path := []string{endWord}
	res = dfs(beginWord, graph, graph[endWord], path, res)
	return res
}

func dumpParents(graph map[string]*graphNode) {
	for _, node := range graph {
		fmt.Printf("%q: ", node.word)
		for _, parent := range node.parents {
			fmt.Printf("%q ", parent.word)
		}
		fmt.Println()
	}
}

func printGraph(graph map[string]*graphNode) {
	for word, node := range graph {
		fmt.Printf("%q: ", word)
		for _, neighbor := range node.neighbors {
			fmt.Printf("%q ", neighbor.word)
		}
		fmt.Println()
	}
}

func testFindLadders(beginWord, endWord string, wordList []string) {
	fmt.Printf("From %q, to %q, with %v:\n", beginWord, endWord, wordList)
	res := findLadders(beginWord, endWord, wordList)
	for _, path := range res {
		fmt.Printf("\t%v\n", path)
	}
}

func main() {
	words := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	testFindLadders("hit", "cog", words)
	words = []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}
	testFindLadders("red", "tax", words)
}
