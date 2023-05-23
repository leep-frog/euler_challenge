package y2016

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day12() aoc.Day {
	return &day12{}
}

type day12 struct{}

func (d *day12) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, map[string]int{}), d.solve(lines, map[string]int{"c": 1}))
}

func (d *day12) solve(lines []string, registers map[string]int) int {
	numOrReg := func(p string) int {
		v, ok := parse.AtoiOK(p)
		if !ok {
			v = registers[p]
		}
		return v
	}

	allParts := parse.Split(lines, " ")

	for i := 0; i < len(allParts); i++ {
		parts := allParts[i]
		// fmt.Println(registers)
		switch parts[0] {
		case "cpy":
			if _, ok := parse.AtoiOK(parts[2]); !ok {
				registers[parts[2]] = numOrReg(parts[1])
			}
		case "inc":
			if _, ok := parse.AtoiOK(parts[1]); !ok {
				registers[parts[1]]++
			}
		case "dec":
			if _, ok := parse.AtoiOK(parts[1]); !ok {
				registers[parts[1]]--
			}
		case "tgl":
			v := i + numOrReg(parts[1])
			if v < 0 || v >= len(lines) {
				continue
			}

			oParts := allParts[v]
			if len(oParts) == 2 {
				if oParts[0] == "inc" {
					oParts[0] = "dec"
				} else {
					oParts[0] = "inc"
				}
			} else {
				if oParts[0] == "jnz" {
					oParts[0] = "cpy"
				} else {
					oParts[0] = "jnz"
				}
			}
		case "jnz":
			x := numOrReg(parts[1])
			y := numOrReg(parts[2])
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
