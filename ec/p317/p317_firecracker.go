package p317

import (
	"fmt"
	"math"
	"math/big"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
)

const (
	acceleration = -9.81
)

func P317() *ecmodels.Problem {
	return ecmodels.IntInputNode(317, func(o command.Output, n int) {

		// better(float64(n))
		// incrementer(n)
		// fmt.Println(pi)
		bigIncrementer(n)
		return
		// Elipses that intersect the point
		// points := []*point.Point[float64]{
		// 	point.New(0.0, 0.0),
		// }

		outOf := float64(n)
		var ch *point.ConvexHull[float64]
		for i := 0.0; i <= outOf; i++ {
			// angle := rand.Float64() * math.Pi / 2.0
			angle := (i / outOf) * math.Pi / 2.0
			// points = append(points, getTime(angle)...)

			ps := getTime(angle)

			if ch == nil {
				ch = point.ConvexHullFromPoints(append(ps, point.Origin[float64]())...)
			} else {
				ch = point.ConvexHullFromPoints(append(ps, ch.Points...)...)
			}
		}

		var plottables []point.Plottable
		var originIdx int
		// ch := point.ConvexHullFromPoints(points...)
		for i, p := range ch.Points {
			plottables = append(plottables, p)
			// fmt.Println(p)
			if p.Eq(point.Origin[float64]()) {
				originIdx = i
			}
		}

		var lowVolume, highVolume float64
		for j := 1; j < len(ch.Points)-1; j++ {
			a := ch.Points[(originIdx+j)%len(ch.Points)]
			b := ch.Points[(originIdx+j+1)%len(ch.Points)]

			// Calculate volume of the cylinder
			height := a.Y - b.Y
			if height < 0 {
				height = -height
			}

			lowRadius := maths.Min(a.X, b.X)
			highRadius := maths.Max(a.X, b.X)

			lowVolume += math.Pi * lowRadius * lowRadius * height
			highVolume += math.Pi * highRadius * highRadius * height
		}

		fmt.Printf("LOW: %0.6f\n", lowVolume)
		fmt.Printf("HIGH: %0.6f\n", highVolume)

		// for x := 0.0; x <= 125; x += 0.001 {
		// 	y := x*x*(-120.38/(99.0*99.0)) + 120.38
		// 	plottables = append(plottables, point.New(x, y))
		// }

		if err := point.CreatePlot("hello.png", 800, 800, plottables...); err != nil {
			fmt.Println("OOPS:", err)
		}

	}, []*ecmodels.Execution{
		{
			Want:     "1856532.8455",
			Estimate: 250,
		},
	})
}

var (
	bestX, bestY = 99.08300000010905, 120.38734722616469
	// BestY calculation:
	// 0 = 20 + acc * t
	// t = - 20 / acc
	// bestY = 100 + 20 * t + acc * t * t / 2.0
	bestYT = -20.0 / acceleration
	bestYY = 100.0 + 20.0*bestYT + acceleration*bestYT*bestYT/2.0
)

func getTime(angle float64) []*point.Point[float64] {
	// angle := rand.Float64() * math.Pi / 2.0
	v_horz, v_vert := 20*math.Sin(angle), 20*math.Cos(angle)

	// fmt.Println(angle, v_horz, v_vert)
	// disp = v_0 * t + acc * t^2 / 2
	// a = acc/2, b = v_0, c = -disp
	a := acceleration / 2.0
	b := v_vert
	c := 100.0 // -(-100)

	roots := maths.QuadraticRoots(a, b, c)
	if (roots[0] < 0) == (roots[1] < 0) {
		panic("Both negative roots?!")
	}
	time := roots[0]
	if time < 0 {
		time = roots[1]
	}
	x_ground := v_horz * time
	// fmt.Println("Hits ground at", x_ground)

	// Intersection points are
	// (0, 100), (x_ground, 0)

	// Now, find the vertex (where y velocity is 0)
	// v = v_0 + acc * t
	// 0 = v_vert + acc * t
	// t = - v_vert / acc
	t_vertex := -v_vert / acceleration

	x_vertex := v_horz * t_vertex
	y_vertex := 100.0 + v_vert*t_vertex + acceleration*t_vertex*t_vertex/2
	// fmt.Printf("Vertex: t=%f (%f, %f)\n", t_vertex, x_vertex, y_vertex)

	// Parabola equation:
	// y = a(x-h)^2 + k
	// h = x_vertex
	// k = y_vertex
	// Plug in (x_ground, 0) to find a:
	// 0 = a(x_ground-x_vertex)^2 + y_vertex
	// a = - y_vertex / (x_ground - x_vertex) ^ 2
	a_coef := -y_vertex / ((x_ground - x_vertex) * (x_ground - x_vertex))

	// fmt.Printf("y = %f * (x - %f)^2 + %f\n", a_coef, x_vertex, y_vertex)

	var plottables []*point.Point[float64]
	for x := x_vertex; ; x += 0.005 {
		y := a_coef*(x-x_vertex)*(x-x_vertex) + y_vertex
		if y < 0.0 {
			break
		}

		p := point.New(x, y)

		// y_curve :=
		if p.Dist(point.New(x, f(x))) < 1.0 {
			plottables = append(plottables, p)
		}

		if p.X > bestX {
			bestX = p.X
			fmt.Println("BEST", bestX)
		}

		if p.Y > bestY {
			bestY = p.Y
			fmt.Println("BEST", bestY)
		}

	}

	return plottables

}

