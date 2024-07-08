package y2020

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
)

func Day09() aoc.Day {
	return &day09{}
}

type day09 struct{}

func (d *day09) Solve(lines []string, o command.Output) {

	preamble := 25
	if len(lines) < 100 {
		preamble = 5
	}
	numbers := functional.Map(lines, parse.Atoi)
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
}

func (d *day09) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"127 62",
			},
		},
		{
			ExpectedOutput: []string{
				"3199139634 438559930",
			},
		},
	}
}
