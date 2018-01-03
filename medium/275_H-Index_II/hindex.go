package main

import "fmt"
import "sort"
import "math/rand"
import "time"

func hIndex(c []int) int {
	n := len(c)
	low, high := 0, n-1
	for low < high {
		mid := low + (high-low)/2
		if (n - mid) <= c[mid] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if low > high {
		return 0
	} else if (n - low) <= c[low] {
		return n - low
	} else {
		return 0
	}
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
	sort.Ints(c)
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
