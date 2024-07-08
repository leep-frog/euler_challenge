package y2017

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
)

func Day08() aoc.Day {
	return &day08{}
}

type day08 struct{}

func (d *day08) Solve(lines []string, o command.Output) {
	registers := map[string]int{}
	max := maths.Largest[int, int]()
	for _, parts := range parse.SplitWhitespace(lines) {
		left, right := registers[parts[4]], parse.Atoi(parts[6])
		var yes bool
		switch parts[5] {
		case "==":
			yes = left == right
		case "!=":
			yes = left != right
		case ">":
			yes = left > right
		case "<":
			yes = left < right
		case ">=":
			yes = left >= right
		case "<=":
			yes = left <= right
		default:
			panic(fmt.Sprintf("Unknown op: %q", parts[5]))
		}
		if yes {
			switch parts[1] {
			case "inc":
				registers[parts[0]] += parse.Atoi(parts[2])
			case "dec":
				registers[parts[0]] -= parse.Atoi(parts[2])
			}
			max.Check(registers[parts[0]])
		}
	}
	o.Stdoutln(maths.Max(maps.Values(registers)...), max.Best())
}

func (d *day08) Cases() []*aoc.Case {
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
