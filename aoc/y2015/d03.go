package y2015

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/point"
)

func Day03() aoc.Day {
	return &day03{}
}

type day03 struct{}

func (d *day03) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, 1), d.solve(lines, 2))
}

func (d *day03) solve(lines []string, numSantas int) int {
	var santas []*point.Point[int]
	for i := 0; i < numSantas; i++ {
		santas = append(santas, point.Origin[int]())
	}
	houses := map[string]bool{}
	houses[santas[0].String()] = true
	for i, c := range lines[0] {
		santa := santas[i%numSantas]
		switch c {
		case '<':
			santa.X--
		case '>':
			santa.X++
		case 'v':
			santa.Y--
		case '^':
			santa.Y++
		}
		houses[santa.String()] = true
	}
	return len(houses)
}

func (d *day03) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"4 3",
			},
		},
		{
			ExpectedOutput: []string{
				"2572 2631",
			},
		},
	}
}
