package leetcode

import "math"

func RainwaterSolve(height []int) int {

	var (
		waterStored int
		maxEleLeft  int = 0
		maxEleRight int = 0
	)
	for index := 1; index < len(height); index++ {
		maxEleLeft = height[index]
		maxEleRight = height[index]
		tempIndex := index - 1
		// iterate to left side to find the maxElement
		for tempIndex >= 0 {
			if height[tempIndex] > maxEleLeft {
				maxEleLeft = height[tempIndex]
			}
			tempIndex--
		}
		// if no element is greater than the current element
		// then continue to next iteration
		if maxEleLeft <= height[index] {
			continue
		}

		tempIndex = index + 1
		for tempIndex < len(height) {
			if height[tempIndex] > maxEleRight {
				maxEleRight = height[tempIndex]
			}
			tempIndex++
		}
		// if no element is greater than the current element
		// then continue to next iteration
		if maxEleRight <= height[index] {
			continue
		}
		waterStored += int(math.Min(float64(maxEleLeft), float64(maxEleRight))) - height[index]
	}

	return waterStored
}
