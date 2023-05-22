package y2016

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/linkedlist"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day19() aoc.Day {
	return &day19{}
}

type day19 struct{}

type elf struct {
	next *elf
}

func (d *day19) Solve(lines []string, o command.Output) {
	n := parse.Atoi(lines[0])
	start := linkedlist.CircularNumbered(n)
	for numElves := n; numElves > 1; numElves-- {
		start.PopNext()
		start = start.Next
	}

	part1 := start.Value + 1

	// Part 2
	start = linkedlist.CircularNumbered(n)
	opposite := start
	for i := 0; i < n/2; i++ {
		opposite = opposite.Next
	}
	opposite = opposite.Prev

	for numElves := n; numElves > 1; numElves-- {
		opposite.PopNext()

		if (numElves-1)%2 == 0 {
			opposite = opposite.Next
		}
		start = start.Next
	}

	o.Stdoutln(part1, start.Value+1)
}

func (d *day19) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3 2",
			},
		},
		{
			ExpectedOutput: []string{
				"1815603 1410630",
			},
		},
	}
}
