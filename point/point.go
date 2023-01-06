package point

import (
	"fmt"
	"math"

	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/slices"
	"gonum.org/v1/plot/plotter"
)

type Points[T maths.Mathable] []*Point[T]

func (p *Point[T]) Quadrant() int {
	if p.X < 0 && p.Y >= 0 {
		return 0
	}

	if p.X >= 0 && p.Y > 0 {
		return 1
	}

	if p.X > 0 && p.Y <= 0 {
		return 2
	}

	if p.X <= 0 && p.Y < 0 {
		return 3
	}

	return -1
}

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

func NewLineSegmentInt(a, b *Point[int]) *LineSegmentInt {
	return &LineSegmentInt{NewLineSegment(a, b)}
}

func (rls *RationalLineSegment) Copy() *RationalLineSegment {
	return &RationalLineSegment{rls.A.Copy(), rls.B.Copy()}
}

type LineSegmentInt struct {
	*LineSegment[int]
}

type BigLineSegment struct {
	*LineSegment[int]
}

func NewRationalPointI(x, y int) *RationalPoint {
	return NewRationalPoint(fraction.NewRational(x, 1), fraction.NewRational(y, 1))
}

func NewRationalPoint(x, y *fraction.Rational) *RationalPoint {
	return &RationalPoint{x, y}
}

func NewRationalLineSegment(a, b *RationalPoint) *RationalLineSegment {
	return &RationalLineSegment{a, b}
}

type RationalPoint struct {
	X, Y *fraction.Rational
}

func (r *RationalPoint) String() string {
	return fmt.Sprintf("(%v, %v)", r.X, r.Y)
}

func (r *RationalPoint) Copy() *RationalPoint {
	return &RationalPoint{r.X.Copy(), r.Y.Copy()}
}

func (r *RationalPoint) EQ(q *RationalPoint) bool {
	return r.X.EQ(q.X) && r.Y.EQ(q.Y)
}

type RationalLineSegment struct {
	A, B *RationalPoint
}

func (ls *LineSegment[T]) EquationMB() (T, T) {
	// y1 = m*x1 + b
	// y2 = m*x2 + b
	// b = y1 - m*x1 = y2 - m*x2
	// y1 - m*x1 = y2 - m*x2
	// y1 - y2 = m*(x1 - x2)
	// m = (y1 - y2) / (x1 - x2)
	x1, y1 := ls.A.X, ls.A.Y
	x2, y2 := ls.B.X, ls.B.Y
	m := (y1 - y2) / (x1 - x2)
	// b = y1 - m*x1
	b := y1 - (m * x1)
	return m, b
}

// TODO: Combine RationalLineSegment with LineSegment (do like bfs.Int class)
func (ls *LineSegment[T]) Intersect(that *LineSegment[T]) *Point[T] {
	ls1, ls2 := ls, that
	m1, b1 := ls1.EquationMB()
	m2, b2 := ls2.EquationMB()
	// m_1 * x + b_1 = m_2 * x + b_2
	// x * (m_1 - m_2) = (b_2 - b_1)
	// x = (b_2 - b_1) / (m_1 - m_2)

	// If slopes are equal, then return
	//fmt.Println("IF")
	if m1 == m2 {
		return nil
	}
	// Either slopes are (veritcal, horizontal), (horizontal, K) (vertical, K), or (K, K)

	var x, y T
	switch true {
	case math.IsNaN(float64(m2)) && m1 == 0:
		m1, m2 = m2, m1
		b1, b2 = b2, b1
		ls1, ls2 = ls2, ls1
		fallthrough
	case math.IsNaN(float64(m1)) && m2 == 0:
		x = ls1.A.X
		y = ls2.A.Y
		break
	case math.IsNaN(float64(m2)):
		m1, m2 = m2, m1
		b1, b2 = b2, b1
		ls1, ls2 = ls2, ls1
		fallthrough
	case math.IsNaN(float64(m1)):
		x = ls1.A.X
		y = m2*x + b2
		break
	case m2 == 0:
		m1, m2 = m2, m1
		b1, b2 = b2, b1
		ls1, ls2 = ls2, ls1
		fallthrough
	case m1 == 0:
		y = ls1.A.Y
		x = (y - b2) / m2
		break
	default:
		x = (b2 - b1) / (m1 - m2)
		y = x*m1 + b1
	}

	return New(x, y)

	// Now verify it's between them by verifying it's inside the box of
	// (minX, minY), (maxX, maxY)
	//p := &RationalPoint{x, y}
	// TODO: Use OnSegmentExclusive??
	/*if ls1.InBoxInclusive(p) && ls2.InBoxInclusive(p) && !ls1.HasVertex(p) && !ls2.HasVertex(p) {
		return p
	}
	return nil*/
}

