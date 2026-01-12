package y2025

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
	dial := 50
	var atZeroCount, passZeroCount int

	for _, line := range lines {
		direction, number := line[0], parse.Atoi(line[1:])

		passZeroCount += number / 100
		number = number % 100

		if direction == 'L' {
			original := dial
			dial -= number
			if dial < 0 {
				dial += 100
				if original != 0 {
					passZeroCount++
				}
			} else if dial == 0 {
				passZeroCount++
			}
		} else {
			dial += number
			for dial >= 100 {
				dial -= 100
				passZeroCount++
			}
		}

		if dial == 0 {
			atZeroCount++
		}
	}
	o.Stdoutln(atZeroCount, passZeroCount)
}

func (d *day01) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3",
			},
		},
		{
			ExpectedOutput: []string{
				"1029",
			},
		},
	}
}
