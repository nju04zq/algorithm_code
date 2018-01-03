package main

import (
	"fmt"
	"math/rand"
	"time"
)

type NumArray struct {
	buf [][]int
}

func copyArray(nums []int) []int {
	nums1 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums1[i] = nums[i]
	}
	return nums1
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	a := new(NumArray)
	a.buf = [][]int{copyArray(nums)}
	prev := nums
	for x := 2; x <= n; x *= 2 {
		m := len(prev)
		b := make([]int, (m+1)/2)
		for i := 0; i < len(b); i++ {
			if 2*i+1 < len(prev) {
				b[i] = prev[2*i] + prev[2*i+1]
			} else {
				b[i] = prev[2*i]
			}
		}
		a.buf = append(a.buf, b)
		prev = b
	}
	return *a
}

func (a *NumArray) sum(i int) int {
	if i <= 0 {
		return 0
	}
	h := len(a.buf)
	low, high, total := 0, i, 0
	for h > 0 && low < high {
		x := 1 << uint(h-1)
		if high >= (low + x) {
			total += a.buf[h-1][low/x]
			low += x
		}
		h--
	}
	return total
}

func (a *NumArray) SumRange(i int, j int) int {
	return a.sum(j+1) - a.sum(i)
}

func (a *NumArray) Update(i int, val int) {
	diff := val - a.buf[0][i]
	x := 1
	for j := 0; j < len(a.buf); j++ {
		k := i / x
		a.buf[j][k] += diff
		x *= 2
	}
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func MakeRandArray() []int {
	maxLen, maxElement := 1000, 10
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int()%maxLen + 1
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func MakeRandInt(a int) int {
	if a == 0 {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int() % a
}

func bf(nums []int, i, j int) int {
	total := 0
	for x := i; x <= j; x++ {
		total += nums[x]
	}
	return total
}

func testSum() {
	nums := MakeRandArray()
	j := MakeRandInt(len(nums))
	i := MakeRandInt(j + 1)
	//nums, i, j = []int{0, 9}, 0, 1
	obj := Constructor(nums)
	res := obj.SumRange(i, j)
	ans := bf(nums, i, j)
	if res != ans {
		panic(fmt.Errorf("Fail on %v, i %d, j %d, get %d, ans %d",
			nums, i, j, res, ans))
	}
	k := MakeRandInt(j - i)
	nums[k] = MakeRandInt(100)
	obj.Update(k, nums[k])
	res = obj.SumRange(i, j)
	ans = bf(nums, i, j)
	if res != ans {
		panic(fmt.Errorf("After update, fail on %v, i %d, j %d, get %d, ans %d",
			nums, i, j, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testSum()
	}
}
