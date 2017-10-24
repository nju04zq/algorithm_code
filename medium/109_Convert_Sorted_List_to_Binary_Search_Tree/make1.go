package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func makeBST(head **ListNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	left := makeBST(head, start, mid-1)
	root := &TreeNode{Val: (*head).Val}
	*head = (*head).Next
	right := makeBST(head, mid+1, end)
	root.Left = left
	root.Right = right
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	size := 0
	for cur := head; cur != nil; cur = cur.Next {
		size++
	}
	return makeBST(&head, 0, size-1)
}

func numArrayToList(nums []int) *ListNode {
	var l, prev *ListNode
	for _, n := range nums {
		node := &ListNode{Val: n}
		if prev == nil {
			l = node
		} else {
			prev.Next = node
		}
		prev = node
	}
	return l
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

func testMake(nums []int) {
	head := numArrayToList(nums)
	root := sortedListToBST(head)
	fmt.Printf("%v, get: ", nums)
	dumpTree(root)
}

func main() {
	testMake([]int{1, 2, 3, 4, 5})
}
