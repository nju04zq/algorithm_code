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

func dfs(graph map[int][]int, i int, visited map[int]bool, path map[int]bool) bool {
	visited[i] = true
	path[i] = true
	for _, neighbor := range graph[i] {
		if _, ok := path[neighbor]; ok {
			return true
		}
		if _, ok := visited[neighbor]; ok {
			continue
		}
		rc := dfs(graph, neighbor, visited, path)
		if rc {
			return true
		}
	}
	delete(path, i)
	return false
}

func canFinish(n int, prerequisites [][]int) bool {
	graph := makeGraph(n, prerequisites)
	visited := make(map[int]bool)
	path := make(map[int]bool)
	for i, _ := range graph {
		rc := dfs(graph, i, visited, path)
		if rc {
			return false
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
