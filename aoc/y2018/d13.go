package y2018

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/walker"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/maps"
)

func Day13() aoc.Day {
	return &day13{}
}

type day13 struct{}

type track int

const (
	emptyPath track = iota
	straightPath
	forwardSlash
	backwardSlash
	intersection
)

type cart struct {
	idx int
	w   *walker.Walker[walker.CardinalDirection]
}

func (c *cart) String() string {
	return c.w.Position().String()
}

func (d *day13) Solve(lines []string, o command.Output) {
	// width := functional.Reduce[int, string](0, lines, func(i int, s string) int {
	// return maths.Max(i, len(s))
	// })
	var grid [][]track
	var carts []*cart
	for y, line := range lines {
		var row []track
		for x, c := range line {
			switch c {
			case 'v':
				carts = append(carts, &cart{0, walker.CardinalWalkerAt(walker.Down, x, y, true)})
				row = append(row, straightPath)
			case '>':
				carts = append(carts, &cart{0, walker.CardinalWalkerAt(walker.Right, x, y, true)})
				row = append(row, straightPath)
			case '<':
				carts = append(carts, &cart{0, walker.CardinalWalkerAt(walker.Left, x, y, true)})
				row = append(row, straightPath)
			case '^':
				carts = append(carts, &cart{0, walker.CardinalWalkerAt(walker.Up, x, y, true)})
				row = append(row, straightPath)
			case '|', '-':
				row = append(row, straightPath)
			case '+':
				row = append(row, intersection)
			case '/':
				row = append(row, forwardSlash)
			case '\\':
				row = append(row, backwardSlash)
			default:
				row = append(row, emptyPath)
			}
		}
		grid = append(grid, row)
	}

	var part1 string

	for {
		functional.SortFunc(carts, func(a, b *cart) bool {
			ap, bp := a.w.Position(), b.w.Position()
			if ap.Y != bp.Y {
				return ap.Y < bp.Y
			}
			return ap.X < bp.X
		})

		cartMap := map[string]*cart{}
		var keys []string
		for _, c := range carts {
			cartMap[c.String()] = c
			keys = append(keys, c.String())
		}

		newCartMap := map[string]*cart{}
		for _, k := range keys {
			c := cartMap[k]
			delete(cartMap, k)
			if c == nil {
				continue
			}

			cp := c.w.Position()
			switch grid[cp.Y][cp.X] {

			case straightPath:
			case forwardSlash:
				if c.w.Direction() == walker.Up || c.w.Direction() == walker.Down {
					c.w.Right()
				} else {
					c.w.Left()
				}
			case backwardSlash:
				if c.w.Direction() == walker.Up || c.w.Direction() == walker.Down {
					c.w.Left()
				} else {
					c.w.Right()
				}
			case intersection:
				if c.idx == 0 {
					c.w.Left()
				} else if c.idx == 2 {
					c.w.Right()
				}
				c.idx = (c.idx + 1) % 3
			default:
				fmt.Println("UNKNOWN GRID CELL", cp.Y, cp.X, grid[cp.Y][cp.X])
				panic("NOPE")
			}
			c.w.Walk(1)

			if _, ok := cartMap[c.String()]; ok {
				// Check if a cart that hasn't moved yet is in the way.
				delete(cartMap, c.String())
				if part1 == "" {
					part1 = fmt.Sprintf("%d,%d", c.w.Position().X, c.w.Position().Y)
				}
			} else if _, ok := newCartMap[c.String()]; ok {
				// Check if a cart that has already moved is in the way.
				delete(newCartMap, c.String())
				if part1 == "" {
					part1 = fmt.Sprintf("%d,%d", c.w.Position().X, c.w.Position().Y)
				}
			} else {
				// Otherwise, the cart survived
				newCartMap[c.String()] = c
			}
		}

		// Update carts
		carts = maps.Values(newCartMap)
		if len(carts) <= 1 {
			c := carts[0]
			o.Stdoutf("%s %d,%d", part1, c.w.Position().X, c.w.Position().Y)
			return
		}
	}
}

func (d *day13) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"1,0 6,4",
			},
		},
		{
			ExpectedOutput: []string{
				"91,69 44,87",
			},
		},
	}
}
