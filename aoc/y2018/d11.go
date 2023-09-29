package y2018

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
)

func Day11() aoc.Day {
	return &day11{}
}

type day11 struct{}

func (d *day11) power(x, y, serial int) int {
	if x < 1 || y < 1 || x > 300 || y > 300 {
		fmt.Println("INVALID X,Y:", x, y)
		panic("NOPE")
	}
	rack := x + 10
	power := (rack*y + serial) * rack
	return ((power / 100) % 10) - 5
}

func (d *day11) Solve(lines []string, o command.Output) {
	// serial := 18
	serial := 5468

	best1 := maths.Largest[string, int]()
	best2 := maths.Largest[string, int]()
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			var sum int

			// Iterate over size squares
			for size := 0; size <= maths.Min(300-x, 300-y); size++ {
				// Just need to add the perimeter of the square incrementally

				// Add the right side
				for xx, yd := x+size, 0; yd <= size; yd++ {
					sum += d.power(xx, y+yd, serial)
				}

				// Add the bottom (but not the end)
				for yy, xd := y+size, 0; xd < size; xd++ {
					sum += d.power(x+xd, yy, serial)
				}

				// Check sums
				if size == 2 {
					best1.IndexCheck(fmt.Sprintf("%d,%d", x, y), sum)
				}
				best2.IndexCheck(fmt.Sprintf("%d,%d,%d", x, y, size+1), sum)
			}
		}
	}

	o.Stdoutln(best1.BestIndex(), best2.BestIndex())
}

func (d *day11) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"243,64 90,101,15",
			},
		},
		{
			ExpectedOutput: []string{
				"243,64 90,101,15",
			},
		},
	}
}
