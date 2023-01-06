package y2020

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day09() aoc.Day {
	return &day09{}
}

type day09 struct{}

func (d *day09) Solve(lines []string, o command.Output) {

	preamble := 25
	numbers := parse.Map(lines, parse.Atoi)
	var invalid int
	for i, k := range numbers {
		if i >= preamble {
			if _, _, ok := maths.TwoSum(k, numbers[i-preamble:i]); !ok {
				invalid = k
				break
			}
		}
	}

	var start, sum int
	for end, k := range numbers {
		sum += k
		for sum > invalid {
			sum -= numbers[start]
			start++
		}
		if sum == invalid {
			rnge := numbers[start : end+1]
			o.Stdoutln(invalid, maths.Max(rnge...)+maths.Min(rnge...))
			return
		}
	}
	// for
}

func (d *day09) Cases() []*aoc.Case {
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
