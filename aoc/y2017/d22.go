package y2017

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/euler_challenge/walker"
)

func Day22() aoc.Day {
	return &day22{}
}

type day22 struct{}

type virusState int

const (
	clean = iota
	weakened
	infected
	flagged
	numStates
)

func (d *day22) Solve(lines []string, o command.Output) {

	part1 := func(w *walker.Walker[walker.CardinalDirection], code string, states map[string]virusState) bool {
		if states[code] == infected {
			w.Right()
			delete(states, code)
			return false
		}

		w.Left()
		states[code] = infected
		return true
	}

	part2 := func(w *walker.Walker[walker.CardinalDirection], code string, states map[string]virusState) bool {
		state := states[code]
		switch state {
		case clean:
			w.Left()
		case weakened:
		case infected:
			w.Right()
		case flagged:
			w.Left()
			w.Left()
		}

		newState := (state + 1) % numStates
		states[code] = newState
		return newState == infected
	}

	o.Stdoutln(d.solve(lines, 10_000, part1), d.solve(lines, 10_000_000, part2))
}

func (d *day22) solve(lines []string, times int, f func(w *walker.Walker[walker.CardinalDirection], code string, states map[string]virusState) bool) int {
	grid := parse.AOCGrid(lines, false, true)
	states := map[string]virusState{}
	for y, row := range grid {
		for x, c := range row {
			if c {
				states[point.New(x, y).String()] = infected
			}
		}
	}
	w := walker.CardinalWalker(walker.Up, true)
	w.MoveTo(point.New(len(grid)/2, len(grid)/2))

	var infCount int

	for i := 0; i < times; i++ {
		code := w.Position().String()
		if f(w, code, states) {
			infCount++
		}
		w.Walk(1)
	}

	return infCount
}

func (d *day22) Cases() []*aoc.Case {
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
