package y2020

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/slices"
)

func Day10() aoc.Day {
	return &day10{}
}

type day10 struct{}

func (d *day10) numPaths(numbers []int) int {
	numbers = bread.Reverse(numbers)
	sums := make([]int, len(numbers), len(numbers))
	sums[0] = 1
	for i := 1; i < len(numbers); i++ {
		v := numbers[i]
		for j := i - 1; j >= 0 && v <= numbers[j] && numbers[j] <= v+3; j-- {
			sums[i] += sums[j]
		}
	}
	return sums[len(sums)-1]
}

func (d *day10) Solve(lines []string, o command.Output) {
	numbers := parse.AtoiArray(lines)
	numbers = append(numbers, 0, maths.Max(numbers...)+3)
	slices.Sort(numbers)
	diffs := make([]int, 4, 4)
	for i, k := range numbers[1:] {
		diffs[k-numbers[i]]++
	}
	o.Stdoutln(diffs[1]*diffs[3], d.numPaths(numbers))
}

func (d *day10) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"220 19208",
			},
		},
		{
			ExpectedOutput: []string{
				"2482 96717311574016",
			},
		},
	}
}
