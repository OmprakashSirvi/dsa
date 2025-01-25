package leetcode

import (
	"fmt"
	"sort"
)

func SolveMinSwapToSort(arr []int) int {
	// arr: 7, 1, 3, 2, 4, 5, 6

	// Know the actual positions:
	pos := make(map[int]int)
	for i, ele := range arr {
		pos[ele] = i
	}

	// Solve the arr
	sort.Ints(arr)

	// Find the number of swaps
	swaps := 0
	visited := make([]bool, len(arr))

	index := i
	for i, ele := range arr {
		// If the current index is already processed then skip
		if visited[i] {
			continue
		}
		for {
			visited[index] = true
			index = pos[arr[index]]
			
		}

	}

	// [1 2 3 4 5 6 7]
	fmt.Printf("sortedArr: %v\n", arr)
	fmt.Printf("Previous positions: %v\n", pos)

	return 0
}
