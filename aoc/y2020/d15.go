package y2020

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day15() aoc.Day {
	return &day15{}
}

type day15 struct{}

func (d *day15) Solve(lines []string, o command.Output) {
	startingNumbers := parse.AtoiArray(strings.Split(lines[0], ","))
	m := map[int]int{}
	for i, k := range startingNumbers {
		m[k] = i
	}

	// Note: we assume nextNum is zero, but that isn't true if
	// there are duplicates in the starting numbers
	var nextNum int

	var part1 int

	for idx := len(startingNumbers); idx < 30000000-1; idx++ {
		if k, ok := m[nextNum]; ok {
			nextNextNum := idx - k
			m[nextNum] = idx
			nextNum = nextNextNum
		} else {
			m[nextNum] = idx
			nextNum = 0
		}
		// Subtract 1 for 0-index and another 1 because current idx is for the number ahead
		if idx == 2020-2 {
			part1 = nextNum
		}
	}
	o.Stdoutln(part1, nextNum)
}

func (d *day15) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"436 175594",
			},
		},
		{
			ExpectedOutput: []string{
				"763 1876406",
			},
		},
	}
}
