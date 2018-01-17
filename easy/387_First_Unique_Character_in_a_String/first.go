package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	tbl := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := tbl[s[i]]; ok {
			tbl[s[i]] = -1
		} else {
			tbl[s[i]] = i
		}
	}
	idx := len(s)
	for _, val := range tbl {
		if val != -1 {
			idx = min(idx, val)
		}
	}
	if idx == len(s) {
		return -1
	} else {
		return idx
	}
}

func testFirst(s string) {
	fmt.Printf("%q, get %d\n", s, firstUniqChar(s))
}

func main() {
	testFirst("")
	testFirst("leetcode")
	testFirst("loveleetcode")
}
