package main

// Splay tree search O(nlgn), the constant is about 2.6 - 2.7
// with 10,000 ints
// Search avg. depth 19.387873
// Insert avg. depth 23.571000
// Remove avg. depth 17.414300

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type SplayTreeNode struct {
	parent *SplayTreeNode
	lchild *SplayTreeNode
	rchild *SplayTreeNode
	val    int
}

type SplayTree struct {
	root *SplayTreeNode
}

func (t *SplayTree) zig(node *SplayTreeNode) {
	//fmt.Printf("zig %d: ", node.val)
	parent := node.parent
	parent.lchild = node.rchild
	if node.rchild != nil {
		node.rchild.parent = parent
	}
	node.rchild = parent
	node.parent = parent.parent
	if parent.parent != nil {
		if parent == parent.parent.lchild {
			parent.parent.lchild = node
		} else {
			parent.parent.rchild = node
		}
	}
	parent.parent = node
	if node.parent == nil {
		t.root = node
	}
	//t.dump()
}

func (t *SplayTree) zag(node *SplayTreeNode) {
	parent := node.parent
	//fmt.Printf("zag %d, parent %d: ", node.val, node.parent.val)
	parent.rchild = node.lchild
	if node.lchild != nil {
		node.lchild.parent = parent
	}
	node.lchild = parent
	node.parent = parent.parent
	if parent.parent != nil {
		if parent == parent.parent.lchild {
			parent.parent.lchild = node
		} else {
			parent.parent.rchild = node
		}
	}
	parent.parent = node
	if node.parent == nil {
		t.root = node
	}
	//t.dump()
}

func (t *SplayTree) splay(node *SplayTreeNode) {
	//fmt.Printf("splay %v, ", node.val)
	//t.dump()
	for {
		//fmt.Printf("root %v, node %v, node.parent %v\n", t.root, node, node.parent)
		if node == t.root {
			break
		} else if node.parent == t.root {
			if node == t.root.lchild {
				t.zig(node)
			} else {
				t.zag(node)
			}
		} else {
			var pl, nl bool
			if node.parent == node.parent.parent.lchild {
				pl = true
			}
			if node == node.parent.lchild {
				nl = true
			}
			if pl && nl {
				t.zig(node.parent)
				t.zig(node)
			} else if pl && !nl {
				t.zag(node)
				t.zig(node)
			} else if !pl && nl {
				t.zig(node)
				t.zag(node)
			} else {
				t.zag(node.parent)
				t.zag(node)
			}
		}
		//fmt.Printf("$$$ ")
		//t.dump()
	}
}

func (t *SplayTree) insert(val int) {
	var prev *SplayTreeNode
	for p := t.root; p != nil; {
		prev = p
		if val <= p.val {
			p = p.lchild
		} else {
			p = p.rchild
		}
	}
	node := &SplayTreeNode{parent: prev, val: val}
	if prev == nil {
		t.root = node
	} else if val <= prev.val {
		prev.lchild = node
	} else {
		prev.rchild = node
	}
	//fmt.Printf("After insert %d: ", val)
	//t.dump()
	t.splay(node)
}

func (t *SplayTree) transplant(u, v *SplayTreeNode) {
	if v != nil {
		v.parent = u.parent
	}
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.lchild {
		u.parent.lchild = v
	} else {
		u.parent.rchild = v
	}
}

func (t *SplayTree) remove(val int) {
	//fmt.Printf("Before remove %d: ", val)
	//t.dump()
	node := t.searchNode(val)
	if node == nil {
		//fmt.Printf("Could not remove %d, not exist\n", val)
		return
	}
	t.splay(node)
	//fmt.Printf("After splay %d: ", val)
	//t.dump()
	if node.lchild == nil {
		t.transplant(node, node.rchild)
	} else if node.rchild == nil {
		t.transplant(node, node.lchild)
	} else {
		p := t.treeMin(node.rchild)
		if p.parent != node {
			t.transplant(p, p.rchild)
			p.rchild = node.rchild
			p.rchild.parent = p
		}
		p.lchild = node.lchild
		p.lchild.parent = p
		t.transplant(node, p)
	}
	//fmt.Printf("After remove %d: ", val)
	//t.dump()
}

