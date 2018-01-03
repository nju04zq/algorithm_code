package main

import "fmt"

// Given a pattern and a string str, find if str follows the same pattern.
//
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty substring in str.
//
// Examples:
//
// pattern = "abab", str = "redblueredblue" should return true.
// pattern = "aaaa", str = "asdasdasdasd" should return true.
// pattern = "aabb", str = "xyzabcxzyabc" should return false.
//
//
// Notes:
// You may assume both pattern and str contains only lowercase letters.

func dfs(ptn, str string, t1 map[byte]string, t2 map[string]byte) bool {
	if len(ptn) == 0 && len(str) == 0 {
		return true
	} else if len(ptn) == 0 || len(str) == 0 {
		return false
	}
	s, ok := t1[ptn[0]]
	if ok {
		if len(s) > len(str) || str[:len(s)] != s {
			return false
		} else {
			return dfs(ptn[1:], str[len(s):], t1, t2)
		}
	}
	for i := 1; i <= len(str); i++ {
		s := str[:i]
		if _, ok := t2[s]; ok {
			continue
		}
		t1[ptn[0]] = s
		t2[s] = ptn[0]
		if dfs(ptn[1:], str[i:], t1, t2) {
			return true
		}
		delete(t1, ptn[0])
		delete(t2, s)
	}
	return false
}

func wordPattern(ptn string, str string) bool {
	t1 := make(map[byte]string)
	t2 := make(map[string]byte)
	return dfs(ptn, str, t1, t2)
}

func testWord(ptn, str string) {
	fmt.Printf("%q, %q, get %t\n", ptn, str, wordPattern(ptn, str))
}

func main() {
	testWord("abab", "redblueredblue")
	testWord("aaaa", "asdasdasdasd")
	testWord("aabb", "xyzabcxzyabc")
}
