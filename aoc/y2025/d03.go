package y2025

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day03() aoc.Day {
	return &day03{}
}

type day03 struct{}

func findMaxJoltage(batteries []int, remainingDigits int) (int, bool) {
	if remainingDigits == 0 {
		return 0, true
	}

	// Get the left most positions of all numbers
	leftMostPositions := map[int]int{}
	for i, b := range batteries {
		if _, ok := leftMostPositions[b]; ok {
			continue
		} else {
			leftMostPositions[b] = i
		}
	}

	// Try from largest digit to smallest
	for i := 9; i >= 0; i-- {
		pos, ok := leftMostPositions[i]
		if !ok {
			continue
		}

		got, solnExists := findMaxJoltage(batteries[pos+1:], remainingDigits-1)
		if solnExists {
			return i*maths.Pow(10, remainingDigits-1) + got, true
		}
	}
	return 0, false
}

func (d *day03) Solve(lines []string, o command.Output) {

	for _, size := range []int{2, 12} {
		var sum int
		for _, batteries := range parse.ToGrid(lines, "") {
			joltage, _ := findMaxJoltage(batteries, size)
			sum += joltage
		}
		o.Stdoutln(size, sum)
	}
}

func (d *day03) Cases() []*aoc.Case {
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
