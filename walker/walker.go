package walker

import "github.com/leep-frog/euler_challenge/point"

type CardinalDirection int

const (
	North CardinalDirection = iota
	East
	South
	West
	cardinalDirectionCount
)

const (
	Up    = North
	Right = East
	Down  = South
	Left  = West
)

// CardinalDirections returns NESW vectors as points
func CardinalDirections(grid bool) []*point.Point[int] {
	yCoef := 1
	if grid {
		yCoef = -1
	}
	return []*point.Point[int]{
		point.New(0, 1*yCoef),
		point.New(1, 0),
		point.New(0, -1*yCoef),
		point.New(-1, 0),
	}
}

type Directionable interface {
	~int
}

type Walker[T ~int] struct {
	directionIdx T
	directions   []*point.Point[int]
}

func CardinalWalker(direction CardinalDirection, grid bool) *Walker[CardinalDirection] {
	return &Walker[CardinalDirection]{direction, CardinalDirections(grid)}
}

func (w *Walker[T]) Right() {
	w.directionIdx = T((int(w.directionIdx) + 1) % len(w.directions))
}

func (w *Walker[T]) Left() {
	w.directionIdx = T((int(w.directionIdx) + len(w.directions) - 1) % len(w.directions))
}

func (w *Walker[T]) GetVector(d T) *point.Point[int] {
	return w.directions[d]
}

func (w *Walker[T]) CurrentVector() *point.Point[int] {
	return w.GetVector(w.directionIdx)
}
