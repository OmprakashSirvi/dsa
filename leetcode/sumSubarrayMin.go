package leetcode

import (
	"dsa/stack"
	"math"
)

func SumSubarrayMins(arr []int) int {
	// Get the next smallest element for each of these entries.
	
	// Example:           [ 1, 4, 6, 7, 3, 7, 8,  1]
	// eg: nextMinNums =  [-1, 3, 3, 3, 1, 1, 1, -1]
	// eg: prevMinNums =  [-1, 1, 4, 6, 1, 3, 7, -1]
	// Indexes:           [ 0, 1, 2, 3, 4, 5, 6,  7]
	// eg: nextMinIndex = [ 7, 4, 4, 4, 7, 7, 7,  8]
	// eg: prevMinIndex = [-1, 0, 1, 2, 0, 4, 5, -1]
	// We start from 0th index.

	// We ask a simple question for the current element.
	// How many time can it contribute to the sum
	// we can find this out by finding out how
	// far is the next or prev smallest element from current index

	// Calculating next smallest number index
	nextMinNumIndex := make([]int, len(arr))
	s := stack.InitStack()
	for i := len(arr) - 1; i >= 0; i-- {
		smallest := -1
		for s.GetStackSize() != 0 {
			if arr[(s.Peek())] <= arr[i] {
				smallest = s.Peek()
				break
			}
			// Topmost element in stack is greater than current element
			// Remove this
			s.Pop()
		}
		if smallest == -1 {
			smallest = len(arr)
		}
		nextMinNumIndex[i] = smallest
		s.Push(i)
	}

	// Reusing the stack
	s = stack.InitStack()

	// Calculate previous smallest number index
	prevMinNumIndex := make([]int, len(arr))
	for i, num := range arr {
		smallest := -1
		for s.GetStackSize() != 0 {
			if arr[s.Peek()] < num {
				smallest = s.Peek()
				break
			}
			s.Pop()
		}
		prevMinNumIndex[i] = smallest
		s.Push(i)
	}

	// our next smallest element is ready
	sum := 0
	mod := int(math.Pow10(9) + 7)
	for i, num := range arr {
		// Number of elements in left that the current num is smallest
		left := i - prevMinNumIndex[i]

		// Number of elements in right that the current num is smallest
		right := nextMinNumIndex[i] - i

		// Total contribution of the current element
		sum += (left * right % mod) * num % mod
		sum %= mod

	}

	return sum
}
