package main

import (
	"fmt"
	"math/rand"
	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findLength(a []int, b []int) int {
	m, n := len(a), len(b)
	if m == 0 || n == 0 {
		return 0
	}
	maxLen := 0
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if a[0] == b[i] {
			dp[i] = 1
			maxLen = 1
		}
	}
	//fmt.Println("  ", b)
	//fmt.Println("", a[0], dp)
	for i := 1; i < m; i++ {
		prev := dp[0]
		if a[i] == b[0] {
			dp[0] = 1
			maxLen = max(maxLen, dp[0])
		} else {
			dp[0] = 0
		}
		for j := 1; j < n; j++ {
			temp := dp[j]
			if a[i] == b[j] {
				dp[j] = prev + 1
				maxLen = max(maxLen, dp[j])
			} else {
				dp[j] = 0
			}
			prev = temp
		}
		//fmt.Println("", a[i], dp)
	}
	return maxLen
}

func bf(a []int, b []int) int {
	maxLen := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			cnt := 0
			for k := 0; i+k < len(a) && j+k < len(b); k++ {
				if a[i+k] != b[j+k] {
					break
				}
				cnt++
			}
			maxLen = max(maxLen, cnt)
		}
	}
	return maxLen
}

func MakeRandInt() int {
	maxNum := 40
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int() % maxNum
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func testFind() {
	a := MakeRandArray()
	b := MakeRandArray()
	res := findLength(a, b)
	ans := bf(a, b)
	if res != ans {
		fmt.Printf("a %v, b %v, get %d, expect %d\n", a, b, res, ans)
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testFind()
	}
}
