package recursion

import "fmt"

// Print in stdout from 0-n, if rev i true then it prints in reverse order
func PrintNum(num int, rev bool){
	defer fmt.Println("done")
	if rev {
		printNumsRev(num, num)
		return
	}

	printNums(1, num)
}

func printNums(index int, num int) {
	if index > num {
		return
	}
	fmt.Printf("%v | ", index)
	index += 1
	printNums(index, num)
}

func printNumsRev(index int, num int) {
	if index == 0 {
		return
	}

	fmt.Printf("%v | ", index)
	index -= 1
	printNumsRev(index, num)
}
