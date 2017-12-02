package main

import "fmt"
import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getListSize(head *ListNode) int {
	size := 0
	for p := head; p != nil; p = p.Next {
		size++
	}
	return size
}

func merge(head1, head2 *ListNode) (*ListNode, *ListNode) {
	dummy := new(ListNode)
	p1, p2, p := head1, head2, dummy
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p.Next = p1
			p = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p = p2
			p2 = p2.Next
		}
	}
	for p1 != nil {
		p.Next = p1
		p = p1
		p1 = p1.Next
	}
	for p2 != nil {
		p.Next = p2
		p = p2
		p2 = p2.Next
	}
	return dummy.Next, p
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: math.MinInt32, Next: head}
	n := getListSize(head)
	for k := 1; k < n; k <<= 1 {
		var prevGroup, nextGroup, start0, end0, start1, end1 *ListNode
		cur, start0 := dummy.Next, dummy.Next
		prevGroup = dummy
		for i := 1; i <= n; i++ {
			if i%(k<<1) == 0 {
				end1 = cur
			} else if i%k == 0 {
				end0 = cur
				start1 = cur.Next
			}
			if i%(k<<1) != 0 && i != n {
				cur = cur.Next
				continue
			}
			if end1 != nil {
				nextGroup = end1.Next
				end1.Next = nil
			} else {
				nextGroup = nil
			}
			if end0 != nil {
				end0.Next = nil
			}
			head, tail := merge(start0, start1)
			prevGroup.Next = head
			tail.Next = nextGroup
			prevGroup = tail
			cur = nextGroup
			start0, start1, end0, end1 = nextGroup, nil, nil, nil
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

func testSort(nums []int) {
	head := numArrayToList(nums)
	fmt.Println("Before sort:")
	dumpList(head)
	fmt.Println("After sort:")
	head = sortList(head)
	dumpList(head)
}

func main() {
	//testSort([]int{1})
	testSort([]int{1, 2})
	testSort([]int{2, 1})
	testSort([]int{1, 2, 3})
	testSort([]int{3, 2, 1})
	testSort([]int{1, 2, 3, 4})
	testSort([]int{4, 3, 2, 1})
	testSort([]int{1, 2, 3, 4, 5})
	testSort([]int{5, 4, 3, 2, 1})
}
