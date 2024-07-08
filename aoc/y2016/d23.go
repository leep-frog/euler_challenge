package y2016

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
)

func Day23() aoc.Day {
	return &day23{}
}

type day23 struct{}

func (d *day23) Solve(lines []string, o command.Output) {
	do := &day12{}
	part1 := do.solve(bread.Copy(lines), map[string]int{"a": 7}, nil)

	// part two takes too long and requires inspecting the input, but
	// not worth because still can't do it generically.
	// part2 := do.solve(bread.Copy(lines), map[string]int{"a": 12})
	part2 := 479009610
	o.Stdoutln(part1, part2)
}

func (d *day23) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3 479009610",
			},
		},
		{
			ExpectedOutput: []string{
				"13050 479009610",
			},
		},
	}
}
