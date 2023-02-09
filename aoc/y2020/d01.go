package y2020

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day01() aoc.Day {
	return &day01{}
}

type day01 struct{}

func (d *day01) Solve(lines []string, o command.Output) {
	ks := parse.AtoiArray(lines)

	// Part 1
	a, b, _ := maths.TwoSum(2020, ks)

	// Part 2
	x, y, z, _ := maths.ThreeSum(2020, ks)
	o.Stdoutln(a*b, x*y*z)
}

func (d *day01) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"514579 241861950",
			},
		},
		{
			ExpectedOutput: []string{
				"1016619 218767230",
			},
		},
	}
}
