package y2018

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day12() aoc.Day {
	return &day12{}
}

type day12 struct{}

type plantRule struct {
	input  []bool
	output bool
}

func (d *day12) Solve(lines []string, o command.Output) {
	state := map[int]bool{}
	for i, c := range strings.Split(strings.TrimPrefix(lines[0], "initial state: "), "") {
		if c == "#" {
			state[i] = true
		}
	}

	for g := 0; g < 20; g++ {

	}

}

func (d *day12) Cases() []*aoc.Case {
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
