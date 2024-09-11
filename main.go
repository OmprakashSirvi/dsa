package main

import (
	"dsa/leetcode"
	"fmt"
)

func main() {
	// arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// leetcode.RainWaterSolve(arr)

	list1 := []int{1, 2, 4}
	list2 := []int{2, 10, 22}

	fmt.Printf("Merged arrays : %v", leetcode.Merge_two_sorted_lists(list1, list2))
}
