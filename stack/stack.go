package stack

import "math"

type stack struct {
	nums []int
}

func InitStack() stack {
	return stack{nums: make([]int, 0)}
}

func (s *stack) Push(num int) {
	s.nums = append(s.nums, num)
}

func (s *stack) Pop() int {
	if len(s.nums) == 0 {
		return math.MinInt32
	}
	last := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return last
}

func (s *stack) Peek() int {
	if len(s.nums) == 0 {
		return math.MinInt32
	}
	return s.nums[len(s.nums)-1]
}

func (s *stack) GetStackSize() int {
	return len(s.nums)
}

func (s *stack) GetStackArr() []int {
	return s.nums
}
