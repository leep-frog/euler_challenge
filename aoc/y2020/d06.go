package y2020

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
)

func Day06() aoc.Day {
	return &day06{}
}

type day06 struct{}

func (d *day06) Solve(lines []string, o command.Output) {
	yeses := map[rune]bool{}
	var yeses2 map[string]bool
	var sum1, sum2 int
	for _, line := range append(lines, "") {
		if line == "" {
			sum1 += len(yeses)
			sum2 += len(yeses2)
			yeses = map[rune]bool{}
			yeses2 = nil
		} else {
			if yeses2 == nil {
				yeses2 = maths.NewSimpleSet(strings.Split(line, "")...)
			} else {
				yeses2 = maths.Intersection(yeses2, maths.NewSimpleSet(strings.Split(line, "")...))
			}
			for _, c := range line {
				yeses[c] = true
			}
		}
	}
	o.Stdoutln(sum1, sum2)
}

func (d *day06) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"11 6",
			},
		},
		{
			ExpectedOutput: []string{
				"6273 3254",
			},
		},
	}
}
