package main

import "fmt"

func isVowel(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		b = b - 'A' + 'a'
	}
	if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
		return true
	} else {
		return false
	}
}

func reverseVowels(s string) string {
	n := len(s)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = s[i]
	}
	i, j := 0, n-1
	for i < j {
		if !isVowel(buf[i]) {
			i++
			continue
		}
		if !isVowel(buf[j]) {
			j--
			continue
		}
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func testReverse(s string) {
	fmt.Printf("%q, get %q\n", s, reverseVowels(s))
}

func main() {
	testReverse("aA")
}
