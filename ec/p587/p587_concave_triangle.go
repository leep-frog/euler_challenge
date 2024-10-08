package p587

import (
	"math"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
)

var (
	pi = math.Pi
)

func P587() *ecmodels.Problem {
	return ecmodels.IntInputNode(587, func(o command.Output, pow int) {
		target := math.Pow(10, -float64(pow))
		res, _ := bfs.UnboundedBinarySearch(1, -target, func(i int) float64 { return -calculate(i) })
		o.Stdoutln(res)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "15",
		},
		{
			Args: []string{"3"},
			Want: "2240",
		},
		{
			Args: []string{"8"},
			Want: "232951565",
		},
	})
}

// We place the drawing such that the left-most circle has a radius of one and
// is centered at the origin. Therefore, the bottom left corner is at (-1, -1)
func calculate(ni int) float64 {
	n := float64(ni)

	squareArea := 4.0 // square with size length 2
	circleArea := pi  // circle with radius 1
	sectionArea := (squareArea - circleArea) / 4.0

	// Get the equation for the line that goes through (-1, -1) and (2n-1, 1)
	// slope = (1 - (-1)) / (2n - 1 - (-1)) = 2 / 2n = 1 / n
	m := (1.0 / n)
	// plug in (-1, -1) to get the y-offset:
	// y = mx + b
	// -1 = -1/n + b
	// b = 1/n - 1
	b := (1.0 / n) - 1.0

	// The equation for the circle is y^2 + x^2 = 1, so solve for the intersection points
	// y^2 = 1 - x^2   |   y = mx + b
	// y^2 = 1 - x^2   |   y^2 = m^2 * x^2 + 2 * m * x * b + b^2
	// 1 - x^2 = m^2*x^2 + 2mb*x + b^2
	// 0 = (m^2 + 1)*x^2 + 2mb*x + b^2 - 1
	A := m*m + 1
	B := 2.0 * m * b
	C := b*b - 1.0

	// The two intersection points will be one to the left of the origin, and one
	// to the right. We just need the one on the left
	x := minusQuadratic(A, B, C)
	y := m*x + b

	// Finally, we calculate the desired area via the following:
	// L_A = Area of the left half of the section (a right triangle)
	// R_A = Area the right half of the section (a right, concave triangle)
	//     = (Area of a regular right triangle with the same points) - (the pizza crust)
	//     = R_A' - crust

	// Calculate L_A and R_A'
	leftTriangleArea := (1.0 + y) * (1.0 + x) / 2.0
	rightTriangleArea := ((1.0 + y) * (-x) / 2.0)

	// Now calculate the area of the pizza crust
	// Area of pizze crust is (area of circle) * (pizza slice's proportion of the circle)
	circleAngle := math.Acos(-y)
	circleProportion := circleAngle / (2.0 * pi)
	pizzaArea := circleProportion * circleArea

	// Calculate the area of the non-crust part of the pizza (just an isosceles triangle)
	circleRadius := 1.0
	innerDiagonalDist := dist(0, -1, x, y) // line segment
	nonCrustArea := isoArea(circleRadius, innerDiagonalDist)

	// Finally, get the area of the crust
	crustArea := pizzaArea - nonCrustArea

	// outerDiagonalDist := dist(-1, -1, x, y)
	// outerTriArea := isoArea(outerDiagonalDist, 1.0)

	finalArea := leftTriangleArea + rightTriangleArea - crustArea

	return finalArea / sectionArea
}

func minusQuadratic(a, b, c float64) float64 {
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
