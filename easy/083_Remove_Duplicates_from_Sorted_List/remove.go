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
	var prev, cur *ListNode
	for cur = head; cur != nil; cur = cur.Next {
		if prev == nil {
			head = cur
		} else if cur.Val == prev.Val {
			continue
		} else {
			prev.Next = cur
		}
		prev = cur
	}
	if prev != nil {
		prev.Next = nil
	}
	return head
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
}
