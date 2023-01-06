package y2022

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day20() aoc.Day {
	return &day20{}
}

type day20 struct{}

type simpleNode struct {
	prev, next *simpleNode
	value      int
}

func (d *day20) toList(nodes []*simpleNode) []int {
	var start *simpleNode
	for _, n := range nodes {
		if n.value == 0 {
			start = n
			break
		}
	}

	var order []int
	for i := 0; i < len(nodes); i++ {
		order = append(order, start.value)
		start = start.next
	}
	return order
}

func (d *day20) brute(lines []string, iterations, coefficient int) int {
	// Convert values to nodes
	nodes := functional.Map(lines, func(line string) *simpleNode {
		return &simpleNode{nil, nil, parse.Atoi(line) * coefficient}
	})

	// Create linked list
	for i, n := range nodes {
		n.next = nodes[(i+1)%len(nodes)]
		n.prev = nodes[(i+len(nodes)-1)%len(nodes)]
	}

	for i := 0; i < iterations; i++ {
		// Iterate over nodes in pre-set order
		for _, n := range nodes {
			// Determine number of moves
			moves := n.value % (len(nodes) - 1)
			if moves < 0 {
				moves = (moves) + (len(nodes) - 1)
			}

			// Move the node
			for i := 0; i < moves; i++ {
				n.prev.next, n.prev, n.next, n.next.prev, n.next.next, n.next.next.prev = n.next, n.next, n.next.next, n.prev, n, n
			}
		}
	}

	ordered := d.toList(nodes)
	var zeroIdx int
	for i, n := range ordered {
		if n == 0 {
			zeroIdx = i
			break
		}
	}

	return functional.Reduce(0, []int{1, 2, 3}, func(base, i int) int {
		return base + ordered[(zeroIdx+i*1000)%len(ordered)]
	})
}

func (d *day20) Solve(lines []string, o command.Output) {
	o.Stdoutln(
		d.brute(lines, 1, 1),
		d.brute(lines, 10, 811589153),
	)
}

func (d *day20) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3 1623178306",
			},
		},
		{
			ExpectedOutput: []string{
				"8302 656575624777",
			},
		},
	}
}
