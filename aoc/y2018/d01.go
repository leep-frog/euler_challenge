package y2018

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day01() aoc.Day {
	return &day01{}
}

type day01 struct{}

func (d *day01) Solve(lines []string, o command.Output) {
	var freq int
	has := map[int]bool{0: true}
	var part1, part2 int
	var gotPart2 bool
	for !gotPart2 {
		for _, line := range lines {
			freq += parse.Atoi(line)
			if has[freq] && !gotPart2 {
				gotPart2 = true
				part2 = freq
			}
			has[freq] = true
		}
		if part1 == 0 {
			part1 = freq
		}
	}
	o.Stdoutln(part1, part2)

}

func (d *day01) Cases() []*aoc.Case {
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
