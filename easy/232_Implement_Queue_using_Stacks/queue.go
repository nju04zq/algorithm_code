package main

import "fmt"

type Stack struct {
	buf []int
}

func initStack() *Stack {
	stack := new(Stack)
	stack.buf = make([]int, 0)
	return stack
}

func (stack *Stack) push(x int) {
	stack.buf = append(stack.buf, x)
}

func (stack *Stack) pop() int {
	if len(stack.buf) == 0 {
		return -1
	}
	var x int
	size := len(stack.buf)
	x, stack.buf = stack.buf[size-1], stack.buf[:size-1]
	return x
}

func (stack *Stack) top() int {
	if len(stack.buf) == 0 {
		return -1
	} else {
		return stack.buf[len(stack.buf)-1]
	}
}

func (stack *Stack) empty() bool {
	return len(stack.buf) == 0
}

type MyQueue struct {
	stack *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	q := new(MyQueue)
	q.stack = initStack()
	return *q
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	stack0 := initStack()
	for !q.stack.empty() {
		stack0.push(q.stack.pop())
	}
	q.stack.push(x)
	for !stack0.empty() {
		q.stack.push(stack0.pop())
	}
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	if q.stack.empty() {
		return -1
	} else {
		return q.stack.pop()
	}
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	return q.stack.top()
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.stack.empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	queue := Constructor()
	q := &queue
	q.Push(1)
	q.Push(2)
	q.Push(3)
	for !q.Empty() {
		fmt.Printf("%d, front %d\n", q.Pop(), q.Peek())
	}
}
