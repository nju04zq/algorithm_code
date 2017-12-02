package main

import "fmt"
import "sort"

type maxHeap []int

func (h *maxHeap) left(i int) int {
	return 2*i + 1
}

func (h *maxHeap) right(i int) int {
	return 2*i + 2
}

func (h *maxHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *maxHeap) length() int {
	return len(*h)
}

func (h *maxHeap) push(num int) {
	*h = append(*h, num)
	i := len(*h) - 1
	for i > 0 {
		parent := h.parent(i)
		if (*h)[i] < (*h)[parent] {
			break
		}
		(*h)[i], (*h)[parent] = (*h)[parent], (*h)[i]
		i = parent
	}
}

func (h *maxHeap) pop() int {
	n := len(*h)
	if n == 0 {
		return 0
	}
	ret := (*h)[0]
	(*h)[0] = (*h)[n-1]
	(*h) = (*h)[:n-1]
	i := 0
	n--
	for i < n {
		largest := i
		left := h.left(i)
		right := h.right(i)
		if left < n && (*h)[left] > (*h)[largest] {
			largest = left
		}
		if right < n && (*h)[right] > (*h)[largest] {
			largest = right
		}
		if largest == i {
			break
		}
		(*h)[i], (*h)[largest] = (*h)[largest], (*h)[i]
		i = largest
	}
	return ret
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] < courses[j][1] {
			return true
		} else {
			return false
		}
	})
	h := new(maxHeap)
	total := 0
	for i := 0; i < len(courses); i++ {
		t, end := courses[i][0], courses[i][1]
		total += courses[i][0]
		h.push(t)
		for total > end {
			total -= h.pop()
		}
	}
	return h.length()
}

func testSchedule(courses [][]int) {
	fmt.Printf("courses %v, get %d\n", courses, scheduleCourse(courses))
}

func main() {
	courses := [][]int{
		[]int{100, 200},
		[]int{200, 1300},
		[]int{1000, 1250},
		[]int{2000, 3200},
	}
	testSchedule(courses)
	courses = [][]int{
		[]int{1, 2},
		[]int{2, 3},
	}
	testSchedule(courses)
	courses = [][]int{
		[]int{5, 5},
		[]int{4, 6},
		[]int{2, 6},
	}
	testSchedule(courses)
	courses = [][]int{
		[]int{7, 17},
		[]int{3, 12},
		[]int{10, 20},
		[]int{9, 10},
		[]int{5, 20},
		[]int{10, 19},
		[]int{4, 18},
	}
	testSchedule(courses)
}
