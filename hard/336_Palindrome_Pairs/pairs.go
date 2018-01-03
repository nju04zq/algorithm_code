package main

import "fmt"

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reversed(s string) string {
	n := len(s)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[n-1-i] = s[i]
	}
	return string(buf)
}

func palindromePairs(words []string) [][]int {
	res := make([][]int, 0)
	tbl := make(map[string]int)
	for i, w := range words {
		tbl[w] = i
	}
	for i, w := range words {
		for j := 0; j <= len(w); j++ {
			s1, s2 := w[:j], w[j:]
			if isPalindrome(s1) {
				if k, ok := tbl[reversed(s2)]; ok && k != i {
					res = append(res, []int{k, i})
				}
			}
			if len(s2) > 0 && isPalindrome(s2) {
				if k, ok := tbl[reversed(s1)]; ok && k != i {
					res = append(res, []int{i, k})
				}
			}
		}
	}
	return res
}

func testPalindrome(words []string) {
	fmt.Printf("%q, get %v\n", words, palindromePairs(words))
}

func main() {
	words := []string{
		"bat", "tab", "cat",
	}
	testPalindrome(words)
	words = []string{
		"abcd", "dcba", "lls", "s", "sssll",
	}
	testPalindrome(words)
	words = []string{
		"", "323",
	}
	testPalindrome(words)
}
