package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/point"
)

func P184() *problem {
	return intInputNode(181, func(o command.Output, r int) {
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

		o.Stdoutln(originTriangles456(pts))
	}, []*execution{
		{
			args: []string{"2"},
			want: "8",
		},
		{
			args: []string{"3"},
			want: "360",
		},
		{
			args: []string{"5"},
			want: "10600",
		},
		{
			args: []string{"105"},
			want: "1725323624056",
		},
	})
}
