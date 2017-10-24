package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] < '1' || s[0] > '9' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		j := (s[i-1]-'0')*10 + (s[i] - '0')
		if j > 9 && j <= 26 {
			dp[i+1] += dp[i-1]
		}
		if s[i] != '0' {
			dp[i+1] += dp[i]
		}
		if dp[i+1] == 0 {
			return 0
		}
	}
	return dp[len(s)]
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
	testDecode("100")
}
