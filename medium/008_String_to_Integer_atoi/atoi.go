package main

import "fmt"

import "math"
import "unicode"

func skipSpaces(s string, start int) int {
	var i int
	for i = start; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	return i
}

func readSign(s string, start int) (bool, int) {
	if start >= len(s) {
		return false, start
	}
	if s[start] == '+' {
		return false, start + 1
	} else if s[start] == '-' {
		return true, start + 1
	} else {
		return false, start
	}
}

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	i := skipSpaces(str, 0)
	neg, i := readSign(str, i)
	var num int64
	for ; i < len(str); i++ {
		if !unicode.IsDigit(rune(str[i])) {
			break
		}
		x := (int64(str[i]) - int64('0'))
		if neg {
			x = -x
		}
		num = (num*10 + x)
		if num > math.MaxInt32 {
			num = math.MaxInt32
			break
		} else if num < math.MinInt32 {
			num = math.MinInt32
			break
		}
	}
	return int(num)
}

func testMyAtoi(testcases map[string]int) {
	for s, ans := range testcases {
		res := myAtoi(s)
		if ans != res {
			panic(fmt.Sprintf("%q, get %d, expect %d\n", s, res, ans))
		}
	}
}

func main() {
	testcases := map[string]int{
		"":         0,
		"0":        0,
		"1":        1,
		"-1":       -1,
		"123":      123,
		"-123":     -123,
		" ":        0,
		" 0":       0,
		"  1":      1,
		"  -1":     -1,
		"  123":    123,
		"  -123":   -123,
		"#0123":    0,
		"1#123":    1,
		"-1#123":   -1,
		"123#123":  123,
		"-123#123": -123,
	}
	testMyAtoi(testcases)
}
