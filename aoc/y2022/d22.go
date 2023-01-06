package y2022

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/point"
)

func Day22() aoc.Day {
	return &day22{}
}

type day22 struct{}

type direction int

const (
	rightTurn = -2
	leftTurn  = -1
)

const (
	gridRight direction = iota
	gridDown
	gridLeft
	gridUp
	directionCount
)

type gridNode struct {
	pos *point.Point[int]

	wall bool

	left, right, up, down *gridNode
}

func (gn *gridNode) String() string {
	return gn.pos.String()
}

func (gn *gridNode) connectHorizontally(that *gridNode) {
	gn.right = that
	that.left = gn
}

func (gn *gridNode) connectVertically(that *gridNode) {
	gn.down = that
	that.up = gn
}

func (d *day22) travel(gn *gridNode, path []int) int {
	direction := gridRight

	for _, p := range path {
		if p == rightTurn {
			direction = (direction + 1) % directionCount
		} else if p == leftTurn {
			direction = (direction + directionCount - 1) % directionCount
		} else {
			for i := 0; i < p; i++ {
				var next *gridNode
				switch direction {
				case gridUp:
					next = gn.up
				case gridLeft:
					next = gn.left
				case gridDown:
					next = gn.down
				case gridRight:
					next = gn.right
				}

				if next.wall {
					continue
				}
				// Update direction if changed side we're on
				if gn.pos.Eq(next.left.pos) {
					direction = gridRight
				} else if gn.pos.Eq(next.right.pos) {
					direction = gridLeft
				} else if gn.pos.Eq(next.down.pos) {
					direction = gridUp
				} else if gn.pos.Eq(next.up.pos) {
					direction = gridDown
				} else if !gn.pos.Eq(next.pos) {
					panic("AHAH")
				}
				if !next.wall {
					gn = next
				}
			}
		}
	}

	// Solution is 1-indexed
	return 1000*(gn.pos.Y+1) + 4*(gn.pos.X+1) + int(direction)
}

type squareSide struct {
	cells [][]*gridNode
}

func (ss *squareSide) top() []*gridNode {
	return ss.cells[0]
}

func (ss *squareSide) bottom() []*gridNode {
	return ss.cells[len(ss.cells)-1]
}

func (ss *squareSide) left() []*gridNode {
	var r []*gridNode
	for i := 0; i < len(ss.cells); i++ {
		r = append(r, ss.cells[i][0])
	}
	return r
}

func (ss *squareSide) right() []*gridNode {
	var r []*gridNode
	for i := 0; i < len(ss.cells); i++ {
		r = append(r, ss.cells[i][len(ss.cells)-1])
	}
	return r
}

func (ss *squareSide) getSide(d direction) ([]*gridNode, func(a, b *gridNode)) {
	switch d {
	case gridUp:
		return ss.top(), func(a, b *gridNode) {
			if a.up != nil {
				panic("NOPE")
			}
			a.up = b
		}
	case gridDown:
		return ss.bottom(), func(a, b *gridNode) {
			if a.down != nil {
				panic("NOPE")
			}
			a.down = b
		}
	case gridRight:
		return ss.right(), func(a, b *gridNode) {
			if a.right != nil {
				panic("NOPE")
			}
			a.right = b
		}
	case gridLeft:
		return ss.left(), func(a, b *gridNode) {
			if a.left != nil {
				panic("NOPE")
			}
			a.left = b
		}
	}
	panic("Unknown side")
}

func (ss *squareSide) stitch(that *squareSide, d, thatD direction, reverse bool) {
	side1, f1 := ss.getSide(d)
	side2, f2 := that.getSide(thatD)

	for i, s1 := range side1 {
		s2 := side2[i]
		if reverse {
			s2 = side2[len(side2)-1-i]
		}
		f1(s1, s2)
		f2(s2, s1)
	}
}

type stitch struct {
	a, b    *point.Point[int]
	aD, bD  direction
	reverse bool
}

// This stitch was used to test stitching logic with example part 1
func (d *day22) exampleStitchPart1(squareLength int, squareSides [][]*squareSide) {
	// Should be 12
	stitches := []*stitch{
		// Top square
		{
			point.New(0, 2), point.New(1, 2),
			gridDown, gridUp,
			false,
		},
		{
			point.New(0, 2), point.New(2, 2),
			gridUp, gridDown,
			false,
		},
		{
			point.New(0, 2), point.New(0, 2),
			gridLeft, gridRight,
			false,
		},
		// Right middle
		{
			point.New(1, 2), point.New(2, 2),
			gridDown, gridUp,
			false,
		},
		{
			point.New(1, 2), point.New(1, 0),
			gridRight, gridLeft,
			false,
		},
		{
			point.New(1, 2), point.New(1, 1),
			gridLeft, gridRight,
			false,
		},
		// Midle middle
		{
			point.New(1, 1), point.New(1, 0),
			gridLeft, gridRight,
			false,
		},
		{
			point.New(1, 1), point.New(1, 1),
			gridUp, gridDown,
			false,
		},
		// Left middle
		{
			point.New(1, 0), point.New(1, 0),
			gridUp, gridDown,
			false,
		},
		// Bottom Left
		{
			point.New(2, 2), point.New(2, 3),
			gridRight, gridLeft,
			false,
		},
		{
			point.New(2, 2), point.New(2, 3),
			gridLeft, gridRight,
			false,
		},
		// Bottom right
		{
			point.New(2, 3), point.New(2, 3),
			gridUp, gridDown,
			false,
		},
	}

	for _, stitch := range stitches {
		squareSides[stitch.a.X][stitch.a.Y].stitch(squareSides[stitch.b.X][stitch.b.Y], stitch.aD, stitch.bD, stitch.reverse)
	}
}

