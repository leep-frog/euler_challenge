package y2020

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
)

func Day17() aoc.Day {
	return &day17{}
}

type day17 struct{}

func (d *day17) neighbors(p *point.Point3D, pm map[string]bool) int {
	offsets := combinatorics.GenerateCombos(&combinatorics.Combinatorics[int]{
		Parts:            []int{-1, 0, 1},
		MinLength:        3,
		MaxLength:        3,
		AllowReplacement: true,
		OrderMatters:     true,
	})

	var count int
	for _, o := range offsets {
		if o[0] == 0 && o[1] == 0 && o[2] == 0 {
			continue
		}
		n := point.New3D(p.X+o[0], p.Y+o[1], p.Z+o[2])
		if pm[n.String()] {
			count++
		}
	}
	return count
}

func (d *day17) neighbors4d(p *point.Point4D, pm map[string]bool) int {
	offsets := combinatorics.GenerateCombos(&combinatorics.Combinatorics[int]{
		Parts:            []int{-1, 0, 1},
		MinLength:        4,
		MaxLength:        4,
		AllowReplacement: true,
		OrderMatters:     true,
	})

	var count int
	for _, o := range offsets {
		if o[0] == 0 && o[1] == 0 && o[2] == 0 && o[3] == 0 {
			continue
		}
		n := point.New4D(p.W+o[0], p.X+o[1], p.Y+o[2], p.Z+o[3])
		if pm[n.String()] {
			count++
		}
	}
	return count
}

func (d *day17) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve3D(lines), d.solve4D(lines))
}

func getPoints[T fmt.Stringer](lines []string, f func(int, int) T) ([]T, map[string]bool) {
	var points []T
	pointMap := map[string]bool{}
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				p := f(x, y)
				points = append(points, p)
				pointMap[p.String()] = true
			}
		}
	}
	return points, pointMap
}

func (d *day17) solve3D(lines []string) int {
	points, pointMap := getPoints(lines, func(x, y int) *point.Point3D {
		return point.New3D(x, y, 0)
	})

	for i := 0; i < 6; i++ {
		var newPoints []*point.Point3D
		newPM := map[string]bool{}
		minX, maxX, minY, maxY, minZ, maxZ := points[0].X, points[0].X, points[0].Y, points[0].Y, points[0].Z, points[0].Z
		for _, p := range points {
			minX = maths.Min(minX, p.X)
			maxX = maths.Max(maxX, p.X)

			minY = maths.Min(minY, p.Y)
			maxY = maths.Max(maxY, p.Y)

			minZ = maths.Min(minZ, p.Z)
			maxZ = maths.Max(maxZ, p.Z)
		}

		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for z := minZ - 1; z <= maxZ+1; z++ {
					p := point.New3D(x, y, z)
					ns := d.neighbors(p, pointMap)
					active := pointMap[p.String()]
					if active && (ns == 2 || ns == 3) {
						newPoints = append(newPoints, p)
						newPM[p.String()] = true
					} else if !active && ns == 3 {
						newPoints = append(newPoints, p)
						newPM[p.String()] = true
					}
				}
			}
		}
		points = newPoints
		pointMap = newPM
	}
	return len(points)
}

func (d *day17) solve4D(lines []string) int {
	points, pointMap := getPoints(lines, func(x, y int) *point.Point4D {
		return point.New4D(0, x, y, 0)
	})

	for i := 0; i < 6; i++ {
		var newPoints []*point.Point4D
		newPM := map[string]bool{}
		minW, maxW, minX, maxX, minY, maxY, minZ, maxZ := points[0].W, points[0].W, points[0].X, points[0].X, points[0].Y, points[0].Y, points[0].Z, points[0].Z
		for _, p := range points {
			minW = maths.Min(minW, p.W)
			maxW = maths.Max(maxW, p.W)

			minX = maths.Min(minX, p.X)
			maxX = maths.Max(maxX, p.X)

			minY = maths.Min(minY, p.Y)
			maxY = maths.Max(maxY, p.Y)

			minZ = maths.Min(minZ, p.Z)
			maxZ = maths.Max(maxZ, p.Z)
		}

		for w := minW - 1; w <= maxW+1; w++ {
			for x := minX - 1; x <= maxX+1; x++ {
				for y := minY - 1; y <= maxY+1; y++ {
					for z := minZ - 1; z <= maxZ+1; z++ {
						p := point.New4D(w, x, y, z)
						ns := d.neighbors4d(p, pointMap)
						active := pointMap[p.String()]
						if active && (ns == 2 || ns == 3) {
							newPoints = append(newPoints, p)
							newPM[p.String()] = true
						} else if !active && ns == 3 {
							newPoints = append(newPoints, p)
							newPM[p.String()] = true
						}
					}
				}
			}
		}
		points = newPoints
		pointMap = newPM
	}
	return len(points)
}

func (d *day17) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"112 848",
			},
		},
		{
			ExpectedOutput: []string{
				"346 1632",
			},
		},
	}
}
