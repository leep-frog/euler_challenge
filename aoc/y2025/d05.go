package y2025

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day05() aoc.Day {
	return &day05{}
}

type day05 struct{}

func (d *day05) Solve(lines []string, o command.Output) {
	var freshRange *maths.Range

	var ids []int

	addingRanges := true
	for _, line := range lines {
		if line == "" {
			addingRanges = false
			continue
		}

		if addingRanges {
			parts := strings.Split(line, "-")
			newRange := maths.NewRange(parse.Atoi(parts[0]), parse.Atoi(parts[1]))
			if freshRange == nil {
				freshRange = newRange
			} else {
				freshRange = freshRange.Merge(newRange)
			}
		} else {
			id := parse.Atoi(line)
			ids = append(ids, id)
		}
	}

	var fresh int
	for _, id := range ids {
		if freshRange.Contains(id) {
			fresh++
		}
	}
	o.Stdoutln(fresh, freshRange.Size())
}

func (d *day05) Cases() []*aoc.Case {
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
