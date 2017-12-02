package main

import "fmt"

type Queue struct {
	buf []int
}

func initQueue() *Queue {
	q := new(Queue)
	q.buf = make([]int, 0)
	return q
}

func (q *Queue) enqueue(x int) {
	q.buf = append(q.buf, x)
}

func (q *Queue) dequeue() int {
	if len(q.buf) == 0 {
		return -1
	}
	x := q.buf[0]
	q.buf = q.buf[1:]
	return x
}

func (q *Queue) size() int {
	return len(q.buf)
}

func (q *Queue) peek() int {
	if len(q.buf) == 0 {
		return -1
	} else {
		return q.buf[0]
	}
}

type MyStack struct {
	q *Queue
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	stack := new(MyStack)
	stack.q = initQueue()
	return *stack
}

/** Push element x onto stack. */
func (stack *MyStack) Push(x int) {
	stack.q.enqueue(x)
}

/** Removes the element on top of the stack and returns that element. */
func (stack *MyStack) Pop() int {
	temp := make([]int, 0)
	for stack.q.size() > 1 {
		y := stack.q.dequeue()
		temp = append(temp, y)
	}
	x := stack.q.dequeue()
	for _, y := range temp {
		stack.q.enqueue(y)
	}
	return x
}

/** Get the top element. */
func (stack *MyStack) Top() int {
	if stack.q.size() == 0 {
		return -1
	}
	temp := make([]int, 0)
	for stack.q.size() > 1 {
		y := stack.q.dequeue()
		temp = append(temp, y)
	}
	x := stack.q.peek()
	temp = append(temp, stack.q.dequeue())
	for _, y := range temp {
		stack.q.enqueue(y)
	}
	return x
}

/** Returns whether the stack is empty. */
func (stack *MyStack) Empty() bool {
	return stack.q.size() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	stack := Constructor()
	s := &stack
	s.Push(1)
	s.Push(2)
	s.Push(3)
	for !s.Empty() {
		fmt.Printf("%d, top %d\n", s.Pop(), s.Top())
	}
}
