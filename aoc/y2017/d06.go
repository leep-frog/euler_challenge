package y2017

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day06() aoc.Day {
	return &day06{}
}

type day06 struct{}

func (d *day06) Solve(lines []string, o command.Output) {
	elements := parse.AtoiArray(parse.SplitWhitespace(lines)[0])

	var answers []int
	for i := 0; i < 2; i++ {
		seen := map[string]bool{}

		for code := fmt.Sprintf("%v", elements); !seen[code]; code = fmt.Sprintf("%v", elements) {
			seen[code] = true

			var maxI, max int
			for i, v := range elements {
				if v > max {
					maxI, max = i, v
				}
			}

			div, rem := max/len(elements), max%len(elements)
			elements[maxI] = 0
			for i := 0; i < len(elements); i++ {
				jdx := (maxI + i + 1) % len(elements)
				elements[jdx] += div
				if rem > 0 {
					elements[jdx]++
					rem--
				}
			}
		}
		answers = append(answers, len(seen))
	}
	o.Stdoutln(answers)
}

func (d *day06) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"[5 4]",
			},
		},
		{
			ExpectedOutput: []string{
				"[4074 2793]",
			},
		},
	}
}
