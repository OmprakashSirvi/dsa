package main

import (
	"dsa/leetcode"
	"fmt"
)

func main() {
	arr := []int{2, 6, 4, 8, 10, 9, 15}
	// arr := []int{1, 3, 2, 3, 3}
	// leetcode.RainWaterSolve(arr)

	fmt.Printf("Result: %v", leetcode.SubArraySort(arr))

	// list1 := []int{1, 2, 4}
	// list2 := []int{2, 10, 22}

	// fmt.Printf("Merged arrays : %v", leetcode.Merge_two_sorted_lists(list1, list2))
}
