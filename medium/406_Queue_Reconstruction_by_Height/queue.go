package main

import "fmt"
import "strings"

type Pair struct {
	h, i int
}

type MinHeap struct {
	buf []*Pair
}

func (h *MinHeap) init(buf []*Pair) *MinHeap {
	h.buf = buf
	for i := len(buf) / 2; i >= 0; i-- {
		h.minHeapify(i)
	}
	return h
}

func (h *MinHeap) size() int {
	return len(h.buf)
}

func (h *MinHeap) lchild(i int) int {
	return 2*i + 1
}

func (h *MinHeap) rchild(i int) int {
	return 2*i + 2
}

func (h *MinHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) smaller(i, j int) bool {
	pi, pj := h.buf[i], h.buf[j]
	if pi.h < pj.h {
		return true
	} else if pi.h > pj.h {
		return false
	} else if pi.i < pj.i {
		return true
	} else {
		return false
	}
}

func (h *MinHeap) minHeapify(i int) {
	n := len(h.buf)
	for i < n {
		least := i
		lchild := h.lchild(i)
		rchild := h.rchild(i)
		if lchild < n && h.smaller(lchild, least) {
			least = lchild
		}
		if rchild < n && h.smaller(rchild, least) {
			least = rchild
		}
		if least == i {
			break
		}
		h.buf[i], h.buf[least] = h.buf[least], h.buf[i]
		i = least
	}
}

func (h *MinHeap) pop() *Pair {
	if len(h.buf) == 0 {
		return nil
	}
	res := h.buf[0]
	n := len(h.buf)
	h.buf[0] = h.buf[n-1]
	h.buf = h.buf[:n-1]
	h.minHeapify(0)
	return res
}

func (h *MinHeap) dump() {
	for i := 0; i < len(h.buf); i++ {
		fmt.Printf("%d,%d ", h.buf[i].h, h.buf[i].i)
	}
	fmt.Println()
}

type BSTreeNode struct {
	val    int
	size   int
	left   *BSTreeNode
	right  *BSTreeNode
	parent *BSTreeNode
}

type BST struct {
	root *BSTreeNode
}

func (bst *BST) insert(i int) {
	var p, parent *BSTreeNode
	p, parent = bst.root, nil
	for p != nil {
		p.size++
		parent = p
		if p.val > i {
			p = p.left
		} else {
			p = p.right
		}
	}
	node := &BSTreeNode{val: i, parent: parent}
	if parent == nil {
		bst.root = node
	} else if parent.val > i {
		parent.left = node
	} else {
		parent.right = node
	}
}

func (bst *BST) transplant(u, v *BSTreeNode) {
	if u.parent == nil {
		bst.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (bst *BST) treeMin(root *BSTreeNode) *BSTreeNode {
	p := root
	for p.left != nil {
		p = p.left
	}
	return p
}

func (bst *BST) decreaseFrom(node *BSTreeNode) {
	for p := node; p != nil; p = p.parent {
		p.size--
	}
}

func (bst *BST) remove(node *BSTreeNode) {
	if node.left == nil {
		bst.transplant(node, node.right)
		bst.decreaseFrom(node.parent)
	} else if node.right == nil {
		bst.transplant(node, node.left)
		bst.decreaseFrom(node.parent)
	} else {
		y := bst.treeMin(node.right)
		bst.decreaseFrom(y.parent)
		if y.parent != node {
			bst.transplant(y, y.right)
			y.right = node.right
			y.right.parent = y
		}
		bst.transplant(node, y)
		y.left = node.left
		y.left.parent = y
		y.size = node.size
	}
}

func (bst *BST) getIthFrom(p *BSTreeNode, i int) *BSTreeNode {
	if i >= p.size {
		fmt.Println(p.size, i)
		return nil
	} else if p.left != nil && i < p.left.size {
		return bst.getIthFrom(p.left, i)
	} else if p.left == nil {
		if i == 0 {
			return p
		} else {
			return bst.getIthFrom(p.right, i-1)
		}
	} else {
		if i == p.left.size {
			return p
		} else {
			return bst.getIthFrom(p.right, i-p.left.size-1)
		}
	}
}

func (bst *BST) removeIth(i int) int {
	//fmt.Println("remove", i)
	//bst.dump()
	node := bst.getIthFrom(bst.root, i)
	bst.remove(node)
	return node.val
}

func (bst *BST) dump() {
	if bst.root == nil {
		fmt.Println("<empty tree>")
		return
	}
	queue := []*BSTreeNode{bst.root}
	res := make([]string, 0, bst.root.size)
	for len(queue) > 0 {
		cnt := len(queue)
		for i := 0; i < cnt; i++ {
			p := queue[i]
			if p == nil {
				res = append(res, "#")
				continue
			}
			var parent int
			if p.parent == nil {
				parent = -1
			} else {
				parent = p.parent.val
			}
			res = append(res, fmt.Sprintf("(%d %d, %d)", p.val, p.size, parent))
			queue = append(queue, p.left)
			queue = append(queue, p.right)
		}
		queue = queue[cnt:]
	}
	i := len(res) - 1
	for ; i > 0; i-- {
		if res[i] != "#" {
			break
		}
	}
	fmt.Println(strings.Join(res[:i+1], ", "))
}

func makePairs(people [][]int) []*Pair {
	pairs := make([]*Pair, len(people))
	for i := 0; i < len(people); i++ {
		pair := &Pair{people[i][0], people[i][1]}
		pairs[i] = pair
	}
	return pairs
}

func makeBSTFrom(parent *BSTreeNode, start, end int) *BSTreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	root := &BSTreeNode{val: mid, parent: parent}
	root.left = makeBSTFrom(root, start, mid-1)
	root.right = makeBSTFrom(root, mid+1, end)
	root.size = 1
	if root.left != nil {
		root.size += root.left.size
	}
	if root.right != nil {
		root.size += root.right.size
	}
	return root
}

func makeBST(start, end int) *BST {
	bst := new(BST)
	bst.root = makeBSTFrom(nil, start, end)
	return bst
}

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, len(people))
	pairs := makePairs(people)
	heap := new(MinHeap).init(pairs)
	bst := makeBST(0, len(people)-1)
	//bst.dump()
	prevHeight := -1
	subCnt := 0
	for heap.size() > 0 {
		pair := heap.pop()
		//fmt.Printf("pop %d, %d\n", pair.h, pair.i)
		idx := pair.i
		if pair.h == prevHeight {
			subCnt++
		} else {
			subCnt = 0
		}
		idx -= subCnt
		j := bst.removeIth(idx)
		res[j] = []int{pair.h, pair.i}
		prevHeight = pair.h
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
