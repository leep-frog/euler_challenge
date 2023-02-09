package y2015

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day10() aoc.Day {
	return &day10{}
}

type day10 struct{}

func (d *day10) process(code string) string {
	var r []string
	var prevChar string
	var charCount int
	for idx, c := range strings.Split(code, "") {
		if idx == 0 {
			prevChar = c
			charCount++
			continue
		}

		if c == prevChar {
			charCount++
		} else {
			r = append(r, parse.Itos(charCount), prevChar)
			prevChar = c
			charCount = 1
		}
	}
	r = append(r, parse.Itos(charCount), prevChar)
	return strings.Join(r, "")
}

func (d *day10) Solve(lines []string, o command.Output) {
	code := lines[0]
	var part1 int
	for i := 0; i < 50; i++ {
		if i == 40 {
			part1 = len(code)
		}
		code = d.process(code)
	}
	o.Stdoutln(part1, len(code))
}

func (d *day10) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"82350 1166642",
			},
		},
		{
			ExpectedOutput: []string{
				"252594 3579328",
			},
		},
	}
}
