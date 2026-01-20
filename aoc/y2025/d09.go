package y2025

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day09() aoc.Day {
	return &day09{}
}

type day09 struct{}

func isValid(point []int, points [][]int) bool {
	// Shape is basically a diamond. Ensuring that all points have at least one
	// other point in each of the four quadrants around it, then that's a "good enough"
	// validity check that the corner is in the shape.
	topLeft, topRight, bottomLeft, bottomRight := false, false, false, false
	for _, otherPoint := range points {

		if otherPoint[0] <= point[0] && otherPoint[1] <= point[1] {
			topLeft = true
		}
		if otherPoint[0] > point[0] && otherPoint[1] <= point[1] {
			topRight = true
		}
		if otherPoint[0] <= point[0] && otherPoint[1] >= point[1] {
			bottomLeft = true
		}
		if otherPoint[0] >= point[0] && otherPoint[1] >= point[1] {
			bottomRight = true
		}
	}

	return topLeft && topRight && bottomLeft && bottomRight
}

func (d *day09) Solve(lines []string, o command.Output) {
	points := [][]int{}
	var maxX, maxY int
	for _, line := range lines {
		parts := strings.Split(line, ",")
		point := []int{parse.Atoi(parts[0]), parse.Atoi(parts[1])}
		points = append(points, point)
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[1] > maxY {
			maxY = point[1]
		}
	}

	// Check the rectangle made by each unique pair of points
	partOneBest := maths.Largest[int, int]()
	partTwoBest := maths.Largest[int, int]()
	for i, point := range points {
		for _, otherPoint := range points[i+1:] {

			// Shape ultimately is a diamond with a big line through the middle.
			// This check ensures that both points are on the same side of the line.
			leftSide := point[1] <= 48378
			otherLeftSide := otherPoint[1] <= 48378
			validSide := leftSide == otherLeftSide

			// Get all corners of the triangle
			corners := [][]int{
				point,
				otherPoint,
				{point[0], otherPoint[1]},
				{otherPoint[0], point[1]},
			}

			// Check all corners are valid
			validRectangle := true
			for _, corner := range corners {
				if !isValid(corner, points) {
					validRectangle = false
					break
				}
			}

			x := maths.Abs(point[0]-otherPoint[0]) + 1
			y := maths.Abs(point[1]-otherPoint[1]) + 1
			partOneBest.Check(x * y)
			if validSide && validRectangle {
				partTwoBest.Check(x * y)
			}
		}
	}

	o.Stdoutln(partOneBest.Best(), partTwoBest.Best())
}

func (d *day09) Cases() []*aoc.Case {
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
