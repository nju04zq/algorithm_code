package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var i int
	endLoop := false
	for i = 0; ; i++ {
		for j, s := range strs {
			if i >= len(s) {
				endLoop = true
				break
			} else if j == 0 {
				continue
			} else if s[i] != strs[j-1][i] {
				endLoop = true
				break
			}
		}
		if endLoop {
			break
		}
	}
	return strs[0][:i]
}

func testLongestCommonPrefix(strs []string) {
	s := longestCommonPrefix(strs)
	fmt.Printf("%v, %q\n", strs, s)
}

func main() {
	testLongestCommonPrefix([]string{"", "abc"})
	testLongestCommonPrefix([]string{"abc", ""})
	testLongestCommonPrefix([]string{"abc", "123"})
	testLongestCommonPrefix([]string{"abc", "a"})
	testLongestCommonPrefix([]string{"abc", "ab"})
	testLongestCommonPrefix([]string{"abc", "abc"})
	testLongestCommonPrefix([]string{"abc", "abcd"})
}
