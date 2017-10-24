package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minDistance(word1 string, word2 string) int {
	prev := make([]int, len(word2)+1)
	cur := make([]int, len(word2)+1)
	for i, _ := range prev {
		prev[i] = i
	}
	fmt.Println(prev)
	for i := 1; i <= len(word1); i++ {
		cur[0] = i
		for j := 1; j <= len(word2); j++ {
			cur[j] = cur[j-1] + 1
			cur[j] = min(cur[j], prev[j]+1)
			if word1[i-1] == word2[j-1] {
				cur[j] = min(cur[j], prev[j-1])
			} else {
				cur[j] = min(cur[j], prev[j-1]+1)
			}
		}
		prev, cur = cur, prev
	}
	return prev[len(word2)]
}

func testMinDistance(word1, word2 string) {
	fmt.Printf("word1 %q, word2 %q, min edit distance %d\n",
		word1, word2, minDistance(word1, word2))
}

func main() {
	testMinDistance("abc", "abc")
	testMinDistance("abc", "adc")
	testMinDistance("abc", "ac")
	testMinDistance("ac", "abc")
}
