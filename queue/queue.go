package queue

import "math"

type Queue struct {
	nums []int
}

func NewQueue() Queue {
	return Queue{nums: make([]int, 0)}
}

func (q *Queue) Push(num int) {
	q.nums = append(q.nums, num)
}

func (q *Queue) Pop() int {
	if len(q.nums) == 0 {
		// queue is empty
		return math.MinInt32
	}
	top := q.nums[0]
	q.nums = q.nums[1:len(q.nums)]
	return top
}

func (q *Queue) Top() int {
	if len(q.nums) == 0 {
		return math.MinInt32
	}
	return q.nums[0]
}

func (q *Queue) Empty() bool {
	return len(q.nums) == 0
}

func (q *Queue) Len() int {
	return len(q.nums)
}
