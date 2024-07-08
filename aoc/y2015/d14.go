package y2015

import (
	"fmt"
	"regexp"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
)

func Day14() aoc.Day {
	return &day14{}
}

type reindeer struct {
	name       string
	speed      int
	flyingTime int
	restTime   int
}

func (r *reindeer) String() string {
	return fmt.Sprintf("%s: %d km/s %d s %d", r.name, r.speed, r.flyingTime, r.restTime)
}

func (r *reindeer) dist(seconds int) int {
	iterations := seconds / (r.flyingTime + r.restTime)
	partial := seconds % (r.flyingTime + r.restTime)
	return r.speed * (iterations*r.flyingTime + maths.Min(partial, r.flyingTime))
}

type day14 struct{}

func (d *day14) Solve(lines []string, o command.Output) {
	// Comet
	seconds := parse.Atoi(lines[0])
	var deer []*reindeer
	r := regexp.MustCompile("^(.*) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.$")
	for _, line := range lines[1:] {
		m := r.FindStringSubmatch(line)
		deer = append(deer, &reindeer{
			m[1],
			parse.Atoi(m[2]),
			parse.Atoi(m[3]),
			parse.Atoi(m[4]),
		})
	}

	// Part 1
	best := maths.Largest[string, int]()
	for _, d := range deer {
		best.IndexCheck(d.name, d.dist(seconds))
	}
	part1 := best.Best()

	// Part 2
	points := map[string]int{}
	for i := 1; i <= seconds; i++ {
		bestDeer := []string{}
		var bestDist int
		for _, d := range deer {
			curDist := d.dist(i)
			if curDist == bestDist {
				bestDeer = append(bestDeer, d.name)
			} else if curDist > bestDist {
				bestDeer = []string{d.name}
				bestDist = curDist
			}
		}
		for _, d := range bestDeer {
			points[d]++
		}
	}
	o.Stdoutln(part1, maths.Max(maps.Values(points)...))
}

func (d *day14) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"1120 689",
			},
		},
		{
			ExpectedOutput: []string{
				"2640 1102",
			},
		},
	}
}
