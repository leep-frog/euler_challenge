package y2020

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
)

func Day05() aoc.Day {
	return &day05{}
}

type day05 struct{}

func (d *day05) Solve(lines []string, o command.Output) {
	var max int
	has := map[int]bool{}
	for _, line := range lines {
		rv := 64
		var row, column int
		for _, fb := range line[:7] {
			if fb == 'B' {
				row += rv
			}
			rv /= 2
		}

		rv = 4
		for _, rl := range line[7:] {
			if rl == 'R' {
				column += rv
			}
			rv /= 2
		}
		id := row*8 + column
		max = maths.Max(max, id)
		has[id] = true
	}
	for i := 0; i < max; i++ {
		if has[i-1] && has[i+1] && !has[i] {
			o.Stdoutln(max, i)
			break
		}
	}
}

func (d *day05) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"823 821",
			},
		},
		{
			ExpectedOutput: []string{
				"989 548",
			},
		},
	}
}
