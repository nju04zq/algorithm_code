package main

import "fmt"
import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func validate(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 == nil || root2 == nil {
		return false
	} else if root1.Val != root2.Val {
		return false
	}
	if !validate(root1.Left, root2.Right) {
		return false
	}
	if !validate(root1.Right, root2.Left) {
		return false
	}
	return true
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return validate(root.Left, root.Right)
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

func testIsSymmetric(vals []string) {
	root := makeTree(vals)
	fmt.Printf("%q, %t\n", vals, isSymmetric(root))
}

func main() {
	testIsSymmetric([]string{"1", "2", "2", "3", "4", "4", "3"})
	testIsSymmetric([]string{"1", "2", "2", "#", "3", "#", "3"})
}
