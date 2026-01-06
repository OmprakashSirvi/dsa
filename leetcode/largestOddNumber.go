package leetcode

import "strings"

func LargestOddNumber(num string) string {
	var retStr strings.Builder
	for _, r := range num {
		// n := int(r)
		if retStr.Len() == 0 {
			retStr.WriteRune(r)
			continue
		}

		
	}

	return retStr.String()
}
