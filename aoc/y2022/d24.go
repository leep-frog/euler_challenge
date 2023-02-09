package y2022

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/bread"
)

func Day24() aoc.Day {
	return &day24{}
}

type blizzardDirection rune

const (
	bdEmpty blizzardDirection = '.'
	bdRight blizzardDirection = '>'
	bdLeft  blizzardDirection = '<'
	bdUp    blizzardDirection = '^'
	bdDown  blizzardDirection = 'v'
)

type day24 struct{}

func (d *day24) Solve(lines []string, o command.Output) {
	bm := &blizzardMap{}
	for _, line := range lines[1 : len(lines)-1] {
		var row []blizzardDirection
		for _, c := range line[1 : len(line)-1] {
			row = append(row, blizzardDirection(c))
		}
		bm.grid = append(bm.grid, row)
	}

	start, end := []int{-1, 0}, []int{len(bm.grid), len(bm.grid[0]) - 1}
	solutions := []int{
		bm.solve(start, end),
		bm.solve(end, start),
		bm.solve(start, end),
	}
	o.Stdoutln(solutions[0], bread.Sum(solutions)+len(solutions)-1)
}

type blizzardMap struct {
	grid       [][]blizzardDirection
	start, end []int
	minutes    int
}

func (bm *blizzardMap) solve(start, end []int) int {
	bm.start, bm.end = start, end
	path, _ := bfs.ContextPathSearch[string](bm, []*blizzardLocation{{bm.start[0], bm.start[1], bm.minutes}})
	bm.minutes += len(path)
	return len(path) - 1
}

func (bm *blizzardMap) draw(bl *blizzardLocation, minutes int) {
	fmt.Println(minutes, "============================", bl)
	for i, row := range bm.grid {
		for j := range row {
			if i == bl.i && j == bl.j {
				fmt.Printf("E")
			} else if bm.open(i, j, minutes) {
				fmt.Printf(".")
			} else {
				fmt.Printf("X")
			}
		}
		fmt.Println()
	}
}

// open returns whether the given cell is open after the provided number of minutes.
func (bm *blizzardMap) open(i, j, minutes int) bool {
	if (i == bm.start[0] && j == bm.start[1]) || (i == bm.end[0] && j == bm.end[1]) {
		return true
	}
	if i < 0 || j < 0 || i >= len(bm.grid) || j >= len(bm.grid[i]) {
		return false
	}
	// Check if a right blizzard
	toCheck := [][]int{
		{i, j - minutes, int(bdRight)},
		{i, j + minutes, int(bdLeft)},
		{i - minutes, j, int(bdDown)},
		{i + minutes, j, int(bdUp)},
	}

	for _, tc := range toCheck {
		a, b, bd := tc[0], tc[1], tc[2]
		for a < 0 {
			a += len(bm.grid)
		}
		a = a % len(bm.grid)
		for b < 0 {
			b += len(bm.grid[a])
		}
		b = b % len(bm.grid[a])
		if int(bm.grid[a][b]) == bd {
			return false
		}
	}
	return true
}

var (
	blizzardMoves = [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{0, 0},
	}
)

type blizzardLocation struct {
	i, j, minutes int
}

func (bl *blizzardLocation) Code(bm *blizzardMap, path bfs.Path[*blizzardLocation]) string {
	return bl.String()
}

func (bl *blizzardLocation) String() string {
	return fmt.Sprintf("(%d,%d,%d)", bl.i, bl.j, bl.minutes)
}

func (bl *blizzardLocation) Done(bm *blizzardMap, path bfs.Path[*blizzardLocation]) bool {
	return bl.i == bm.end[0] && bl.j == bm.end[1]
}

func (bl *blizzardLocation) AdjacentStates(bm *blizzardMap, path bfs.Path[*blizzardLocation]) []*blizzardLocation {
	var r []*blizzardLocation
	for _, m := range blizzardMoves {
		neighbor := &blizzardLocation{bl.i + m[0], bl.j + m[1], bl.minutes + 1}
		if bm.open(neighbor.i, neighbor.j, neighbor.minutes) {
			r = append(r, neighbor)
		}
	}
	return r
}

func (d *day24) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"18 54",
			},
		},
		{
			ExpectedOutput: []string{
				"266 853",
			},
		},
	}
}
