package main

import "fmt"
import (
	"math/rand"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func jump(nums []int) int {
	lastArray := make([]int, len(nums))
	for i := 0; i < len(lastArray); i++ {
		lastArray[i] = -1
	}
	for i := 0; i < len(nums); i++ {
		next := nums[i] + i
		if next >= len(nums) {
			next = len(nums) - 1
		}
		if lastArray[next] == -1 {
			lastArray[next] = i
		}
	}
	steps, prev := 0, -1
	for i := len(nums) - 1; i > 0; {
		steps++
		last := lastArray[i]
		for j := i + 1; j < prev; j++ {
			if lastArray[j] == -1 {
				continue
			} else if last == -1 {
				last = lastArray[j]
			} else {
				last = min(last, lastArray[j])
			}
		}
		prev = i
		i = last
	}
	return steps
}

func jumpDp(nums []int) int {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 0
	for i := len(nums) - 2; i >= 0; i-- {
		dp[i] = -1
		for j := i + 1; j < len(nums) && j <= nums[i]+i; j++ {
			if dp[j] == -1 {
				continue
			}
			if dp[i] == -1 {
				dp[i] = dp[j] + 1
			} else {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[0]
}

func testJump(nums []int) {
	ans1 := jump(nums)
	ans2 := jumpDp(nums)
	if ans1 != ans2 {
		panic(fmt.Errorf("nums %v, get %d, should be %d\n", nums, ans1, ans2))
	}
}

func MakeRandArray() []int {
	maxLen, maxElement := 20, 10
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := r.Int()%maxLen + 1
	a := make([]int, len)
	for i := 0; i < len; i++ {
		a[i] = r.Int() % maxElement
	}
	return a
}

func canJump(nums []int) bool {
	next := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= next {
			next = i
		}
	}
	if next != 0 {
		return false
	} else {
		return true
	}
}

func generateArray() []int {
	for {
		a := MakeRandArray()
		if canJump(a) {
			return a
		}
	}

}

func main() {
	for i := 0; i < 10000; i++ {
		a := generateArray()
		testJump(a)
	}
}
