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
func reverseList(head *ListNode) *ListNode {
	var prev, cur, next *ListNode
	for cur = head; cur != nil; {
		next = cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
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

func testReverse(nums []int) {
	head := numArrayToList(nums)
	fmt.Printf("Before: ")
	dumpList(head)
	head = reverseList(head)
	fmt.Printf("After: ")
	dumpList(head)
}

func main() {
	testReverse([]int{})
	testReverse([]int{1})
	testReverse([]int{1, 2, 3, 4, 5})
}
