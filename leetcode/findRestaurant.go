package leetcode

import "math"

func FindRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int)

	for i, l := range list1 {
		m[l] = i + 1
	}

	minIndex := math.MaxInt32
	for i, s := range list2 {
		if m[s] > 0 {
			// the current string is present in the first list
			index := i + (m[s] - 1)
			if index < minIndex {
				minIndex = index
			}
		}
	}

	if minIndex == math.MaxInt32 {
		// no common strings were found
		return []string{}
	}

	// at this point we would have the min index
	retList := make([]string, 0)
	for i, s := range list2 {
		if m[s] > 0 {
			index := i + (m[s] - 1)
			if index == minIndex {
				retList = append(retList, s)
			}
		}
	}

	return retList
}
