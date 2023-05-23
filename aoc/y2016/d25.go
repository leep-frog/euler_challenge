package y2016

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day25() aoc.Day {
	return &day25{}
}

type day25 struct{}

func (d *day25) Solve(lines []string, o command.Output) {
	do := &day12{}
	for i := 0; i < 1000; i++ {
		wantZero := true
		var cnt int
		var success bool
		outFunc := func(k int) bool {
			if k != 0 && k != 1 {
				return false
			}
			if wantZero != (k == 0) {
				return false
			}
			wantZero = !wantZero
			cnt++
			if cnt > 100 {
				success = true
				return false
			}
			return true
		}
		do.solve(lines, map[string]int{"a": i}, outFunc)

		if success {
			o.Stdoutln(i)
			return
		}
	}
}

func (d *day25) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"198",
			},
		},
		{
			ExpectedOutput: []string{
				"198",
			},
		},
	}
}