func area() {
	// (99, 0)
	// Vertex : (0, 120.38)

	// y = a(x-h)^2 + k
	// h = x_vertex = 0
	// k = y_vertex = 120.38
	// y = ax^2 + 120.38
	// Plug in (x_ground, 0) to find a:
	// 0 = a(99^2) + 120.38
	// a = -120.38 / 99*99

	// a := -120.38 / (99.0*99.0)
	// a = 0.01228242016
	// k := 120.38

	// Integral from 0 to 99 of
	// y = pi * x * [-120.38 / (99.0*99.0) x^2 + 120.38]

	// Invert
	// y = ax^2 + k
	// sqrt((y - k) / a) = x
	// sqrt((120.38 - y) / 0.01228242016) = x
	// Integral from 0 to 120.38 of pi * (y - 120.38) / 0.01228242016

}

func f(x float64) float64 {
	return x*x*(-120.38/(99.0*99.0)) + 120.38
}

/*

When shoot directly up, we want to find when velocity is zero
0 = v_0 + acceleration * t
t = -v_0 / acceleration
  = -20 / -9.81
	= 20 / 9.81

Then, find the displacement
disp = 100 + v_0 * t + acc * t * t / 2
     = 100 + 20 * 20 / 9.81 - 9.81 * 20 * 20 / 9.81 / 9.81 / 2
		 = 100 + 400 / 9.81 - 20 * 10 / 9.81
		 = 100 + 400 / 9.81 - 200 / 9.81
		 = 120.387359837

////////////////////////////////////////////////////////////////////////////////

Next, find the farthest x point, which occurs when velocity is totally horizontal

First, find the amount of time it takes to hit the ground
0 = 100 + v_h0 * t + acc * t * t / 2
  = 100 + acc * t * t / 2
	= 200 + acc * t * t

-200 / acc = t*t

t = sqrt(-200 / -9.81)
t = sqrt(200 / 9.81)

disp = 20 * t
     = 20 * sqrt(200 / 9.81)


////////////////////////////////////////////////////////////////////////////////

Maximize displacement with angle as an input

disp_h = sin(angle) * t

Amount of time it takes to hit the ground
e0 = 100 +




1856532.8455
*/

