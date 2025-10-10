package leetcode

import (
	"dsa/stack"
)

func AsteroidCollision(asteroids []int) []int {
	s := stack.InitStack()
	freeAst := make([]int, 0)

	for _, asteroid := range asteroids {
		// If the asteroid is moving right, just push that to defense stack
		if asteroid > 0 {
			s.Push(asteroid)
			continue
		}

		// when will the two asteroids collide ?

		// If the incoming asteroid is negative i.e. moving left,
		// and the top of the asteroid stack is moving positive i.e. moving right

		// we loop till the two asteroids are colliding
		absAsteroid := -asteroid
		for asteroid < 0 && s.Peek() > 0 && absAsteroid > 0{
			// Incoming will destroy the top
			if absAsteroid > s.Peek() {
				s.Pop()
				continue
			}

			// both asteroids will destroy itself
			if absAsteroid == s.Peek() {
				absAsteroid = 0
				s.Pop()
				break
			}

			// Incoming will be destroyed by defense
			if absAsteroid < s.Peek() {
				absAsteroid = 0
			}
		}

		// Did the incoming asteroid survive the defense?
		if absAsteroid > 0 {
			freeAst = append(freeAst, asteroid)
		}
	}

	return append(freeAst, s.GetStackArr()...)
}
