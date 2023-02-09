package y2022

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
)

func Day14() aoc.Day {
	return &day14{}
}

type day14 struct{}

type cave struct {
	cave      [][]caveCell
	xOffset   int
	sandCount int
}

type caveCell int

const (
	emptyCell caveCell = iota
	rockCell
	sandCell
)

func (c *cave) String() string {
	var r []string
	for _, row := range c.cave {
		var codes []string
		for _, b := range row {
			switch b {
			case emptyCell:
				codes = append(codes, ".")
			case sandCell:
				codes = append(codes, "o")
			case rockCell:
				codes = append(codes, "#")
			}
		}
		r = append(r, strings.Join(codes, ""))
	}
	return strings.Join(r, "\n")
}

func (c *cave) add(p *point.Point[int], v caveCell) {
	c.cave[p.Y][p.X-c.xOffset] = v
}

func (c *cave) at(p *point.Point[int]) caveCell {
	return c.cave[p.Y][p.X-c.xOffset]
}

func (c *cave) addSand() bool {
	moves := []*point.Point[int]{
		point.New(0, 1),
		point.New(-1, 1),
		point.New(1, 1),
	}
	sand := point.New(500, 0)
	if c.at(sand) != emptyCell {
		return false
	}
	for {
		for _, move := range moves {
			next := sand.Plus(move)

			if next.X < c.xOffset || next.Y >= len(c.cave) {
				return false
			}
			// fmt.Println("NEXT", next)
			if c.at(next) == emptyCell {
				sand = next
				goto RELOOP
			}
		}
		// Sand can't move any further
		c.sandCount++
		c.add(sand, sandCell)
		return true
	RELOOP:
	}
}

func (d *day14) solve(lines []string, part1 bool) int {
	rockWalls := functional.Map(lines, func(line string) []*point.Point[int] {
		return functional.Map(strings.Split(line, " -> "), func(s string) *point.Point[int] {
			parts := strings.Split(s, ",")
			return point.New(parse.Atoi(parts[0]), parse.Atoi(parts[1]))
		})
	})

	minX := rockWalls[0][0].X
	maxX := minX
	maxY := rockWalls[0][0].Y
	for _, wall := range rockWalls {
		for _, rock := range wall {
			minX = maths.Min(minX, rock.X)
			maxX = maths.Max(maxX, rock.X)
			maxY = maths.Max(maxY, rock.Y)
		}
	}

	if !part1 {
		maxY += 2
		minX = maths.Min(minX, 500-maxY)
		maxX = maths.Max(maxX, 500+maxY)
		rockWalls = append(rockWalls, []*point.Point[int]{
			point.New(minX, maxY),
			point.New(maxX, maxY),
		})
	}

	var grid [][]caveCell
	for i := 0; i <= maxY; i++ {
		grid = append(grid, make([]caveCell, maxX-minX+1, maxX-minX+1))
	}
	c := &cave{grid, minX, 0}

	for _, wall := range rockWalls {
		for i := 0; i < len(wall)-1; i++ {
			from, to := wall[i], wall[i+1]
			var vector *point.Point[int]
			if from.X != to.X {
				if from.X < to.X {
					vector = point.New(1, 0)
				} else {
					vector = point.New(-1, 0)
				}
			} else {
				if from.Y < to.Y {
					vector = point.New(0, 1)
				} else {
					vector = point.New(0, -1)
				}
			}
			for p := from; !p.Eq(to.Plus(vector)); p = p.Plus(vector) {
				c.add(p, rockCell)
			}
		}
	}

	for c.addSand() {
	}

	return c.sandCount
}

func (d *day14) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, true), d.solve(lines, false))

}

func (d *day14) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"24 93",
			},
		},
		{
			ExpectedOutput: []string{
				"862 28744",
			},
		},
	}
}
