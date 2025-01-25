package leetcode

import (
	"fmt"
	"sort"
)

func SolveMinSwapToSort(arr []int) int {
	// Know the actual positions:
	pos := make(map[int]int)
	for i, ele := range arr {
		pos[ele] = i
	}
	fmt.Printf("Original array: %v\n", arr)
	fmt.Printf("Positions: %v\n", pos)

	// Solve the arr
	sort.Ints(arr)
	fmt.Printf("Sorted array: %v\n", arr)

	// Find the number of swaps
	visited := make([]bool, len(arr))

	totalSwaps := 0
	for i := range arr {
		swaps := 0
		index := i
		fmt.Printf("Current node: %v\n", i)
		// If the current index is already processed then skip
		if visited[i] {
			fmt.Println("This node is already visited")
			continue
		}
		for {
			fmt.Println("marking this node as visited")
			visited[index] = true
			index = pos[arr[index]]
			swaps += 1
			fmt.Printf("next index: %v\n", index)

			if index == i {
				fmt.Println("Completed a cycle")
				break
			}
		}
		totalSwaps += swaps - 1
	}
	return totalSwaps
}
