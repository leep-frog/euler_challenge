package y2020

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/point"
)

func Day24() aoc.Day {
	return &day24{}
}

type day24 struct{}

func (d *day24) Solve(lines []string, o command.Output) {
	blackTiles := map[string]*point.Point[int]{}
	for _, line := range lines {
		tile := point.Origin[int]()
		for i := 0; i < len(line); i++ {
			if line[i] == 'e' {
				tile.X++
			} else if line[i] == 'w' {
				tile.X--
			} else if line[i:i+2] == "se" {
				tile.Y--
				i++
			} else if line[i:i+2] == "sw" {
				tile.Y--
				tile.X--
				i++
			} else if line[i:i+2] == "ne" {
				tile.Y++
				tile.X++
				i++
			} else if line[i:i+2] == "nw" {
				tile.Y++
				i++
			} else {
				fmt.Println(line, i)
				panic("No clue")
			}
		}
		if _, ok := blackTiles[tile.String()]; ok {
			delete(blackTiles, tile.String())
		} else {
			blackTiles[tile.String()] = tile
		}
	}
	part1 := len(blackTiles)

	neighborMoves := []*point.Point[int]{
		point.New(0, 1),
		point.New(1, 1),
		point.New(-1, -1),
		point.New(0, -1),
		point.New(1, 0),
		point.New(-1, 0),
	}

	for i := 0; i < 100; i++ {
		tileCounts := map[string]*tileCount{}
		for _, t := range blackTiles {
			for _, m := range neighborMoves {
				n := t.Plus(m)
				if v, ok := tileCounts[n.String()]; ok {
					v.count++
				} else {
					tileCounts[n.String()] = &tileCount{n, 1}
				}
			}
		}

		newBlackTiles := map[string]*point.Point[int]{}
		// Add black tiles
		for _, bt := range blackTiles {
			if tc := tileCounts[bt.String()]; tc != nil && (tc.count == 1 || tc.count == 2) {
				newBlackTiles[bt.String()] = bt
			}
		}

		// Add white tiles
		for _, tc := range tileCounts {
			// Ignore black tiles since already checked those
			if _, ok := blackTiles[tc.tile.String()]; ok {
				continue
			}

			if tc.count == 2 {
				newBlackTiles[tc.tile.String()] = tc.tile
			}
		}
		blackTiles = newBlackTiles
	}
	o.Stdoutln(part1, len(blackTiles))
}

type tileCount struct {
	tile  *point.Point[int]
	count int
}

func (d *day24) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"10 2208",
			},
		},
		{
			ExpectedOutput: []string{
				"339 3794",
			},
		},
	}
}
