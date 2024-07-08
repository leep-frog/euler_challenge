package eulerchallenge

import (
	"math"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
)

func P222() *problem {
	return noInputNode(222, func(o command.Output) {
		best := maths.Smallest[string, float64]()

		// Sequential ball ordering
		var sequentialOrder []float64
		for x := 30.0; x <= 50.0; x++ {
			sequentialOrder = append(sequentialOrder, x)
		}
		best.IndexCheck("sequential", pipeLength(sequentialOrder))

		// Alternate between maximum and minimum
		var maxMinOrder []float64
		var minMaxOrder []float64
		for i := 0.0; i <= 9; i++ {
			maxMinOrder = append(maxMinOrder, 50-i, 30+i)
			minMaxOrder = append(minMaxOrder, 30+i, 50-i)
		}
		maxMinOrder = append(maxMinOrder, 40)
		minMaxOrder = append(minMaxOrder, 40)
		best.IndexCheck("maxMin", pipeLength(maxMinOrder))
		best.IndexCheck("minMax", pipeLength(minMaxOrder))

		// Smallest in the middle and then grow as you go out
		// [50 ... 34 32 30 31 33 ... 49]
		var firstHalf, secondHalf []float64
		for i := 0.0; i <= 9; i++ {
			firstHalf = append(firstHalf, 50-2*i)
			secondHalf = append(secondHalf, 49-2*i)
		}
		middleOutOrder := append(append(firstHalf, 30), bread.Reverse(secondHalf)...)
		best.IndexCheck("middleOut", pipeLength(middleOutOrder))

		o.Stdoutf("%s %.0f\n", best.BestIndex(), 1000.0*best.Best())
	}, &execution{
		want: "middleOut 1590933",
	})
}

func pipeLength(sequentialOrder []float64) float64 {
	totalLength := 0.0
	for _, r := range sequentialOrder {
		totalLength += 2 * r
	}

	totalOverlap := 0.0
	for i := 0; i < len(sequentialOrder)-1; i++ {
		totalOverlap += circleOverlap(sequentialOrder[i], sequentialOrder[i+1])
	}

	return (totalLength - totalOverlap)
}

func circleOverlap(r_1, r_2 float64) float64 {
	// Inersect each circle with the same angle
	// Make two right triangles. Intersection point at intersection angle, theta, with the vertical for each circle.
	// This makes two congruent circles where the sides are:
	// 1: hypotenuse = 2(r_1)*c, vertical = k
	// 2: hypotenuse = 2(r_2)*c, vertical = 100 - k
	// where 'c' is some constant, and k is a variable.
	// Since the angles will be the same:
	// 1: cos(T) = k/(2(r_1)*c)
	// 2: cos(T) = (100-k)/(2(r_2)*c)
	// k/(2(r_1)*c) = (100-k)/(2(r_2)*c)
	// Cancel out 1/(2c)
	// k/(r_1) = (100-k)/(r_2)
	// Solve for k
	// k(r_2) = (100-k)(r_1)
	// k(r_2) = 100(r_1) - k(r_1)
	// k(r_1 + r_2) = 100(r_1)
	// k = (100 * r_1) / (r_1 + r_2)
	k := (100.0 * r_1) / (r_1 + r_2)
	// Now make another triangle inside the circle.
	// Where the hypotenuse is the radius (r_1) and the veritcal side is (k - r_1)
	// aka the part of k that is above the hieght of the middle of the circle.
	// The horizontal part is sqrt(r_1^2 - (k - r_1)^2)
	hyp1 := r_1
	vert1 := k - r_1
	horz1 := math.Sqrt((hyp1 * hyp1) - (vert1 * vert1))
	// Do the same thing with the other triangle (using r_2 and 100-k)
	// The horizontal part is sqrt(r_2^2 - ((100-k) - r_2)^2)
	hyp2 := r_2
	vert2 := (100.0 - k) - r_2
	horz2 := math.Sqrt((hyp2 * hyp2) - (vert2 * vert2))

	// Finally, the overlap length is
	overlap1 := r_1 - horz1
	overlap2 := r_2 - horz2
	return overlap1 + overlap2
}
