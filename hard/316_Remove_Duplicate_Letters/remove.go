package main

import "fmt"
import "math/rand"
import "sort"
import "strings"
import "time"

func removeDuplicateLetters(s string) string {
	if len(s) == 0 {
		return ""
	}
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		ch := int(s[i] - 'a')
		cnt[ch]++
	}
	pos := 0
	for i := 0; i < len(s); i++ {
		if s[i] < s[pos] {
			pos = i
		}
		ch := int(s[i] - 'a')
		cnt[ch]--
		if cnt[ch] == 0 {
			break
		}
	}
	s0 := fmt.Sprintf("%c", s[pos])
	s = strings.Replace(s[pos+1:], s0, "", -1)
	return s0 + removeDuplicateLetters(s)
}

func smaller(s string, res string, path []int) bool {
	if len(res) == 0 {
		return true
	}
	for i := 0; i < len(path); i++ {
		if s[path[i]] < res[i] {
			return true
		} else if s[path[i]] > res[i] {
			return false
		}
	}
	return false
}

func dfs(s string, t byte, path []int, res []string) {
	if (t - 'a') == 26 {
		temp := make([]int, len(path))
		for i := 0; i < len(path); i++ {
			temp[i] = path[i]
		}
		sort.Ints(temp)
		if smaller(s, res[0], temp) {
			buf := make([]byte, len(path))
			for i := 0; i < len(path); i++ {
				buf[i] = s[temp[i]]
			}
			res[0] = string(buf)
		}
		return
	}
	seen := false
	for i := 0; i < len(s); i++ {
		if s[i] != t {
			continue
		}
		seen = true
		path = append(path, i)
		dfs(s, t+1, path, res)
		path = path[:len(path)-1]
	}
	if !seen {
		dfs(s, t+1, path, res)
	}
}

func bf(s string) string {
	t := byte('a')
	path := make([]int, 0)
	res := make([]string, 1)
	dfs(s, t, path, res)
	return res[0]
}

func MakeRandString() string {
	maxLen, maxElement := 10, 10
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]byte, len)
	for i := 0; i < len; i++ {
		a[i] = 'a' + byte(r.Int()%maxElement)
	}
	return string(a)
}

func testRemove() {
	s := MakeRandString()
	res := removeDuplicateLetters(s)
	ans := bf(s)
	if res != ans {
		panic(fmt.Errorf("%q, get %q, expect %q\n", s, res, ans))
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		testRemove()
	}
}
