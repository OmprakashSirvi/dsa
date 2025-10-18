package leetcode

func FindTheDifference(s string, t string) byte {
	l := make([]int, 26)

	for _, c := range s {
		l[c-'a']++
	}

	for _, c := range t {
		if l[c-'a'] == 0 {
			return byte(c)
		}
		l[c-'a']--
	}

	return 0
}
