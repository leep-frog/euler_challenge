package y2018

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/maps"
)

func Day06() aoc.Day {
	return &day06{}
}

type day06 struct{}

func (d *day06) Solve(lines []string, o command.Output) {
	var ps []*point.Point[int]
	for _, parts := range parse.Split(lines, ", ") {
		ps = append(ps, point.New(parse.Atoi(parts[0]), parse.Atoi(parts[1])))
	}

	ch := point.ConvexHullFromPoints(ps...)
	inCH := map[string]bool{}
	for _, p := range ch.Points {
		inCH[p.String()] = true
	}

	minX, maxX, minY, maxY := maths.Smallest[int, int](), maths.Largest[int, int](), maths.Smallest[int, int](), maths.Largest[int, int]()
	for _, p := range ps {
		minX.Check(p.X)
		maxX.Check(p.X)
		minY.Check(p.Y)
		maxY.Check(p.Y)
	}

	counts := map[string]int{}
	for x := minX.Best(); x <= maxX.Best(); x++ {
		for y := minY.Best(); y <= maxY.Best(); y++ {
			c := point.New(x, y)

			dists := map[string]int{}
			for _, p := range ps {
				dists[p.String()] = p.ManhattanDistance(c)
			}

			best := maths.Min(maps.Values(dists)...)
			var bestPoint string
			for p, v := range dists {
				if v == best {
					if bestPoint == "" {
						bestPoint = p
					} else {
						// already set
						bestPoint = ""
						break
					}
				}
			}

			if bestPoint != "" {
				counts[bestPoint]++
			}
		}
	}

	bestPoint := maths.Largest[string, int]()
	for p, c := range counts {
		if inCH[p] {
			continue
		}
		bestPoint.IndexCheck(p, c)
	}

	// Part 2
	var regionCnt int
	for x := minX.Best(); x <= maxX.Best(); x++ {
		for y := minY.Best(); y <= maxY.Best(); y++ {
			c := point.New(x, y)
			var sum int
			for _, p := range ps {
				sum += c.ManhattanDistance(p)
			}

			if sum < 10000 {
				regionCnt++
			}

		}
	}
	o.Stdoutln(bestPoint.Best(), regionCnt)
}

func (d *day06) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"17 72",
			},
		},
		{
			ExpectedOutput: []string{
				"4771 39149",
			},
		},
	}
}
