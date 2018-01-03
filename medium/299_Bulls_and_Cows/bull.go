package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getHint(secret string, guess string) string {
	cntA := 0
	n := len(secret)
	if n == 0 {
		return ""
	}
	sTbl := make([]int, 10)
	gTbl := make([]int, 10)
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			cntA++
		} else {
			j := int(secret[i] - '0')
			sTbl[j]++
			j = int(guess[i] - '0')
			gTbl[j]++
		}
	}
	cntB := 0
	for i := 0; i < len(sTbl); i++ {
		cntB += min(sTbl[i], gTbl[i])
	}
	return fmt.Sprintf("%dA%dB", cntA, cntB)
}

func testGame(s, g string) {
	fmt.Printf("%q, %q, get %q\n", s, g, getHint(s, g))
}

func main() {
	testGame("1807", "7810")
	testGame("1123", "0111")
}
