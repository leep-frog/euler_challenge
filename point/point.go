package point

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
	"gonum.org/v1/plot/plotter"
)

type Points[T maths.Mathable] []*Point[T]

func (pts Points[T]) Plot(p *Plot) ([]Plottable, error) {
	var plt []Plottable
	for _, p := range pts {
		plt = append(plt, p)
	}
	return plt, nil
}

func eh() {
	k := []*Point[int]{}
	CreatePlot("", 1, 2, Points[int](k))
}

type Triangle[T maths.Mathable] struct {
	a, b, c *Point[T]
}

type LineSegment[T maths.Mathable] struct {
	a, b *Point[T]
}

func NewLineSegment[T maths.Mathable](a, b *Point[T]) *LineSegment[T] {
	return &LineSegment[T]{a, b}
}

func (t *LineSegment[T]) Plot(p *Plot) ([]Plottable, error) {
	ab, err := plotter.NewLine(plotter.XYs{
		{X: float64(t.a.X), Y: float64(t.a.Y)},
		{X: float64(t.b.X), Y: float64(t.b.Y)},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to plot line segment: %v", err)
	}
	p.P.Add(ab)
	return nil, nil
}

func (t *Triangle[T]) Plot(p *Plot) ([]Plottable, error) {
	return []Plottable{
		t.a, t.b, t.c,
		NewLineSegment(t.a, t.b),
		NewLineSegment(t.b, t.c),
		NewLineSegment(t.c, t.a),
	}, nil
}

func Origin[T maths.Mathable]() *Point[T] {
	return New[T](0, 0)
}

func NewTriangle[T maths.Mathable](a, b, c *Point[T]) *Triangle[T] {
	ps := []*Point[T]{a, b, c}
	slices.SortFunc(ps, func(this, that *Point[T]) bool {
		if this.X != that.X {
			return this.X < that.X
		}
		return this.Y < that.Y
	})
	return &Triangle[T]{ps[0], ps[1], ps[2]}
}

func (t *Triangle[T]) Contains(p *Point[T]) bool {
	ch := &ConvexHull[T]{[]*Point[T]{t.a, t.b, t.c}}
	return ch.Contains(p)
}

func (t *Triangle[T]) String() string {
	return fmt.Sprintf("[%v, %v, %v]", t.a, t.b, t.c)
}

type Point[T maths.Mathable] struct {
	X T
	Y T
}

func (p *Point[T]) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func New[T maths.Mathable](x, y T) *Point[T] {
	return &Point[T]{x, y}
}

func (p *Point[T]) Plot(plt *Plot) ([]Plottable, error) {
	sc, err := plotter.NewScatter(plotter.XYs{
		plotter.XY{X: float64(p.X), Y: float64(p.Y)},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to plot point: %v", err)
	}
	plt.P.Add(sc)
	return nil, nil
}

func (p *Point[T]) Eq(that *Point[T]) bool {
	return p.X == that.X && p.Y == that.Y
}

// Returns true if q is between p and p2
func (p *Point[T]) Between(q, p2 *Point[T]) bool {
	if p.Eq(q) || p2.Eq(q) {
		return true
	}

	if p.HalfPlane(p2, q) != 0 {
		return false
	}

	// Now verify it's between them by verifying it's inside the box of
	// (minX, minY), (maxX, maxY)
	minX := maths.Min(p.X, p2.X)
	maxX := maths.Max(p.X, p2.X)
	minY := maths.Min(p.Y, p2.Y)
	maxY := maths.Max(p.Y, p2.Y)
	return q.X >= minX && q.X <= maxX && q.Y >= minY && q.Y <= maxY
}

type ConvexHull[T maths.Mathable] struct {
	Points []*Point[T]
}

func (ch *ConvexHull[T]) Contains(p *Point[T]) bool {
	sign := ch.Points[0].HalfPlane(ch.Points[1], p) > 0
	for i := 1; i < len(ch.Points); i++ {
		s := ch.Points[i].HalfPlane(ch.Points[(i+1)%len(ch.Points)], p)
		// s is zero if it's on the line.
		if (s > 0) != sign && s != 0 {
			return false
		}
	}
	return true
}

func (p *Point[T]) Minus(that *Point[T]) *Point[T] {
	return New(p.X-that.X, p.Y-that.Y)
}

func (p *Point[T]) Cross(that *Point[T]) T {
	return p.X*that.Y - p.Y*that.X
}

func (p *Point[T]) HalfPlane(p2, p3 *Point[T]) T {
	return p2.Minus(p).Cross(p2.Minus(p3))
}

// Returns a sorted thing of points
func ConvexHullFromPoints[T maths.Mathable](points ...*Point[T]) *ConvexHull[T] {
	if len(points) < 3 {
		panic("Need at least 3 points to compute the convex hull")
	}

	// Sort by *ascending* Y coordinate for the bottom hull.
	slices.SortFunc(points, func(this, that *Point[T]) bool {
		if this.X != that.X {
			return this.X < that.X
		}
		return this.Y > that.Y
	})

	var top []*Point[T]
	for _, p := range points {
		// Sort upper hull
		for (len(top) > 0 && top[len(top)-1].Eq(p)) || (len(top) >= 2 && p.HalfPlane(top[len(top)-1], top[len(top)-2]) <= 0) {
			top = top[:len(top)-1]
		}
		top = append(top, p)
	}

	var bottom []*Point[T]
	for _, p := range points {
		// Sort upper hull
		for (len(bottom) > 0 && bottom[len(bottom)-1].Eq(p)) || (len(bottom) >= 2 && p.HalfPlane(bottom[len(bottom)-1], bottom[len(bottom)-2]) >= 0) {
			bottom = bottom[:len(bottom)-1]
		}
		bottom = append(bottom, p)
	}

	return &ConvexHull[T]{append(top, maths.Reverse(bottom)[1:len(bottom)-1]...)}
}
