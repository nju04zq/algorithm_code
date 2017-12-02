package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	tbl := make(map[string]bool)
	for _, word := range wordDict {
		tbl[word] = true
	}
	dp := make([]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j <= len(s); j++ {
			if _, ok := tbl[s[i:j]]; !ok {
				continue
			}
			if j == len(s) || dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[0]
}

func testBreak(s string, wordDict []string) {
	fmt.Printf("%q, dict %v, get %t\n", s, wordDict, wordBreak(s, wordDict))
}

func main() {
	testBreak("leetcode", []string{"leet", "code"})
}
