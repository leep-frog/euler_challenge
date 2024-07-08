package y2017

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
)

func Day13() aoc.Day {
	return &day13{}
}

type day13 struct{}

type depthRange struct {
	d  int
	r  int
	dr int
}

func (d *day13) Solve(lines []string, o command.Output) {
	drs := functional.Map(lines, func(line string) *depthRange {
		parts := strings.Split(line, ": ")
		d, r := parse.Atoi(parts[0]), parse.Atoi(parts[1])
		return &depthRange{d, r, d * r}
	})

	part1, _ := d.severity(0, drs)
	for ps := 0; ; ps++ {
		if _, ok := d.severity(ps, drs); ok {
			o.Stdoutln(part1, ps)
			return
		}
	}
}

func (d *day13) severity(ps int, drs []*depthRange) (int, bool) {
	var severity int
	ok := true
	for _, dr := range drs {
		if (ps+dr.d)%((dr.r-1)*2) == 0 {
			severity += dr.dr
			ok = false
		}
	}
	return severity, ok
}

func (d *day13) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"24 10",
			},
		},
		{
			ExpectedOutput: []string{
				"1504 3823370",
			},
		},
	}
}
