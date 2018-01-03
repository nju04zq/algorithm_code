package main

import "fmt"
import "sort"
import "math/rand"
import "time"

func hIndex(c []int) int {
	n := len(c)
	bucket := make([]int, n+1)
	for _, num := range c {
		if num >= n {
			bucket[n]++
		} else {
			bucket[num]++
		}
	}
	cnt := 0
	for i := n; i >= 0; i-- {
		cnt += bucket[i]
		if cnt >= i {
			return i
		}
	}
	return 0
}

func bf(c []int) int {
	n := len(c)
	sort.Ints(c)
	for i := 0; i < n; i++ {
		if (n - i) <= c[i] {
			return n - i
		}
	}
	return 0
}

func MakeRandArray() []int {
	maxLen, maxElement := 20, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func testHindex() {
	c := MakeRandArray()
	res := hIndex(c)
	ans := bf(c)
	if res != ans {
		panic(fmt.Errorf("%v, get %d, expect %d\n", c, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testHindex()
	}
}