// returns m, b
func (rls *RationalLineSegment) EquationMB() (*fraction.Rational, *fraction.Rational) {
	// y1 = m*x1 + b
	// y2 = m*x2 + b
	// b = y1 - m*x1 = y2 - m*x2
	// y1 - m*x1 = y2 - m*x2
	// y1 - y2 = m*(x1 - x2)
	// m = (y1 - y2) / (x1 - x2)
	x1, y1 := rls.A.X, rls.A.Y
	x2, y2 := rls.B.X, rls.B.Y
	m := y1.Minus(y2).Div(x1.Minus(x2))
	// b = y1 - m*x1
	b := y1.Minus(m.Times(x1))
	return m, b
}

// Note doesn't include edge points

func (rls *RationalLineSegment) Intersect(that *RationalLineSegment) *RationalPoint {
	ls1, ls2 := rls, that
	m1, b1 := ls1.EquationMB()
	m2, b2 := ls2.EquationMB()
	// m_1 * x + b_1 = m_2 * x + b_2
	// x * (m_1 - m_2) = (b_2 - b_1)
	// x = (b_2 - b_1) / (m_1 - m_2)

	// If slopes are equal, then return
	//fmt.Println("IF")
	if m1.EQ(m2) {
		return nil
	}
	// Either slopes are (veritcal, horizontal), (horizontal, K) (vertical, K), or (K, K)

	zero := fraction.NewRational(0, 1)
	var x, y *fraction.Rational
	switch true {
	case m2.Undefined() && maths.EQ(m1, zero):
		m1, m2 = m2, m1
		b1, b2 = b2, b1
		ls1, ls2 = ls2, ls1
		fallthrough
	case m1.Undefined() && maths.EQ(m2, zero):
		x = ls1.A.X
		y = ls2.A.Y
		break
	case m2.Undefined():
		m1, m2 = m2, m1
		b1, b2 = b2, b1
		ls1, ls2 = ls2, ls1
		fallthrough
	case m1.Undefined():
		x = ls1.A.X
		y = m2.Times(x).Plus(b2)
		break
	case maths.EQ(m2, zero):
		m1, m2 = m2, m1
		b1, b2 = b2, b1
		ls1, ls2 = ls2, ls1
		fallthrough
	case maths.EQ(m1, zero):
		y = ls1.A.Y
		x = y.Copy().Minus(b2).Div(m2)
		break
	default:
		x = (b2.Minus(b1)).Div(m1.Minus(m2))
		y = x.Times(m1).Plus(b1)
	}

	// Now verify it's between them by verifying it's inside the box of
	// (minX, minY), (maxX, maxY)
	p := &RationalPoint{x, y}
	// TODO: Use OnSegmentExclusive??
	if ls1.InBoxInclusive(p) && ls2.InBoxInclusive(p) && !ls1.HasVertex(p) && !ls2.HasVertex(p) {
		return p
	}
	return nil
}

// Excludes border of box
func (rls *RationalLineSegment) InBoxExclusive(p *RationalPoint) bool {
	minX := maths.MinT(rls.A.X, rls.B.X)
	maxX := maths.MaxT(rls.A.X, rls.B.X)
	minY := maths.MinT(rls.A.Y, rls.B.Y)
	maxY := maths.MaxT(rls.A.Y, rls.B.Y)

	inX := minX.LT(p.X) && p.X.LT(maxX)
	inY := minY.LT(p.Y) && p.Y.LT(maxY)
	return inX && inY
}

// Includes border of box
func (rls *RationalLineSegment) InBoxInclusive(p *RationalPoint) bool {
	minX := maths.MinT(rls.A.X, rls.B.X)
	maxX := maths.MaxT(rls.A.X, rls.B.X)
	minY := maths.MinT(rls.A.Y, rls.B.Y)
	maxY := maths.MaxT(rls.A.Y, rls.B.Y)

	inX := maths.LTE(minX, p.X) && maths.LTE(p.X, maxX)
	inY := maths.LTE(minY, p.Y) && maths.LTE(p.Y, maxY)
	return inX && inY
}

func (rls *RationalLineSegment) HasVertex(p *RationalPoint) bool {
	return p.EQ(rls.A) || p.EQ(rls.B)
}

func (p *RationalPoint) Cross(that *RationalPoint) *fraction.Rational {
	return p.X.Times(that.Y).Minus(p.Y.Times(that.X))
}

func (p *RationalPoint) Minus(that *RationalPoint) *RationalPoint {
	return NewRationalPoint(p.X.Minus(that.X), p.Y.Minus(that.Y))
}

func (p *RationalPoint) HalfPlane(p2, p3 *RationalPoint) *fraction.Rational {
	return p2.Minus(p).Cross(p2.Minus(p3))
}

func (rls *RationalLineSegment) OnSegmentExclusive(p *RationalPoint) bool {
	return rls.A.BetweenExclusive(p, rls.B)
}

func (rls *RationalLineSegment) OnSegmentInclusive(p *RationalPoint) bool {
	return rls.A.BetweenInclusive(p, rls.B)
}

func (p *RationalPoint) BetweenInclusive(q, p2 *RationalPoint) bool {
	return p.EQ(q) || p2.EQ(q) || p.BetweenExclusive(q, p2)
}

