package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([]bool, len(s2)+1)
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 && j == 0 {
				dp[j] = true
				continue
			}
			if i >= 1 && s1[i-1] == s3[i+j-1] && dp[j] {
				dp[j] = true
			} else {
				dp[j] = false
			}
			if dp[j] {
				continue
			}
			if j >= 1 && s2[j-1] == s3[i+j-1] {
				dp[j] = dp[j-1]
			}
		}
	}
	return dp[len(s2)]
}

func test(s1, s2, s3 string) {
	rc := isInterleave(s1, s2, s3)
	fmt.Printf("%q, %q, %q, %t\n", s1, s2, s3, rc)
}

func main() {
	test("aabcc", "dbbca", "aadbbcbcac")
	test("aabcc", "dbbca", "aadbbbaccc")
	test("a", "b", "a")
	test("ab", "bc", "babc")
}
