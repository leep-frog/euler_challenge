package eulerchallenge

import (
	"fmt"
	"math"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
)

// Use type so helper functions are scoped to this problem only (and not
// included in autocomplete suggestions when coding other problems).
type problem177 struct{}

func (p *problem177) almostEqual(a, b float64) bool {
	epsilon := 0.000000001
	return maths.Abs(a-b) < epsilon
}

func (p *problem177) cos(degree float64) float64 {
	// math.Cos(degree * 2 * math.Pi / 360.0)
	return math.Cos(degree * math.Pi / 180.0)
}

func (p *problem177) acos(v float64) float64 {
	return 180.0 * math.Acos(v) / math.Pi
}

func (p *problem177) sin(degree float64) float64 {
	// math.Cos(degree * 2 * math.Pi / 360.0)
	return math.Sin(degree * math.Pi / 180.0)
}

func (p *problem177) integerAngle(p1, p2, p3 *point.Point[float64]) (int, bool) {
	ax, ay := p1.X-p2.X, p1.Y-p2.Y
	bx, by := p3.X-p2.X, p3.Y-p2.Y
	a := math.Sqrt(ax*ax + ay*ay)
	b := math.Sqrt(bx*bx + by*by)
	dot := ax*bx + ay*by
	angle := p.acos(dot / (a * b))
	rem := maths.Abs(math.Remainder(angle, 1))
	return int(math.Round(angle)), p.almostEqual(rem, 0)
}

func (p *problem177) valid(theta1, theta2, theta3, theta4 int, m map[string]bool) bool {
	if theta2 >= 180-(theta3+theta4) {
		return false
	}

	if theta3 >= 180-(theta1+theta2) {
		return false
	}

	A := point.Origin[float64]()
	B := point.New(1.0, 0.0)

	BAD := float64(theta1 + theta2)
	AD := point.NewLineSegment(A, point.New(p.cos(BAD), p.sin(BAD)))

	BCD := float64(theta2)
	AC := point.NewLineSegment(A, point.New(p.cos(BCD), p.sin(BCD)))

	ABD := float64(theta3)
	BD := point.NewLineSegment(B, point.New(B.X-p.cos(ABD), p.sin(ABD)))

	ABC := float64(theta3 + theta4)
	BC := point.NewLineSegment(B, point.New(B.X-p.cos(ABC), p.sin(ABC)))

	D := AD.Intersect(BD)
	C := AC.Intersect(BC)

	ACD, ok := p.integerAngle(A, C, D)
	if !ok {
		return false
	}

	// M = middle, intersection point
	AMB := 180 - theta2 - theta3
	AMD := 180 - AMB
	BMC := AMD
	CMD := AMB
	// Angle order
	angleOrder := []int{
		theta1, theta2, theta3, theta4,
		180 - BMC - theta4,
		ACD,
		180 - ACD - CMD,
		180 - theta1 - AMD,
	}
	if bread.Sum(angleOrder) != 360 {
		panic("ANGLE ORDER")
	}
	// Unique ID
	m[maths.Min(p.quadCode(angleOrder), p.quadCode(bread.Reverse(angleOrder)))] = true
	return true
}

func (p *problem177) quadCode(angles []int) string {
	r := fmt.Sprintf("%v", angles)
	for i := 2; i < len(angles); i += 2 {
		var shifted []int
		for j := 0; j < len(angles); j++ {
			shifted = append(shifted, angles[(i+j)%len(angles)])
		}
		r = maths.Min(r, fmt.Sprintf("%v", shifted))
	}
	return r
}

func P177() *problem {
	return noInputNode(177, func(o command.Output) {
		p := &problem177{}
		m := map[string]bool{}
		for theta1 := 1; theta1 < 180; theta1++ {
			for theta2 := 1; theta1+theta2 < 180; theta2++ {
				// To ensure the intersection for point D is above the x-axis,
				// theta3 < 180 - (theta1 + theta2)
				for theta3 := 1; theta3 < 180-(theta1+theta2); theta3++ {
					// To ensure the intersection for point C is above the x-axis,
					// theta2 < 180 - (theta3 + theta4)
					// theta2 < 180 - theta3 - theta4
					// theta4 < 180 - (theat3 + theta2)
					for theta4 := 1; theta4 < 180-(theta3+theta2); theta4++ {
						p.valid(theta1, theta2, theta3, theta4, m)
					}
				}
			}
		}
		o.Stdoutln(len(m))
	}, &execution{
		want:     "129325",
		estimate: 30,
	})
}
