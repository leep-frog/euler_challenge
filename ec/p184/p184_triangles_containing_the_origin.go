package p184

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p456"
	"github.com/leep-frog/euler_challenge/point"
)

func P184() *ecmodels.Problem {
	return ecmodels.IntInputNode(184, func(o command.Output, r int) {
		var pts []*point.Point[int]
		for x := 1; x < r; x++ {
			// Add points on axis
			pts = append(pts,
				point.New(x, 0),
				point.New(-x, 0),
				point.New(0, x),
				point.New(0, -x),
			)

			// Ad other points
			for y := 1; x*x+y*y < r*r; y++ {
				pts = append(pts,
					point.New(x, y),
					point.New(x, -y),
					point.New(-x, y),
					point.New(-x, -y),
				)
			}
		}

		o.Stdoutln(p456.OriginTriangles456(pts))
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "8",
		},
		{
			Args: []string{"3"},
			Want: "360",
		},
		{
			Args: []string{"5"},
			Want: "10600",
		},
		{
			Args: []string{"105"},
			Want: "1725323624056",
		},
	})
}