func (t *SplayTree) treeMin(p *SplayTreeNode) *SplayTreeNode {
	for p.lchild != nil {
		p = p.lchild
	}
	return p
}

func (t *SplayTree) searchNode(val int) *SplayTreeNode {
	for p := t.root; p != nil; {
		if p.val < val {
			p = p.rchild
		} else if p.val > val {
			p = p.lchild
		} else {
			return p
		}
	}
	return nil
}

func (t *SplayTree) search(val int) bool {
	//t.dump()
	var p *SplayTreeNode
	for p = t.root; p != nil; {
		if p.val < val {
			p = p.rchild
		} else if p.val > val {
			p = p.lchild
		} else {
			break
		}
	}
	if p != nil {
		t.splay(p)
		return true
	} else {
		return false
	}
}

func (t *SplayTree) lowerBound(val int) (bool, int) {
	found, res := false, 0
	for p := t.root; p != nil; {
		if p.val < val {
			p = p.rchild
		} else if val == p.val {
			found, res = true, p.val
			break
		} else {
			found, res = true, p.val
			p = p.lchild
		}
	}
	return found, res
}

func (t *SplayTree) upperBound(val int) (bool, int) {
	found, res := false, 0
	for p := t.root; p != nil; {
		if p.val <= val {
			p = p.rchild
		} else {
			found, res = true, p.val
			p = p.lchild
		}
	}
	return found, res
}

func (t *SplayTree) splayPerformance(node *SplayTreeNode) int {
	//fmt.Printf("splay %v, ", node.val)
	//t.dump()
	depth := 0
	for {
		depth++
		//fmt.Printf("root %v, node %v, node.parent %v\n", t.root, node, node.parent)
		if node == t.root {
			break
		} else if node.parent == t.root {
			if node == t.root.lchild {
				t.zig(node)
			} else {
				t.zag(node)
			}
		} else {
			var pl, nl bool
			if node.parent == node.parent.parent.lchild {
				pl = true
			}
			if node == node.parent.lchild {
				nl = true
			}
			if pl && nl {
				t.zig(node.parent)
				t.zig(node)
			} else if pl && !nl {
				t.zag(node)
				t.zig(node)
			} else if !pl && nl {
				t.zig(node)
				t.zag(node)
			} else {
				t.zag(node.parent)
				t.zag(node)
			}
		}
		//fmt.Printf("$$$ ")
		//t.dump()
	}
	return depth
}

func (t *SplayTree) searchPerformance(val int) int {
	var p *SplayTreeNode
	depth := 0
	for p = t.root; p != nil; {
		depth++
		if p.val < val {
			p = p.rchild
		} else if p.val > val {
			p = p.lchild
		} else {
			break
		}
	}
	if p != nil {
		depth += t.splayPerformance(p)
	}
	return depth
}

func (t *SplayTree) insertPerformance(val int) int {
	var prev *SplayTreeNode
	depth := 0
	for p := t.root; p != nil; {
		depth++
		prev = p
		if val <= p.val {
			p = p.lchild
		} else {
			p = p.rchild
		}
	}
	node := &SplayTreeNode{parent: prev, val: val}
	if prev == nil {
		t.root = node
	} else if val <= prev.val {
		prev.lchild = node
	} else {
		prev.rchild = node
	}
	//fmt.Printf("After insert %d: ", val)
	//t.dump()
	depth += t.splayPerformance(node)
	return depth
}

func (t *SplayTree) treeMinPerformance(p *SplayTreeNode) (*SplayTreeNode, int) {
	depth := 0
	for p.lchild != nil {
		depth++
		p = p.lchild
	}
	return p, depth
}

