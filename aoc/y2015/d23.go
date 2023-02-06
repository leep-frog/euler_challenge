package y2015

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day23() aoc.Day {
	return &day23{}
}

type day23 struct{}

func (d *day23) Solve(lines []string, o command.Output) {
	var parts []string
	for _, aStart := range []int{0, 1} {
		register := map[string]int{
			"a": aStart,
		}
		for i := 0; i < len(lines); i++ {
			parts := strings.Split(strings.ReplaceAll(lines[i], ",", ""), " ")
			switch parts[0] {
			case "hlf":
				register[parts[1]] /= 2
			case "tpl":
				register[parts[1]] *= 3
			case "inc":
				register[parts[1]]++
			case "jmp":
				i += (parse.Atoi(strings.ReplaceAll(parts[1], "+", "")) - 1)
			case "jio":
				if register[parts[1]] == 1 {
					i += (parse.Atoi(strings.ReplaceAll(parts[2], "+", "")) - 1)
				}
			case "jie":
				if register[parts[1]]%2 == 0 {
					i += (parse.Atoi(strings.ReplaceAll(parts[2], "+", "")) - 1)
				}
			default:
				panic(fmt.Sprintf("Unknown command: %v", parts))
			}
		}
		parts = append(parts, parse.Itos(register["b"]))
	}
	o.Stdoutln(strings.Join(parts, " "))
}

func (d *day23) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"4 2",
			},
		},
		{
			ExpectedOutput: []string{
				"184 231",
			},
		},
	}
}
