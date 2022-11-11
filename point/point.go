package point

import (
	"fmt"
	"math"

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

type Triangle[T maths.Mathable] struct {
	A, B, C *Point[T]
}

type LineSegment[T maths.Mathable] struct {
	A, B *Point[T]
}

func NewLineSegment[T maths.Mathable](a, b *Point[T]) *LineSegment[T] {
	if b.X < a.X || (b.X == a.X && b.Y < a.Y) {
		a, b = b, a
	}
	return &LineSegment[T]{a, b}
}

func (ls *LineSegment[T]) Plot(p *Plot) ([]Plottable, error) {
	ab, err := plotter.NewLine(plotter.XYs{
		{X: float64(ls.A.X), Y: float64(ls.A.Y)},
		{X: float64(ls.B.X), Y: float64(ls.B.Y)},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to plot line segment: %v", err)
	}
	p.P.Add(ab)
	return nil, nil
}

func (ls *LineSegment[T]) Code() string {
	return ls.String()
}

func (ls *LineSegment[T]) OnSegment(p *Point[T]) bool {
	return ls.A.Between(p, ls.B)
}

func (ls *LineSegment[T]) HalfPlane(p *Point[T]) bool {
	return ls.A.HalfPlane(ls.B, p) > 0
}

func (ls *LineSegment[T]) String() string {
	return fmt.Sprintf("[%v-%v]", ls.A, ls.B)
}

func (t *Triangle[T]) Points() []*Point[T] {
	return []*Point[T]{t.A, t.B, t.C}
}

// LineSegments returns the triangle's line segments in relative order to the Points function.
// Specifically, The triangle can be reconstructed from LineSegments()[i] and Points()[i]
func (t *Triangle[T]) LineSegments() []*LineSegment[T] {
	return []*LineSegment[T]{
		NewLineSegment(t.B, t.C),
		NewLineSegment(t.A, t.C),
		NewLineSegment(t.A, t.B),
	}
}

func (t *Triangle[T]) Plot(p *Plot) ([]Plottable, error) {
	var r []Plottable
	for _, p := range t.Points() {
		r = append(r, p)
	}
	for _, ls := range t.LineSegments() {
		r = append(r, ls)
	}
	return r, nil
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

func (t *Triangle[T]) Area() float64 {
	a := t.A.Dist(t.B)
	b := t.B.Dist(t.C)
	c := t.C.Dist(t.A)
	s := (a + b + c) / 2.0
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return area
}

func (t *Triangle[T]) Contains(p *Point[T]) bool {
	ch := &ConvexHull[T]{[]*Point[T]{t.A, t.B, t.C}}
	return ch.Contains(p)
}

// Contains, but not on edge
func (t *Triangle[T]) ContainsExclusive(p *Point[T]) bool {
	for _, ls := range t.LineSegments() {
		if ls.OnSegment(p) {
			return false
		}
	}
	ch := &ConvexHull[T]{[]*Point[T]{t.A, t.B, t.C}}
	return ch.Contains(p)
}

func (t *Triangle[T]) String() string {
	return fmt.Sprintf("[%v, %v, %v]", t.A, t.B, t.C)
}

type Point[T maths.Mathable] struct {
	X T
	Y T
}

// Implemented for maths.Mappable interface
func (p *Point[T]) Code() string {
	return p.String()
}

func (p *Point[T]) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func (p *Point[T]) Dist(that *Point[T]) float64 {
	x := p.X - that.X
	y := p.Y - that.Y
	return math.Sqrt(float64(x*x + y*y))
}

// Check if points are collinear
func (p *Point[T]) Colinear(pts ...*Point[T]) bool {
	if len(pts) <= 1 {
		return true
	}
	q := pts[0]
	for _, pt := range pts[1:] {
		if p.HalfPlane(q, pt) != 0 {
			return false
		}
	}

	return true
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

func (ch *ConvexHull[T]) Area() float64 {
	if len(ch.Points) < 3 {
		return 0
	}
	sum := 0.0
	a := ch.Points[0]
	for i := 1; i < len(ch.Points)-1; i++ {
		sum += NewTriangle(a, ch.Points[i], ch.Points[i+1]).Area()
	}
	return sum
}

func (ch *ConvexHull[T]) Plot(p *Plot) ([]Plottable, error) {
	var r []Plottable
	for i, p := range ch.Points {
		r = append(r, p)
		r = append(r, NewLineSegment(p, ch.Points[(i+1)%len(ch.Points)]))
	}
	return r, nil
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

// IsConvex returns whether or not the set of points are all *corners* on the
// convex hull created by the points.
// Note: if there are duplicate points or any points on the edges of the hull (i.e. not a corner),
// then this will return false.
/*func IsConvex[T maths.Mathable](points ...*Point[T]) bool {
	ch := ConvexHullFromPoints(points...)
	return len(ch.Points) == len(points)
}*/

func IsConvex[T maths.Mathable](points ...*Point[T]) bool {
	if len(points) < 3 {
		return true
	}

	sign := points[0].HalfPlane(points[1], points[2]) > 0
	for i := 1; i <= len(points)-1; i++ {
		hp := points[i].HalfPlane(points[(i+1)%len(points)], points[(i+2)%len(points)])
		if (hp > 0) != sign || hp == 0 {
			return false
		}
	}
	return true
}

// Returns a sorted thing of points
func ConvexHullFromPoints[T maths.Mathable](points ...*Point[T]) *ConvexHull[T] {
	points = maths.CopySlice(points)
	if len(points) < 3 {
		return &ConvexHull[T]{Points: points}
		//panic("Need at least 3 points to compute the convex hull")
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
