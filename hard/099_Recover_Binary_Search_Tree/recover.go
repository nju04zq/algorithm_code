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
func recoverTree(root *TreeNode) {
	var last, cur, n0, n0Next, n1, n1Next *TreeNode
	cur = root
	for cur != nil {
		node := inorderPredecessor(cur)
		if node != nil && node.Right == nil {
			node.Right = cur
			cur = cur.Left
		} else {
			if last != nil && cur.Val < last.Val {
				if n0 == nil {
					n0, n0Next = last, cur
				} else {
					n1, n1Next = last, cur
				}
			}
			if node != nil {
				node.Right = nil
			}
			last = cur
			cur = cur.Right
		}
	}
	if n0 != nil && n1 != nil {
		n0.Val, n1Next.Val = n1Next.Val, n0.Val
	} else {
		n0.Val, n0Next.Val = n0Next.Val, n0.Val
	}
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

func dumpTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Nil")
	}
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		cnt := len(nodes)
		allNil := true
		for i := 0; i < cnt; i++ {
			if nodes[i] == nil {
				fmt.Printf("# ")
				continue
			} else {
				fmt.Printf("%d ", nodes[i].Val)
			}
			nodes = append(nodes, nodes[i].Left)
			nodes = append(nodes, nodes[i].Right)
			if nodes[i].Left != nil || nodes[i].Right != nil {
				allNil = false
			}
		}
		if allNil {
			break
		}
		nodes = nodes[cnt:]
	}
	fmt.Println()
}

func testRecover(vals []string) {
	root := makeTree(vals)
	fmt.Println("Before recover:")
	dumpTree(root)
	recoverTree(root)
	fmt.Println("After recover:")
	dumpTree(root)
}

func main() {
	testRecover([]string{"4", "2", "3", "1", "6", "5", "7"})
	testRecover([]string{"3", "2", "6", "1", "4", "5", "7"})
}
