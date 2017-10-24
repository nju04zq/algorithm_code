package main

import "fmt"
import (
	"math/rand"
	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSub, maxProduct, minProduct := nums[0], nums[0], nums[0]
	for i, num := range nums {
		if i == 0 {
			continue
		}
		if num == 0 {
			maxProduct, minProduct = 0, 0
		} else if num > 0 {
			maxProduct = max(maxProduct*num, num)
			minProduct = min(minProduct*num, num)
		} else {
			tmp1 := max(minProduct*num, num)
			tmp2 := min(maxProduct*num, num)
			maxProduct, minProduct = tmp1, tmp2
		}
		maxSub = max(maxSub, maxProduct)
	}
	return maxSub
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int()%maxLen + 1
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int()%maxElement - 8
	}
	return a
}

func maxProductBF(nums []int) int {
	maxSub := nums[0]
	for i := 0; i < len(nums); i++ {
		product := 1
		for j := i; j < len(nums); j++ {
			product *= nums[j]
			maxSub = max(maxSub, product)
		}
	}
	return maxSub
}

func testMaxProduct(nums []int) {
	res := maxProduct(nums)
	ans := maxProductBF(nums)
	if res != ans {
		panic(fmt.Errorf("nums %v, get %d, expect %d", nums, res, ans))
	}
}

func main() {
	for i := 0; i < 100000; i++ {
		testMaxProduct(MakeRandArray())
	}
}
