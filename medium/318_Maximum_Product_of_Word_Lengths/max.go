package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getHash(s string) int {
	x := 0
	for i := 0; i < len(s); i++ {
		j := uint(s[i] - 'a')
		x |= (1 << j)
	}
	return x
}

func maxProduct(words []string) int {
	n := len(words)
	hash := make([]int, n)
	lens := make([]int, n)
	for i, word := range words {
		hash[i] = getHash(word)
		lens[i] = len(word)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if hash[i]&hash[j] == 0 {
				res = max(res, lens[i]*lens[j])
			}
		}
	}
	return res
}

func testMaxProduct(words []string) {
	fmt.Printf("%v, get %d\n", words, maxProduct(words))
}

func main() {
	words := []string{}
	words = []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	testMaxProduct(words)
	words = []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	testMaxProduct(words)
	words = []string{"a", "aa", "aaa", "aaaa"}
	testMaxProduct(words)
}
