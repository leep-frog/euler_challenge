package eulerchallenge

import (
	"fmt"
	"math"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/slices"
)

type p177Point struct {
	p      *point.Point[float64]
	theta  int
	oTheta int
	xTheta int
}

func (p *p177Point) String() string {
	return fmt.Sprintf("[%d %d]: %v", p.oTheta, p.xTheta, p.p)
}

func P177() *problem {
	return intInputNode(177, func(o command.Output, n int) {

		epsilon := 0.000000001
		cos := func(degree float64) float64 {
			// math.Cos(degree * 2 * math.Pi / 360.0)
			return math.Cos(degree * math.Pi / 180.0)
		}
		acos := func(v float64) float64 {
			return 180.0 * math.Acos(v) / math.Pi
		}
		sin := func(degree float64) float64 {
			// math.Cos(degree * 2 * math.Pi / 360.0)
			return math.Sin(degree * math.Pi / 180.0)
		}
		integerAngle := func(p1, p2, p3 *point.Point[float64]) (int, bool) {
			// a = (p1.x - p2.x, p1.y - p2.y)
			// b = (p1.x - p3.x, p1.y - p3.y)
			// p2 is the origin
			// a = p1
			// b = p2
			// TODO: Reverse minus order
			ax, ay := p1.X-p2.X, p1.Y-p2.Y
			bx, by := p1.X-p3.X, p1.Y-p3.Y
			a := math.Sqrt(ax*ax + ay*ay)
			b := math.Sqrt(bx*bx + by*by)
			dot := ax*bx + ay*by
			angle := acos(dot / (a * b))
			rem := maths.Abs(math.Remainder(angle, 1))
			if rem > epsilon {
				return 0, false
			}
			return int(math.Round(angle)), true
		}

		// Points will be (0, 0), (0, 1) and then generate all other points from that
		o.Stdoutln(n)
		offsets := []*point.Point[float64]{nil}
		for theta := 1.0; theta < 180; theta++ {
			offsets = append(offsets, point.New(cos(theta), sin(theta)))
		}

		// TODO: plot points and make sure they're a semi-circle

		origin := point.Origin[float64]()
		xPoint := point.New(1.0, 0.0)

		var positiveYpoints []*p177Point
		for theta1 := 1; theta1 < 180; theta1++ {
			slope1 := offsets[theta1]
			// y = m1*x (no b since at the origin)
			m1 := slope1.Y / slope1.X
			for theta2 := theta1 + 1; theta2 < 180; theta2++ {
				slope2 := offsets[theta2]
				// y = m2*x + b
				// Point is (1, 0)
				// 0 = m2 + b
				// b = -m2
				m2 := slope2.Y / slope2.X
				b := -m2

				// Get intersection of points
				// Equation one: y = m1*x
				// Equation two: y = m2*x + b
				// m1 * x = m2 * x + b
				// (m1 - m2) * x = b
				// x = b / (m1 - m2)
				x := b / (m1 - m2)
				y := m1 * x
				p := point.New(x, y)
				midTheta, ok := integerAngle(origin, p, xPoint)
				if !ok {
					continue
				}
				if theta1 == 45 && theta2 == 135 {
					fmt.Println("THEYO", p)
				}
				positiveYpoints = append(positiveYpoints, &p177Point{
					p, midTheta, theta1, 180 - theta2,
				})
			}
		}

		fmt.Println(len(positiveYpoints))

		cnt := 0
		fmt.Println("START", len(positiveYpoints))
		angle := 0
		ids := map[string]bool{}
		// TOOD: switch back to:
		for i, pyp177 := range positiveYpoints {
			if pyp177.oTheta != angle {
				angle = pyp177.oTheta
				fmt.Println("A", angle)
			}
			_ = i
			pyp := pyp177.p
			// TODO: Use range here since j starts at i (and not i + 1)
			for j := i; j < len(positiveYpoints); j++ {
				nyp177 := positiveYpoints[j]
				good := pyp177.oTheta%45 == 0 && pyp177.xTheta%45 == 0 && nyp177.oTheta%45 == 0 && nyp177.xTheta%45 == 0

				nyp := nyp177.p.Copy()
				nyp.Y = -nyp.Y

				ps := []*point.Point[float64]{
					origin,
					pyp,
					xPoint,
					nyp,
				}
				if good {
					fmt.Println("HEYO", pyp177, nyp177, ps, point.IsConvex(ps...))
				}

				if !point.IsConvex(ps...) {
					continue
				}

				if _, ok := integerAngle(origin, nyp, pyp); !ok {
					continue
				}

				fmt.Println("YUP")
				cnt++
				// ids
				angles := []int{
					pyp177.theta,
					pyp177.oTheta + nyp177.oTheta,
					pyp177.xTheta + nyp177.xTheta,
					nyp177.theta,
				}
				slices.Sort(angles)
				if maths.SumSys(angles...) != 360 {
					fmt.Println(angles)
					panic("OOPS")
				}
				ids[fmt.Sprintf("%v", angles)] = true
			}
		}
		fmt.Println(cnt, len(ids))
	}, []*execution{
		{
			args: []string{"1"},
			want: "",
		},
	})
}
