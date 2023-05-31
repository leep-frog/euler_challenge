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

func NeighborsWithDiagonals() []*point.Point[int] {
	return []*point.Point[int]{
		point.New(0, 1),
		point.New(0, -1),
		point.New(1, 1),
		point.New(1, 0),
		point.New(1, -1),
		point.New(-1, 1),
		point.New(-1, 0),
		point.New(-1, -1),
	}
}

type Directionable interface {
	~int
}

type Walker[T ~int] struct {
	directionIdx T
	directions   []*point.Point[int]
	position     *point.Point[int]
}

func CardinalWalker(direction CardinalDirection, grid bool) *Walker[CardinalDirection] {
	return CardinalWalkerAt(direction, 0, 0, grid)
}

func CardinalWalkerAt(direction CardinalDirection, x, y int, grid bool) *Walker[CardinalDirection] {
	return &Walker[CardinalDirection]{
		direction,
		CardinalDirections(grid),
		point.New(x, y),
	}
}

func (w *Walker[T]) Position() *point.Point[int] {
	return w.position.Copy()
}

func (w *Walker[T]) Direction() T {
	return w.directionIdx
}

func (w *Walker[T]) Move(d CardinalDirection, steps int) {
	w.position = w.position.Plus(w.directions[d].Times(steps))
}

func (w *Walker[T]) MoveTo(p *point.Point[int]) {
	w.position = p.Copy()
}

func (w *Walker[T]) Walk(steps int) {
	w.position = w.position.Plus(w.CurrentVector().Times(steps))
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
