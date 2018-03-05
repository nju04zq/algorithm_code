package main

import "fmt"

func countSubstrings(s string) int {
	n := len(s)
	countAPI := func(left, right int) int {
		cnt := 0
		for left >= 0 && right < n {
			if s[left] != s[right] {
				break
			}
			cnt++
			left--
			right++
		}
		return cnt
	}
	cnt := 0
	for i := 0; i < n; i++ {
		cnt += countAPI(i, i)
		cnt += countAPI(i, i+1)
	}
	return cnt
}

func testCount(s string) {
	fmt.Printf("%q, get %d\n", s, countSubstrings(s))
}

func main() {
	testCount("a")
	testCount("abc")
	testCount("aaa")
	testCount("aaaaa")
}
