package leetcode

import (
	"dsa/stack"
	"fmt"
	"math"
)

func Attempt1(nums1 []int, nums2 []int) []int {
	m := make(map[int]int, 0)
	largest := math.MinInt32
	for i := len(nums2) - 1; i >= 0; i-- {
		// If the current element is smaller than the largest element yet found,
		// We add an entry to the map.
		if nums2[i] < largest {
			m[nums2[i]] = largest
			continue
		}
		largest = nums2[i]
		// There is not greater element yet found for this entry
		m[nums2[i]] = -1
	}
	fmt.Printf("map: %v\n", m)

	// The map should be prepared by now..
	retArr := make([]int, 0)
	for _, num := range(nums1) {
		retArr = append(retArr, m[num])
	}

	return retArr
}

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
