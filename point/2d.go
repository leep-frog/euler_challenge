package point

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
)

type Triangle2D[T maths.Mathable] struct {
	a, b, c *Point2D[T]
}

func Origin2D[T maths.Mathable]() *Point2D[T] {
	return New2D[T](0, 0)
}

func NewTriangle2D[T maths.Mathable](a, b, c *Point2D[T]) *Triangle2D[T] {
	ps := []*Point2D[T]{a, b, c}
	slices.SortFunc(ps, func(this, that *Point2D[T]) bool {
		if this.X != that.X {
			return this.X < that.X
		}
		return this.Y < that.Y
	})
	return &Triangle2D[T]{ps[0], ps[1], ps[2]}
}

func (t *Triangle2D[T]) Contains(p *Point2D[T]) bool {
	ch := &ConvexHull2D[T]{[]*Point2D[T]{t.a, t.b, t.c}}
	return ch.Contains(p)
}

func (t *Triangle2D[T]) String() string {
	return fmt.Sprintf("[%v, %v, %v]", t.a, t.b, t.c)
}

type Point2D[T maths.Mathable] struct {
	X T
	Y T
}

func (p *Point2D[T]) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func New2D[T maths.Mathable](x, y T) *Point2D[T] {
	return &Point2D[T]{x, y}
}

func (p *Point2D[T]) Eq(that *Point2D[T]) bool {
	return p.X == that.X && p.Y == that.Y
}

// Returns true if q is between p and p2
func (p *Point2D[T]) Between(q, p2 *Point2D[T]) bool {
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

type ConvexHull2D[T maths.Mathable] struct {
	Points []*Point2D[T]
}

func (ch *ConvexHull2D[T]) Contains(p *Point2D[T]) bool {
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

func (p *Point2D[T]) Minus(that *Point2D[T]) *Point2D[T] {
	return New2D(p.X-that.X, p.Y-that.Y)
}

func (p *Point2D[T]) Cross(that *Point2D[T]) T {
	return p.X*that.Y - p.Y*that.X
}

func (p *Point2D[T]) HalfPlane(p2, p3 *Point2D[T]) T {
	return p2.Minus(p).Cross(p2.Minus(p3))
}

// Returns a sorted thing of points
func ConvexHull2DFromPoints[T maths.Mathable](points ...*Point2D[T]) *ConvexHull2D[T] {
	if len(points) < 3 {
		panic("Need at least 3 points to compute the convex hull")
	}

	// Sort by *ascending* Y coordinate for the bottom hull.
	slices.SortFunc(points, func(this, that *Point2D[T]) bool {
		if this.X != that.X {
			return this.X < that.X
		}
		return this.Y > that.Y
	})

	var top []*Point2D[T]
	for _, p := range points {
		// Sort upper hull
		for (len(top) > 0 && top[len(top)-1].Eq(p)) || (len(top) >= 2 && p.HalfPlane(top[len(top)-1], top[len(top)-2]) <= 0) {
			top = top[:len(top)-1]
		}
		top = append(top, p)
	}

	var bottom []*Point2D[T]
	for _, p := range points {
		// Sort upper hull
		for (len(bottom) > 0 && bottom[len(bottom)-1].Eq(p)) || (len(bottom) >= 2 && p.HalfPlane(bottom[len(bottom)-1], bottom[len(bottom)-2]) >= 0) {
			bottom = bottom[:len(bottom)-1]
		}
		bottom = append(bottom, p)
	}

	return &ConvexHull2D[T]{append(top, maths.Reverse(bottom)[1:len(bottom)-1]...)}
}
