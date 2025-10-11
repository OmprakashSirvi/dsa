package leetcode

import "dsa/queue"

func FirstUniqChar(s string) int {
	// If a char is in string.
	// Its map representation will be at least 1.
	// For more than 1 appearances, it will be 2
	m := make(map[byte]int)
	q := queue.NewQueue()

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 0 {
			// If the current element is not present in the map
			// Add that element to the map
			m[s[i]] = 1
			q.Push(i)
			continue
		}
		m[s[i]] = 2
		continue
	}

	for !q.Empty() {
		// We have an element which has appeared 1 time
		if m[s[q.Top()]] == 1 {
			return q.Top()
		}
		q.Pop()
	}

	return -1
}
