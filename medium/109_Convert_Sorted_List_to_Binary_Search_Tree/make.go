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

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := len(nums) / 2
	root := &TreeNode{Val: nums[i]}
	root.Left = sortedArrayToBST(nums[:i])
	if i+1 < len(nums) {
		root.Right = sortedArrayToBST(nums[i+1:])
	}
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	nums := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		nums = append(nums, cur.Val)
	}
	return sortedArrayToBST(nums)
}

func main() {
	fmt.Println("vim-go")
}
