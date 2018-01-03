package main

import "fmt"
import "bytes"

var tbl1 = []string{"", "One", "Two", "Three", "Four", "Five",
	"Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve",
	"Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}

var tbl2 = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty",
	"Sixty", "Seventy", "Eighty", "Ninety"}

func translateSlice(num int, buf *bytes.Buffer) string {
	if num == 0 {
		return ""
	}
	parts := make([]string, 0)
	if num/100 > 0 {
		parts = append(parts, fmt.Sprintf("%s Hundred", tbl1[num/100]))
		num %= 100
	}
	if num >= 20 {
		parts = append(parts, fmt.Sprintf("%s", tbl2[num/10]))
		num %= 10
	}
	if num > 0 {
		parts = append(parts, tbl1[num])
	}
	buf.Reset()
	for i, part := range parts {
		buf.WriteString(part)
		if i < len(parts)-1 {
			buf.WriteByte(' ')
		}
	}
	return buf.String()
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	units := []string{"Thousand", "Million", "Billion"}
	base := 1000
	i := -1
	parts := make([]string, 0)
	buf := bytes.NewBuffer(nil)
	for num > 0 {
		if num%base == 0 {
			num /= base
			i++
			continue
		}
		s := translateSlice(num%base, buf)
		unit := ""
		if i >= 0 {
			unit = " " + units[i]
		}
		parts = append(parts, fmt.Sprintf("%s%s", s, unit))
		num /= base
		i++
	}
	buf.Reset()
	for i := len(parts) - 1; i >= 0; i-- {
		buf.WriteString(parts[i])
		if i > 0 {
			buf.WriteByte(' ')
		}
	}
	return buf.String()
}

func testWords(num int) {
	fmt.Printf("%d: %q\n", num, numberToWords(num))
}

func main() {
	for i := 0; i < 30; i++ {
		testWords(i)
	}
	testWords(55)
	testWords(155)
	testWords(2155)
	testWords(32155)
	testWords(432155)
	testWords(5432155)
}
