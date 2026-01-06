package queue

import (
	"math"
)

type Queue struct {
	nums     []int
	smallest int
	largest  int
}

func NewQueue() Queue {
	return Queue{nums: make([]int, 0)}
}

func (q *Queue) Enqueue(num int) {
	if q.Len() == 0 {
		q.largest = num
		q.smallest = num
	}

	if num > q.largest {
		q.largest = num
	} else {
		q.smallest = num
	}

	q.nums = append(q.nums, num)
}

func (q *Queue) Dequeue() int {
	if len(q.nums) == 0 {
		// queue is empty
		return math.MinInt32
	}
	top := q.nums[0]
	q.nums = q.nums[1:len(q.nums)]
	return top
}

func (q *Queue) First() int {
	if len(q.nums) == 0 {
		return math.MinInt32
	}
	return q.nums[0]
}

func (q *Queue) Last() int {
	if len(q.nums) == 0 {
		return math.MinInt32
	}
	return q.nums[q.Len()-1]
}

func (q *Queue) Empty() bool {
	return len(q.nums) == 0
}

func (q *Queue) Len() int {
	return len(q.nums)
}

func (q *Queue) GetSmallest() int {
	return q.smallest
}

func (q *Queue) GetLargest() int {
	return q.largest
}
