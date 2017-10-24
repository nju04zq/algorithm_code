package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func trap1(height []int) int {
	findLeftMax := func(i int) int {
		res := 0
		for j := 0; j < i; j++ {
			res = max(height[j], res)
		}
		return res
	}
	findRightMax := func(i int) int {
		res := 0
		for j := i + 1; j < len(height); j++ {
			res = max(height[j], res)
		}
		return res
	}
	sum := 0
	for i, h := range height {
		if i == 0 || i == len(height)-1 {
			continue
		}
		left := findLeftMax(i)
		right := findRightMax(i)
		barrier := min(left, right)
		if barrier > h {
			sum += (barrier - h)
		}
	}
	return sum
}

func trap2(height []int) int {
	left := make([]int, len(height))
	for i, h := range height {
		if i == 0 {
			left[0] = height[0]
		} else {
			left[i] = max(left[i-1], h)
		}
	}
	rightMax, sum := 0, 0
	for i := len(height) - 1; i > 0; i-- {
		if i == len(height)-1 {
			rightMax = height[i]
		} else {
			rightMax = max(rightMax, height[i])
		}
		barrier := min(left[i], rightMax)
		if barrier > height[i] {
			sum += (barrier - height[i])
		}
	}
	return sum
}

func trap(height []int) int {
	sum, barrier := 0, 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] <= height[right] {
			barrier = height[left]
			left++
			for height[left] < barrier {
				sum += (barrier - height[left])
				left++
			}
		} else {
			barrier = height[right]
			right--
			for height[right] < barrier {
				sum += (barrier - height[right])
				right--
			}
		}
	}
	return sum
}

func testTrap(height []int, ans int) {
	res := trap(height)
	if res != ans {
		panic(fmt.Sprintf("%v, get %d, should be %d\n", height, res, ans))
	}
}

func main() {
	testTrap([]int{9, 6, 8, 8, 5, 6, 3}, 3)
	testTrap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)
	testTrap([]int{2, 0, 2}, 2)
	testTrap([]int{5, 2, 1, 2, 1, 5}, 14)
	testTrap([]int{9, 6, 8, 8, 5, 6, 3}, 3)
}
