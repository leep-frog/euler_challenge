// Package hexgrid defines objects to represent a hexagon
// in a hexagonal grid where each spot has exactly six neighbors.

/*
//      \ n  /
//    nw +--+ ne
//      /    \
//    -+      +-
//      \    /
//    sw +--+ se
//      / s  \
*/

package hexgrid

import (
	"log"
	"strings"

	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
)

// Origin returns a tile at the center of the hexagonal grid.
func Origin() *Tile {
	return &Tile{point.Origin[int]()}
}

// Tile represents a single tile in the hexagonal grid.
type Tile struct {
	// p.X represents the 'column' of the tile.
	// The 0-axis for this runs along the sw-ne line
	// p.Y represents the position of the tile in the column.
	// The 0-axis for this runs along the n-s line
	p *point.Point[int]
}

var (
	codes = map[string]*point.Point[int]{
		"n":  point.New(0, 1),
		"s":  point.New(0, -1),
		"ne": point.New(1, 0),
		"sw": point.New(-1, 0),
		"nw": point.New(-1, 1),
		"se": point.New(1, -1),
	}
)

func (t *Tile) MoveCode(code string) {
	p, ok := codes[strings.ToLower(code)]
	if !ok {
		log.Fatalf("Invalid move code: %q", code)
	}
	t.p.X += p.X
	t.p.Y += p.Y
}

// Distance returns the Tile's distance from the origin.
func (t *Tile) Distance() int {
	if t.p.X == 0 || t.p.Y == 0 {
		return maths.Abs(t.p.X + t.p.Y)
	}

	// Quadrants 1 and 3 are additive
	if (t.p.X > 0) == (t.p.Y > 0) {
		return maths.Abs(t.p.X) + maths.Abs(t.p.Y)
	}

	// Quadrants 2 and 4 are redundant steps
	// ne + s = se
	// By induction, (where b < a)
	// a*ne + b*s = (a-b)*ne + b*se = (a-b) + b steps = a
	return maths.Max(maths.Abs(t.p.X), maths.Abs(t.p.Y))
}
