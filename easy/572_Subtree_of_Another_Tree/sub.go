package main

import "fmt"
import "strings"
import "strconv"

type TreeAttrib struct {
	minDepth int
	maxDepth int
	nodesCnt int
}

func (a *TreeAttrib) theSame(a1 *TreeAttrib) bool {
	if a.minDepth == a1.minDepth && a.maxDepth == a1.maxDepth && a.nodesCnt == a1.nodesCnt {
		return true
	} else {
		return false
	}
}

func getTreeAttrib(t *TreeNode) *TreeAttrib {
	if t == nil {
		return nil
	}
	var a *TreeAttrib
	aLeft := getTreeAttrib(t.Left)
	aRight := getTreeAttrib(t.Right)
	if aLeft == nil && aRight == nil {
		return &TreeAttrib{1, 1, 1}
	} else if aLeft == nil {
		a = aRight
	} else if aRight == nil {
		a = aLeft
	} else {
		a = aLeft
		a.minDepth += aRight.minDepth
		a.maxDepth += aRight.maxDepth
		a.nodesCnt += aRight.nodesCnt
	}
	a.minDepth++
	a.maxDepth++
	a.nodesCnt++
	return a
}

func dfs(s, t *TreeNode, targetAttrib *TreeAttrib, res *bool) *TreeAttrib {
	if s == nil {
		return nil
	}
	var a *TreeAttrib
	aLeft := dfs(s.Left, t, targetAttrib, res)
	if *res == true {
		return nil
	}
	aRight := dfs(s.Right, t, targetAttrib, res)
	if *res == true {
		return nil
	}
	if aLeft == nil && aRight == nil {
		a = &TreeAttrib{0, 0, 0}
	} else if aLeft == nil {
		a = aRight
	} else if aRight == nil {
		a = aLeft
	} else {
		a = aLeft
		a.minDepth += aRight.minDepth
		a.maxDepth += aRight.maxDepth
		a.nodesCnt += aRight.nodesCnt
	}
	a.minDepth++
	a.maxDepth++
	a.nodesCnt++
	if targetAttrib.theSame(a) && sameTree(s, t) {
		*res = true
	}
	return a
}

func sameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil {
		return false
	} else if s.Val == t.Val && sameTree(s.Left, t.Left) && sameTree(s.Right, t.Right) {
		return true
	} else {
		return false
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	targetAttrib := getTreeAttrib(t)
	res := false
	dfs(s, t, targetAttrib, &res)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func makeTree(s string) *TreeNode {
	s = strings.Replace(s, " ", "", -1)
	vals := strings.Split(s, ",")
	makeNode := func(val string) *TreeNode {
		if val == "#" {
			return nil
		}
		node := new(TreeNode)
		if i, err := strconv.ParseInt(val, 10, 32); err != nil {
			panic(fmt.Errorf("Fail to parse %q, %v", val, err))
		} else {
			node.Val = int(i)
		}
		return node
	}
	var root, node *TreeNode
	lastLevel := make([]*TreeNode, 0)
	for i := 0; i < len(vals); {
		if len(lastLevel) == 0 {
			if root = makeNode(vals[i]); root == nil {
				return nil
			}
			lastLevel = append(lastLevel, root)
			i++
			continue
		}
		cnt := len(lastLevel)
		for j := 0; j < cnt && i < len(vals); j++ {
			node = makeNode(vals[i])
			lastLevel[j].Left = node
			if node != nil {
				lastLevel = append(lastLevel, node)
			}
			i++
			if i >= len(vals) {
				break
			}
			node = makeNode(vals[i])
			lastLevel[j].Right = node
			if node != nil {
				lastLevel = append(lastLevel, node)
			}
			i++
		}
		lastLevel = lastLevel[cnt:]
	}
	return root
}

func testSubTree(s0, s1 string) {
	s, t := makeTree(s0), makeTree(s1)
	fmt.Printf("s %q, t %q, get %t\n", s0, s1, isSubtree(s, t))
}

func main() {
	testSubTree("3, 4, 5, 1, 2", "4, 1, 2")
	testSubTree("3, 4, 5, 1, 2, #, #, #, #, 0", "4, 1, 2")
	testSubTree("1", "1")
}
