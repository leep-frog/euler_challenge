package eulerchallenge

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
)

var (
	tc = 0
)

type TromBoard struct {
	// first one is row, second is column
	grid        [][]int
	openSquares int
	// Since we build the tower from bottom up, the board will
	// always look like a bar graph (if (x, y) is filled in, then
	// (x, j) is filled in for all 0 <= j <= y).
	// Note: this required modifying one of the tromino shapes.
	// See tromino thing below
	heights []int
}

func NewTB(x, y int) *TromBoard {
	var grid [][]int
	for i := 0; i < y; i++ {
		grid = append(grid, make([]int, x, x))
	}
	return &TromBoard{grid, x * y, make([]int, x, x)}
}

func (tb *TromBoard) Contains(x, y int) bool {
	contains := y >= 0 && y < len(tb.grid) && x >= 0 && x < len(tb.grid[0])
	return contains && tb.grid[y][x] == 0
}

func (tb *TromBoard) ContainsTromino(t *Tromino, x, y int) bool {
	if !tb.Contains(x, y) {
		return false
	}
	for _, p := range t.ps {
		if !tb.Contains(p[0]+x, p[1]+y) {
			return false
		}
	}
	return true
}

func (tb *TromBoard) Add(id int, t *Tromino, x, y int) bool {
	if !tb.ContainsTromino(t, x, y) {
		return false
	}

	tb.add(id, x, y)
	for _, p := range t.ps {
		tb.add(id, p[0]+x, p[1]+y)
	}
	return true
}

func (tb *TromBoard) add(id int, x, y int) {
	tb.grid[y][x] = id
	tb.openSquares--
	tb.heights[x]++
}

func (tb *TromBoard) Remove(t *Tromino, x, y int) {
	tb.remove(x, y)
	for _, p := range t.ps {
		tb.remove(p[0]+x, p[1]+y)
	}
}

func (tb *TromBoard) remove(x, y int) {
	tb.grid[y][x] = 0
	tb.openSquares++
	tb.heights[x]--
}

type Tromino struct {
	// p1 is (0, 0)
	// Assume that (0,0) is include
	ps      [][]int
	special bool
}

var (
	// Assume every square above and to the left is filled
	trominoes = []*Tromino{
		// X
		// X
		// X
		{[][]int{{0, 1}, {0, 2}}, false},
		// XXX
		{[][]int{{1, 0}, {2, 0}}, false},
		// X_
		// XX
		{[][]int{{1, 0}, {0, 1}}, false},
		// To prevent overhang (which would violate our tb.heights caveat)
		// We change this to two different shapes (using the only two
		// shapes that can come after this, given we go from left to right):
		// XX   XXO  XX
		// X_   XOO  XOOO
		{[][]int{{0, 1}, {1, 1}}, true},
		{[][]int{{0, 1}, {1, 1}, {1, 0}, {2, 0}, {2, 1}}, false},
		{[][]int{{0, 1}, {1, 1}, {1, 0}, {2, 0}, {3, 0}}, false},
		// XX
		// _X
		{[][]int{{0, 1}, {-1, 1}}, false},
		// _X
		// XX
		{[][]int{{1, 0}, {1, 1}}, false},
	}
)

func (tb *TromBoard) String() string {
	var ss []string
	for _, row := range tb.grid {
		var s []string
		for _, c := range row {
			if c == 0 {
				s = append(s, "_")
			} else {
				s = append(s, fmt.Sprintf("%d", c))
			}
		}
		ss = append(ss, strings.Join(s, " "))
	}
	return strings.Join(bread.Reverse(ss), "\n")
}

func (tb *TromBoard) Hash() string {
	return fmt.Sprintf("%v", tb.heights)
}

func (tb *TromBoard) NumberOfArrangements(depth int, curX, curY, tromCnt int, m map[string]int) int {
	// See if we already know how many ways we can complete the current shape.
	if v, ok := m[tb.Hash()]; ok {
		return v
	}

	if tb.openSquares == 0 {
		return 1
	}

	for tb.grid[curY][curX] != 0 {
		if curX == len(tb.grid[0])-1 {
			curX = 0
			curY++
		} else {
			curX++
		}
	}
	cnt := 0
	for _, t := range trominoes {
		if t.special {
			// If on the edge or next cell isn't empty
			if curX == len(tb.grid[0])-1 || tb.grid[curY][curX+1] == 0 {
				continue
			}
		}
		if tb.Add(tromCnt, t, curX, curY) {
			//board
			tromCnt++
			cnt += tb.NumberOfArrangements(depth+1, curX, curY, tromCnt, m)
			tromCnt--
			tb.Remove(t, curX, curY)
		}
	}
	m[tb.Hash()] = cnt
	return cnt
}

func P161() *problem {
	return intsInputNode(161, 2, 0, func(o command.Output, dim []int) {
		o.Stdoutln(NewTB(dim[0], dim[1]).NumberOfArrangements(0, 0, 0, 1, map[string]int{}))
	}, []*execution{
		{
			args: []string{"2", "9"},
			want: "41",
		},
		{
			args: []string{"3", "6"},
			want: "170",
		},
		{
			args: []string{"5", "9"},
			want: "1269900",
		},
		{
			args:     []string{"9", "12"},
			want:     "20574308184277971",
			estimate: 0.5,
		},
	})
}
