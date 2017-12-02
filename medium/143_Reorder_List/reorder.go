package main

import "fmt"

func reverse(head *ListNode) *ListNode {
	var prev, cur, next *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

func merge(head1, head2 *ListNode) *ListNode {
	p1, p2 := head1, head2
	for p1 != nil && p2 != nil {
		next1 := p1.Next
		next2 := p2.Next
		p1.Next = p2
		p2.Next = next1
		p1, p2 = next1, next2
	}
	return head1
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if slow == fast {
		return
	}
	head1 := reverse(slow.Next)
	slow.Next = nil
	merge(head, head1)
}

type ListNode struct {
	Val  int
	Next *ListNode
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

func testReorder(nums []int) {
	head := numArrayToList(nums)
	fmt.Println("Before reorder:")
	dumpList(head)
	reorderList(head)
	fmt.Println("After reorder:")
	dumpList(head)
}

func main() {
	testReorder([]int{1})
	testReorder([]int{1, 2})
	testReorder([]int{1, 2, 3})
	testReorder([]int{1, 2, 3, 4})
	testReorder([]int{1, 2, 3, 4, 5})
}