func (t *SplayTree) removePerformance(val int) int {
	//fmt.Printf("Before remove %d: ", val)
	//t.dump()
	depth := 0
	node := t.searchNode(val)
	if node == nil {
		//fmt.Printf("Could not remove %d, not exist\n", val)
		return 0
	}
	depth += t.splayPerformance(node)
	//fmt.Printf("After splay %d: ", val)
	//t.dump()
	if node.lchild == nil {
		t.transplant(node, node.rchild)
	} else if node.rchild == nil {
		t.transplant(node, node.lchild)
	} else {
		p, d := t.treeMinPerformance(node.rchild)
		depth += d
		if p.parent != node {
			t.transplant(p, p.rchild)
			p.rchild = node.rchild
			p.rchild.parent = p
		}
		p.lchild = node.lchild
		p.lchild.parent = p
		t.transplant(node, p)
	}
	//fmt.Printf("After remove %d: ", val)
	//t.dump()
	return depth
}

func (t *SplayTree) dump() {
	if t.root == nil {
		fmt.Println("Nil")
	}
	nodes := make([]*SplayTreeNode, 0)
	nodes = append(nodes, t.root)
	for len(nodes) > 0 {
		cnt := len(nodes)
		allNil := true
		for i := 0; i < cnt; i++ {
			if nodes[i] == nil {
				fmt.Printf("# ")
				continue
			} else {
				fmt.Printf("%d ", nodes[i].val)
			}
			nodes = append(nodes, nodes[i].lchild)
			nodes = append(nodes, nodes[i].rchild)
			if nodes[i].lchild != nil || nodes[i].rchild != nil {
				allNil = false
			}
		}
		if allNil {
			break
		}
		nodes = nodes[cnt:]
	}
	fmt.Println()
	t.sanity()
}

func (t *SplayTree) sanity() {
	if t.root != nil {
		t.sanityTree(t.root)
	}
}

func (t *SplayTree) sanityTree(root *SplayTreeNode) {
	l, r := root.lchild, root.rchild
	if l != nil && l.parent != root {
		panic(fmt.Sprintf("%d parent is %v, should be %v", l.val, l.parent, root))
	}
	if r != nil && r.parent != root {
		panic(fmt.Sprintf("%d parent is %v, should be %v", r.val, r.parent, root))
	}
	if l != nil {
		t.sanityTree(l)
	}
	if r != nil {
		t.sanityTree(r)
	}
}

func MakeRandInt(minNum, maxNum int) int {
	x := maxNum - minNum
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int()%x + minNum
}

func MakeRandArray(size int) []int {
	if size < 0 {
		size = MakeRandInt(1, 500)
	}
	//size := MakeRandInt(1, 10)
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = MakeRandInt(-500, 500)
		//a[i] = MakeRandInt(0, 20)
	}
	return a
}

func searchArray(a []int, target int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == target {
			return true
		}
	}
	return false
}

func arrayLowerBound(a []int, x int) int {
	res := -1
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			return i
		} else if a[i] > x {
			if res == -1 || a[i] < a[res] {
				res = i
			}
		}
	}
	return res
}

func arrayUpperBound(a []int, x int) int {
	res := -1
	for i := 0; i < len(a); i++ {
		if a[i] > x {
			if res == -1 || a[i] < a[res] {
				res = i
			}
		}
	}
	return res
}

