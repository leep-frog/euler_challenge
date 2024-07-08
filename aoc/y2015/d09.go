package y2015

import (
	"regexp"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
)

func Day09() aoc.Day {
	return &day09{}
}

type day09 struct{}

func (d *day09) rec(distances map[string]map[string]int, city string, remaining map[string]bool, dist int, best *maths.Bester[int, int]) {
	if len(remaining) == 1 {
		best.Check(dist)
		return
	}

	delete(remaining, city)
	for _, toCity := range maps.Keys(remaining) {
		d.rec(distances, toCity, remaining, dist+distances[city][toCity], best)
	}
	remaining[city] = true
}

func (d *day09) Solve(lines []string, o command.Output) {
	var parts []int
	for _, best := range []*maths.Bester[int, int]{maths.Smallest[int, int](), maths.Largest[int, int]()} {
		cities := map[string]map[string]int{}
		remaining := map[string]bool{}
		r := regexp.MustCompile(`^(.*) to (.*) = (.*)$`)
		for _, line := range lines {
			m := r.FindStringSubmatch(line)
			ca, cb, dist := m[1], m[2], parse.Atoi(m[3])
			remaining[ca] = true
			remaining[cb] = true
			maths.Insert(cities, ca, cb, dist)
			maths.Insert(cities, cb, ca, dist)
		}

		for city := range cities {
			d.rec(cities, city, remaining, 0, best)
		}
		parts = append(parts, best.Best())
	}
	o.Stdoutln(parts[0], parts[1])
}

func (d *day09) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"605 982",
			},
		},
		{
			ExpectedOutput: []string{
				"117 909",
			},
		},
	}
}
