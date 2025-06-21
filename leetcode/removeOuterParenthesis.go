package leetcode

/*
Create a stack to keep track of the parenthesis depth.
If the stack is empty and we encounter a open parenthesis, then we have a start of a part of primitive.
If we encounter a closed parenthesis, then we pop from the stack, and if the stack is empty,
we append the primitive part to the result string
*/
func RemoveOuterParentheses(s string) string {
	length := 0
	result := ""
	prim := ""

	for _, c := range s {
		if c == '(' {
			// Check if the stack is empty
			length += 1
			if length == 1 {
				continue
			}
			// The stack was not empty, hence this will count as a part of primitive.
			prim += string(c)
			continue
			// We know we would only receive valid strings, so no need to check for validity
		} else {
			prim += string(c)
			// pop the stack
			length -= 1
		}

		// If the stack is empty, then we have reached the end of the part primitive.
		if length == 0 {
			result += prim
			prim = ""
		}
	}

	return result
}