func (d *day22) exampleStitchPart2(squareLength int, squareSides [][]*squareSide) {
	// Should be 12
	stitches := []*stitch{
		// Vanilla moves
		// Top square
		{
			point.New(0, 2), point.New(1, 2),
			gridDown, gridUp,
			false,
		},
		// Right middle
		{
			point.New(1, 2), point.New(2, 2),
			gridDown, gridUp,
			false,
		},
		{
			point.New(1, 2), point.New(1, 1),
			gridLeft, gridRight,
			false,
		},
		// Midle middle
		{
			point.New(1, 1), point.New(1, 0),
			gridLeft, gridRight,
			false,
		},
		// Left middle
		// Bottom Left
		{
			point.New(2, 2), point.New(2, 3),
			gridRight, gridLeft,
			false,
		},
		// Bottom right

		// Chocolate connections
		// Top square
		{
			point.New(0, 2), point.New(1, 1),
			gridLeft, gridUp,
			false,
		},
		{
			point.New(0, 2), point.New(1, 0),
			gridUp, gridUp,
			true,
		},
		{
			point.New(0, 2), point.New(2, 3),
			gridRight, gridRight,
			true,
		},
		// Middle right
		{
			point.New(1, 2), point.New(2, 3),
			gridRight, gridUp,
			true,
		},
		// Middle middle
		{
			point.New(1, 1), point.New(2, 2),
			gridDown, gridLeft,
			true,
		},
		// Middle left
		{
			point.New(1, 0), point.New(2, 2),
			gridDown, gridDown,
			true,
		},
		{
			point.New(1, 0), point.New(2, 3),
			gridLeft, gridDown,
			true,
		},
	}

	for _, stitch := range stitches {
		squareSides[stitch.a.X][stitch.a.Y].stitch(squareSides[stitch.b.X][stitch.b.Y], stitch.aD, stitch.bD, stitch.reverse)
	}
}

func (d *day22) stitchPart2(squareLength int, squareSides [][]*squareSide) {
	//   012
	// 0 .XX 0
	// 1 .X. 1
	// 2 XX. 2
	// 3 X.. 3
	//   012
	// Should be 12
	stitches := []*stitch{
		// Vanilla stitches
		{
			point.New(0, 1), point.New(0, 2),
			gridRight, gridLeft,
			false,
		},
		{
			point.New(0, 1), point.New(1, 1),
			gridDown, gridUp,
			false,
		},
		{
			point.New(1, 1), point.New(2, 1),
			gridDown, gridUp,
			false,
		},
		{
			point.New(2, 0), point.New(2, 1),
			gridRight, gridLeft,
			false,
		},
		{
			point.New(2, 0), point.New(3, 0),
			gridDown, gridUp,
			false,
		},
		// Chocolate stitches
		{
			point.New(0, 2), point.New(1, 1),
			gridDown, gridRight,
			false,
		},
		{
			point.New(0, 2), point.New(2, 1),
			gridRight, gridRight,
			true,
		},
		{
			point.New(0, 2), point.New(3, 0),
			gridUp, gridDown,
			false,
		},
		{
			point.New(0, 1), point.New(2, 0),
			gridLeft, gridLeft,
			true,
		},
		{
			point.New(0, 1), point.New(3, 0),
			gridUp, gridLeft,
			false,
		},
		{
			point.New(1, 1), point.New(2, 0),
			gridLeft, gridUp,
			false,
		},
		{
			point.New(2, 1), point.New(3, 0),
			gridDown, gridRight,
			false,
		},
	}

	for _, stitch := range stitches {
		squareSides[stitch.a.X][stitch.a.Y].stitch(squareSides[stitch.b.X][stitch.b.Y], stitch.aD, stitch.bD, stitch.reverse)
	}
}

func (d *day22) wrapAroundStitch(squareLength int, squareSides [][]*squareSide) {
	// Now stitchpar
	firstCols := make([]*squareSide, squareLength, squareLength)
	for i, row := range squareSides {
		var firstRow *squareSide
		for j, ss := range row {
			if ss == nil {
				continue
			}

			// Connect horizontally
			if firstRow == nil {
				firstRow = ss
			}
			if j+1 < len(row) && row[j+1] != nil {
				ss.stitch(row[j+1], gridRight, gridLeft, false)
			} else {
				ss.stitch(firstRow, gridRight, gridLeft, false)
			}

			// Connect vertically
			if firstCols[j] == nil {
				firstCols[j] = ss
			}
			// fmt.Println(i+1, len(squareSides), j, len(squareSides[i+1]))
			if i+1 < len(squareSides) && j < len(squareSides[i+1]) && squareSides[i+1][j] != nil {
				ss.stitch(squareSides[i+1][j], gridDown, gridUp, false)
			} else {
				ss.stitch(firstCols[j], gridDown, gridUp, false)
			}
		}
	}
}