func testSplayTree() {
	a := MakeRandArray(-1)
	//a = []int{13, 2, 8, 1, 7}
	//a = []int{6, 7, 0, 0}
	//fmt.Printf("Test array %v\n", a)
	tree := new(SplayTree)
	searchTest := func(array []int, cnt int, x int) {
		for i := 0; i < cnt; i++ {
			if cnt > 0 {
				x = MakeRandInt(-1000, 1000)
			}
			//fmt.Printf("$$Search %d\n", x)
			res := tree.search(x)
			ans := searchArray(array, x)
			if res != ans {
				panic(fmt.Sprintf("Search %d, get %t, expect %t", x, res, ans))
			}
		}
	}
	for i := 0; i < len(a); i++ {
		//fmt.Printf("Insert %d\n", a[i])
		tree.insert(a[i])
		//tree.dump()
		searchTest(a[i+1:], 0, a[i])
	}
	for i := 0; i < len(a); i++ {
		res := tree.search(a[i])
		if !res {
			panic(fmt.Sprintf("search failed, %d should exist", a[i]))
		}
	}
	searchTest(a, 1000, 0)
	testLowerBound := func(array []int, x int) {
		exist, res := tree.lowerBound(x)
		ans := arrayLowerBound(array, x)
		if ans == -1 && exist {
			panic(fmt.Sprintf("LowerBound for %d should not exit", x))
		} else if ans != -1 {
			if !exist {
				panic(fmt.Sprintf("LowerBound for %d should exit", x))
			}
			if array[ans] != res {
				panic(fmt.Sprintf("LowerBound for %d, get %d, expect %d", x, res, array[ans]))
			}
		}
	}
	testUpperBound := func(array []int, x int) {
		exist, res := tree.upperBound(x)
		ans := arrayUpperBound(array, x)
		if ans == -1 && exist {
			panic(fmt.Sprintf("UpperBound for %d should not exit", x))
		} else if ans != -1 {
			if !exist {
				panic(fmt.Sprintf("UpperBound for %d should exit", x))
			}
			if array[ans] != res {
				panic(fmt.Sprintf("UpperBound for %d, get %d, expect %d", x, res, array[ans]))
			}
		}
	}
	for i := 0; i < 1000; i++ {
		x := MakeRandInt(-1000, 1000)
		testLowerBound(a, x)
		testUpperBound(a, x)
	}
	for i := 0; i < len(a); i++ {
		//fmt.Printf("Remove %d\n", a[i])
		tree.remove(a[i])
		searchTest(a[i+1:], 0, a[i])
		searchTest(a[i+1:], 10, 0)
	}
}

func correctnessTest() {
	for i := 0; i < 1000; i++ {
		fmt.Printf("\r%d", i)
		testSplayTree()
	}
	fmt.Println()
}

func performanceTest() {
	a := MakeRandArray(10 * 10000)
	tree := new(SplayTree)
	for i := 0; i < len(a); i++ {
		tree.insert(a[i])
	}
	sort.Ints(a)
	b := MakeRandArray(10000)
	testBinarySearch := func(x int) {
		low, high := 0, len(a)-1
		for low <= high {
			mid := low + (high-low)/2
			if a[mid] == x {
				return
			} else if a[mid] > x {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	testCnt := 1000
	t1 := time.Now()
	for i := 0; i < testCnt; i++ {
		for j := 0; j < len(b); j++ {
			testBinarySearch(b[j])
		}
	}
	t2 := time.Now()
	d1 := t2.Sub(t1)
	fmt.Printf("Binary search took %v\n", d1)
	t1 = time.Now()
	for i := 0; i < testCnt; i++ {
		for j := 0; j < len(b); j++ {
			tree.search(b[j])
		}
	}
	t2 = time.Now()
	d2 := t2.Sub(t1)
	fmt.Printf("Splay search took %v\n", d2)
	fmt.Printf("Slower %f\n", float64(d2)/float64(d1))
}

func performanceTest1() {
	a := MakeRandArray(10 * 10000)
	tree := new(SplayTree)
	for i := 0; i < len(a); i++ {
		tree.insert(a[i])
	}
	d := 0
	testCnt := 1000
	b := MakeRandArray(10000)
	for i := 0; i < testCnt; i++ {
		for j := 0; j < len(b); j++ {
			d += tree.searchPerformance(b[j])
		}
	}
	fmt.Printf("Search avg. depth %f\n", float64(d)/float64(len(b)*testCnt))
	d = 0
	c := MakeRandArray(10000)
	for i := 0; i < testCnt; i++ {
		tree := new(SplayTree)
		for j := 0; j < len(b); j++ {
			tree.insert(b[j])
		}
		for j := 0; j < len(c); j++ {
			d += tree.insertPerformance(c[j])
			tree.remove(c[j])
		}
	}
	fmt.Printf("Insert avg. depth %f\n", float64(d)/float64(len(c)*testCnt))
	d = 0
	for i := 0; i < testCnt; i++ {
		tree := new(SplayTree)
		for j := 0; j < len(b); j++ {
			tree.insert(b[j])
		}
		for j := 0; j < len(b); j++ {
			d += tree.removePerformance(b[j])
			tree.insert(b[j])
		}
	}
	fmt.Printf("Remove avg. depth %f\n", float64(d)/float64(len(b)*testCnt))
}

func main() {
	correctnessTest()
	performanceTest()
	performanceTest1()
}
