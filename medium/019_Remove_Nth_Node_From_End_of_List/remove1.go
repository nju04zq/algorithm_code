package main

/*
 * 6ms
 * one pass 3ms
 */

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	cnt := 0
	for p := head; p != nil; p = p.Next {
		cnt++
	}
	n = cnt - n
	var prev *ListNode
	cur, i := head, 0
	for p := head; p != nil; p = p.Next {
		if i == n {
			break
		}
		i++
		prev = cur
		cur = cur.Next
	}
	if prev == nil {
		return cur.Next
	} else {
		prev.Next = cur.Next
		return head
	}
}

func numArrayToList(num []int) *ListNode {
	var l, prev *ListNode
	for _, n := range num {
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
	fmt.Println("")
}

func testRemoveNthFromEnd(nums []int, n int) {
	head := numArrayToList(nums)
	fmt.Println("=========================")
	fmt.Printf("Before remove %d:\n", n)
	dumpList(head)
	head = removeNthFromEnd(head, n)
	fmt.Printf("After remove %d:\n", n)
	dumpList(head)
	fmt.Println("=========================")
}

func main() {
	testRemoveNthFromEnd([]int{1}, 1)
	testRemoveNthFromEnd([]int{1, 2}, 1)
	testRemoveNthFromEnd([]int{1, 2}, 2)
	testRemoveNthFromEnd([]int{1, 2, 3, 4, 5}, 1)
	testRemoveNthFromEnd([]int{1, 2, 3, 4, 5}, 2)
	testRemoveNthFromEnd([]int{1, 2, 3, 4, 5}, 3)
	testRemoveNthFromEnd([]int{1, 2, 3, 4, 5}, 4)
	testRemoveNthFromEnd([]int{1, 2, 3, 4, 5}, 5)
}
