package main

import (
	"dsa/bits"
	"dsa/concurrency"
	"dsa/leetcode"
	linkedlist "dsa/linkedList"
	"dsa/recursion"
	"dsa/stack"
	"fmt"
	"strconv"
	"strings"
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

	fmt.Printf("outputs from the threads: %v\n", outputs)

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
		func() {
			list := linkedlist.New()
			list.Append(1)
			list.Append(2)
			list.Append(4)
			list.Append(6)
			list.Append(3)
			list.Append(5)
			list.Append(10)
			list.RemoveHead()
			err := list.Insert(1, 1)
			if err != nil {
				fmt.Printf("error while inserting element into the list: %v", err)
				return
			}
			if list.IsPresent(4) {
				fmt.Println("element 4 is present")
			}
			fmt.Printf("Elements in linked list: %v\n", list.GetListAsString())
			fmt.Printf("middle element of list: %v\n", list.GetMiddle())
			list.Reverse()
			fmt.Printf("reversed linked list: %v\n", list.GetListAsString())
			list.RemoveHead()
			fmt.Printf("Elements in linked list before reverse: %v\n", list.GetListAsString())
			fmt.Printf("sorting list: %s\n", list.GetListAsString())
			list.Sort()
			fmt.Printf("sorted linked list: %v\n", list.GetListAsString())
			// head := list.GetHeadNode()
			// revListHead := linkedlist.ReverseRec(head)
			// fmt.Printf("reverse linked list through recursion: %v\n", linkedlist.GetListAsStringFromNode(revListHead))
		},
		func() {
			fmt.Println(strings.Repeat("*", 20))
			recursion.PrintNum(6, true)
			fmt.Printf("Sum of first 3 numbers: %v\n", recursion.Sum(3))
			fmt.Printf("10 factorial: %v\n", recursion.Factorial(10))
			li := []int{1, 2, 3, 4, 5, 6}
			fmt.Printf("rev of array: %v, is %v\n", recursion.GetString(li), recursion.RevArray(li))
			fmt.Println(strings.Repeat("*", 20))
		},
		func() {
			fmt.Printf("5 (%s) and 6 (%s): %v (%s)\n", strconv.FormatInt(5, 3), strconv.FormatInt(6, 2), 5&6, strconv.FormatInt(5&6, 2))
			num := int64(127)
			index := 2
			fmt.Printf("%v (%s) %vth bit is set?: %v\n", num, strconv.FormatInt(num, index), index, bits.IsBitSet(int(num), index))

			fmt.Printf("is %v pow of 2? : %v\n", num, bits.IsPow2(int(num)))
			fmt.Printf("the number of set bits in %v is %v\n", num, bits.CountSetBits(int(num)))

			fmt.Printf("dividend: %v, divisor: %v, quotient: %v\n", 10, 3, bits.Divide(10, 3))
		},
		func() {
			fmt.Printf("verifying min stack")
			mstack := stack.Constructor()
			mstack.Push(0)
			mstack.Push(1)
			mstack.Push(0)
			mstack.GetMin()
			mstack.Pop()
			mstack.GetMin()
			mstack.Pop()
			mstack.GetMin()
			mstack.Pop()
			mstack.Push(-1)
			mstack.Push(-2)
			mstack.Push(-1)
			mstack.GetMin()
			mstack.Pop()
			mstack.GetMin()
			mstack.Pop()
		},
		func() {
			ele := leetcode.NextGreaterElement([]int{4,1,2}, []int{1, 3, 4, 2})
			fmt.Printf("ele: %v\n", ele)

			ele = leetcode.NextGreaterElement2([]int{1,2,3,4,3})
			fmt.Printf("ele: %v\n", ele)
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
	for _, output := range outputs {
		fmt.Printf("Problems output: \n\t%v\n", output)
	}
}
