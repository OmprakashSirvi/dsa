package main

import (
	"dsa/concurrency"
	"fmt"
	"sync"
)

func main() {
	// arr1 := []int{23, 5, 10, 17, 30}
	// arr2 := []int{26, 134, 135, 14, 19}
	// arr := []int{1, 3, 2, 3, 3}
	// leetcode.RainWaterSolve(arr)
	// i1, i2 := leetcode.SolveMinDifference(arr1, arr2)
	// fmt.Printf("Result: %v %v", i1, i2)

	fb := concurrency.NewFooBar(1)
	var wg sync.WaitGroup
	wg.Add(2)

	fb.FooChan <- struct{}{}

	// Start two threads
	output := make([]string, 0)

	// Foo chain
	go func() {
		defer wg.Done()
		fb.Foo(func() {
			output = append(output, "foo")
		})
	}()

	// Bar chain
	go func() {
		defer wg.Done()
		fb.Bar(func() {
			output = append(output, "bar")
		})
	}()

	// Wait for the threads to complete execution
	wg.Wait()

	fmt.Printf("Output from the threads: %v", output)

	// list1 := []int{1, 2, 4}
	// list2 := []int{2, 10, 22}

	// fmt.Printf("Merged arrays : %v", leetcode.Merge_two_sorted_lists(list1, list2))
}
