package main

import "fmt"

func getWord(s string, i int) int {
	if i == len(s) {
		return -1
	}
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			break
		}
	}
	return i
}

func skipSpaces(s string, start int) int {
	var i int
	for i = start; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	return i
}

func wordPattern(pattern string, str string) bool {
	t1 := make(map[byte]string)
	t2 := make(map[string]byte)
	var i, start int
	for i, start = 0, 0; i < len(pattern); i++ {
		ch := pattern[i]
		start = skipSpaces(str, start)
		j := getWord(str, start)
		if j == -1 {
			return false
		}
		word := str[start:j]
		w1, ok1 := t1[ch]
		c2, ok2 := t2[word]
		if !ok1 && !ok2 {
			t1[ch], t2[word] = word, ch
		} else if !ok1 || !ok2 {
			return false
		} else if word != w1 || ch != c2 {
			return false
		}
		start = j
	}
	if i < len(pattern) {
		return false
	}
	start = skipSpaces(str, start)
	if start != len(str) {
		return false
	}
	return true
}

func testWord(ptn string, str string) {
	fmt.Printf("%q, %q, get %t\n", ptn, str, wordPattern(ptn, str))
}

func main() {
	testWord("abba", "dog cat cat dog")
	testWord("abba", "dog cat cat fish")
	testWord("aaaa", "dog cat cat dog")
	testWord("abba", "dog dog dog dog")
}
