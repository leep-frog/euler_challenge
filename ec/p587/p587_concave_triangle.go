package p587

import (
	"math"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

var (
	pi = math.Pi
)

func P587() *ecmodels.Problem {
	return ecmodels.IntInputNode(587, func(o command.Output, pow int) {
		target := math.Pow(10, -float64(pow))
		n := 1
		for ; calculate(float64(n)) > target; n++ {
		}
		o.Stdoutln(n)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "15",
		},
		{
			Args: []string{"3"},
			Want: "2240",
		},
	})
}

func calculate(n float64) float64 {
	squareArea := 4.0
	circleArea := pi
	sectionArea := (squareArea - circleArea) / 4.0

	m := (1.0 / n)
	b := (1.0 / n) - 1.0

	k := (1.0 / n) - 1.0
	A := (1.0 / (n * n)) + 1.0
	B := 2.0 * k / n
	C := k*k - 1.0

	x := quadratic(A, B, C)
	y := m*x + b

	circleRadius := 1.0
	innerDiagonalDist := dist(0, -1, x, y)

	circleAngle := math.Acos(-y)
	circleProportion := circleAngle / (2.0 * pi)
	pizzaArea := circleProportion * circleArea

	crustArea := pizzaArea - isoArea(circleRadius, innerDiagonalDist)

	// outerDiagonalDist := dist(-1, -1, x, y)
	// outerTriArea := isoArea(outerDiagonalDist, 1.0)
	leftTri := (1.0 + y) * (1.0 + x) / 2.0
	rightTri := ((1.0 + y) * (-x) / 2.0)

	outerTriArea := leftTri + rightTri
	finalArea := outerTriArea - crustArea

	ratio := finalArea / sectionArea
	return ratio
}

func quadratic(a, b, c float64) float64 {
	determinant := math.Sqrt(b*b - 4.0*a*c)
	return (-b - determinant) / (2.0 * a)
}

func dist(x1, y1, x2, y2 float64) float64 {
	xd, yd := x2-x1, y2-y1
	return math.Sqrt(xd*xd + yd*yd)
}

func isoArea(twoLen, oneLen float64) float64 {
	s := twoLen
	w := oneLen
	h := math.Sqrt(s*s - w*w/4.0)
	return h * w / 2.0
}
