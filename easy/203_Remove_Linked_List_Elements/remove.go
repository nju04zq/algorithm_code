package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
	}
	return dummy.Next
}

func main() {
	fmt.Println("vim-go")
}
