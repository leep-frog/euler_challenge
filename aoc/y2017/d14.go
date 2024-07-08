package y2017

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/euler_challenge/unionfind"
	"github.com/leep-frog/euler_challenge/walker"
)

func Day14() aoc.Day {
	return &day14{}
}

type day14 struct{}

func (d *day14) Solve(lines []string, o command.Output) {
	var count int
	used := maths.NewSet[*point.Point[int]]()
	for x := 0; x < 128; x++ {
		k := fmt.Sprintf("%s-%d", lines[0], x)
		v := maths.FromHex(knotHash(k))
		b := v.ToBinary().String()

		for y, c := range strings.Repeat("0", 128-len(b)) + b {
			if c == '1' {
				count++
				used.Add(point.New(x, y))
			}
		}
	}

	uf := unionfind.New[string]()
	used.For(func(p *point.Point[int]) bool {
		uf.Insert(p.String())
		for _, d := range walker.CardinalDirections(true) {
			q := p.Plus(d)
			if used.Contains(q) {
				uf.Merge(p.String(), q.String())
			}
		}
		return false
	})
	o.Stdoutln(count, uf.NumberOfSets())
}

func (d *day14) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"8108 1242",
			},
		},
		{
			ExpectedOutput: []string{
				"8140 1182",
			},
		},
	}
}
