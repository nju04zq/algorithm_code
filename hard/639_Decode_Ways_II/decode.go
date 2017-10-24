package main

import "fmt"

func valid(c byte) bool {
	if c == '*' || (c >= '0' && c <= '9') {
		return true
	} else {
		return false
	}
}

func numDecodings(s string) int {
	MOD := 1000000000 + 7
	if len(s) == 0 || !valid(s[0]) || s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] == '*' {
		dp[1] = 9
	} else {
		dp[1] = 1
	}
	for i := 1; i < len(s); i++ {
		if !valid(s[i]) {
			return 0
		} else if s[i] == '*' {
			dp[i+1] = dp[i] * 9
			if s[i-1] == '*' {
				dp[i+1] += dp[i-1] * 15
			} else if s[i-1] == '1' {
				dp[i+1] += dp[i-1] * 9
			} else if s[i-1] == '2' {
				dp[i+1] += dp[i-1] * 6
			}
		} else if s[i-1] == '*' {
			if s[i] != '0' {
				dp[i+1] = dp[i]
			}
			if s[i] <= '6' {
				dp[i+1] += (dp[i-1] * 2)
			} else {
				dp[i+1] += dp[i-1]
			}
		} else {
			if s[i] != '0' {
				dp[i+1] = dp[i]
			}
			j := (s[i-1]-'0')*10 + (s[i] - '0')
			if j >= 10 && j <= 26 {
				dp[i+1] += dp[i-1]
			}
		}
		if dp[i+1] == 0 {
			return 0
		}
		dp[i+1] %= MOD
	}
	return dp[len(s)]
}

func testDecode(s string) {
	fmt.Printf("%q, get %d\n", s, numDecodings(s))
}

func main() {
	testDecode("*")
	testDecode("1*")
	testDecode("12")
	testDecode("***")
	testDecode("*********")
	testDecode("********************")
}
