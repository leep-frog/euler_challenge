package y2017

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/slices"
)

func Day04() aoc.Day {
	return &day04{}
}

type day04 struct{}

func (d *day04) Solve(lines []string, o command.Output) {

	o.Stdoutln(functional.Map([]bool{true, false}, func(part2 bool) int {
		return functional.CountFunc(lines, func(s string) bool {

			parts := strings.Split(s, " ")
			if part2 {
				parts = functional.Map(parts, func(s string) string {
					ss := strings.Split(s, "")
					slices.Sort(ss)
					return strings.Join(ss, "")
				})
			}
			slices.Sort(parts)
			m := maths.NewSimpleSet(parts...)
			return len(m) == len(parts)
		})
	}))
}

func (d *day04) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"[0 0]",
			},
		},
		{
			ExpectedOutput: []string{
				"[231 337]",
			},
		},
	}
}
