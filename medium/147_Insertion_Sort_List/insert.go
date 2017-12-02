package main

import "fmt"
import "math"

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
func insertionSortList(head *ListNode) *ListNode {
	var prev, cur, next *ListNode
	dummy := &ListNode{Val: math.MinInt32, Next: head}
	prev, cur = dummy, head
	for cur != nil {
		next = cur.Next
		prev.Next = next
		cur.Next = nil
		prev0, cur0 := dummy, dummy.Next
		for cur0 != next && cur0.Val <= cur.Val {
			prev0 = cur0
			cur0 = cur0.Next
		}
		prev0.Next = cur
		cur.Next = cur0
		if cur0 == next {
			prev = cur
		}
		cur = next
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

func testSort(nums []int) {
	head := numArrayToList(nums)
	fmt.Println("Before sort:")
	dumpList(head)
	head = insertionSortList(head)
	fmt.Println("After sort:")
	dumpList(head)
}

func main() {
	testSort([]int{1, 2, 3, 4, 5})
	testSort([]int{5, 4, 3, 2, 1})
}
