package leetcode

import "strings"

func LengthOfLastWord(s string) int {
	li := strings.Split(s, " ")
	for i := len(li) - 1; i>= 0; i-- {
		if li[i] != "" {
			return len(li[i])
		}
	}
	return 0
}
