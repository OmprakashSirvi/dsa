package recursion

import "fmt"

func RevArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	return append(RevArray(nums[1:]), nums[0])
}

func GetString(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	if len(nums) == 1 {
		return fmt.Sprintf(" %v ", nums[0])
	}

	return fmt.Sprintf("[ %v, %v ]", nums[0], getRest(nums[1:]))
}

func getRest(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	if len(nums) == 1 {
		return fmt.Sprintf(" %v", nums[0])
	}

	return fmt.Sprintf(" %v, %v", nums[0], getRest(nums[1:]))
}
