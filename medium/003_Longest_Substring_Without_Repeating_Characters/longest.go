package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	tbl := make(map[rune]int, 0)
	start, end, maxLen := 0, 0, 0
	for i, r := range s {
		end = i
		j, ok := tbl[r]
		if !ok || j < start {
			tbl[r] = i
			continue
		}
		maxLen = max(maxLen, end-start)
		start = j + 1
		tbl[r] = i
	}
	maxLen = max(maxLen, end-start+1)
	return maxLen
}

func test(s string) {
	maxLen := lengthOfLongestSubstring(s)
	fmt.Printf("%q, %d\n", s, maxLen)
}

func main() {
	test("")         // "", 0
	test("a")        // "a", 1
	test("abc")      // "abc", 3
	test("abba")     // "abba", 22
	test("abcabcbb") // "abc", 3
	test("bbbbb")    // "b", 1
	test("pwwkew")   //"wke", 3
}
