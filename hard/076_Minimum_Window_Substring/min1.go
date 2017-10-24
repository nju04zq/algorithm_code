// Replace hash map with int array, improve 66ms to 6ms

package main

import "fmt"

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var res string
	var total, minLen int
	tbl := make([]int, 256)
	for i, _ := range t {
		tbl[t[i]]++
	}
	start := -1
	occurence := make([]int, 256)
	for i, _ := range s {
		ch := s[i]
		if tbl[ch] == 0 {
			continue
		}
		if start == -1 {
			start = i
		}
		occurence[ch]++
		if occurence[ch] <= tbl[ch] {
			total++
		}
		if total != len(t) {
			continue
		}
		for ; start <= i; start++ {
			curLen := i - start + 1
			if minLen == 0 || curLen < minLen {
				minLen = curLen
				res = s[start : i+1]
			}
			ch := s[start]
			if tbl[ch] == 0 {
				continue
			}
			if occurence[ch] <= tbl[ch] {
				break
			}
			occurence[ch]--
		}
	}
	return res
}

func testMinWindow(s string, t string) {
	ans := minWindow(s, t)
	fmt.Printf("s %q, t %q, ans %q\n", s, t, ans)
}

func main() {
	testMinWindow("1", "1")           //1
	testMinWindow("1", "2")           //""
	testMinWindow("201213", "123")    //213
	testMinWindow("20120213", "1223") //20213
}
