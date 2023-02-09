package y2020

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day08() aoc.Day {
	return &day08{}
}

type day08 struct{}

type instruction struct {
	code  string
	value int
}

func (d *day08) solve1(instructions []*instruction) (int, bool) {
	run := make([]bool, len(instructions), len(instructions))
	acc := 0

	i := 0
	for i < len(instructions) && !run[i] {
		run[i] = true
		instr := instructions[i]
		switch instr.code {
		case "nop":
			i++
		case "acc":
			acc += instr.value
			i++
		case "jmp":
			i += instr.value
		}
	}
	return acc, i == len(instructions)
}

func (d *day08) solve2(instructions []*instruction) int {
	for _, instr := range instructions {
		switch instr.code {
		case "nop":
			instr.code = "jmp"
			v, ok := d.solve1(instructions)
			instr.code = "nop"
			if ok {
				return v
			}
		case "jmp":
			instr.code = "nop"
			v, ok := d.solve1(instructions)
			instr.code = "jmp"
			if ok {
				return v
			}
		}
	}
	return -1
}

func (d *day08) Solve(lines []string, o command.Output) {
	var run []bool
	var instructions []*instruction
	for _, line := range lines {
		run = append(run, false)
		instructions = append(instructions, &instruction{
			line[:3],
			parse.Atoi(strings.Replace(line[4:], "+", "", -1)),
		})
	}

	v, _ := d.solve1(instructions)
	o.Stdoutln(v, d.solve2(instructions))
}

func (d *day08) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"5 8",
			},
		},
		{
			ExpectedOutput: []string{
				"1928 1319",
			},
		},
	}
}
