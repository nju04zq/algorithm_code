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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var tail *ListNode
	var n int
	for p := head; p != nil; p = p.Next {
		tail = p
		n++
	}
	k = k % n
	if k == 0 {
		return head
	}
	var prev, cur *ListNode
	var i int
	for cur = head; cur != nil; cur = cur.Next {
		if i == n-k {
			break
		}
		prev = cur
		i++
	}
	tail.Next = head
	prev.Next = nil
	head = cur
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

func testRotate(nums []int, k int) {
	head := numArrayToList(nums)
	fmt.Printf("Before rotate %d:\n", k)
	dumpList(head)
	head = rotateRight(head, k)
	fmt.Println("After rotate:")
	dumpList(head)
}

func main() {
	testRotate([]int{1}, 0)
	testRotate([]int{1}, 1)
	testRotate([]int{1}, 2)
	testRotate([]int{1, 2, 3, 4, 5}, 0)
	testRotate([]int{1, 2, 3, 4, 5}, 1)
	testRotate([]int{1, 2, 3, 4, 5}, 2)
	testRotate([]int{1, 2, 3, 4, 5}, 3)
	testRotate([]int{1, 2, 3, 4, 5}, 4)
	testRotate([]int{1, 2, 3, 4, 5}, 5)
	testRotate([]int{1, 2, 3, 4, 5}, 6)
}
