package y2020

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day01() aoc.Day {
	return &day01{}
}

type day01 struct{}

func (d *day01) Solve(lines []string, o command.Output) {
	ks := parse.AtoiArray(lines)
	m := map[int]bool{}
	var part1 int
	for _, k := range ks {
		if m[k] {
			part1 = k * (2020 - k)
			break
		}
		m[2020-k] = true
	}

	// Part 2
	for i, a := range ks {
		for j := i + 1; j < len(ks); j++ {
			b := ks[j]
			for k := j + 1; k < len(ks); k++ {
				c := ks[k]
				if a+b+c == 2020 {
					o.Stdoutln(part1, a*b*c)
				}
			}
		}
	}
}

func (d *day01) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"514579 241861950",
			},
		},
		{
			ExpectedOutput: []string{
				"1016619 218767230",
			},
		},
	}
}
