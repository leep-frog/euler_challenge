package p85

import (
	"github.com/leep-frog/command/command"

	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P85() *ecmodels.Problem {
	return ecmodels.NoInputNode(85, func(o command.Output) {
		// number of rectangles is (h * (h + 1) / 2) * (w * (w + 1) / 2)

		// denominator of equation times what we want
		want := 2_000_000 * 4
		best := maths.Closest[int](want)

		var iOverTarget bool
		for i := 1; !iOverTarget; i++ {
			hv := i * (i + 1)
			iOverTarget = hv >= want
			var jOverTarget bool
			for j := i; !jOverTarget; j++ {
				wv := j * (j + 1)
				jOverTarget = wv*hv >= want
				best.IndexCheck(i*j, wv*hv)
			}
		}
		o.Stdoutln(best.BestIndex(), best.Best())
	}, &ecmodels.Execution{
		Want: "2772 7999992",
	})
}
