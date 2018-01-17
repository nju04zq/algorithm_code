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
	n := len(s)
	tbl := make([]int, 255)
	for i := 0; i < len(tbl); i++ {
		tbl[i] = -1
	}
	for i := 0; i < n; i++ {
		j := int(s[i])
		if tbl[j] == -1 {
			tbl[j] = i
		} else {
			tbl[j] = n
		}
	}
	idx := n
	for i := 0; i < len(tbl); i++ {
		if j := tbl[i]; j >= 0 {
			idx = min(idx, j)
		}
	}
	if idx == n {
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
