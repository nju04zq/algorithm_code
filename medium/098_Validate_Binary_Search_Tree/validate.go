package main

import "fmt"
import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderPredecessor(node *TreeNode) *TreeNode {
	cur := node.Left
	if cur == nil {
		return nil
	}
	for cur.Right != nil && cur.Right != node {
		cur = cur.Right
	}
	return cur
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var cur, last *TreeNode
	cur = root
	for cur != nil {
		if cur.Left == nil {
			if last != nil && cur.Val <= last.Val {
				return false
			}
			last, cur = cur, cur.Right
			continue
		}
		node := inorderPredecessor(cur)
		if node.Right == nil {
			node.Right = cur
			cur = cur.Left
		} else {
			if last != nil && cur.Val <= last.Val {
				return false
			}
			node.Right = nil
			last, cur = cur, cur.Right
		}
	}
	return true
}

func makeTree(vals []string) *TreeNode {
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

func testValidate(vals []string) {
	root := makeTree(vals)
	fmt.Printf("%q, %t\n", vals, isValidBST(root))
}

func main() {
	testValidate([]string{"2", "1", "3"})
	testValidate([]string{"1", "2", "3"})
	testValidate([]string{"5", "2", "7", "1", "4", "6", "#", "#", "#", "3"})
}
