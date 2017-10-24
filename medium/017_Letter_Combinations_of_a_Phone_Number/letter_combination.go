package main

import "fmt"

var dMap = map[byte]string{
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func combInternal(digits string, results *[]string, seq []byte, idx int) {
	if idx >= len(digits) {
		*results = append(*results, string(seq))
		return
	}
	letters := dMap[digits[idx]]
	for i := 0; i < len(letters); i++ {
		seq[idx] = letters[i]
		combInternal(digits, results, seq, idx+1)
	}
}

func letterCombinations(digits string) []string {
	results := make([]string, 0)
	if len(digits) == 0 {
		return results
	}
	seq := make([]byte, len(digits))
	combInternal(digits, &results, seq, 0)
	return results
}

func testLetterCombinations(digits string) {
	results := letterCombinations(digits)
	fmt.Printf("digits %q, results %v\n", digits, results)
}

func main() {
	// output []
	testLetterCombinations("")
	// output ["a", "b", "c"]
	testLetterCombinations("2")
	// output ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
	testLetterCombinations("23")
}
