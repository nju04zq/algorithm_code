package main

import "fmt"
import "math"

// This is a follow up of Shortest Word Distance. The only difference is now word1 could be the same as word2.
//
// Given a list of words and two words word1 and word2, return the shortest distance between these two words in the list.
//
// word1 and word2 may be the same and they represent two individual words in the list.
//
// For example,
// Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
//
// Given word1 = “makes”, word2 = “coding”, return 1.
// Given word1 = "makes", word2 = "makes", return 3.
//
// Note:
// You may assume word1 and word2 are both in the list.

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func shortest(word1, word2 string, words []string) int {
	theSame := false
	if word1 == word2 {
		theSame = true
	}
	i1, i2, d := -1, -1, math.MaxInt32
	for i, word := range words {
		if theSame {
			if word != word1 {
				continue
			}
			if i1 == -1 {
				i1 = i
			} else if i2 == -1 {
				i2 = i
			} else {
				i2 = i1
				i1 = i
			}
		} else {
			if word == word1 {
				i1 = i
			} else if word == word2 {
				i2 = i
			} else {
				continue
			}
		}
		if i1 != -1 && i2 != -1 {
			d = min(d, abs(i1-i2))
		}
	}
	return d
}

func testShortest(word1, word2 string, words []string) {
	fmt.Printf("%q, %q, %v, get %d\n", word1, word2, words, shortest(word1, word2, words))
}

func main() {
	words := []string{"practice", "makes", "perfect", "coding", "makes"}
	testShortest("coding", "practice", words)
	testShortest("makes", "coding", words)
	testShortest("makes", "makes", words)
}