func better(n float64) {
	origin := point.Origin[float64]()

	// TODO: increment to find bestX and bestY
	corner := point.New(bestX, 0.0)
	prevPoint := point.New(0, bestYY)
	fmt.Println("BESTY", bestYY)

	points := []*point.Point[float64]{
		// origin, prevPoint,
		prevPoint,
	}
	plottables := []point.Plottable{
		// origin, prevPoint,
		prevPoint,
	}

	for angle, incr := 1.0/n, 1.0/n; ; angle += incr {
		// for i := 1.0; i < n; i++ {
		// fmt.Println("\n\n===================", i/n)
		// angle := (i / n) * math.Pi / 2.0
		gnp := getNextPoint(angle, origin, corner, prevPoint)
		if gnp == nil {
			continue
		}
		if gnp.Y < 0 {
			break
		}
		points = append(points, gnp)
		plottables = append(plottables, gnp)
		prevPoint = gnp
	}
	// fmt.Println(points)
	points = append(points, point.New(bestX, 0))

	if err := point.CreatePlot("hello.png", 800, 800, plottables...); err != nil {
		fmt.Println("OOPS:", err)
	}

	var volume float64
	for j := 1; j < len(points); j++ {
		a := points[j-1]
		b := points[j]

		// Calculate volume of the cylinder
		height := a.Y - b.Y
		if height < 0 {
			height = -height
		}

		radius := (a.X + b.X) / 2.0

		volume += math.Pi * radius * radius * height
	}
	fmt.Printf("VOL: %0.6f\n", volume)

	var lowVolume, highVolume float64
	for j := 1; j < len(points); j++ {
		a := points[j-1]
		b := points[j]

		// Calculate volume of the cylinder
		height := a.Y - b.Y
		if height < 0 {
			height = -height
		}

		lowRadius := maths.Min(a.X, b.X)
		highRadius := maths.Max(a.X, b.X)

		lowVolume += math.Pi * lowRadius * lowRadius * height
		highVolume += math.Pi * highRadius * highRadius * height
	}

	fmt.Printf("LOW: %0.6f\n", lowVolume)
	fmt.Printf("HIGH: %0.6f\n", highVolume)

}

func getNextPoint(angle float64, origin, corner, prevPoint *point.Point[float64]) *point.Point[float64] {
	// angle := rand.Float64() * math.Pi / 2.0
	v_horz, v_vert := 20*math.Sin(angle), 20*math.Cos(angle)
	// fmt.Println("VEL", v_horz, v_vert, angle)

	// fmt.Println(angle, v_horz, v_vert)
	// disp = v_0 * t + acc * t^2 / 2
	// a = acc/2, b = v_0, c = -disp
	a := acceleration / 2.0
	b := v_vert
	c := 100.0 // -(-100)

	roots := maths.QuadraticRoots(a, b, c)
	if (roots[0] < 0) == (roots[1] < 0) {
		panic("Both negative roots?!")
	}
	time := roots[0]
	if time < 0 {
		time = roots[1]
	}
	x_ground := v_horz * time
	// fmt.Println("Hits ground at", x_ground)

	// Intersection points are
	// (0, 100), (x_ground, 0)

	// Now, find the vertex (where y velocity is 0)
	// v = v_0 + acc * t
	// 0 = v_vert + acc * t
	// t = - v_vert / acc
	t_vertex := -v_vert / acceleration

	x_vertex := v_horz * t_vertex
	y_vertex := 100.0 + v_vert*t_vertex + acceleration*t_vertex*t_vertex/2
	// fmt.Printf("Vertex: t=%f (%f, %f)\n", t_vertex, x_vertex, y_vertex)

	// Parabola equation:
	// y = a(x-h)^2 + k
	// h = x_vertex
	// k = y_vertex
	// Plug in (x_ground, 0) to find a:
	// 0 = a(x_ground-x_vertex)^2 + y_vertex
	// a = - y_vertex / (x_ground - x_vertex) ^ 2
	a_coef := -y_vertex / ((x_ground - x_vertex) * (x_ground - x_vertex))

	// fmt.Printf("y = %f * (x - %f)^2 + %f\n", a_coef, x_vertex, y_vertex)

	f := func(x float64) float64 {
		return a_coef*(x-x_vertex)*(x-x_vertex) + y_vertex
	}

	// Finally, do the thing
	var ch *point.ConvexHull[float64]
	var lastCHPoint *point.Point[float64]
	decr := 0.1
	for x := prevPoint.X + 2.0; x >= prevPoint.X; x -= decr {
		// for x := prevPoint.X + 0.05; x >= prevPoint.X; x -= 0.00001 {
		y := f(x)

		p := point.New(x, y)

		if ch == nil {
			ch = point.ConvexHullFromPoints(origin, prevPoint, p)
			lastCHPoint = p
			continue
		}

		prevLen := len(ch.Points)
		ch = point.ConvexHullFromPoints(append(ch.Points, p)...)
		newLen := len(ch.Points)

		// A point wasn't added, meaning we passed the tangent
		if prevLen == newLen {
			if decr == 0.001 {
				fmt.Println(lastCHPoint)
				return lastCHPoint
			} else {
				x += decr
				decr = decr / 10.0
				continue
			}
		}
		lastCHPoint = p
	}

	return nil
	panic("ARGH")

	// ch := point.ConvexHullFromPoints(origin, prevPoint, f(prevPoint.X + 2.0))
	// nextPoint :=
	// var ch *point.ConvexHull[float64]
	/*ch := point.ConvexHullFromPoints(origin, corner, prevPoint)

	// First, get the first point that breaks out of the existing shape
	var nextPoint *point.Point[float64]
	for x := bestX; ; x -= 0.1 {
		y := f(x)
		if y <= 0.0 {
			continue
		}
		if x >= bestX {
			continue
		}

		p := point.New(x, y)

		ch = point.ConvexHullFromPoints(append(ch.Points, p)...)

		if len(ch.Points) > 3 {
			nextPoint = p
			break
		}
	}

	if nextPoint == nil {
		panic("NO")
	}

	// First, get the

	// var prevP *point.Point[float64]
	ch = point.ConvexHullFromPoints(origin, corner, prevPoint, nextPoint)
	// for x := prevPoint.X + 2.0; ; x -= 0.1 {
	for x := nextPoint.X - 0.01; x >= prevPoint.X; x -= 0.01 {
		y := f(x)

		p := point.New(x, y)

		fmt.Println("PREV", ch.Points)
		prevLen := len(ch.Points)
		ch = point.ConvexHullFromPoints(append(ch.Points, p)...)
		fmt.Println("NEW", ch.Points)
		newLen := len(ch.Points)
		if prevLen == newLen {
			fmt.Println(nextPoint)
			panic("AHO")
			return nextPoint
		}
		nextPoint = p
	}

	panic("AH")*/
}

