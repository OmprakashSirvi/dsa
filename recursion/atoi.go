package recursion

import "errors"

func MyAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	retNum := atoi(s, 0)
	return retNum
}

func atoi(s string, num int) int {
	// base condition
	if len(s) == 0 {
		return 0
	}

	n, err := convert(rune(s[0]), num)
	if err != nil {
		return num
	}

	num = (num * 10) + n
	return atoi(s[1:], num)
}

func convert(c rune, num int) (int, error) {
	if c == ' ' {
		return 0, errors.New("empty space detected")
	}

	return 0, nil
}
