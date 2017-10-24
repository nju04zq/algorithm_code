package main

import "fmt"
import "math/rand"
import "time"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Stack struct {
	nums []int
}

func (s *Stack) init() {
	s.nums = make([]int, 0)
}

func (s *Stack) push(i int) {
	s.nums = append(s.nums, i)
}

func (s *Stack) pop() (int, error) {
	if len(s.nums) == 0 {
		return 0, fmt.Errorf("Fail to pop, stack empty.")
	}
	num := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return num, nil
}

func longestValidParentheses(s string) int {
	var maxPairs, start, i int
	stack := &Stack{}
	stack.init()
	for i = 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.push(i)
			continue
		}
		if _, err := stack.pop(); err != nil {
			maxPairs = max(maxPairs, (i-start)/2)
			start = i + 1
		}
	}
	for {
		j, err := stack.pop()
		if err != nil {
			break
		}
		maxPairs = max(maxPairs, (i-j)/2)
		i = j
	}
	maxPairs = max(maxPairs, (i-start)/2)
	return maxPairs * 2
}

func longestBF(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		var stack int
		for j := i; j < len(s); j++ {
			if s[j] == '(' {
				stack++
			} else {
				stack--
			}
			if stack < 0 {
				break
			} else if stack == 0 {
				longest = max(longest, j-i+1)
			}
		}
	}
	return longest
}

func testLongest(s string) {
	res := longestValidParentheses(s)
	ans := longestBF(s)
	if res != ans {
		panic(fmt.Errorf("%q: get %d, expect %d\n", s, res, ans))
	}
}

func makeRandStr() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % 20
	buf := make([]byte, len)
	for i, _ := range buf {
		if r.Int()%2 == 0 {
			buf[i] = '('
		} else {
			buf[i] = ')'
		}
	}
	return string(buf)
}

func main() {
	testLongest("(()")
	testLongest(")()())")
	var i int
	for i = 0; i < 100000; i++ {
		s := makeRandStr()
		testLongest(s)
	}
	fmt.Printf("%d test cases passed!\n", i)
}
