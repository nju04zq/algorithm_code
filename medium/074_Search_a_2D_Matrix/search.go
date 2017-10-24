package main

import "fmt"

func searchInCol(matrix [][]int, target int) int {
	low, high := 0, len(matrix)-1
	for low < high {
		mid := high - (high-low)/2
		if matrix[mid][0] == target {
			return mid
		} else if matrix[mid][0] > target {
			high = mid - 1
		} else {
			low = mid
		}
	}
	if matrix[low][0] <= target {
		return low
	} else {
		return -1
	}
}

func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := searchInCol(matrix, target)
	if row == -1 {
		return false
	} else if matrix[row][0] == target {
		return true
	}
	return search(matrix[row], target)
}

func dump(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func testSearch(matrix [][]int, target int, ans bool) {
	res := searchMatrix(matrix, target)
	if res != ans {
		dump(matrix)
		fmt.Printf("Search %d, get %t, should %t\n", target, res, ans)
	}
}

func main() {
	m := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}
	testSearch(m, 0, false)
	testSearch(m, 4, false)
	testSearch(m, 5, true)
	testSearch(m, 9, false)
	testSearch(m, 10, true)
	testSearch(m, 20, true)
	testSearch(m, 21, false)
	testSearch(m, 23, true)
	testSearch(m, 50, true)
	testSearch(m, 51, false)
	m = [][]int{
		[]int{1},
	}
	testSearch(m, 0, false)
	testSearch(m, 1, true)
	testSearch(m, 2, false)
}
