package y2015

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day20() aoc.Day {
	return &day20{}
}

type day20 struct{}

func (d *day20) part2(p *generator.Prime, k int) int {
	var sum int
	for _, f := range p.Factors(k) {
		if k/f <= 50 {
			sum += f * 11
		}
	}
	return sum
}

func (d *day20) Solve(lines []string, o command.Output) {
	s := parse.Atoi(lines[0])
	p := generator.Primes()

	// Part 1
	a := 1
	for ; bread.Sum(p.Factors(a))*10 < s; a++ {
	}

	// Part 2
	b := 1
	for ; d.part2(p, b) < s; b++ {
	}
	o.Stdoutln(a, b)
}

func (d *day20) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"6 6",
			},
		},
		{
			ExpectedOutput: []string{
				"665280 705600",
			},
		},
	}
}
