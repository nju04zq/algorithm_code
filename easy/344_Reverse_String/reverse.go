package main

import "fmt"

func reverseString(s string) string {
	n := len(s)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[n-1-i] = s[i]
	}
	return string(buf)
}

func main() {
	fmt.Println("vim-go")
}
