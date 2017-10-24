package main

import "fmt"
import "strconv"

func reverse(from, to *TreeNode) {
	var prev, cur, next *TreeNode
	for cur = from; cur != nil; {
		next = cur.Right
		cur.Right = prev
		prev = cur
		cur = next
	}
}

func reverseDump(from, to *TreeNode, res []int) []int {
	reverse(from, to)
	for cur := to; cur != nil; cur = cur.Right {
		res = append(res, cur.Val)
	}
	reverse(to, from)
	return res
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
func postorderTraversal(root *TreeNode) []int {
	dummy := &TreeNode{Left: root}
	res := make([]int, 0)
	for cur := dummy; cur != nil; {
		if cur.Left == nil {
			cur = cur.Right
			continue
		}
		node := inorderPredecessor(cur)
		if node.Right == nil {
			node.Right = cur
			cur = cur.Left
		} else {
			node.Right = nil
			res = reverseDump(cur.Left, node, res)
			cur = cur.Right
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func testPostorder(vals []string) {
	fmt.Printf("Tree: %v\n", vals)
	root := makeTree(vals)
	res := postorderTraversal(root)
	fmt.Println(res)
}

func main() {
	testPostorder([]string{"1", "2", "3", "#", "#", "4", "5", "#", "6", "7"})
}
