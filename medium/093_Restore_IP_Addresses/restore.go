package main

import "fmt"

func toNum(tok string) int {
	num := 0
	for i, _ := range tok {
		num = num*10 + int(tok[i]-'0')
	}
	return num
}

func valid(tok string) bool {
	if tok[0] == '0' && len(tok) > 1 {
		return false
	}
	if toNum(tok) > 255 {
		return false
	} else {
		return true
	}
}

func formatIp(toks []string) string {
	nums := make([]int, 4)
	for i, tok := range toks {
		nums[i] = toNum(tok)
	}
	return fmt.Sprintf("%d.%d.%d.%d", nums[0], nums[1], nums[2], nums[3])
}

func restore(s string, start, idx int, toks []string, res []string) []string {
	if start >= len(s) || idx >= 4 {
		if start == len(s) && idx == 4 {
			res = append(res, formatIp(toks))
		}
		return res
	}
	for i := start; i < len(s) && i < start+4; i++ {
		tok := s[start : i+1]
		if valid(tok) {
			toks[idx] = tok
			res = restore(s, i+1, idx+1, toks, res)
		}
	}
	return res
}

func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	toks := make([]string, 4)
	return restore(s, 0, 0, toks, res)
}

func testRestore(s string) {
	fmt.Printf("%q, get %v\n", s, restoreIpAddresses(s))
}

func main() {
	testRestore("")
	testRestore("25525511135")
	testRestore("010010")
}
