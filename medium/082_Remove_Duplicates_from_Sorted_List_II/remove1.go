package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prev, cur := dummy, head
	duplicate := false
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
			duplicate = true
		} else if duplicate {
			prev.Next = cur.Next
			duplicate = false
			cur = prev.Next
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return dummy.Next
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

func dumpList(head *ListNode) {
	if head == nil {
		fmt.Println("Nil")
		return
	}
	for p := head; p != nil; p = p.Next {
		if p == head {
			fmt.Printf("%d", p.Val)
		} else {
			fmt.Printf(" -> %d", p.Val)
		}
	}
	fmt.Println()
}

func testDelete(nums []int) {
	l := numArrayToList(nums)
	fmt.Printf("Before: ")
	dumpList(l)
	l = deleteDuplicates(l)
	fmt.Printf("After: ")
	dumpList(l)
}

func main() {
	testDelete([]int{})
	testDelete([]int{1})
	testDelete([]int{1, 1})
	testDelete([]int{1, 1, 2, 3, 3})
	testDelete([]int{1, 1, 2, 2})
}
