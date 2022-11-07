package eulerchallenge

import (
	"math"

	"github.com/leep-frog/command"
)

func P222() *problem {
	return noInputNode(222, func(o command.Output) {
		totalLength := 0.0
		for x := 30.0; x <= 50.0; x++ {
			totalLength += x + x
		}

		totalOverlap := 0.0
		// Inersect each circle with the same angle
		// Make two right triangles. Intersection point at intersection angle, theta, with the vertical for each circle.
		// This makes two congruent circles where the sides are:
		// 1: hypotenuse = 2x*c, vertical = k
		// 2: hypotenuse = 2(x + 1)*c, vertical = 100 - k
		// where 'c' is some constant, x is the diameter (we know), and k is a variable.
		// Since the angles will be the same:
		// 1: cos(T) = k/2x*c
		// 2: cos(T) = (100-k)/(2(x+1)*c)
		// k/2x*c = (100-k)/(2(x+1)*c)
		// Cancel out 1/(2c)
		// k/x = (100-k)/(x+1)
		// Solve for k
		// kx + k = 100x - kx
		// 2kx + k = 100x
		// k(2x + 1) = 100x
		// k = 100x/(2x + 1)
		for x := 30.0; x < 50.0; x++ {
			k := (100.0 * x) / (2.0*x + 1.0)
			// Now make another triangle inside the circle.
			// Where the hypotenuse is the radius (x) and the veritcal side is (k - x)
			// aka the part of k that is above the hieght of the middle of the circle.
			// The horizontal part is sqrt(x^2 - (k - x)^2)
			hyp1 := x
			vert1 := k - x
			horz1 := math.Sqrt((hyp1 * hyp1) - (vert1 * vert1))
			// Do the same thing with the other triangle (using x+1 and 50-k)
			// The horizontal part is sqrt((x+1)^2 - ((100-k) - (x+1))^2)
			hyp2 := (x + 1.0)
			vert2 := (100.0 - k) - (x + 1.0)
			horz2 := math.Sqrt((hyp2 * hyp2) - (vert2 * vert2))

			// Finally, the overlap length is
			overlap1 := x - horz1
			overlap2 := (x + 1.0) - horz2
			overlap := overlap1 + overlap2

			totalOverlap += overlap
		}
		totalLength *= 1000.0
		totalOverlap *= 1000.0
		tl := 0.0
		for i := 30.0; i <= 50.0; i++ {
			tl += 2.0 * i
		}
		total := totalLength - totalOverlap
		o.Stdoutln(tl, totalLength, totalOverlap, total)
		o.Stdoutf(" %.0f \n", total)
	}, &execution{
		want: "0",
	})
}
