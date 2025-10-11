package stack

import "math"

/*
Implementing queues using stack..

eg: [1, 2, 3, 4]

pop() -> 1
Implementation : {
	use another stack -> s2
	till s.GetStackSize() == 1{
		s2.push(s.pop()) -> [4, 3, 2]
	}
	top = s.pop()
	now s is empty
	till s2.GetStackSize() == 0 {
		s.push(s2.pop()) -> [2, 3, 4]
	}
}

peek() -> 1
Implementation: {
	Just return top
}

push(5) -> [1, 2, 3, 4, 5]
Implementation: {
	push(x)

	Just push this element into the stack.
	if s.GetStackSize() == 0 {
		top = x
		s.Push(x)
	}
}

*/

type MyQueue struct {
	s   stack
	s2  stack
	top int
}

func Constructor() MyQueue {
	return MyQueue{InitStack(), InitStack(), math.MinInt32}
}

func (this *MyQueue) Push(x int) {
	if this.s.GetStackSize() == 0 {
		this.top = x
	}
	this.s.Push(x)
}

func (this *MyQueue) Pop() int {
	// Invalid operation
	if this.s.GetStackSize() == 0 {
		return math.MinInt32
	}
	for this.s.GetStackSize() != 1 {
		this.s2.Push(this.s.Pop())
	}
	// This will be the new top
	this.top = this.s2.Peek()
	// This is what we are returning
	retNum := this.s.Pop()
	for this.s2.GetStackSize() != 0 {
		this.s.Push(this.s2.Pop())
	}

	return retNum
}

func (this *MyQueue) Peek() int {
	return this.top
}

func (this *MyQueue) Empty() bool {
	return this.s.GetStackSize() == 0
}
