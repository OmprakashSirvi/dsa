package leetcode

func SolveMaxSubarraySum(arr []int) int {
	// -1,2,3,4,-2,6,-8,3
	maxSum := 0
	currentSum := 0
	for _, ele := range arr {
		currentSum += ele
		if currentSum < 0 {
			currentSum = 0
			continue
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
