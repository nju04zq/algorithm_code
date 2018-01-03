package main

import "fmt"

type heap struct {
	buf []string
}

func (h *heap) add(s string) {
	h.buf = append(h.buf, s)
}

func (h *heap) init() {
	for i := len(h.buf) / 2; i >= 0; i-- {
		h.minHeapify(i)
	}
}

func (h *heap) size() int {
	return len(h.buf)
}

func (h *heap) lchild(i int) int {
	return 2*i + 1
}

func (h *heap) rchild(i int) int {
	return 2*i + 2
}

func (h *heap) minHeapify(i int) {
	n := len(h.buf)
	for i < n {
		lchild := h.lchild(i)
		rchild := h.rchild(i)
		least := i
		if lchild < n && h.buf[lchild] < h.buf[least] {
			least = lchild
		}
		if rchild < n && h.buf[rchild] < h.buf[least] {
			least = rchild
		}
		if i == least {
			break
		}
		h.buf[i], h.buf[least] = h.buf[least], h.buf[i]
		i = least
	}
}

func (h *heap) popMin() string {
	if len(h.buf) == 0 {
		return ""
	}
	val := h.buf[0]
	h.buf[0] = h.buf[len(h.buf)-1]
	h.buf = h.buf[:len(h.buf)-1]
	h.minHeapify(0)
	return val
}

func makeGraph(tickets [][]string) map[string]*heap {
	graph := make(map[string]*heap)
	for _, ticket := range tickets {
		from := ticket[0]
		to := ticket[1]
		if _, ok := graph[from]; ok {
			graph[from].add(to)
		} else {
			graph[from] = &heap{[]string{to}}
		}
	}
	for _, h := range graph {
		h.init()
	}
	return graph
}

func findItinerary(tickets [][]string) []string {
	graph := makeGraph(tickets)
	res := make([]string, 0)
	stack := []string{"JFK"}
	for len(stack) > 0 {
		for {
			top := stack[len(stack)-1]
			if _, ok := graph[top]; !ok {
				break
			}
			if graph[top].size() == 0 {
				break
			}
			stack = append(stack, graph[top].popMin())
		}
		res = append(res, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func testFind(tickets [][]string) {
	fmt.Printf("Find in %v, get %s\n", tickets, findItinerary(tickets))
}

func main() {
	tickets := [][]string{
		[]string{"MUC", "LHR"}, []string{"JFK", "MUC"},
		[]string{"SFO", "SJC"}, []string{"LHR", "SFO"},
	}
	testFind(tickets)
	tickets = [][]string{
		[]string{"JFK", "SFO"}, []string{"JFK", "ATL"},
		[]string{"SFO", "ATL"}, []string{"ATL", "JFK"},
		[]string{"ATL", "SFO"},
	}
	testFind(tickets)
	tickets = [][]string{
		[]string{"JFK", "KUL"}, []string{"JFK", "NRT"},
		[]string{"NRT", "JFK"},
	}
	testFind(tickets)
	tickets = [][]string{
		[]string{"EZE", "TIA"}, []string{"EZE", "AXA"},
		[]string{"AUA", "EZE"}, []string{"EZE", "JFK"},
		[]string{"JFK", "ANU"}, []string{"JFK", "ANU"},
		[]string{"AXA", "TIA"}, []string{"JFK", "AUA"},
		[]string{"TIA", "JFK"}, []string{"ANU", "EZE"},
		[]string{"ANU", "EZE"}, []string{"TIA", "AUA"},
	}
	testFind(tickets)
	tickets = [][]string{
		[]string{"JFK", "TIA"},
		[]string{"TIA", "AUA"}, []string{"AUA", "AXA"},
		[]string{"AXA", "TIA"}, []string{"TIA", "ADL"},
	}
	testFind(tickets)
	tickets = [][]string{
		[]string{"JFK", "AUA"}, []string{"JFK", "AUA"},
		[]string{"AUA", "AXA"}, []string{"AUA", "AXB"},
		[]string{"AUA", "AXC"}, []string{"AXC", "AUA"},
		[]string{"AXA", "JFK"}, []string{"AXB", "JFK"},
		[]string{"JFK", "ADL"},
	}
	testFind(tickets)
}
