package y2022

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
)

func Day03() aoc.Day {
	return &day03{}
}

type day03 struct{}

func (d *day03) priority(c rune) int {
	v := c - 'a' + 1
	if v >= 1 && v <= 26 {
		return int(v)
	}
	return int(c - 'A' + 1 + 26)
}

func (d *day03) Solve(lines []string, o command.Output) {
	var sum1, sum2 int

	var badgePotential map[rune]bool
	for i, line := range lines {
		// Part 1
		left, right := line[:len(line)/2], line[len(line)/2:]
		m := map[rune]bool{}
		for _, l := range left {
			m[l] = true
		}
		for _, r := range right {
			if m[r] {
				sum1 += d.priority(r)
				break
			}
		}

		// Part 2
		if i%3 == 0 {
			if i != 0 {
				sum2 += d.priority(maps.Keys(badgePotential)[0])
			}
			badgePotential = maths.NewSimpleSet(parse.ToCharArray(line)...)
		} else {
			badgePotential = maths.Intersection(badgePotential, maths.NewSimpleSet(parse.ToCharArray(line)...))
		}
	}

	sum2 += d.priority(maps.Keys(badgePotential)[0])
	o.Stdoutln(sum1, sum2)
}

func (d *day03) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"157 70",
			},
		},
		{
			ExpectedOutput: []string{
				"8072 2567",
			},
		},
	}
}
