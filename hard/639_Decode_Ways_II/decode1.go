package main

import "fmt"

func buildTable() map[string]int {
	set, s := "1234567890*", ""
	buf := make([]byte, 2)
	tbl := make(map[string]int)
	for i, _ := range set {
		buf[0] = set[i]
		s = string(buf[0:1])
		if buf[0] == '*' {
			tbl[s] = 9
		} else if buf[0] >= '1' && buf[0] <= '9' {
			tbl[s] = 1
		}
		for j, _ := range set {
			buf[1] = set[j]
			s = string(buf)
			if buf[0] == '*' {
				if buf[1] == '*' {
					tbl[s] = 15
				} else if buf[1] >= '0' && buf[1] <= '6' {
					tbl[s] = 2
				} else {
					tbl[s] = 1
				}
			} else if buf[0] == '1' {
				if buf[1] == '*' {
					tbl[s] = 9
				} else {
					tbl[s] = 1
				}
			} else if buf[0] == '2' {
				if buf[1] == '*' {
					tbl[s] = 6
				} else if buf[1] >= '0' && buf[1] <= '6' {
					tbl[s] = 1
				}
			}
		}
	}
	return tbl
}

func numDecodings(s string) int {
	MOD := 1000000000 + 7
	ans, dp1, dp2 := 0, 1, 1
	buf := []byte{'-', ' '}
	tbl := buildTable()
	for i, _ := range s {
		buf[1] = s[i]
		ans = (dp1*tbl[string(buf[1:])] + dp2*tbl[string(buf)]) % MOD
		dp2, dp1 = dp1, ans
		buf[0] = s[i]
	}
	return ans
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
