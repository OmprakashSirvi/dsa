package leetcode

import (
	"dsa/stack"
)

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, 0)
	s := stack.InitStack()
	
	for i := len(nums2) - 1; i >= 0; i-- {
		// Need to check the stack now..

		// If the current number is greater than the top of stack.
		// Pop elements until the top is greater than num or the stack is empty
		largest := -1
		for s.GetStackSize() != 0 {
			if s.Peek() > nums2[i] {
				largest = s.Peek()
				break
			}
			// The current element is larger than the stack top
			s.Pop()
		}
		m[nums2[i]] = largest
		s.Push(nums2[i])
	}

	// At this point we should have the map ready
	retArr := make([]int, 0)
	for _, num := range(nums1) {
		retArr = append(retArr, m[num])
	}

	return retArr
}

func NextGreaterElement2(nums []int) []int {
	retArr := make([]int, len(nums))
	s := stack.InitStack()
	
	for i := len(nums) - 1; i >= 0; i-- {
		// Need to check the stack now..

		// If the current number is greater than the top of stack.
		// Pop elements until the top is greater than num or the stack is empty
		for s.GetStackSize() != 0 {
			if s.Peek() > nums[i] {
				break
			}
			// The current element is larger than the stack top
			s.Pop()
		}
		s.Push(nums[i])
	}

	// once again loop
	for i := len(nums) - 1; i >= 0; i-- {
		largest := -1
		for s.GetStackSize() != 0 {
			if s.Peek() > nums[i] {
				largest = s.Peek()
				break
			}
			s.Pop()
		}

		retArr[i] = largest
		s.Push(nums[i])
	}

	return retArr
}
