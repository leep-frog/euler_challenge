package y2022

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
)

func Day23() aoc.Day {
	return &day23{}
}

type day23 struct{}

type cardinalDirections struct {
	move      *point.Point[int]
	positions []*point.Point[int]
}

var (
	// All moves the elves can make
	allMoves = []*point.Point[int]{
		point.New(1, -1),
		point.New(1, 0),
		point.New(1, 1),
		point.New(0, -1),
		point.New(0, 1),
		point.New(-1, -1),
		point.New(-1, 0),
		point.New(-1, 1),
	}
	// Prioritized directions to consider
	cardinalMoves = []*cardinalDirections{
		// North
		{
			point.New(0, -1),
			[]*point.Point[int]{
				point.New(1, -1),
				point.New(0, -1),
				point.New(-1, -1),
			},
		},
		// South
		{
			point.New(0, 1),
			[]*point.Point[int]{
				point.New(1, 1),
				point.New(0, 1),
				point.New(-1, 1),
			},
		},
		// West
		{
			point.New(-1, 0),
			[]*point.Point[int]{
				point.New(-1, 1),
				point.New(-1, 0),
				point.New(-1, -1),
			},
		},
		// East
		{
			point.New(1, 0),
			[]*point.Point[int]{
				point.New(1, 1),
				point.New(1, 0),
				point.New(1, -1),
			},
		},
	}
)

// Get bounds of rectangle
func (d *day23) bounds(elves []*point.Point[int]) (int, int, int, int) {
	minX, maxX, minY, maxY := elves[0].X, elves[0].X, elves[0].Y, elves[0].Y
	for _, e := range elves {
		minX = maths.Min(minX, e.X)
		maxX = maths.Max(maxX, e.X)
		minY = maths.Min(minY, e.Y)
		maxY = maths.Max(maxY, e.Y)
	}
	return minX, maxX, minY, maxY
}

func (d *day23) Solve(lines []string, o command.Output) {
	// Populate elf locations
	var elves []*point.Point[int]
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				elves = append(elves, point.New(x, y))
			}
		}
	}

	var round int
	var part1 int
	for elfMoved := true; elfMoved; round++ {
		elfMoved = false

		// Map of all spots that are currently occupied
		occupied := map[string]bool{}
		for _, e := range elves {
			occupied[e.String()] = true
		}

		// Planned moves by index and by location
		var plannedMoves []*point.Point[int]
		plannedMovesMap := map[string]int{}

		// Plan all moves
		for _, e := range elves {
			// If no elves nearby, then stay still
			if maths.All(allMoves, func(m *point.Point[int]) bool {
				return !occupied[e.Plus(m).String()]
			}) {
				plannedMoves = append(plannedMoves, nil)
				continue
			}

			// Determine which direction to go.
			var nextSpot *point.Point[int]
			for ci := round; ci < len(cardinalMoves)+round && nextSpot == nil; ci++ {
				valid := true
				cm := cardinalMoves[ci%len(cardinalMoves)]
				for _, move := range cm.positions {
					if occupied[e.Plus(move).String()] {
						valid = false
						break
					}
				}
				if valid {
					nextSpot = e.Plus(cm.move)
				}
			}

			// Add the plan
			plannedMoves = append(plannedMoves, nextSpot)
			if nextSpot != nil {
				plannedMovesMap[nextSpot.String()]++
				elfMoved = true
			}
		}

		// Make planned move if no collisions.
		for i, pm := range plannedMoves {
			if pm == nil || plannedMovesMap[pm.String()] > 1 {
				continue
			}
			elves[i] = pm
		}

		// Part 1
		if round == 10 {
			minX, maxX, minY, maxY := d.bounds(elves)
			area := (maxX - minX + 1) * (maxY - minY + 1)
			part1 = area - len(elves)
		}
	}

	o.Stdoutln(part1, round)
}

func (d *day23) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"110 20",
			},
		},
		{
			ExpectedOutput: []string{
				"3931 944",
			},
		},
	}
}

/*func (d *day23) draw(elves []*point.Point[int]) {
	m := map[string]bool{}
	for _, e := range elves {
		m[e.String()] = true
	}

	minX, maxX, minY, maxY := d.bounds(elves)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if m[point.New(x, y).String()] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("========================")
}
*/
