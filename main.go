package main

import (
	"dsa/concurrency"
	"dsa/leetcode"
	"fmt"
	"sync"
)

func main() {
	fb := concurrency.NewFooBar(1)
	var wg sync.WaitGroup
	wg.Add(2)

	// Start two threads
	outputs := make([]string, 0)

	// Foo chain
	go func() {
		defer wg.Done()
		arr := []int{1, 3, 2, 3, 3}
		fmt.Printf("Rainwater solution for: %v is : %v\n", arr, leetcode.RainwaterSolve(arr))
		fb.Foo(func() {
			outputs = append(outputs, "foo")
		})
	}()

	// Bar chain
	go func() {
		defer wg.Done()
		arr1 := []int{23, 5, 10, 17, 30}
		arr2 := []int{26, 134, 135, 14, 19}
		i1, i2 := leetcode.SolveMinDifference(arr1, arr2)
		fmt.Printf("Min difference between: %v, and %v is: %v & %v\n", arr1, arr2, i1, i2)
		fb.Bar(func() {
			outputs = append(outputs, "bar")
		})
	}()

	// Wait for the threads to complete execution
	wg.Wait()

	fmt.Printf("outputs from the threads: %v", outputs)

	list1 := []int{1, 2, 4}
	list2 := []int{2, 10, 22}

	fmt.Printf("Merged arrays : %v\n", leetcode.Merge_two_sorted_lists(list1, list2))

	// Reset the outputs as it has served its purpose already
	outputs = make([]string, 0)

	mu := sync.Mutex{}

	// A list of functions which will run concurrently
	problems := []func(){
		func() {
			arr := []int{3, 4, 5, 6, 1, 2}
			flag := leetcode.SolveSortedOrRotated(arr)
			result := fmt.Sprintf("Is the array: %v sorted and rotated: %v", arr, flag)
			mu.Lock()
			outputs = append(outputs, result)
			mu.Unlock()
		},
	}

	wg.Add(len(problems))

	for _, problem := range problems {
		go func(p func()) {
			defer wg.Done()
			p()
		}(problem)
	}

	// Wait for all the problems to finish execution
	wg.Wait()
	for _, output := range(outputs) {
		fmt.Printf("Problems output: \n\t%v\n", output)
	}
}
