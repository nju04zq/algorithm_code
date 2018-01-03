package main

import "fmt"
import "strings"

func isValidSerialization(preorder string) bool {
	toks := strings.Split(preorder, ",")
	diff := 1
	for _, tok := range toks {
		diff--
		if diff < 0 {
			return false
		}
		if tok != "#" {
			diff += 2
		}
	}
	return diff == 0
}

func testIsValid(s string) {
	fmt.Printf("%q, get %t\n", s, isValidSerialization(s))
}

func main() {
	testIsValid("9,3,4,#,#,1,#,#,2,#,6,#,#")
	testIsValid("1,#")
	testIsValid("1,#,#")
	testIsValid("1,#,#,#")
	testIsValid("9,#,#,1")
	testIsValid("#")
}
