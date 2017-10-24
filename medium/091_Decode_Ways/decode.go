package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] == '0' {
				return 0
			}
			dp[i] = 1
			continue
		}
		if s[i] == '0' {
			if s[i-1] > '2' || s[i-1] == '0' {
				return 0
			}
			if i < 2 {
				dp[i] = 1
			} else {
				dp[i] = dp[i-2]
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			if i < 2 {
				dp[i] = dp[i-1] + 1
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(s)-1]
}

func testDecode(s string) {
	fmt.Printf("%q, %d\n", s, numDecodings(s))
}

func main() {
	testDecode("")
	testDecode("12")
	testDecode("1212")
	testDecode("0")
	testDecode("1010")
	testDecode("1060")
	testDecode("27")
}
