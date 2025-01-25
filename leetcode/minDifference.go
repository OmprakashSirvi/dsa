package leetcode

import (
	"math"
	"sort"
)

func SolveMinDifference(arr1 []int, arr2 []int) (int, int) {
	// 26, 134, 135, 14, 19
	// 23, 5, 10, 17, 30
	sort.Ints(arr1)
	sort.Ints(arr2)
	var pair []int
	minSum := math.MaxFloat32

	for _, e1 := range arr1 {
		for _, e2 := range arr2 {
			diff := math.Abs(float64(e1 - e2))
			if minSum > diff {
				minSum = diff
				pair = []int{e1, e2}
			} else {
				break
			}
		}
	}

	if len(pair) == 2 {
		return pair[0], pair[1]
	}
	return 0, 0
}
