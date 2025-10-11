package stack

import (
	"dsa/queue"
	"math"
)

// We need this -> Last in First Out

// We have this -> First in First Out

/*
push(1)
push(2)
push(3)
push(4)

[1, 2, 3, 4]

pop()
Implementation : {
	popQ() -> [2, 3, 4] (1)

	till we get last element do this (len(queue) - 1) -> {
		push(popQ()) -> [2, 3, 4, 1]
	}
	at the end we should have this: [4, 1, 2, 3]

	return popQ() -> (4)
}

[1, 2, 3]
push(5)
Implementation: {
	Keeping this simple: just perform push operation in queue
	pushQ(5) -> ([1, 2, 3, 5])
	update the top variable
}

top()
Implementation: {
	One way we can do this efficiently, is keeping track of the recently pushed number
	return last element
}

[1, 2, 3]
*/

// Push() ->

type MyStack struct {
	q   queue.Queue
	top int
}

func NewQStack() MyStack {
	return MyStack{queue.NewQueue(), math.MinInt32}
}

func (this *MyStack) Push(x int) {
	this.q.Push(x)
	this.top = x
}

func (this *MyStack) Pop() int {
	// Invalid operation
	if this.q.Len() == 0 {
		return math.MinInt32
	}

	if this.q.Len() == 1 {
		this.top = math.MinInt32
		return this.q.Pop()
	}
	for i := 0; i < this.q.Len()-1; i++ {
		ele := this.q.Pop()
		this.q.Push(ele)
		if i == this.q.Len()-2 {
			// The last element to be push will be assigned as top of the stack
			this.top = ele
		}
	}

	return this.q.Pop()
}

func (this *MyStack) Top() int {
	return this.top
}

func (this *MyStack) Empty() bool {
	return this.q.Len() == 0
}
