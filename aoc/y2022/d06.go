package y2022

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
)

func Day06() aoc.Day {
	return &day06{}
}

type day06 struct{}

func (d *day06) Solve(lines []string, o command.Output) {
	d.solve(lines, o, 4)
	d.solve(lines, o, 14)
}

func (d *day06) solve(lines []string, o command.Output, needLength int) {
	// A single line
	for _, line := range lines {
		m := map[rune]int{}
		length := 1
		for i, c := range line {
			length = maths.Min(length, i-m[c])
			m[c] = i

			if length == needLength {
				o.Stdoutln(i + 1)
				break
			}
			length++
		}
	}
}

func (d *day06) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"7",
				"5",
				"6",
				"10",
				"11",
				"19",
				"23",
				"23",
				"29",
				"26",
			},
		},
		{
			ExpectedOutput: []string{
				"1658",
				"2260",
			},
		},
	}
}
