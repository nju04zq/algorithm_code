package main

import "fmt"

func dump(dp [][]bool) {
	n := len(dp)
	fmt.Printf("  ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < n; j++ {
			if dp[i][j] {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Println()
	}
}

func countSubstrings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	cnt := 0
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if s[j] != s[i] {
				continue
			}
			if j+1 >= i-1 || dp[j+1][i-1] {
				dp[j][i] = true
				cnt++
			}
		}
	}
	//dump(dp)
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
