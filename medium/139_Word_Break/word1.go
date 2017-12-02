package main

import "fmt"
import "strings"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		for _, word := range wordDict {
			if !strings.HasPrefix(s[i:], word) {
				continue
			}
			j := i + len(word)
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
