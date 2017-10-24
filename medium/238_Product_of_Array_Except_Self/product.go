package main

import "fmt"
import (
	"math"
	"math/rand"
	"time"
)

func product(x int, y int) int {
	z := int64(x)
	z *= int64(y)
	if z > math.MaxInt32 {
		z = math.MaxInt32
	} else if x < math.MinInt32 {
		z = math.MinInt32
	}
	return int(z)
}

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	} else if len(nums) == 1 {
		return []int{0}
	}
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = num
	}
	var x, y int
	for i, num := range res {
		if i == 0 {
			x = num
			continue
		}
		x = product(x, num)
		nums[i] = x
	}
	for i := len(res) - 1; i >= 0; i-- {
		if i == len(res)-1 {
			x = res[i]
			res[i] = nums[i-1]
			continue
		} else if i == 0 {
			res[i] = x
			continue
		}
		y = product(x, nums[i-1])
		x = product(x, res[i])
		res[i] = y
	}
	return res
}

func productBF(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	} else if len(nums) == 1 {
		return []int{0}
	}
	res := make([]int, len(nums))
	for i := 0; i < len(res); i++ {
		x := 1
		for j := 0; j < len(res); j++ {
			if j == i {
				continue
			}
			x *= nums[j]
		}
		if x > math.MaxInt32 {
			x = math.MaxInt32
		} else if x < math.MinInt32 {
			x = math.MinInt32
		}
		res[i] = x
	}
	return res
}

func testProduct(nums []int) {
	nums1 := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		nums1[i] = nums[i]
	}
	ans := productBF(nums)
	res := productExceptSelf(nums)
	err := fmt.Errorf("nums %v, get %v, expect %v", nums1, res, ans)
	if len(res) != len(ans) {
		panic(err)
	}
	for i := 0; i < len(res); i++ {
		if ans[i] != res[i] {
			panic(err)
		}
	}
}

func MakeRandArray() []int {
	maxLen, maxElement := 10, 20
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int() % maxLen
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func main() {
	for i := 0; i < 10000; i++ {
		a := MakeRandArray()
		testProduct(a)
	}
}
