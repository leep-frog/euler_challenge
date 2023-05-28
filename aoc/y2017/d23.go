package y2017

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day23() aoc.Day {
	return &day23{}
}

type day23 struct{}

type cmd struct {
}

func (d *day23) Solve(lines []string, o command.Output) {
	registers := map[string]int{}

	numOrReg := func(k string) int {
		if v, ok := parse.AtoiOK(k); ok {
			return v
		}
		return registers[k]
	}

	registers["a"] = 1

	mc := 0
	partsArr := parse.SplitWhitespace(lines)
	for i := 0; i < len(partsArr); i++ {
		parts := partsArr[i]
		a, b := numOrReg(parts[1]), numOrReg(parts[2])
		switch parts[0] {
		case "set":
			registers[parts[1]] = b
		case "mul":
			registers[parts[1]] = a * b
			mc++
		case "sub":
			registers[parts[1]] = a - b
		case "jnz":
			if a != 0 {
				i = i + b - 1
			}
		}
	}
	o.Stdoutln(mc)
}

func (d *day23) Cases() []*aoc.Case {
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
