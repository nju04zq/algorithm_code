package main

/*
 * ======== TLE ========
 */

import "fmt"
import "math/rand"
import "time"

var lastTbl = map[int]int{}

func countArray(a []int) map[int]int {
	tbl := make(map[int]int)
	for i := 0; i < len(a); i++ {
		tbl[a[i]] += 1
	}
	return tbl
}

func twoSumCount(c, d []int, target int) int {
	cnt := 0
	for i := 0; i < len(c); i++ {
		need := target - c[i]
		if get, ok := lastTbl[need]; ok {
			cnt += get
		}
	}
	return cnt
}

func threeSumCount(b, c, d []int, target int) int {
	cnt, prevCnt := 0, 0
	for i := 0; i < len(b); i++ {
		if i > 0 && b[i] == b[i-1] {
			cnt += prevCnt
		} else {
			prevCnt = twoSumCount(c, d, target-b[i])
			cnt += prevCnt
		}
	}
	return cnt
}

func fourSumCount(a, b, c, d []int) int {
	lastTbl = countArray(d)
	cnt, prevCnt := 0, 0
	for i := 0; i < len(a); i++ {
		target := -a[i]
		if i > 0 && a[i] == a[i-1] {
			cnt += prevCnt
		} else {
			prevCnt = threeSumCount(b, c, d, target)
			cnt += prevCnt
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
