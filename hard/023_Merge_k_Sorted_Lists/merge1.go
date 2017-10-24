// merge with heap 36ms

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type heapNode struct {
	lNode *ListNode
	i     int
}

type minHeap struct {
	nodes []*heapNode
}

func (h *minHeap) init(nodes []*heapNode) {
	if len(nodes) == 0 {
		return
	}
	h.nodes = nodes
	last := len(h.nodes) - 1
	parent := h.parent(last)
	for i := parent; i >= 0; i-- {
		h.minHeapify(i)
	}
}

func (h *minHeap) minHeapify(i int) {
	lChild := h.leftChild(i)
	rChild := h.rightChild(i)
	var min int
	if lChild < len(h.nodes) && h.val(lChild) < h.val(i) {
		min = lChild
	} else {
		min = i
	}
	if rChild < len(h.nodes) && h.val(rChild) < h.val(min) {
		min = rChild
	}
	if min != i {
		h.nodes[min], h.nodes[i] = h.nodes[i], h.nodes[min]
		h.minHeapify(min)
	}
}

func (h *minHeap) val(i int) int {
	return h.nodes[i].lNode.Val
}

func (h *minHeap) parent(child int) int {
	if child%2 == 0 {
		return child/2 - 1
	} else {
		return (child - 1) / 2
	}
}

func (h *minHeap) leftChild(parent int) int {
	return 2*parent + 1
}

func (h *minHeap) rightChild(parent int) int {
	return 2*parent + 2
}

func (h *minHeap) pop() *heapNode {
	if len(h.nodes) == 0 {
		return nil
	}
	node := h.nodes[0]
	h.nodes[0] = h.nodes[len(h.nodes)-1]
	h.nodes = h.nodes[:len(h.nodes)-1]
	h.minHeapify(0)
	return node
}

func (h *minHeap) push(node *heapNode) {
	h.nodes = append(h.nodes, node)
	i := len(h.nodes) - 1
	for i > 0 {
		parent := h.parent(i)
		if h.val(parent) < h.val(i) {
			break
		}
		h.nodes[parent], h.nodes[i] = h.nodes[i], h.nodes[parent]
		i = parent
	}
}

func createHeap(lists []*ListNode) *minHeap {
	heap := &minHeap{}
	nodes := make([]*heapNode, 0)
	for i, _ := range lists {
		node := popListNode(lists, i)
		if node != nil {
			nodes = append(nodes, &heapNode{node, i})
		}
	}
	heap.init(nodes)
	return heap
}

func popListNode(lists []*ListNode, i int) *ListNode {
	if lists[i] == nil {
		return nil
	}
	node := lists[i]
	lists[i] = lists[i].Next
	return node
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var head, prev, cur *ListNode
	heap := createHeap(lists)
	for {
		hNode := heap.pop()
		if hNode == nil {
			break
		}
		cur = hNode.lNode
		if prev == nil {
			head = cur
		} else {
			prev.Next = cur
		}
		prev = cur
		if lNode := popListNode(lists, hNode.i); lNode != nil {
			heap.push(&heapNode{lNode, hNode.i})
		}
	}
	return head
}

func numArrayToList(num []int) *ListNode {
	var l, prev *ListNode
	for _, n := range num {
		node := &ListNode{Val: n}
		if prev == nil {
			l = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return l
}

func dumpList(head *ListNode) {
	if head == nil {
		fmt.Println("Nil")
		return
	}
	for p := head; p != nil; p = p.Next {
		if p == head {
			fmt.Printf("%d", p.Val)
		} else {
			fmt.Printf(" -> %d", p.Val)
		}
	}
	fmt.Println("")
}

func testMerge(arrays [][]int) {
	fmt.Println("=====merge=====")
	lists := make([]*ListNode, len(arrays))
	for i, _ := range arrays {
		lists[i] = numArrayToList(arrays[i])
		dumpList(lists[i])
	}
	head := mergeKLists(lists)
	fmt.Println("after merge:")
	dumpList(head)
}

func main() {
	testMerge([][]int{[]int{}, []int{}, []int{}})
	testMerge([][]int{[]int{}, []int{1, 3, 5}, []int{2, 4, 6}})
	testMerge([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})
	testMerge([][]int{[]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9}})
}
