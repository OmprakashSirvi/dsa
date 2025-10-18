package leetcode

func SearchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := len(nums) - 1
	var mid = left + ((right - left) / 2)
	for left <= right {
		mid = left + ((right - left) / 2)
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
