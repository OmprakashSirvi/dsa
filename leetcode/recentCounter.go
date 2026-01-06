package leetcode

import "dsa/queue"

type RecentCounter struct {
	q      queue.Queue
	window int
}

func Constructor() RecentCounter {
	q := queue.NewQueue()
	return RecentCounter{
		q,
		3000,
	}
}

func (rc *RecentCounter) Ping(t int) int {
	bottom := t - rc.window
	if bottom > 0{
		for rc.q.First() < bottom && rc.q.Len() > 0 {
			rc.q.Dequeue()
		}
	}

	rc.q.Enqueue(t)
	return rc.q.Len()
}
