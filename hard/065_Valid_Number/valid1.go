package main

import "fmt"

func stripSpaces(s string) string {
	start, end, i := 0, 0, 0
	for i = 0; i < len(s); i++ {
		if s[i] != ' ' {
			start = i
			break
		}
	}
	start = i
	for i = len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			end = i
			break
		}
	}
	end = i
	if start >= len(s) {
		return ""
	} else {
		return s[start : end+1]
	}
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}

func getInt(s string) (string, bool) {
	var i int
	for i = 0; i < len(s); i++ {
		if !isDigit(s[i]) {
			break
		}
	}
	if i == 0 {
		return "", false
	} else if i >= len(s) {
		return "", true
	} else {
		return s[i:], true
	}
}

func getNum(s string) (string, bool) {
	var rc, intFound bool
	if len(s) == 0 {
		return "", false
	}
	if isDigit(s[0]) {
		s, rc = getInt(s)
		if !rc {
			return "", false
		}
		intFound = true
	}
	if len(s) == 0 {
		return "", true
	}
	if s[0] != '.' {
		if intFound {
			return s, true
		} else {
			return "", false
		}
	}
	s = s[1:]
	if len(s) == 0 || !isDigit(s[0]) {
		if intFound {
			return s, true
		} else {
			return "", false
		}
	}
	s, rc = getInt(s)
	if !rc {
		return "", false
	}
	return s, true
}

func isNumber(s string) bool {
	var rc bool
	s = stripSpaces(s)
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	s, rc = getNum(s)
	if !rc {
		return false
	}
	if len(s) == 0 {
		return true
	} else if s[0] != 'e' {
		return false
	}
	s = s[1:]
	if len(s) == 0 {
		return false
	} else if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	s, rc = getInt(s)
	if !rc {
		return false
	} else if len(s) != 0 {
		return false
	} else {
		return true
	}
}

func testValid(s string, ans bool) {
	res := isNumber(s)
	if res != ans {
		panic(fmt.Errorf("%q, get %t, should %t", s, res, ans))
	}
}

func main() {
	testValid(" 123  ", true)
	testValid(" 1  ", true)
	testValid(" +1  ", true)
	testValid(" -1  ", true)
	testValid(".1", true)
	testValid("1.1", true)
	testValid("-1.1", true)
	testValid("1.", true)
	testValid(".", false)
	testValid("1e1", true)
	testValid("1e-1", true)
	testValid("1e-1.0", false)
	testValid("+1e-1", true)
	testValid("e1", false)
}
