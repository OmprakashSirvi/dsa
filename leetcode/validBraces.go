package leetcode

func ValidBraces(s string) bool {
	var l = make([]rune, 0)
	var comp = map[rune]rune{')': '(', '}': '{', ']': '['}

	for i, char := range s {
		// For first character insert it into stack
		if i == 0 {
			// If the first character is close braces then return false
			if comp[char] == 0 {
				l = append(l, char)
				continue
			}
			return false
		}

		// If current element is open brace,
		// check whether the last element in the list is complement to the current element or not
		if comp[char] == 0 {
			l = append(l, char)
			continue
		} else {
			// else the current element is close brace, now check for its complementary
			if len(l) > 0 && l[len(l)-1] == comp[char] {
				// Pop the last element
				l = l[:(len(l) - 1)]
				continue
			}
			return false

		}

	}

	return len(l) <= 0
}
