package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day01() aoc.Day {
	return &day01{}
}

type day01 struct{}

func (d *day01) Solve(lines []string, o command.Output) {
	var sums []int
	for _, offset := range []int{1, len(lines[0]) / 2} {
		var sum int
		for i, c := range lines[0] {
			next := (i + offset) % len(lines[0])
			if c == rune(lines[0][next]) {
				sum += parse.Atoi(string(c))
			}
		}
		sums = append(sums, sum)
	}
	o.Stdoutln(sums[0], sums[1])
}

func (d *day01) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3 4",
			},
		},
		{
			ExpectedOutput: []string{
				"1223 1284",
			},
		},
	}
}
