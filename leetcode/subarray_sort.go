package leetcode

import (
	"fmt"
	"math"
)

// Find the smallest number which is out of sync, and the largest number which is out of sync
// Then find the correct positions for those two numbers

// eg: [2, 6, 4, 8, 10, 9, 15]

func SubArraySort(nums []int) int {
	// Normally we want that all the current num should be greater or equal than the previous element.
	sm := math.MaxInt32
	lg := math.MinInt32
	isSorted := true
	/*
		Detect if the array is already sorted.
		If array is already sorted then we just return.

		We need to identify the unsorted part,
		Iterate through the array to find out the places where the order is violated.

		Update the sm and lg (smallest and the largest value in unsorted region)

		Find the correct position of the values.

		The smallest one should go to the first position from left where sm < nums[i]
		The largest one should go to the last position from right where lg > nums[i]

		The sub-array starts at the sm_i (sm index) and it ends at the position where lg is located.
	*/
	for i, num := range nums {
		// We do not want to process the last entry as we have nums[i+1] operations in the code and that would lead to an error
		if i == len(nums)-1 {
			continue
		}
		fmt.Printf("Analyzing %vth position, value: %v\n", i, num)
		if num <= nums[i+1] {
			continue
		}

		// Array is not sorted
		if num > nums[i+1] {
			isSorted = false
			fmt.Println("This element is not in right position")
			// not in right position

			// Reassign the sm and lg if needed.
			if nums[i+1] < sm {
				fmt.Println("This is the smallest number so far")
				sm = nums[i+1]
			}
			if num > lg {
				fmt.Println("This is the largest number so far")
				lg = num
			}
		}
	}
	if isSorted {
		return 0
	}
	fmt.Printf("The smallest and largest numbers are: (%v, %v)\n", sm, lg)

	sm_i := 0
	// Now how do we find the correct positions of the scrambled values?
	for i, num := range(nums) {
		if sm < num {
			sm_i = i
			fmt.Printf("the smallest number should be in position: %v\n", sm_i)
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if lg > nums[i] {
			fmt.Printf("the largest number should be in position: %v\n", i)
			return (i - sm_i) + 1

		}
	}

	// If the largest number index is still not set then it should be at last
	return (len(nums) - sm_i)
}
