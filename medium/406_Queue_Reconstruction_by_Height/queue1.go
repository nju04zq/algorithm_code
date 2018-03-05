package main

import "fmt"
import "sort"

func insert(res [][]int, person []int, i int) [][]int {
	res = append(res, person)
	if i < len(res)-1 {
		for j := len(res) - 1; j > i; j-- {
			res[j] = res[j-1]
		}
		res[i] = person
	}
	return res
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		pi, pj := people[i], people[j]
		if pi[0] > pj[0] {
			return true
		} else if pi[0] < pj[0] {
			return false
		} else if pi[1] < pj[1] {
			return true
		} else {
			return false
		}
	})
	res := make([][]int, 0, len(people))
	for i := 0; i < len(people); i++ {
		res = insert(res, people[i], people[i][1])
	}
	return res
}

func testQueue(people [][]int) {
	res := reconstructQueue(people)
	fmt.Printf("%v, get %v\n", people, res)
}

func main() {
	people := [][]int{[]int{7, 0}, []int{4, 4}, []int{7, 1}, []int{5, 0},
		[]int{6, 1}, []int{5, 2}}
	//testQueue(people)
	people = [][]int{[]int{8, 2}, []int{4, 2}, []int{4, 5}, []int{2, 0}, []int{7, 2}, []int{1, 4}, []int{9, 1}, []int{3, 1}, []int{9, 0}, []int{1, 0}}
	//testQueue(people)
	people = [][]int{[]int{0, 0}, []int{6, 2}, []int{5, 5}, []int{4, 3}, []int{5, 2}, []int{1, 1}, []int{6, 0}, []int{6, 3}, []int{7, 0}, []int{5, 1}}
	testQueue(people)
}
