package y2025

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/euler_challenge/unionfind"
	"github.com/leep-frog/functional"
)

func Day08() aoc.Day {
	return &day08{}
}

type day08 struct{}

type pointDistance struct {
	A, B *point.Point3D
	Dist float64
}

func (d *day08) Solve(lines []string, o command.Output) {
	var points []*point.Point3D
	for _, line := range lines {
		parts := strings.Split(line, ",")
		points = append(points, &point.Point3D{
			parse.Atoi(parts[0]),
			parse.Atoi(parts[1]),
			parse.Atoi(parts[2]),
		})
	}

	isExample := len(points) <= 20

	var pointDistances []*pointDistance
	for i, point := range points {
		for _, otherPoint := range points[i+1:] {
			dist := point.Distance(otherPoint)
			pointDistances = append(pointDistances, &pointDistance{
				A:    point,
				B:    otherPoint,
				Dist: dist,
			})
		}
	}

	functional.SortFunc[*pointDistance](pointDistances, func(a, b *pointDistance) bool {
		return a.Dist < b.Dist
	})

	var partOne, partTwo int

	times := 1000
	if isExample {
		times = 10
	}

	uf := unionfind.New[string]()
	for idx, pd := range pointDistances {
		if idx == times {
			var sizes []int
			for _, set := range uf.Sets() {
				sizes = append(sizes, len(set))
			}
			functional.SortFunc[int](sizes, func(a, b int) bool {
				return a > b
			})
			partOne = bread.Product(sizes[:3])
		}
		uf.Merge(pd.A.String(), pd.B.String())
		if uf.LargestSetSize() >= len(points) {
			partTwo = pd.A.X * pd.B.X
			break
		}
	}

	o.Stdoutln(partOne, partTwo)
}

func (d *day08) Cases() []*aoc.Case {
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
