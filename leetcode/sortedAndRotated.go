package leetcode

import "fmt"

func SolveSortedOrRotated(nums []int) bool {
	// Find the dip point of the array. (Rotation value)
	for i, num := range nums {
		// Since we deal with i+1 elements, we need to terminal this loop when it reaches the last element
		if i == len(nums)-1 {
			break
		}

		dips := 0
		rotation := 0
		if num > nums[i+1] {
			dips += 1
			rotation = i + 1
		}

		if dips == 0 {
			// The array is sorted
			return true
		}
		if dips > 1 {
			// The array is nor sorted nor rotated.
			return false
		}
		fmt.Println("Rotation: ", rotation)
	}

	// Loop from the rotation point and check if is sorted through the rotation point
	for i := range(len(nums)) {
		fmt.Println(i)
	}

	return true
}
