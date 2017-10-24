package main

import "fmt"

func lengthOfLastWord(s string) int {
	prev, cur := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if cur != 0 {
				prev = cur
			}
			cur = 0
		} else {
			cur++
		}
	}
	if cur == 0 {
		return prev
	} else {
		return cur
	}
}

func testLast(s string) {
	fmt.Printf("s %q, %d\n", s, lengthOfLastWord(s))
}

func main() {
	testLast("")
	testLast("abc")
	testLast("ab abc")
	testLast("ab abc  ")
}
