package leetcode

import (
	"dsa/stack"
)

func SubArrayRanges(nums []int) int64 {
	// Get the next smallest element for each of these entries.

	// Example:           [ 1, 4, 6, 7, 3, 7, 8,  1]
	// eg: nextMinNums =  [-1, 3, 3, 3, 1, 1, 1, -1] -- Done
	// eg: nextMaxNums =  [ 4, 6, 7, 8, 7, 8,-1, -1] --
	// eg: prevMaxNums =  [-1,-1,-1,-1, 7,-1,-1,  8]
	// eg: prevMinNums =  [-1, 1, 4, 6, 1, 3, 7, -1] -- Done
	// Indexes:           [ 0, 1, 2, 3, 4, 5, 6,  7] -- Done
	// eg: nextMinIndex = [ 8, 4, 4, 4, 7, 7, 7,  8] -- Done
	// eg: prevMinIndex = [-1, 0, 1, 2, 0, 4, 5,  0]
	// We start from 0th index.

	// We ask a simple question for the current element.
	// How many time can it contribute to the sum
	// we can find this out by finding out how
	// far is the next or prev smallest element from current index

	// Calculating next smallest number index

	// NOTE TODO: Can we somehow use this same loop to calculate both NSI, NGI (next smaller index, next greater index?)
	nextMinNumIndex := make([]int, len(nums))
	nextBigNumIndex := make([]int, len(nums))
	minNumStack := stack.InitStack()
	bigNumStack := stack.InitStack()

	// find NSI and NGI
	for i := len(nums) - 1; i >= 0; i-- {
		smallest := -1
		largest := -1
		for minNumStack.GetStackSize() != 0 {
			if nums[(minNumStack.Peek())] <= nums[i] {
				smallest = minNumStack.Peek()
				break
			}
			// Topmost element in stack is greater than current element
			// Remove this
			minNumStack.Pop()
		}

		for bigNumStack.GetStackSize() != 0 {
			if nums[bigNumStack.Peek()] >= nums[i] {
				largest = bigNumStack.Peek()
				break
			}

			// Topmost element index in the stack is smaller than the current element
			bigNumStack.Pop()
		}
		if smallest == -1 {
			smallest = len(nums)
		}
		if largest == -1 {
			largest = len(nums)
		}

		nextBigNumIndex[i] = largest
		bigNumStack.Push(i)
		nextMinNumIndex[i] = smallest
		minNumStack.Push(i)
	}

	// Reusing the stack
	minNumStack = stack.InitStack()
	bigNumStack = stack.InitStack()

	// Calculate previous smallest number index and previous greater element
	prevMinNumIndex := make([]int, len(nums))
	prevBigNumIndex := make([]int, len(nums))
	for i, num := range nums {
		smallest := -1
		largest := -1

		for minNumStack.GetStackSize() != 0 {
			if nums[minNumStack.Peek()] < num {
				smallest = minNumStack.Peek()
				break
			}
			minNumStack.Pop()
		}

		for bigNumStack.GetStackSize() != 0 {
			if nums[bigNumStack.Peek()] > nums[i] {
				largest = bigNumStack.Peek()
				break
			}

			// The current element is greater than or equal to the current head index of the stack
			bigNumStack.Pop()
		}

		// We do not care if the largest or smallest is set or not..
		// if these are not set, then that means we did not find any greater element
		prevBigNumIndex[i] = largest
		bigNumStack.Push(i)
		prevMinNumIndex[i] = smallest
		minNumStack.Push(i)
	}

	// our next smallest element is ready
	minSubArraySum := 0
	maxSubArraySum := 0

	for i, num := range nums {
		// PART 1: Calculate the sum of subArray minimums
		// Number of elements in left that the current num is smallest
		left := i - prevMinNumIndex[i]

		// Number of elements in right that the current num is smallest
		right := nextMinNumIndex[i] - i

		// Total contribution of the current element
		minSubArraySum += (left * right) * num

		// PART 2: Calculate the sum of subArray maximums
		// Number of elements in the left that the current num is largest
		left = i - prevBigNumIndex[i]

		// Number of elements in the right that the current num is largest
		right = nextBigNumIndex[i] - i
		maxSubArraySum += (left * right) * num
	}

	return int64(maxSubArraySum) - int64(minSubArraySum)
}
