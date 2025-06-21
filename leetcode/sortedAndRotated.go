package leetcode

import (
	"dsa/utils"
)

func SolveSortedOrRotated(nums []int) bool {
	logger := utils.InitLoggerWithName("SolveSortedOrRotated", "debug")
	logger.Infof("Solving sorted or rotated for array: %v", nums)
	// Find the dip point of the array. (Rotation value)
	dips := 0
	for i, num := range nums {
		// Since we deal with i+1 elements, we need to terminate this loop when it reaches the last element
		if i == len(nums)-1 {
			break
		}

		if num > nums[i+1] {
			dips += 1
		}
	}

	if dips == 0 {
		// The array is sorted
		return true
	}
	if dips > 1 {
		// The array is nor sorted nor rotated.
		return false
	}

	if nums[len(nums)-1] <= nums[0] {
		return true
	}

	return false

	// Loop from the rotation point and check if is sorted through the rotation point
	// for i := range len(nums) {
	// 	logger.Debugf("Operating on: %v", i)
	// }
}
