// merge with recursive, 19ms

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	var head, prev, cur *ListNode
	lp1, lp2 := list1, list2
	for lp1 != nil || lp2 != nil {
		if lp1 == nil {
			cur = lp2
			lp2 = lp2.Next
		} else if lp2 == nil {
			cur = lp1
			lp1 = lp1.Next
		} else if lp1.Val < lp2.Val {
			cur = lp1
			lp1 = lp1.Next
		} else {
			cur = lp2
			lp2 = lp2.Next
		}
		if prev == nil {
			head = cur
		} else {
			prev.Next = cur
		}
		prev = cur
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	mid := len(lists) / 2
	head1 := mergeKLists(lists[:mid])
	head2 := mergeKLists(lists[mid:])
	return mergeLists(head1, head2)
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

func testMerge(arrays [][]int) {
	fmt.Println("=====merge=====")
	lists := make([]*ListNode, len(arrays))
	for i, _ := range arrays {
		lists[i] = numArrayToList(arrays[i])
		dumpList(lists[i])
	}
	head := mergeKLists(lists)
	fmt.Println("after merge:")
	dumpList(head)
}

func main() {
	testMerge([][]int{[]int{}, []int{}, []int{}})
	testMerge([][]int{[]int{}, []int{1, 3, 5}, []int{2, 4, 6}})
	testMerge([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})
	testMerge([][]int{[]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9}})
}
