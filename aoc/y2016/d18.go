package y2016

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day18() aoc.Day {
	return &day18{}
}

type day18 struct{}

func (d *day18) toString(traps []bool) string {
	var r []string
	for _, t := range traps {
		if t {
			r = append(r, "^")
		} else {
			r = append(r, ".")
		}
	}
	return strings.Join(r, "")
}

func (d *day18) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines[0], parse.Atoi(lines[1])), d.solve(lines[0], parse.Atoi(lines[2])))
}

func (d *day18) solve(start string, rows int) int {
	var safeCount int
	var traps []bool
	for _, c := range start {
		traps = append(traps, c == '^')
		if c != '^' {
			safeCount++
		}
	}

	for i := 0; i < rows-1; i++ {
		var nt []bool
		for ti, t := range traps {
			left := ti > 0 && traps[ti-1]
			middle := t
			right := ti < len(traps)-1 && traps[ti+1]

			isTrap := (left && middle && !right) || (!left && middle && right) || (left && !middle && !right) || (!left && !middle && right)
			nt = append(nt, isTrap)

			if !isTrap {
				safeCount++
			}
		}
		traps = nt
	}
	return safeCount
}

func (d *day18) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"38 38",
			},
		},
		{
			ExpectedOutput: []string{
				"2005 20008491",
			},
		},
	}
}
