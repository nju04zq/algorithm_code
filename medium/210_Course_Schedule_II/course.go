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

func dfs(graph map[int][]int, i int, visited, path map[int]bool, res []int) []int {
	visited[i] = true
	path[i] = true
	for _, neighbor := range graph[i] {
		if _, ok := path[neighbor]; ok {
			return []int{}
		}
		if _, ok := visited[neighbor]; ok {
			continue
		}
		res = dfs(graph, neighbor, visited, path, res)
		if len(res) == 0 {
			return res
		}
	}
	delete(path, i)
	res = append(res, i)
	return res
}

func reverse(res []int) {
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
}

func findOrder(n int, prerequisites [][]int) []int {
	if n == 0 {
		return []int{}
	}
	graph := makeGraph(n, prerequisites)
	res := make([]int, 0)
	visited := make(map[int]bool)
	path := make(map[int]bool)
	for i := 0; i < n; i++ {
		if _, ok := visited[i]; ok {
			continue
		}
		res = dfs(graph, i, visited, path, res)
		if len(res) == 0 {
			return res
		}
	}
	reverse(res)
	return res
}

func testFindOrder(n int, prerequisites [][]int) {
	fmt.Printf("%d, %v, get %v\n", n, prerequisites, findOrder(n, prerequisites))
}

func main() {
	testFindOrder(2, [][]int{[]int{1, 0}})
	testFindOrder(4, [][]int{[]int{1, 0}, []int{2, 0}, []int{3, 1}, []int{3, 2}})
	testFindOrder(2, [][]int{[]int{1, 0}, []int{0, 1}})
}
