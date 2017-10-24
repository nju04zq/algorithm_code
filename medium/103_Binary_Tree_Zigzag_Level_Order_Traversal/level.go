package main

import "fmt"
import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverse(vals []int) {
	i, j := 0, len(vals)-1
	for i < j {
		vals[i], vals[j] = vals[j], vals[i]
		i++
		j--
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	nodes := []*TreeNode{root}
	doReverse := false
	for len(nodes) > 0 {
		level := make([]int, 0)
		cnt := len(nodes)
		for i := 0; i < cnt; i++ {
			if nodes[i] != nil {
				level = append(level, nodes[i].Val)
				nodes = append(nodes, nodes[i].Left)
				nodes = append(nodes, nodes[i].Right)
			}
		}
		if len(level) > 0 {
			if doReverse {
				reverse(level)
			}
			res = append(res, level)
		}
		nodes = nodes[cnt:]
		doReverse = !doReverse
	}
	return res
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

func testLevel(vals []string) {
	root := makeTree(vals)
	res := zigzagLevelOrder(root)
	fmt.Printf("%q\n%v\n", vals, res)
}

func main() {
	testLevel([]string{"3", "9", "20", "#", "#", "15", "7"})
}
