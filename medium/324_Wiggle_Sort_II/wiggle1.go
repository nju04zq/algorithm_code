package main

import "fmt"

func split(nums []int, low, high int) int {
	pilot := nums[low]
	j := low + 1
	for i := low + 1; i < high; i++ {
		if nums[i] < pilot {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	nums[low], nums[j-1] = nums[j-1], nums[low]
	return j - 1 - low
}

func getKthDfs(nums []int, k, low, high int) int {
	j := split(nums, low, high)
	jAbs := j + low
	if k == j {
		return nums[jAbs]
	} else if k < j {
		return getKthDfs(nums, k, low, jAbs)
	} else {
		return getKthDfs(nums, k-j-1, jAbs+1, high)
	}
}

func getKth(nums []int, k int) int {
	return getKthDfs(nums, k, 0, len(nums))
}

func partition(nums []int, val int) (int, int) {
	j, k := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < val {
			temp := nums[j]
			nums[j] = nums[i]
			nums[k] = temp
			j++
			k++
		} else if nums[i] == val {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
	return j, k
}

func wiggleSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		temp[i] = nums[i]
	}
	//fmt.Println("temp", temp)
	medianIdx := (n + 1) / 2
	median := getKth(temp, (n+1)/2)
	//fmt.Println("median", median)
	partition(temp, median)
	//fmt.Println("partition", temp)
	i, j, k := medianIdx-1, n-1, 0
	for i >= 0 || j >= medianIdx {
		if i >= 0 {
			nums[k] = temp[i]
			i--
			k++
		}
		if j >= medianIdx {
			nums[k] = temp[j]
			j--
			k++
		}
	}
}

func verify(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if i%2 == 0 && nums[i] >= nums[i+1] {
			return false
		}
		if i%2 == 1 && nums[i] <= nums[i+1] {
			return false
		}
	}
	return true
}

func testWiggle(nums []int) {
	fmt.Printf("%v\n", nums)
	wiggleSort(nums)
	fmt.Printf("get %v\n", nums)
	if !verify(nums) {
		panic("Not qualified")
	}
}

func main() {
	testWiggle([]int{1, 5, 1, 1, 6, 4})
	testWiggle([]int{1, 3, 2, 2, 3, 1})
	testWiggle([]int{1, 5, 1, 1, 6, 4, 7})
	testWiggle([]int{1, 1, 2, 1, 2, 2, 1})
	testWiggle([]int{5, 3, 1, 2, 6, 7, 8, 5, 5})
	testWiggle([]int{4, 5, 5, 6})
	testWiggle([]int{4})
}
