package main

import "fmt"
import "bytes"

func zeroArray(a []int) {
	for i, num := range a {
		if num != 0 {
			a[i] = 0
		}
	}
}

func makeKey(tbl []int, buf *bytes.Buffer, s string) string {
	for i, _ := range s {
		c := s[i]
		tbl[c]++
	}
	for i, cnt := range tbl {
		if cnt == 1 {
			buf.WriteByte(byte(i))
		} else if cnt > 1 {
			buf.WriteString(fmt.Sprintf("%c%d", byte(i), cnt))
		}
	}
	return buf.String()
}

func groupAnagrams(strs []string) [][]string {
	tbl := make([]int, 256)
	buf := bytes.NewBuffer(nil)
	res := make(map[string][]string)
	for _, s := range strs {
		zeroArray(tbl)
		buf.Reset()
		key := makeKey(tbl, buf, s)
		if _, ok := res[key]; !ok {
			res[key] = []string{s}
		} else {
			res[key] = append(res[key], s)
		}
	}
	ans := make([][]string, 0)
	for _, a := range res {
		ans = append(ans, a)
	}
	return ans
}

func testGroup(strs []string) {
	ans := groupAnagrams(strs)
	fmt.Printf("strs: %v\n", strs)
	fmt.Printf("get: %v\n", ans)
}

func main() {
	testGroup([]string{"eeat", "teea", "tan", "atee", "nat", "bat"})
}
