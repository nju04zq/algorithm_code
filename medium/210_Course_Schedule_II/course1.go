package main

import "fmt"

func makeGraph(n int, edges [][]int) map[int][]int {
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range edges {
		v0, v1 := edge[0], edge[1]
		graph[v1] = append(graph[v1], v0)
	}
	return graph
}

func findOrder(n int, prerequisites [][]int) []int {
	if n == 0 {
		return []int{}
	}
	graph := makeGraph(n, prerequisites)
	indegrees := make([]int, n)
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			indegrees[neighbor]++
		}
	}
	toVisit := make([]int, 0)
	for i, indegree := range indegrees {
		if indegree == 0 {
			toVisit = append(toVisit, i)
		}
	}
	visited := 0
	res := make([]int, 0)
	for len(toVisit) > 0 {
		cnt := len(toVisit)
		for _, i := range toVisit {
			for _, neighbor := range graph[i] {
				indegrees[neighbor]--
				if indegrees[neighbor] == 0 {
					toVisit = append(toVisit, neighbor)
				}
			}
			res = append(res, i)
			visited++
		}
		toVisit = toVisit[cnt:]
	}
	if visited == n {
		return res
	} else {
		return []int{}
	}
}

func testFindOrder(n int, prerequisites [][]int) {
	fmt.Printf("%d, %v, get %v\n", n, prerequisites, findOrder(n, prerequisites))
}

func main() {
	testFindOrder(2, [][]int{[]int{1, 0}})
	testFindOrder(4, [][]int{[]int{1, 0}, []int{2, 0}, []int{3, 1}, []int{3, 2}})
	testFindOrder(2, [][]int{[]int{1, 0}, []int{0, 1}})
}
