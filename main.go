package main

import (
	"dsa/leetcode"
	"fmt"
)

func main() {
	arr := []int{7, 1, 3, 2, 4, 5, 6}
	// arr2 := []int{2,1,4,3,9,6}
	// arr := []int{1, 3, 2, 3, 3}
	// leetcode.RainWaterSolve(arr)

	fmt.Printf("Result: %v", leetcode.SolveMinSwapToSort(arr))

	// list1 := []int{1, 2, 4}
	// list2 := []int{2, 10, 22}

	// fmt.Printf("Merged arrays : %v", leetcode.Merge_two_sorted_lists(list1, list2))
}
