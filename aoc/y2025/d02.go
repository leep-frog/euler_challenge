package y2025

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day02() aoc.Day {
	return &day02{}
}

type day02 struct{}

func join(a int, times int) int {
	return parse.Atoi(strings.Repeat(fmt.Sprintf("%d", a), times))
}

func (d *day02) Solve(lines []string, o command.Output) {
	var ranges []*maths.Range
	var largest int
	for _, _range := range strings.Split(lines[0], ",") {
		parts := strings.Split(_range, "-")
		left, right := parse.Atoi(parts[0]), parse.Atoi(parts[1])
		ranges = append(ranges, maths.NewRange(left, right))
		if right > largest {
			largest = right
		}
		if left > largest {
			largest = left
		}
	}

	finalRange := ranges[0]
	for _, r := range ranges[1:] {
		finalRange = finalRange.Merge(r)
	}
	fmt.Println("=======")
	fmt.Println(finalRange)

	alreadyFound := map[int]bool{}
	var sum int
	for joinSize := 2; ; joinSize++ {

		first := true
		for i := 1; ; i++ {
			joined := join(i, joinSize)
			if joined > largest {
				if first {
					fmt.Println(sum)
					return
				}
				break
			}
			first = false

			// Ignore ones we've already checked
			if alreadyFound[joined] {
				continue
			}
			alreadyFound[joined] = true

			// Check if in range
			if finalRange.Contains(joined) {
				sum += joined
			}
		}
	}

	fmt.Println(sum)

}

func (d *day02) Cases() []*aoc.Case {
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
