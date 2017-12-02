package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type MinStack struct {
	valStack []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	s := &MinStack{
		valStack: make([]int, 0),
		minStack: make([]int, 0),
	}
	return *s
}

func (s *MinStack) Push(x int) {
	top := len(s.valStack) - 1
	s.valStack = append(s.valStack, x)
	if top == -1 {
		s.minStack = append(s.minStack, x)
	} else {
		s.minStack = append(s.minStack, min(x, s.minStack[top]))
	}
}

func (s *MinStack) Pop() {
	if len(s.valStack) > 0 {
		top := len(s.valStack) - 1
		s.valStack = s.valStack[:top]
		s.minStack = s.minStack[:top]
	}
}

func (s *MinStack) Top() int {
	if len(s.valStack) > 0 {
		top := len(s.valStack) - 1
		return s.valStack[top]
	} else {
		return 0
	}
}

func (s *MinStack) GetMin() int {
	if len(s.valStack) > 0 {
		top := len(s.valStack) - 1
		return s.minStack[top]
	} else {
		return 0
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	stack := Constructor()
	s := &stack
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.GetMin())
	s.Pop()
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
}
