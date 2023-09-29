package y2022

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/functional"
)

func Day17() aoc.Day {
	return &day17{}
}

type day17 struct{}

type rock struct {
	parts []*point.Point[int]
}

func printChamber(chamber [][]bool, coord *point.Point[int], rock *rock) string {

	coords := map[string]bool{}
	for _, r := range rock.parts {
		coords[r.Plus(coord).String()] = true
	}

	var c []string
	c = append(c, "************************")
	for y, row := range chamber {
		var r []string
		for x, b := range row {
			if coords[point.New(x, y).String()] {
				r = append(r, "@")
			} else if b {
				r = append(r, "#")
			} else {
				r = append(r, ".")
			}
		}
		c = append(c, strings.Join(r, ""))
	}
	return strings.Join(bread.Reverse(c), "\n")
}

func (r *rock) place(coord *point.Point[int], chamber [][]bool) {
	for _, p := range r.parts {
		to := p.Plus(coord)
		chamber[to.Y][to.X] = true
	}
}

func (r *rock) move(coord, move *point.Point[int], chamber [][]bool) bool {
	for _, p := range r.parts {
		to := p.Plus(move).Plus(coord)
		if to.Y < 0 || to.Y >= len(chamber) || to.X < 0 || to.X >= len(chamber[to.Y]) || chamber[to.Y][to.X] {
			return false
		}
	}
	return true
}

const (
	xStart17 = 2
)

var (
	down = point.New(0, -1)
)

func chamberHeight(chamber [][]bool) int {
	for i := len(chamber) - 1; i >= 0; i-- {
		for _, v := range chamber[i] {
			if v {
				return i + 1
			}
		}
	}
	return 0
}

// Returns new jetIdx
func rec17jet(jetIdx int, rock *rock, coord *point.Point[int], chamber [][]bool, jets []bool) int {
	for ; jetIdx >= len(jets); jetIdx -= len(jets) {
	}

	move := point.New(1, 0)
	if jets[jetIdx] {
		move = point.New(-1, 0)
	}

	// If we can move left or right, then do so.
	if rock.move(coord, move, chamber) {
		coord = coord.Plus(move)
	}

	// If we can move down, then do so
	if rock.move(coord, down, chamber) {
		coord = coord.Plus(down)
	} else {
		// Otherwise, we are done
		// fmt.Println(printChamber(chamber, coord, rock))
		rock.place(coord, chamber)
		return jetIdx + 1
	}

	return rec17jet(jetIdx+1, rock, coord, chamber, jets)
}

func hasPattern(heights []int, patternLength int) (int, int, bool) {
	for pMult := 1; len(heights)-1-pMult*2*patternLength >= 0; pMult++ {
		pattern := true

		start1 := len(heights) - pMult*patternLength
		start2 := len(heights) - pMult*2*patternLength
		for pos1, pos2 := start1, start2; pos1 < len(heights); pos1, pos2 = pos1+1, pos2+1 {
			if heights[pos1]-heights[pos1-1] != heights[pos2]-heights[pos2-1] {
				pattern = false
				break
			}
		}
		if pattern {
			return pMult * patternLength, heights[len(heights)-1] - heights[start1-1], true
		}
	}
	return 0, 0, false
}

func rec17(remRocks int, rocks []*rock, jets []bool) int {
	jetIdx := 0
	var chamber [][]bool
	var heights []int

	minPatternLength := len(rocks) * len(jets)

	for r := 0; r < remRocks; r++ {

		// First, look for a pattern
		// Now look for a pattern
		if r%minPatternLength == 0 {
			patternLength, patternHeight, ok := hasPattern(heights, minPatternLength)
			if ok {
				remaining := remRocks - r

				height := chamberHeight(chamber)

				// Add full pattern heights
				fullPatterns := remaining / patternLength
				height += fullPatterns * patternHeight
				remainderPattern := remaining % patternLength

				// Add height of remainder
				patternStartIdx := len(heights) - 1 - patternLength
				height += heights[patternStartIdx+remainderPattern] - heights[patternStartIdx]

				return height
			}
		}

		rock := rocks[r%len(rocks)]

		// Get starting position
		yStart := 3
		for i := len(chamber) - 1; i >= 0 && yStart == 3; i-- {
			for _, v := range chamber[i] {
				if v {
					yStart = i + 4
				}
			}
		}

		// Increase chamber height
		for len(chamber) < yStart+4 {
			chamber = append(chamber, make([]bool, 7, 7))
		}

		// Now move the rock
		jetIdx = rec17jet(jetIdx, rock, point.New(xStart17, yStart), chamber, jets)

		// rec17(remRocks-1, jetIdx, (rockIdx+1)%len(rocks), chamber, rocks, jets, heights)
		heights = append(heights, chamberHeight(chamber))
	}
	return heights[len(heights)-1]
}

func (d *day17) Solve(lines []string, o command.Output) {
	jets := functional.Map(strings.Split(lines[0], ""), func(s string) bool { return s == "<" })
	// Rocks are going to be falling, but the bottom will be the 0th row
	// .....#. 4
	// .....#. 3
	// ...###. 2
	// ..##... 1
	// ..##... 0
	// ------- floor
	rocks := []*rock{
		// Bottom left point is always considered to be (0,0)
		// ####
		{
			[]*point.Point[int]{
				point.New(0, 0),
				point.New(1, 0),
				point.New(2, 0),
				point.New(3, 0),
			},
		},
		// .#.
		// ### // Possibly don't need to consider the middle of this rock
		// .#.
		{
			[]*point.Point[int]{
				point.New(1, 0),
				point.New(1, 1),
				point.New(1, 2),
				point.New(0, 1),
				point.New(2, 1),
			},
		},
		// ..#
		// ..#
		// ###
		{
			[]*point.Point[int]{
				point.New(0, 0),
				point.New(1, 0),
				point.New(2, 0),
				point.New(2, 1),
				point.New(2, 2),
			},
		},
		// #
		// #
		// #
		// #
		{
			[]*point.Point[int]{
				point.New(0, 0),
				point.New(0, 1),
				point.New(0, 2),
				point.New(0, 3),
			},
		},
		// ##
		// ##
		{
			[]*point.Point[int]{
				point.New(0, 0),
				point.New(1, 0),
				point.New(0, 1),
				point.New(1, 1),
			},
		},
	}

	o.Stdoutln(rec17(2022, rocks, jets), rec17(1_000_000_000_000, rocks, jets))
}

func (d *day17) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"3068 1514285714288",
			},
		},
		{
			ExpectedOutput: []string{
				// Pattern is detected at 34_158_035
				"3055 1507692307690",
			},
		},
	}
}
