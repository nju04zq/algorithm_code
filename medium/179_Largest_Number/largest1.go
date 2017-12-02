package main

import "fmt"
import "sort"
import "strconv"
import "strings"

func largestNumber(nums []int) string {
	less := func(i, j int) bool {
		s1, s2 := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		s12 := s1 + s2
		s21 := s2 + s1
		if strings.Compare(s12, s21) < 0 {
			return true
		} else {
			return false
		}
	}
	sort.Slice(nums, less)
	buf := make([]string, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if i > 0 && nums[i] == 0 && len(buf) == 0 {
			continue
		}
		buf = append(buf, strconv.Itoa(nums[i]))
	}
	return strings.Join(buf, "")
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
