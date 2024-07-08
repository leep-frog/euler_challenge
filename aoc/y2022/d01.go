package y2022

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/slices"
)

func Day01() aoc.Day {
	return &day01{}
}

type day01 struct{}

func (d *day01) Solve(lines []string, o command.Output) {
	// Part 1
	sums := functional.Map(parse.SplitOnLines(lines, ""), func(group []string) int {
		return bread.Sum(functional.Map(group, parse.Atoi))
	})
	o.Stdoutln(maths.Max(sums...))

	// Part 2
	slices.Sort(sums)
	sums = bread.Reverse(sums)
	o.Stdoutln(bread.Sum(sums[:3]))
}

func (d *day01) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"24000",
				"45000",
			},
		},
		{
			ExpectedOutput: []string{
				"68442",
				"204837",
			},
		},
	}
}
