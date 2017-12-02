package main

import "fmt"
import "sort"

func splitToDigits(num int) []byte {
	digits := make([]byte, 0)
	if num == 0 {
		digits = append(digits, '0')
	}
	for num > 0 {
		digits = append(digits, byte(num%10)+'0')
		num /= 10
	}
	i, j := 0, len(digits)-1
	for i < j {
		digits[i], digits[j] = digits[j], digits[i]
		i++
		j--
	}
	return digits
}

func largestNumber(nums []int) string {
	tbl := make(map[int][]byte)
	for _, num := range nums {
		if _, ok := tbl[num]; ok {
			continue
		} else {
			tbl[num] = splitToDigits(num)
		}
	}
	less := func(i, j int) bool {
		d1, d2 := tbl[nums[i]], tbl[nums[j]]
		for i = 0; i < len(d1) && i < len(d2); i++ {
			if d1[i] < d2[i] {
				return true
			} else if d1[i] > d2[i] {
				return false
			}
		}
		if len(d1) == len(d2) {
			return true
		}
		d, compD1 := d1, true
		if i < len(d2) {
			d, compD1 = d2, false
		}
		for j := 0; j < len(d); j++ {
			k := (j + i) % len(d)
			if d[j] < d[k] {
				if compD1 {
					return false
				} else {
					return true
				}
			} else if d[j] > d[k] {
				if compD1 {
					return true
				} else {
					return false
				}
			}
		}
		return true
	}
	sort.Slice(nums, less)
	buf := make([]byte, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if i > 0 && nums[i] == 0 && len(buf) == 0 {
			continue
		}
		digits := tbl[nums[i]]
		for _, digit := range digits {
			buf = append(buf, digit)
		}
	}
	return string(buf)
}

func testLargest(nums []int) {
	fmt.Printf("%v, get %q\n", nums, largestNumber(nums))
}

func main() {
	testLargest([]int{3, 30, 34, 5, 9})
	testLargest([]int{8, 878})
	testLargest([]int{8, 898})
	testLargest([]int{1, 1, 1})
	testLargest([]int{1, 128})
	testLargest([]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512})
	testLargest([]int{0, 0})
	testLargest([]int{0, 1})
}
