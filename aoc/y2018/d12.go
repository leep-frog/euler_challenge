package y2018

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
)

func Day12() aoc.Day {
	return &day12{}
}

type day12 struct{}

type plantRule struct {
	input  []bool
	output bool
}

func (d *day12) Solve(lines []string, o command.Output) {
	state := map[int]bool{}
	init := strings.Split(strings.TrimPrefix(lines[0], "initial state: "), "")
	for i, c := range init {
		if c == "#" {
			state[i] = true
		}
	}

	var rules []*plantRule
	rm := map[string]bool{}
	for _, parts := range parse.Split(lines[2:], " => ") {
		r := &plantRule{}
		for _, c := range parts[0] {
			r.input = append(r.input, c == '#')
		}
		r.output = parts[1] == "#"
		rules = append(rules, r)
		rm[parts[0]] = parts[1] == "#"
	}

	left, right := -2, len(init)+1

	values := []int{
		bread.Sum(maps.Keys(state)),
	}

	for {
		min, max := maths.Smallest[int, int](), maths.Largest[int, int]()
		newState := map[int]bool{}
		for i := left; i < right; i++ {
			var s []string
			for j := i - 2; j <= i+2; j++ {
				if state[j] {
					s = append(s, "#")
				} else {
					s = append(s, ".")
				}
			}
			if rm[strings.Join(s, "")] {
				newState[i] = true
				min.Check(i)
				max.Check(i)
			}
		}

		state = newState
		left, right = min.Best()-2, max.Best()+2
		values = append(values, bread.Sum(maps.Keys(state)))

		idx := len(values) - 1

		for dist := 50; idx-3*dist >= 0; dist++ {
			diff1 := values[idx] - values[idx-dist]
			diff2 := values[idx-dist] - values[idx-2*dist]
			diff3 := values[idx-2*dist] - values[idx-3*dist]
			if diff1 == diff2 && diff2 == diff3 {
				x1, x2 := idx, idx-dist
				y1, y2 := values[idx], values[idx-dist]
				// y = mx + b
				// y1 = m*x1 + b
				// y2 = m*x2 + b
				// y1 - m*x1 = b = y2 - m*x2
				// y1 - y2 = m*(x1 - x2)
				// m = (y1 -y2)/(x1 - x2)
				m := fraction.New(y1-y2, x1-x2)
				// y1 = m*x1 + b
				// y1 - m*x1 = b
				b := fraction.New(y1, 1).Minus(m.Times(fraction.New(x1, 1)))
				wantX := fraction.New(50_000_000_000, 1)
				wantY := m.Times(wantX).Plus(b)
				o.Stdoutln(values[20], wantY.Simplify(generator.Primes()))
				return
			}
		}
	}
}

func (d *day12) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"325 999999999374/1",
			},
		},
		{
			ExpectedOutput: []string{
				"1623 1600000000401/1",
			},
		},
	}
}
