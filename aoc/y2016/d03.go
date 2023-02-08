package y2016

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day03() aoc.Day {
	return &day03{}
}

type day03 struct{}

func (d *day03) Solve(lines []string, o command.Output) {
	var cnt int
	for _, tri := range parse.ToGrid(lines, " ") {
		if tri[0]+tri[1] > tri[2] && tri[1]+tri[2] > tri[0] && tri[0]+tri[2] > tri[1] {
			cnt++
		}
	}

	var cnt2 int
	for _, tri := range maths.SimpleTranspose(parse.ToGrid(lines, " ")) {
		for i := 0; i < len(tri); i += 3 {
			if tri[i]+tri[i+1] > tri[i+2] && tri[i+1]+tri[i+2] > tri[i] && tri[i]+tri[i+2] > tri[i+1] {
				cnt2++
			}
		}
	}
	o.Stdoutln(cnt, cnt2)
}

func (d *day03) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3 6",
			},
		},
		{
			ExpectedOutput: []string{
				"1032 1838",
			},
		},
	}
}
