package leetcode

import (
	"dsa/stack"
	"strings"
)

func RemoveOuterParentheses(s string) string {
	st := stack.InitStack()
	var retStr strings.Builder

	for _, r := range s {
		// if it is opening braces, and there is nothing 
		// in the stack then we just push this and move forward.
		if r == '(' {
			st.Push(1)
			if st.GetStackSize() != 1 {
				retStr.WriteRune(r)
			}
			continue
		} 

		// closing braces are there..

		if r == ')' {
			st.Pop()
			if st.GetStackSize() != 0 {
				retStr.WriteRune(r)
			}
			continue
		}
	}

	return retStr.String()
}
