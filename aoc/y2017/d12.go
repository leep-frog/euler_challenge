package y2017

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/unionfind"
)

func Day12() aoc.Day {
	return &day12{}
}

type day12 struct{}

func (d *day12) Solve(lines []string, o command.Output) {
	uf := unionfind.New()

	for _, line := range lines {
		sides := strings.Split(line, " <-> ")
		left := parse.Atoi(sides[0])
		for _, right := range strings.Split(sides[1], ", ") {
			uf.Merge(left, parse.Atoi(right))
		}
	}

	o.Stdoutln(functional.CountFunc(uf.Elements(), func(k int) bool {
		return uf.Connected(0, k)
	}), uf.NumberOfSets())
}

func (d *day12) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"6 2",
			},
		},
		{
			ExpectedOutput: []string{
				"134 193",
			},
		},
	}
}
