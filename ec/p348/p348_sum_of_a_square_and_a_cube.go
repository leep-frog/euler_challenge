package p348

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P348() *ecmodels.Problem {
	return ecmodels.IntInputNode(348, func(o command.Output, n int) {

		var count, sum int

		// Consider the smallest values at a time
		h := maths.NewHeap[*node](func(n1, n2 *node) bool {
			return n1.value < n2.value
		})
		h.Push(newNode(2, 2))

		var palinCount int
		var palinValue int
		for count < n {
			cubeSquare := h.Pop()

			if cubeSquare.value == palinValue {
				palinCount++
			} else {
				// Check the previous value
				if maths.Palindrome(palinValue) && palinCount == 4 {
					count++
					sum += palinValue
				}
				palinValue = cubeSquare.value
				palinCount = 1
			}

			for _, neighbor := range cubeSquare.nexts() {
				h.Push(neighbor)
			}
		}

		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1"},
			Want:     "5229225",
			Estimate: 0.5,
		},
		{
			Args:     []string{"5"},
			Want:     "1004195061",
			Estimate: 6,
		},
	})
}

type node struct {
	cube   int
	square int
	value  int
}

func newNode(cube, square int) *node {
	return &node{
		cube:   cube,
		square: square,
		value:  cube*cube*cube + square*square,
	}
}

func (n *node) nexts() []*node {
	var ns []*node
	if n.cube == n.square {
		ns = append(ns, newNode(n.cube+1, n.square+1))
	}

	if n.cube >= n.square {
		ns = append(ns, newNode(n.cube+1, n.square))
	}

	if n.square >= n.cube {
		ns = append(ns, newNode(n.cube, n.square+1))
	}

	return ns
}
