package leetcode

import "dsa/stack"

func DailyTemperatures(temperatures []int) []int {
	// [73,74,75,71,69,72,76,73]

	// This will store the indexes of the elements..
	s := stack.InitStack()
	nextGreaterElements := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		// loop till there is something in the stack..
		tempFound := false
		for s.GetStackSize() != 0 {
			// We found temperature which is greater than the current temperature..
			if temperatures[i] < temperatures[s.Peek()] {
				nextGreaterElements[i] = s.Peek()
				tempFound = true
				break
			}

			// Else if the current temperature is greater than the topmost element in the stack
			// pop that..
			s.Pop()
		}

		if !tempFound {
			// Not necessary as go will default unassigned to 0..
			// But for the sake of clarity
			nextGreaterElements[i] = 0
		}

		s.Push(i)
	}

	for i, next := range nextGreaterElements {
		if next == 0 {
			continue
		}
		nextGreaterElements[i] = nextGreaterElements[i] - i
	}

	return nextGreaterElements
}
