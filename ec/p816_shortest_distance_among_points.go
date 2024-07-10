package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/functional"
)

func P816() *problem {
	return intInputNode(816, func(o command.Output, n int) {
		s := 290797
		mod := 50515093

		var pts []*point.Point[int]
		for len(pts) < n {
			x := s
			s = (s * s) % mod
			y := s
			s = (s * s) % mod
			pts = append(pts, point.New(x, y))
		}

		// o.Stdoutf("%.9f\n", elegant816One(pts))
		o.Stdoutf("%.9f\n", elegant816Two(pts))
	}, []*execution{
		{
			args: []string{"14"},
			want: "546446.466846479",
		},
		{
			args:     []string{"2_000_000"},
			want:     "20.880613018",
			estimate: 1.25,
		},
	})
}

func elegant816Two(pts []*point.Point[int]) float64 {
	functional.SortFunc(pts, func(a, b *point.Point[int]) bool {
		return a.X < b.X
	})

	best := maths.Smallest[any, float64]()
	for idx, p := range pts {
		for _, q := range pts[idx+1:] {
			if best.Set() && best.Best() < maths.Abs[float64](float64(p.X-q.X)) {
				break
			}
			best.Check(p.Dist(q))
		}
	}
	return best.Best()
}

func elegant816One(pts []*point.Point[int]) float64 {
	rc := point.NewRectangularContainer[int](pts)
	best := maths.Smallest[any, float64]()

	for _, q := range pts[1:] {
		best.Check(pts[0].Dist(q))
	}

	for _, p := range pts {
		rc.ShortestDistance(p, best)
	}
	return best.Best()
}
