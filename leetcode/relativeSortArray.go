package leetcode

import "sort"

func RelativeSortSolve(arr1 []int, arr2 []int) []int {
	arr2Map := make(map[int]int)
	for _, e2 := range arr2 {
		arr2Map[e2] = 1
	}

	dumpArr := make([]int, 0)
	for _, e1 := range arr1 {
		if arr2Map[e1] > 0 {
			arr2Map[e1] += 1
			continue
		}
		dumpArr = append(dumpArr, e1)
	}
	
	sort.Ints(dumpArr)

	retArr := make([]int, 0)
	for _, e2 := range arr2 {
		for i := 1; i < arr2Map[e2]; i++ {
			retArr = append(retArr, e2)
		}
	}
	return append(retArr, dumpArr...)
}
