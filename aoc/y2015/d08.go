package y2015

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
)

func Day08() aoc.Day {
	return &day08{}
}

type day08 struct{}

type charState int

const (
	basic charState = iota
	singleBackslash
)

func (d *day08) Solve(lines []string, o command.Output) {
	var mem, chars, mem2 int

	for _, line := range lines {
		mem += len(line)
		mem2 += len(line) + 4 // Plus four for beginning and ending quote
		line = line[1 : len(line)-1]

		state := basic
		for i := 0; i < len(line); i++ {
			c := line[i : i+1]
			if state == basic {
				// Adding a regular character or starting an escape sequence (which equates to one character)
				chars++
				if c == "\\" {
					state = singleBackslash
				}
			} else {
				state = basic
				if c == "x" {
					i += 2
					mem2++
				} else {
					mem2 += 2
				}
			}
		}
	}
	o.Stdoutln(mem-chars, mem2-mem)
}

func (d *day08) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"12 19",
			},
		},
		{
			ExpectedOutput: []string{
				"1350 2085",
			},
		},
	}
}
