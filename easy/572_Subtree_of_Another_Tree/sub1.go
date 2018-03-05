package main

import "fmt"
import "strings"
import "strconv"

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
	if s == nil {
		return false
	}
	if sameTree(s, t) {
		return true
	} else if isSubtree(s.Left, t) || isSubtree(s.Right, t) {
		return true
	} else {
		return false
	}
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
