package y2017

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day15() aoc.Day {
	return &day15{}
}

type day15 struct{}

type valueGenerator struct {
	value  int
	factor int
	rem    int
	divBy  int
}

func (g *valueGenerator) Next() int {
	g.value = (g.value * g.factor) % g.rem
	for (g.value % g.divBy) != 0 {
		g.value = (g.value * g.factor) % g.rem
	}
	return g.value
}

func (d *day15) solve(aStart, aDiv, bStart, bDiv, times int) int {
	rem := 2147483647
	a, b := &valueGenerator{aStart, 16807, rem, aDiv}, &valueGenerator{bStart, 48271, rem, bDiv}
	var count int
	for i := 0; i < times; i++ {
		av, bv := a.Next(), b.Next()
		if av%65536 == bv%65536 {
			count++
		}
	}
	return count
}

func (d *day15) Solve(lines []string, o command.Output) {

	// startA, startB := 65, 8921
	// startA, startB := 783, 325
	parts := strings.Split(lines[0], ",")
	startA, startB := parse.Atoi(parts[0]), parse.Atoi(parts[1])

	o.Stdoutln(d.solve(startA, 1, startB, 1, 40_000_000), d.solve(startA, 4, startB, 8, 5_000_000))
}

func (d *day15) Cases() []*aoc.Case {
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
