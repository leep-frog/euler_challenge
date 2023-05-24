package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day05() aoc.Day {
	return &day05{}
}

type day05 struct{}

func (d *day05) Solve(lines []string, o command.Output) {
	o.Stdoutln(functional.Map([]bool{false, true}, func(part2 bool) int {
		jumps := functional.Map(lines, parse.Atoi)
		var steps int
		for i := 0; i >= 0 && i < len(jumps); {
			steps++
			newI := i + jumps[i]
			if jumps[i] >= 3 && part2 {
				jumps[i]--
			} else {
				jumps[i]++
			}
			i = newI
		}
		return steps
	}))
}

func (d *day05) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"[0 0]",
			},
		},
		{
			ExpectedOutput: []string{
				"[325922 24490906]",
			},
		},
	}
}
