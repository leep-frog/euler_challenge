package y2025

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day07() aoc.Day {
	return &day07{}
}

type day07 struct{}

func (d *day07) Solve(lines []string, o command.Output) {
	grid := parse.MapToGrid(lines, map[rune]int{
		'S': 2,
		'^': 1,
		'.': 0,
	})

	beamCounts := []int{}
	for _, cell := range grid[0] {
		if cell == 2 {
			beamCounts = append(beamCounts, 1)
		} else {
			beamCounts = append(beamCounts, 0)
		}
	}

	var splitCount int
	for _, row := range grid[1:] {
		newBeamCounts := make([]int, len(beamCounts))
		for x, cell := range row {
			if beamCounts[x] == 0 {
				continue
			}

			if cell == 0 {
				newBeamCounts[x] += beamCounts[x]
			} else {
				splitCount++
				if x > 0 {
					newBeamCounts[x-1] += beamCounts[x]
				}
				if x < len(beamCounts)-1 {
					newBeamCounts[x+1] += beamCounts[x]
				}
			}
		}
		beamCounts = newBeamCounts
	}

	o.Stdoutln(splitCount, bread.Sum(beamCounts))
}

func (d *day07) Cases() []*aoc.Case {
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
