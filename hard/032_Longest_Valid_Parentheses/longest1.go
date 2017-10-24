package main

import "fmt"
import "math/rand"
import "time"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	dp := make([]int, len(s))
	longest := 0
	for i, c := range s {
		if i == 0 || c == '(' {
			continue
		}
		j := i - dp[i-1] - 1
		if j < 0 || s[j] != '(' {
			continue
		} else {
			dp[i] = dp[i-1] + 2
			if j-1 >= 0 {
				dp[i] += dp[j-1]
			}
		}
		longest = max(longest, dp[i])
	}
	return longest
}

func longestBF(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		var stack int
		for j := i; j < len(s); j++ {
			if s[j] == '(' {
				stack++
			} else {
				stack--
			}
			if stack < 0 {
				break
			} else if stack == 0 {
				longest = max(longest, j-i+1)
			}
		}
	}
	return longest
}

func testLongest(s string) {
	res := longestValidParentheses(s)
	ans := longestBF(s)
	if res != ans {
		fmt.Printf("%q: get %d, expect %d\n", s, res, ans)
	}
}

func makeRandStr() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % 20
	buf := make([]byte, len)
	for i, _ := range buf {
		if r.Int()%2 == 0 {
			buf[i] = '('
		} else {
			buf[i] = ')'
		}
	}
	return string(buf)
}

func main() {
	testLongest("(()")
	testLongest(")()())")
	var i int
	for i = 0; i < 100000; i++ {
		s := makeRandStr()
		testLongest(s)
	}
	fmt.Printf("%d test cases passed!\n", i)
}