func incrementer(ni int) {
	// point.New()
	/*
		- Given an angle a/b
		- Increment over a curve
	*/

	points := []point.Plottable{
		point.New(0.0, bestY),
	}

	incr := 0.0001
	angleIncr := 0.00000001
	angle := angleIncr
	for atX := incr; ; atX += incr {

		// find the best angle
		height := getHeight(angle, atX)
		for {
			nextAngle := angle + angleIncr
			nextHeight := getHeight(nextAngle, atX)
			if nextHeight >= height {
				angle, height = nextAngle, nextHeight
			} else {
				break
			}
		}

		if height < 0.0 {
			break
		}

		points = append(points, point.New(atX, height))

		// fmt.Println(angle, height)
	}

	if err := point.CreatePlot("hello.png", 800, 800, points...); err != nil {
		fmt.Println("OOPS:", err)
	}

	var lowVolume, volume, highVolume float64
	for j := 1; j < len(points)-1; j++ {
		a := points[j-1].(*point.Point[float64])
		b := points[j].(*point.Point[float64])

		// Calculate volume of the cylinder
		height := a.Y - b.Y
		if height < 0 {
			height = -height
		}

		lowRadius := maths.Min(a.X, b.X)
		radius := (a.X + b.X) / 2.0
		highRadius := maths.Max(a.X, b.X)

		lowVolume += math.Pi * lowRadius * lowRadius * height
		volume += math.Pi * radius * radius * height
		highVolume += math.Pi * highRadius * highRadius * height
	}

	fmt.Printf("LOW:  %0.6f\n", lowVolume)
	fmt.Printf("MID:  %0.6f\n", volume)
	fmt.Printf("HIGH: %0.6f\n", highVolume)

}

func getHeight(angle float64, atX float64) float64 {
	v_horz, v_vert := 20*math.Sin(angle), 20*math.Cos(angle)

	// time from x=0 to x=atX
	t := atX / v_horz

	// height at t
	// disp = v_0 * t + acc * t^2 / 2
	return 100.0 + v_vert*t + acceleration*t*t/2.0
}

var (
	zero = newFlote(0.0)
)

