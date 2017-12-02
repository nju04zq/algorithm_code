package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func missingNumber(nums []int) int {
	xor := 0
	for i, num := range nums {
		xor = xor ^ i ^ num
	}
	return xor ^ len(nums)
}

func bf(nums []int) int {
	nums1 := make([]int, len(nums))
	for i, num := range nums {
		nums1[i] = num
	}
	sort.Ints(nums1)
	for i, num := range nums1 {
		if i != num {
			return i
		}
	}
	return len(nums)
}

func MakeRandArray() []int {
	maxLen := 100
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	size := r.Int()%maxLen + 1
	a := r.Perm(size)
	a = a[:len(a)-1]
	return a
}

func testMissing() {
	nums := MakeRandArray()
	res := missingNumber(nums)
	ans := bf(nums)
	if res != ans {
		panic(fmt.Errorf("%v, get %d, expect %d\n", nums, res, ans))
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		testMissing()
	}
}
