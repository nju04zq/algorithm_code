package main

import "fmt"
import "strconv"
import "strings"

func split(version string) []int {
	toks := strings.Split(version, ".")
	nums := make([]int, len(toks))
	for i := 0; i < len(nums); i++ {
		j, _ := strconv.ParseInt(toks[i], 10, 32)
		nums[i] = int(j)
	}
	return nums
}

func compareVersion(version1 string, version2 string) int {
	var num1, num2 int
	nums1 := split(version1)
	nums2 := split(version2)
	for i := 0; i < len(nums1) || i < len(nums2); i++ {
		if i < len(nums1) {
			num1 = nums1[i]
		} else {
			num1 = 0
		}
		if i < len(nums2) {
			num2 = nums2[i]
		} else {
			num2 = 0
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}

func testCompare(version1, version2 string) {
	fmt.Printf("%q, %q, get %d\n", version1, version2, compareVersion(version1, version2))
}

func main() {
	testCompare("0.1", "1.1")
	testCompare("1.1", "1.2")
	testCompare("1.2", "13.37")
	testCompare("01", "1")
	testCompare("1", "1.1")
}