func bigIncrementer(ni int) {
	// point.New()
	/*
		- Given an angle a/b
		- Increment over a curve
	*/

	prevPoint := newPoint(newFlote(0.0), newFlote(bestY))
	volume := newFlote(0.0)

	points := []*bigPoint{
		prevPoint,
	}

	// TODO: instead of incrementing, binary search up to a threshold
	incr := newFlote(0.0001)
	// angleIncr := newFlote(0.000001)
	// angle := angleIncr
	/// angle := &bigAngle{newFlote(1.0)}
	sine := newFlote(0.0000000000001)
	for atX := incr; ; atX = atX.Plus(incr) {

		// find the best angle
		/*//height := getBigHeight(angle, atX)
		for {
			nextAngle := angle.next()
			nextHeight := getBigHeight(nextAngle, atX)
			if nextHeight.GTE(height) {
				angle, height = nextAngle, nextHeight
			} else {
				break
			}
		}*/
		var height *flote
		sine, height = getBestHeight2(sine, atX)

		if height.LT(zero) {
			break
		}

		// Calc volume
		curPoint := newPoint(atX, height)
		// midPoint := newPoint(curPoint.X.Plus(prevPoint.X).Times(newFlote(0.5)), curPoint.Y.Plus(prevPoint.Y).Times(newFlote(0.5)))
		// volume = volume.Plus(calcBigVolume(curPoint, midPoint)).Plus(calcBigVolume(midPoint, prevPoint))

		volume = volume.Plus(calcBigVolume(atX, curPoint, prevPoint))
		prevPoint = curPoint
		points = append(points, prevPoint)
	}

	fmt.Printf("MID:  %0.6f\n", volume.float)

	for i := 0; i < 10; i++ {
		fmt.Println(points[i])
	}

	fmt.Println("========")

	for i := 0; i < 10; i++ {
		fmt.Println(points[len(points)-i-1])
	}

	var plts []point.Plottable

	for _, p := range points {
		x, _ := p.X.float.Float64()
		y, _ := p.Y.float.Float64()
		plts = append(plts, point.New(x, y))
	}

	if err := point.CreatePlot("hello.png", 800, 800, plts...); err != nil {
		fmt.Println("OOPS:", err)
	}

	/*lowVolume, volume, highVolume := newFlote(0), newFlote(0), newFlote(0)
	_ = lowVolume
	_ = highVolume
	for j := 1; j < len(points)-1; j++ {
		a := points[j-1]
		b := points[j]

		// Calculate volume of the cylinder
		height := a.Y.Minus(b.Y)
		if height.LT(zero) {
			// height = height.Times()
			// fmt.Println(height, a.Y, b.Y)
			// panic("noo")
			height = height.Times(newFlote(-1))
		}

		// lowRadius := maths.Min(a.X, b.X)
		radius := (a.X.Plus(b.X)).Div(newFlote(2.0))
		// highRadius := maths.Max(a.X, b.X)

		// lowVolume += math.Pi * lowRadius * lowRadius * height
		volume = volume.Plus(newFlote(math.Pi).Times(radius).Times(radius).Times(height))
		// highVolume += math.Pi * highRadius * highRadius * height
	}*/

	// fmt.Printf("LOW:  %0.6f\n", lowVolume)
	// fmt.Printf("MID:  %0.6f\n", volume.float)
	// fmt.Printf("HIGH: %0.6f\n", highVolume)

}

var (
	tenth = newFlote(0.1)
	coef  = newFlote(0.1).Times(tenth).Times(tenth).Times(tenth).Times(tenth).Times(tenth).Times(tenth).Times(tenth)
)

type bigAngle struct {
	idx *flote
}

func (ba *bigAngle) next() *bigAngle {
	return &bigAngle{ba.idx.Plus(newFlote(1.0))}
}

func (ba *bigAngle) sineCosine() (*flote, *flote) {
	s := ba.idx.Times(coef)
	c := newFlote(1.0).Minus(s.Times(s)).Sqrt()
	return s, c
}

func getBigHeight(angle *bigAngle, atX *flote) *flote {

	// floatAngle, _ := angle.float.Float64()

	s, c := angle.sineCosine()

	v_horz, v_vert := s.Times(newFlote(20.0)), c.Times(newFlote(20.0))

	// time from x=0 to x=atX
	t := atX.Div(v_horz)

	// height at t
	// disp = v_0 * t + acc * t^2 / 2
	return newFlote(100.0).Plus(v_vert.Times(t)).Plus(newFlote(acceleration).Times(t).Times(t).Div(newFlote(2.0)))
}

func calcBigVolume(atX *flote, pointA, pointB *bigPoint) *flote {
	dHeight := pointA.Y.Minus(pointB.Y)
	if dHeight.LT(zero) {
		// height = height.Times()
		// fmt.Println(height, a.Y, b.Y)
		// panic("noo")
		dHeight = dHeight.Times(newFlote(-1))
	}
	radius := (pointA.X.Plus(pointB.X)).Div(newFlote(2.0))

	return pi.Times(radius).Times(radius).Times(dHeight)
}

