package y2020

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/linkedlist"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day23() aoc.Day {
	return &day23{}
}

type day23 struct{}

func (d *day23) Solve(lines []string, o command.Output) {
	part1 := d.solve(lines, 9, 100)
	var r []string
	for part1.Value != 0 {
		r = append(r, parse.Itos(part1.Value+1))
		part1 = part1.Next
	}

	part2 := d.solve(lines, 1_000_000, 10_000_000)
	o.Stdoutln(strings.Join(r, ""), (part2.Value+1)*(part2.Next.Value+1))
}

func (d *day23) getCups(lines []string, numCups int) (*linkedlist.Node[int], []*linkedlist.Node[int]) {
	cupMap := make([]*linkedlist.Node[int], numCups, numCups)

	var orderedCups []*linkedlist.Node[int]
	for _, r := range lines[0] {
		id := parse.Atoi(string(r)) - 1
		c := &linkedlist.Node[int]{Value: id}
		cupMap[id] = c
		orderedCups = append(orderedCups, c)
	}

	for len(orderedCups) < numCups {
		id := len(orderedCups)
		c := &linkedlist.Node[int]{Value: id}
		cupMap[id] = c
		orderedCups = append(orderedCups, c)
	}

	for i, c := range orderedCups {
		c.Next = orderedCups[(i+1)%numCups]
		c.Prev = orderedCups[(i+numCups-1)%numCups]
	}
	return orderedCups[0], cupMap
}

func (d *day23) solve(lines []string, nc, numMoves int) *linkedlist.Node[int] {
	cur, cupMap := d.getCups(lines, nc)

	for i := 0; i < numMoves; i++ {
		removed := map[int]bool{}
		poppedEnd := cur.Next
		poppedStart := poppedEnd
		for k := 0; k < 3; k++ {
			removed[poppedEnd.Value] = true
			poppedEnd = poppedEnd.Next
		}
		poppedEnd = poppedEnd.Prev
		cur.Next, poppedEnd.Next.Prev = poppedEnd.Next, cur

		// Get the index
		index := (cur.Value + nc - 1) % nc
		for removed[index] {
			index = (index + nc - 1) % nc
		}
		dest := cupMap[index]
		dest.Next, poppedStart.Prev, poppedEnd.Next, dest.Next.Prev = poppedStart, dest, dest.Next, poppedEnd
		cur = cur.Next
	}

	for cur.Value != 0 {
		cur = cur.Next
	}
	return cur.Next
}

func (d *day23) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}
