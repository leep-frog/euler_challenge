package y2016

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day12() aoc.Day {
	return &day12{}
}

type day12 struct{}

func (d *day12) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, false), d.solve(lines, true))
}

func (d *day12) solve(lines []string, part2 bool) int {
	registers := map[string]int{}
	if part2 {
		registers["c"] = 1
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		parts := strings.Split(line, " ")
		switch parts[0] {
		case "cpy":
			v, ok := parse.AtoiOK(parts[1])
			if !ok {
				v = registers[parts[1]]
			}
			registers[parts[2]] = v
		case "inc":
			registers[parts[1]]++
		case "dec":
			registers[parts[1]]--
		case "jnz":
			x, ok := parse.AtoiOK(parts[1])
			if !ok {
				x = registers[parts[1]]
			}
			y, ok := parse.AtoiOK(parts[2])
			if !ok {
				y = registers[parts[2]]
			}
			if x != 0 {
				// -1 since i++
				i += y - 1
			}
		}
	}
	return registers["a"]
}

func (d *day12) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"42 42",
			},
		},
		{
			ExpectedOutput: []string{
				"318007 9227661",
			},
		},
	}
}
