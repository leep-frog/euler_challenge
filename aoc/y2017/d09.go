package y2017

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day09() aoc.Day {
	return &day09{}
}

type day09 struct{}

func (d *day09) Solve(lines []string, o command.Output) {
	puzzle := lines[0]
	var garbage bool
	var depth, sum, garbageCount int
	for i := 0; i < len(puzzle); i++ {
		if garbage {
			switch puzzle[i] {
			case '!':
				i++
			case '>':
				garbage = false
			default:
				garbageCount++
			}
		} else {
			switch puzzle[i] {
			case '{':
				depth++
			case '}':
				sum += depth
				depth--
			case '<':
				garbage = true
			}
		}
	}
	o.Stdoutln(sum, garbageCount)
}

func (d *day09) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3 17",
			},
		},
		{
			ExpectedOutput: []string{
				"11347 5404",
			},
		},
	}
}
