// Time Limit Exceeded

package main

import "fmt"
import "math"

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
	if node == nil {
		return
	}
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

func (t *SplayTree) lowerBound(val int) (bool, int) {
	var p1 *SplayTreeNode
	found, res := false, 0
	for p := t.root; p != nil; {
		if p.val < val {
			p = p.rchild
		} else if val == p.val {
			found, res = true, p.val
			p1 = p
			break
		} else {
			found, res = true, p.val
			p1 = p
			p = p.lchild
		}
	}
	t.splay(p1)
	return found, res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findMax(a []int, sum []int, k int) int {
	tree := new(SplayTree)
	m, maxSum := len(a), math.MinInt32
	tree.insert(sum[0])
	for i := 1; i <= m; i++ {
		sum[i] = sum[i-1] + a[i-1]
		exist, val := tree.lowerBound(sum[i] - k)
		if exist {
			maxSum = max(maxSum, sum[i]-val)
			//fmt.Printf("sum %v, i %d, get %d, %d\n", sum, i, sum[i], val)
		}
		tree.insert(sum[i])
	}
	return maxSum
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	maxSum := math.MinInt32
	m, n := len(matrix), len(matrix[0])
	col := make([]int, m+1)
	for i := 0; i < n; i++ {
		jSum := make([]int, m)
		for j := i; j < n; j++ {
			for k := 0; k < m; k++ {
				jSum[k] += matrix[k][j]
			}
			//fmt.Println("jSum ", i, j, jSum)
			tempSum := findMax(jSum, col, k)
			maxSum = max(maxSum, tempSum)
		}
	}
	return maxSum
}

func testMax(matrix [][]int, k int) {
	fmt.Printf("%v, k %d, get %d\n", matrix, k, maxSumSubmatrix(matrix, k))
}

func main() {
	matrix := [][]int{
		[]int{1, 0, 1},
		[]int{0, -2, 3},
	}
	testMax(matrix, 2)
	testMax(matrix, 3)
	matrix = [][]int{
		[]int{5, -4, -3, 4},
		[]int{-3, -4, 4, 5},
		[]int{5, 1, 5, -4},
	}
	testMax(matrix, 8)
}
