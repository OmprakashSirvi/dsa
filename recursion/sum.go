package recursion

func Sum(num int) int {
	if num == 0 {
		return num
	}
	return num + Sum(num-1)
}
