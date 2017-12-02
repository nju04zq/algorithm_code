package main

import "fmt"

func makeGraph(n int, prerequisites [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range prerequisites {
		v0, v1 := edge[0], edge[1]
		graph[v1] = append(graph[v1], v0)
	}
	return graph
}

func canFinish(n int, prerequisites [][]int) bool {
	graph := makeGraph(n, prerequisites)
	indegrees := make([]int, n)
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			indegrees[neighbor]++
		}
	}
	var i, j int
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if indegrees[j] == 0 {
				break
			}
		}
		if j == n {
			return false
		}
		indegrees[j] = -1
		for _, neighbor := range graph[j] {
			indegrees[neighbor]--
		}
	}
	return true
}

func testCanFinish(n int, prerequistes [][]int) {
	fmt.Printf("%d, %v, get %t\n", n, prerequistes, canFinish(n, prerequistes))
}

func main() {
	testCanFinish(2, [][]int{[]int{1, 0}})
	testCanFinish(2, [][]int{[]int{1, 0}, []int{0, 1}})
}
