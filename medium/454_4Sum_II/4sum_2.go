package main

import "fmt"
import "math/rand"
import "time"

func makeSumMap(a, b []int) map[int]int {
	m := make(map[int]int)
	for _, numA := range a {
		for _, numB := range b {
			m[numA+numB] += 1
		}
	}
	return m
}

func fourSumCount(a, b, c, d []int) int {
	cnt := 0
	sumMap := makeSumMap(a, b)
	for _, numC := range c {
		for _, numD := range d {
			cnt += sumMap[-(numC + numD)]
		}
	}
	return cnt
}

func MakeRandArray() []int {
	maxLen, maxElement := 20, 100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int()%maxElement - maxElement/2
	}
	return a
}

func fourSumCountBf(a, b, c, d []int) int {
	cnt := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			for k := 0; k < len(c); k++ {
				for l := 0; l < len(d); l++ {
					if a[i]+b[j]+c[k]+d[l] == 0 {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}

func testFourSumCount(a, b, c, d []int) bool {
	ans := fourSumCountBf(a, b, c, d)
	res := fourSumCount(a, b, c, d)
	if ans != res {
		fmt.Printf("a %v, b %v, c %v, d %v, get %d, expect %d\n", a, b, c, d, res, ans)
		return false
	}
	return true
}

func main() {
	for i := 0; i < 1000; i++ {
		a := MakeRandArray()
		b := MakeRandArray()
		c := MakeRandArray()
		d := MakeRandArray()
		res := testFourSumCount(a, b, c, d)
		if !res {
			break
		}
	}
}
