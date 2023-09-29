package y2018

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/slices"
)

func Day14() aoc.Day {
	return &day14{}
}

type day14 struct{}

func (d *day14) Solve(lines []string, o command.Output) {
	recipes := []int{3, 7}
	elves := []int{0, 1}
	need := parse.Atoi(lines[0])

	needDigits := maths.Digits(need)
	var part1 string
	var part2 int
	for part1 == "" || part2 == 0 {
		var sum int
		for _, e := range elves {
			sum += recipes[e]
		}
		ds := maths.Digits(sum)
		recipes = append(recipes, ds...)
		for i, e := range elves {
			elves[i] = (e + recipes[e] + 1) % len(recipes)
		}

		if part2 == 0 {
			for k := 0; k < len(ds) && len(recipes)-k-len(needDigits) >= 0; k++ {
				if slices.Equal(needDigits, recipes[len(recipes)-k-len(needDigits):len(recipes)-k]) {
					part2 = len(recipes) - k - len(needDigits)
				}
			}
		}

		if part1 == "" && len(recipes) >= need+10 {
			parts := functional.Map(recipes[need:need+10], parse.Itos)
			part1 = strings.Join(parts, "")
		}
	}
	o.Stdoutln(part1, part2)
}

func (d *day14) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"5941429882 86764",
			},
		},
		{
			ExpectedOutput: []string{
				"1776718175 20220949",
			},
		},
	}
}