func (p *RationalPoint) BetweenExclusive(q, p2 *RationalPoint) bool {
	if p.EQ(q) || p2.EQ(q) {
		return false
	}

	if p.HalfPlane(p2, q).NEQ(fraction.NewRational(0, 1)) {
		return false
	}

	// Now verify it's between them by verifying it's inside the box of
	// (minX, minY), (maxX, maxY)
	minX := maths.MinT(p.X, p2.X)
	maxX := maths.MaxT(p.X, p2.X)
	minY := maths.MinT(p.Y, p2.Y)
	maxY := maths.MaxT(p.Y, p2.Y)
	return q.X.GTE(minX) && q.X.LTE(maxX) && q.Y.GTE(minY) && q.Y.LTE(maxY)
}

func (ls *RationalLineSegment) HalfPlane(p *RationalPoint) bool {
	return ls.A.HalfPlane(ls.B, p).GT(fraction.NewRational(0, 1))
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

func (ls *LineSegment[T]) OnSegmentExclusive(p *Point[T]) bool {
	return ls.A.BetweenExclusive(p, ls.B)
}

func (ls *LineSegment[T]) HasVertex(p *Point[T]) bool {
	return p.Eq(ls.A) || p.Eq(ls.B)
}

func (ls *LineSegment[T]) OnSegmentInclusive(p *Point[T]) bool {
	return ls.A.BetweenInclusive(p, ls.B)
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
	return ch.ContainsExclusive(p)
}

// Contains, but not on edge
func (t *Triangle[T]) ContainsExclusive(p *Point[T]) bool {
	for _, ls := range t.LineSegments() {
		if ls.OnSegmentExclusive(p) {
			return false
		}
	}
	ch := &ConvexHull[T]{[]*Point[T]{t.A, t.B, t.C}}
	return ch.ContainsExclusive(p)
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

func (p *Point[T]) Copy() *Point[T] {
	return New(p.X, p.Y)
}

func (p *Point[T]) ManhattanDistance(that *Point[T]) T {
	return maths.Abs(p.X-that.X) + maths.Abs(p.Y-that.Y)
}

func (p *Point[T]) ManhattanDistanceWithDiagonals(that *Point[T]) T {
	return maths.Max(maths.Abs(p.X-that.X), maths.Abs(p.Y-that.Y))
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

type Rectangle[T maths.Mathable] struct {
	MinX, MinY, MaxX, MaxY T
}

func NewRectangle[T maths.Mathable](MinX, MinY, MaxX, MaxY T) *Rectangle[T] {
	return &Rectangle[T]{MinX, MinY, MaxX, MaxY}
}

func (r *Rectangle[T]) Eq(q *Rectangle[T]) bool {
	return r.MinX == q.MinX && r.MinY == q.MinY && r.MaxX == q.MaxX && r.MaxY == q.MaxY
}

func (r *Rectangle[T]) String() string {
	return fmt.Sprintf("[(%v, %v), (%v, %v)]", r.MinX, r.MinY, r.MaxX, r.MaxY)
}

func (r *Rectangle[T]) Corners() []*Point[T] {
	m := map[string]*Point[T]{}
	ps := []*Point[T]{
		New(r.MinX, r.MinY),
		New(r.MinX, r.MaxY),
		New(r.MaxX, r.MinY),
		New(r.MaxX, r.MaxY),
	}
	for _, p := range ps {
		m[p.String()] = p
	}

	var ret []*Point[T]
	for _, p := range m {
		ret = append(ret, p)
	}
	return ret
}

func (r *Rectangle[T]) Contains(p *Point[T]) bool {
	return r.MinX <= p.X && p.X <= r.MaxX && r.MinY <= p.Y && p.Y <= r.MaxY
}

func (p *Point[T]) BetweenInclusive(q, p2 *Point[T]) bool {
	return p.Eq(q) || p2.Eq(q) || p.BetweenExclusive(q, p2)
}

// Returns true if q is between p and p2
func (p *Point[T]) BetweenExclusive(q, p2 *Point[T]) bool {
	if p.Eq(q) || p2.Eq(q) {
		return false
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

// Returns whether or not the point is in the convex hull, but false if it is on the boundary
func (ch *ConvexHull[T]) ContainsExclusive(p *Point[T]) bool {
	hp := ch.Points[0].HalfPlane(ch.Points[1], p)
	if hp == 0 {
		return false
	}
	sign := hp > 0
	for i := 1; i < len(ch.Points); i++ {
		s := ch.Points[i].HalfPlane(ch.Points[(i+1)%len(ch.Points)], p)
		if s == 0 {
			return false
		}
		// s is zero if it's on the line.
		if (s > 0) != sign {
			return false
		}
	}
	return true
}

func (p *Point[T]) Minus(that *Point[T]) *Point[T] {
	return New(p.X-that.X, p.Y-that.Y)
}

func (p *Point[T]) Plus(that *Point[T]) *Point[T] {
	return New(p.X+that.X, p.Y+that.Y)
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
	points = bread.Copy(points)
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

	return &ConvexHull[T]{append(top, bread.Reverse(bottom)[1:len(bottom)-1]...)}
}
