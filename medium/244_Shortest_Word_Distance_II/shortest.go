package main

import "fmt"
import "math"

// This is a follow up of Shortest Word Distance. The only difference is now you are given the list of words and your method will be called repeatedly many times with different parameters. How would you optimize it?
//
// Design a class which receives a list of words in the constructor, and implements a method that takes two words word1 and word2 and return the shortest distance between these two words in the list.
//
// For example,
// Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
//
// Given word1 = “coding”, word2 = “practice”, return 3.
// Given word1 = "makes", word2 = "coding", return 1.
//
// Note:
// You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

type Shortest struct {
	tbl map[string][]int
}

func (s *Shortest) Init(words []string) *Shortest {
	s.tbl = make(map[string][]int)
	for i, word := range words {
		if _, ok := s.tbl[word]; ok {
			s.tbl[word] = append(s.tbl[word], i)
		} else {
			s.tbl[word] = []int{i}
		}
	}
	return s
}

func (s *Shortest) Calc(word1, word2 string) int {
	idxs1, idxs2 := s.tbl[word1], s.tbl[word2]
	i, j, d := 0, 0, math.MaxInt32
	for i < len(idxs1) && j < len(idxs2) {
		d = min(d, abs(idxs1[i]-idxs2[j]))
		if idxs1[i] < idxs2[j] {
			i++
		} else {
			j++
		}
	}
	return d
}

func testShortest(s *Shortest, word1, word2 string) {
	fmt.Printf("word1 %q, word2 %q, get %d\n", word1, word2, s.Calc(word1, word2))
}

func main() {
	words := []string{"practice", "makes", "perfect", "coding", "makes"}
	s := new(Shortest).Init(words)
	testShortest(s, "coding", "practice")
	testShortest(s, "makes", "coding")
}
