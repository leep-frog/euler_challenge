package y2018

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/euler_challenge/rgx"
)

func Day03() aoc.Day {
	return &day03{}
}

type day03 struct{}

func (d *day03) Solve(lines []string, o command.Output) {
	r := rgx.New(`^#([1-9][0-9]*) @ ([0-9]*),([0-9]*): ([1-9][0-9]*)x([1-9][0-9]*)$`)
	spots := map[string]int{}
	for _, line := range lines {
		vs := r.MatchInts(line)
		for x := vs[1]; x < vs[1]+vs[3]; x++ {
			for y := vs[2]; y < vs[2]+vs[4]; y++ {
				p := point.New(x, y)
				spots[p.String()]++
			}
		}
	}

	var part1 int
	for _, v := range spots {
		if v > 1 {
			part1++
		}
	}

	for _, line := range lines {
		vs := r.MatchInts(line)
		for x := vs[1]; x < vs[1]+vs[3]; x++ {
			for y := vs[2]; y < vs[2]+vs[4]; y++ {
				p := point.New(x, y)
				if spots[p.String()] > 1 {
					goto INVALID
				}
			}
		}
		o.Stdoutln(part1, vs[0])
		return
	INVALID:
	}
}

func (d *day03) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"4 3",
			},
		},
		{
			ExpectedOutput: []string{
				"111630 724",
			},
		},
	}
}
