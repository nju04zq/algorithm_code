package main

import "fmt"
import "sort"

// Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), determine if a person could attend all meetings.
//
// For example,
// Given [[0, 30],[5, 10],[15, 20]],
// return false.

func meetingRooms(a [][]int) bool {
	sort.Slice(a, func(i, j int) bool {
		if a[i][0] < a[j][0] {
			return true
		} else if a[i][0] > a[j][0] {
			return false
		} else if a[i][1] < a[j][1] {
			return true
		} else {
			return false
		}
	})
	end := 0
	for i := 0; i < len(a); i++ {
		if a[i][0] < end {
			return false
		}
		end = a[i][1]
	}
	return true
}

func main() {
	a := [][]int{[]int{3, 10}, []int{5, 10}, []int{15, 20}}
	fmt.Printf("a %d, get %t\n", a, meetingRooms(a))
	a = [][]int{[]int{3, 10}, []int{15, 20}}
	fmt.Printf("a %d, get %t\n", a, meetingRooms(a))
	a = [][]int{[]int{3, 10}}
	fmt.Printf("a %d, get %t\n", a, meetingRooms(a))
}
