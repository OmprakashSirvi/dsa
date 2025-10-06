package stack

type MinStack struct {
	min  []int
	nums []int
}

func Constructor() MinStack {
	return MinStack{min: make([]int, 0), nums: make([]int, 0)}
}

func (s *MinStack) Push(val int) {
	s.nums = append(s.nums, val)
	if len(s.min) == 0 {
		s.min = append(s.min, val)
		return
	}
	minVal := min(val, s.min[len(s.min)-1])
	s.min = append(s.min, minVal)
}

func (s *MinStack) Pop() {
	s.nums = s.nums[:len(s.nums)-1]
	s.min = s.min[:len(s.min)-1]
}

func (s *MinStack) Top() int {
	return s.nums[len(s.nums)-1]
}

func (s *MinStack) GetMin() int {
	return s.min[len(s.min)-1]
}