func getBestHeight(angle *bigAngle, atX *flote) (*bigAngle, *flote) {
	height := getBigHeight(angle, atX)
	for {
		nextAngle := angle.next()
		nextHeight := getBigHeight(nextAngle, atX)
		if nextHeight.GTE(height) {
			angle, height = nextAngle, nextHeight
		} else {
			break
		}
	}

	return angle, height
}

var (
	sineOffset    = newFlote(0.0001)
	diffThreshold = newFlote(0.00000000001)
	two           = newFlote(2.0)
	pi            = calcPi()
)

func calcPi() *flote {
	f := newFlote(0.0)
	pif, _, err := f.float.Parse("3.1415926535897932384626433832795028841971", 10)
	if err != nil {
		panic(err)
	}
	return &flote{pif}
}

func getBestHeight2(sine *flote, atX *flote) (*flote, *flote) {
	left, right := sine, sine.Plus(sineOffset)

	// fmt.Println("HEIGHT", sine, atX)

	for leftHeight, rightHeight := getBigHeight2(left, atX), getBigHeight2(right, atX); left.AbsDiff(right).GTE(diffThreshold); {

		midPoint := left.Plus(right).Div(two)

		if leftHeight.GTE(rightHeight) {
			right = midPoint
			rightHeight = getBigHeight2(right, atX)
		} else {
			left = midPoint
			leftHeight = getBigHeight2(left, atX)
		}
	}

	midPoint := left.Plus(right).Div(two)
	return midPoint, getBigHeight2(midPoint, atX)
}

func getBigHeight2(sine *flote, atX *flote) *flote {

	// floatAngle, _ := angle.float.Float64()

	cosine := newFlote(1.0).Minus(sine.Times(sine)).Sqrt()

	v_horz, v_vert := sine.Times(newFlote(20.0)), cosine.Times(newFlote(20.0))

	// time from x=0 to x=atX
	t := atX.Div(v_horz)

	// height at t
	// disp = v_0 * t + acc * t^2 / 2
	return newFlote(100.0).Plus(v_vert.Times(t)).Plus(newFlote(acceleration).Times(t).Times(t).Div(newFlote(2.0)))
}

type bigPoint struct {
	X, Y *flote
}

func newPoint(x, y *flote) *bigPoint {
	return &bigPoint{x, y}
}

func (bp *bigPoint) String() string {
	return fmt.Sprintf("(%v, %v)", bp.X, bp.Y)
}

type flote struct {
	float *big.Float
}

const PREC = 500

func newFlote(fl float64) *flote {
	f := big.NewFloat(fl)
	f.SetPrec(PREC)
	return &flote{f}
}

func (f *flote) String() string {
	return fmt.Sprintf("%v", f.float)
}

func (f *flote) Plus(g *flote) *flote {
	r := newFlote(0.0).float
	return &flote{r.Add(f.float, g.float)}
}

func (f *flote) Minus(g *flote) *flote {
	r := newFlote(0.0).float
	return &flote{r.Sub(f.float, g.float)}
}

func (f *flote) AbsDiff(g *flote) *flote {
	r := f.Minus(g)
	if r.LT(zero) {
		r = r.Times(newFlote(-1.0))
	}
	return r
}

func (f *flote) Times(g *flote) *flote {
	r := newFlote(0.0).float
	return &flote{r.Mul(f.float, g.float)}
}

func (f *flote) Div(g *flote) *flote {
	r := newFlote(0.0).float
	return &flote{r.Quo(f.float, g.float)}
}

func (f *flote) Sqrt() *flote {
	r := newFlote(0.0).float
	return &flote{r.Sqrt(f.float)}
}

func (f *flote) GTE(g *flote) bool {
	return f.float.Cmp(g.float) >= 0
}

func (f *flote) GT(g *flote) bool {
	return f.float.Cmp(g.float) > 0
}

func (f *flote) LTE(g *flote) bool {
	return f.float.Cmp(g.float) <= 0
}

func (f *flote) LT(g *flote) bool {
	return f.float.Cmp(g.float) < 0
}
