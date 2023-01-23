package y2020

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day20() aoc.Day {
	return &day20{}
}

type day20 struct{}

func (d *day20) Solve(lines []string, o command.Output) {

	// Parse the input
	var tiles []*tile
	idMap := map[int]bool{}
	tMap := map[int][]*tile{}
	for i := 0; i < len(lines); i += 12 {
		id := parse.Atoi(lines[i][5 : len(lines[i])-1])
		idMap[id] = true
		grid := parse.AOCGrid(lines[i+1:i+11], false, true)
		for range []bool{true, false} {
			for rotation := 0; rotation < 4; rotation++ {
				t := &tile{id, grid}
				tiles = append(tiles, t)
				tMap[t.id] = append(tMap[t.id], t)
				grid = maths.Rotate(grid)
			}
			grid = maths.SimpleTranspose(grid)
		}
	}

	ctx := &tileContext{map[string][]*tile{}, map[string][]*tile{}, map[string][]*tile{}, map[string][]*tile{}}
	for _, t := range tiles {
		ctx.add(t)
	}

	size := maths.Sqrt(len(tiles) / 8)
	var grid [][]*tile
	for i := 0; i < size; i++ {
		grid = append(grid, make([]*tile, size, size))
	}
	if !d.search(tMap, ctx, grid, 0, 0, size, idMap) {
		o.Stdoutln("Search yielded no results :(")
		return
	}

	corners := []int{
		grid[0][0].id, grid[size-1][0].id, grid[0][size-1].id, grid[size-1][size-1].id,
	}

	// Construct the picture
	var picture [][]bool
	for _, row := range grid {
		for crowIdx := 1; crowIdx < len(row[0].cells)-1; crowIdx++ {
			var pictureRow []bool
			for _, c := range row {
				pictureRow = append(pictureRow, c.cells[crowIdx][1:len(c.cells[crowIdx])-1]...)
			}
			// if crowIdx > 0 && crowIdx < len(row[0].cells)-1 {
			picture = append(picture, pictureRow)
			// }
		}
	}

	// Construct the sea monster
	seaMonster := functional.Map([]string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}, func(s string) []bool {
		return functional.Map(strings.Split(s, ""), func(c string) bool {
			return c == "#"
		})
	})

	htCount := functional.Count2D(picture, true)
	smSize := functional.Count2D(seaMonster, true)
	var smCount int

	for range []bool{true, false} {
		for rot := 0; rot < 4; rot++ {

			// Iterate over starting points
			for i := 0; i < len(picture)-(len(seaMonster)-1); i++ {
				for j := 0; j < len(picture[i])-(len(seaMonster[0])-1); j++ {
					// Check for monster
					for a, smRow := range seaMonster {
						for b, v := range smRow {
							if !v {
								continue
							}
							if !picture[i+a][j+b] {
								goto NOT_A_MONSTER
							}
						}
					}
					smCount++

				NOT_A_MONSTER:
				}
			}
			if smCount > 0 {
				o.Stdoutln(bread.Product(corners), htCount-smCount*smSize)
				return
			}
			picture = maths.Rotate(picture)
		}
		picture = maths.SimpleTranspose(picture)
	}

	o.Stderr("No match found")
}

func (d *day20) search(tiles map[int][]*tile, ctx *tileContext, grid [][]*tile, row, col, size int, ids map[int]bool) bool {
	if col == size {
		col = 0
		row++
	}
	if row == size {
		return true
	}

	var options []*tile
	for id := range ids {
		options = append(options, tiles[id]...)
	}

	if row > 0 {
		validTop := grid[row-1][col].bottomCode()
		var os []*tile
		for _, o := range options {
			if o.topCode() == validTop {
				os = append(os, o)
			}
		}
		options = os
	}

	if col > 0 {
		validLeft := grid[row][col-1].rightCode()
		var os []*tile
		for _, o := range options {
			if o.leftCode() == validLeft {
				os = append(os, o)
			}
		}
		options = os
	}

	for _, o := range options {
		grid[row][col] = o
		delete(ids, o.id)
		if d.search(tiles, ctx, grid, row, col+1, size, ids) {
			return true
		}
		ids[o.id] = true
		grid[row][col] = nil
	}

	return false
}

type tileContext struct {
	leftSides, rightSides, bottomSides, topSides map[string][]*tile
}

func (ctx *tileContext) add(t *tile) {
	lc, rc, tc, bc := t.leftCode(), t.rightCode(), t.topCode(), t.bottomCode()
	ctx.leftSides[lc] = append(ctx.leftSides[lc], t)
	ctx.rightSides[rc] = append(ctx.rightSides[rc], t)
	ctx.topSides[tc] = append(ctx.topSides[tc], t)
	ctx.bottomSides[bc] = append(ctx.bottomSides[bc], t)
}

type tile struct {
	id    int
	cells [][]bool
}

func (t *tile) toCode(values []bool) string {
	return strings.Join(functional.Map(values, func(b bool) string {
		if b {
			return "#"
		}
		return "."
	}), "")
}

func (t *tile) topCode() string {
	return t.toCode(t.cells[0])
}

func (t *tile) bottomCode() string {
	return t.toCode(t.cells[len(t.cells)-1])
}

func (t *tile) leftCode() string {
	var bs []bool
	for _, row := range t.cells {
		bs = append(bs, row[0])
	}
	return t.toCode(bs)
}

func (t *tile) rightCode() string {
	var bs []bool
	for _, row := range t.cells {
		bs = append(bs, row[len(row)-1])
	}
	return t.toCode(bs)
}

func (t *tile) String() string {
	var r []string
	for _, row := range t.cells {
		r = append(r, t.toCode(row))
	}
	return strings.Join(r, "\n")
}

func (d *day20) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"20899048083289 273",
			},
		},
		{
			ExpectedOutput: []string{
				"68781323018729 1629",
			},
		},
	}
}