func (d *day22) solveIt(lines []string, stitcher func(int, [][]*squareSide)) int {
	// Parse input
	numberPath := strings.Split(strings.Replace(strings.Replace(lines[len(lines)-1], "L", fmt.Sprintf(" %d ", leftTurn), -1), "R", fmt.Sprintf(" %d ", rightTurn), -1), " ")
	path := functional.Map(numberPath, parse.Atoi)
	lines = lines[:len(lines)-2]

	// Calculate square sizes
	var squareLength int
	for _, line := range lines {
		for _, c := range line {
			if c != ' ' {
				squareLength++
			}
		}
	}
	squareLength = maths.Sqrt(squareLength / 6)

	// The top left node
	var startNode *gridNode

	// Create the face of each side
	var squareSides [][]*squareSide
	for i := 0; i < len(lines); i += squareLength {
		var squareRow []*squareSide
		for j := 0; j < len(lines[i]); j += squareLength {
			// Square is empty
			if lines[i][j] == ' ' {
				squareRow = append(squareRow, nil)
				continue
			}

			// Otherwise, it's a full square
			var grid [][]*gridNode
			for y, line := range lines[i : i+squareLength] {
				var row []*gridNode
				for x, c := range line[j : j+squareLength] {
					gn := &gridNode{
						pos:  point.New(j+x, i+y),
						wall: c == '#',
					}
					if startNode == nil {
						startNode = gn
					}
					row = append(row, gn)
				}
				grid = append(grid, row)
			}
			squareRow = append(squareRow, &squareSide{grid})
		}
		squareSides = append(squareSides, squareRow)
	}

	// Connect all the nodes in a single side
	for _, ssRow := range squareSides {
		for _, ss := range ssRow {
			if ss == nil {
				continue
			}
			for i := 1; i < squareLength; i++ {
				for j := 0; j < squareLength; j++ {
					ss.cells[j][i-1].connectHorizontally(ss.cells[j][i])
					ss.cells[i-1][j].connectVertically(ss.cells[i][j])
				}
			}
		}
	}

	// Stitch accordingly
	stitcher(squareLength, squareSides)

	return d.travel(startNode, path)
}

/*
// This solves the problem by wrapping around edges,
// but the part2 solution also can solve part1
func (d *day22) oldSolve(lines []string) int {
	numberPath := strings.Split(strings.Replace(strings.Replace(lines[len(lines)-1], "L", fmt.Sprintf(" %d ", leftTurn), -1), "R", fmt.Sprintf(" %d ", rightTurn), -1), " ")
	path := functional.Map(numberPath, parse.Atoi)
	lines = lines[:len(lines)-2]

	max := maths.Max(functional.Map(lines, func(line string) int { return len(line) })...)

	var start *gridNode

	var cells []*gridNode

	colStarts := make([]*gridNode, max, max)
	prevColCells := make([]*gridNode, max, max)
	for i, line := range lines {
		var rowStart *gridNode
		var prevCell *gridNode
		for j, c := range line {
			if c == ' ' {
				continue
			}
			gn := &gridNode{
				pos:  point.New(j, i),
				wall: c == '#',
			}
			cells = append(cells, gn)
			// Check if very first node
			if start == nil && !gn.wall {
				start = gn
			}

			// Connect horizontally
			if rowStart == nil {
				rowStart = gn
			}
			if prevCell != nil {
				prevCell.connectHorizontally(gn)
			}
			prevCell = gn

			// Connect vertically
			if colStarts[j] == nil {
				colStarts[j] = gn
			}
			if prevColCells[j] != nil {
				prevColCells[j].connectVertically(gn)
			}
			prevColCells[j] = gn
		}

		prevCell.connectHorizontally(rowStart)
	}

	for c, start := range colStarts {
		prevColCells[c].connectVertically(start)
	}

	for _, c := range cells {
		fmt.Println(c, c.left, c.up, c.right, c.down)
	}

	return d.travel(start, path)
}
/**/

func (d *day22) Solve(lines []string, o command.Output) {
	solutions := []int{d.solveIt(lines, d.wrapAroundStitch)}

	stitch := d.stitchPart2
	if len(lines[0]) < 100 {
		solutions = append(solutions, d.solveIt(lines, d.exampleStitchPart1))
		stitch = d.exampleStitchPart2
	}
	solutions = append(solutions, d.solveIt(lines, stitch))

	o.Stdoutln(solutions)
}

func (d *day22) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"[6032 6032 5031]",
			},
		},
		{
			ExpectedOutput: []string{
				"[95358 144361]",
			},
		},
	}
}
