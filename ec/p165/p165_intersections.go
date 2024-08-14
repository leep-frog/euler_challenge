package p165

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/functional"
)

func P165() *ecmodels.Problem {
	return ecmodels.IntInputNode(165, func(o command.Output, n int) {

		ls := generatePoints165(n)
		functional.SortFunc(ls, func(a, b *point.RationalLineSegment) bool {
			aMinX := maths.MinT(a.A.X, a.B.X)
			bMinX := maths.MinT(b.A.X, b.B.X)
			return aMinX.LT(bMinX)
		})
		intersections := make([]*point.RationalPoint, 0, len(ls)*len(ls))

		var minXs []*fraction.Rational
		for _, l := range ls {
			minXs = append(minXs, maths.MinT(l.A.X, l.B.X))
		}

		for i, l := range ls {
			lMaxX := maths.MaxT(l.A.X, l.B.X)
			for j := i + 1; j < len(ls); j++ {
				o := ls[j]

				if minXs[j].GTE(lMaxX) {
					break
				}

				if intersect := l.Intersect(o); intersect != nil {
					intersections = append(intersections, intersect)
				}
			}
		}

		// Sort the points
		functional.SortFunc(intersections, func(p, q *point.RationalPoint) bool {
			if maths.NEQ(p.X, q.X) {
				return p.X.LT(q.X)
			}
			return p.Y.LT(q.Y)
		})

		// Compare adjacent points and remove duplicates
		var uniq []*point.RationalPoint
		var prev *point.RationalPoint
		for _, i := range intersections {
			if prev == nil || !prev.EQ(i) {
				uniq = append(uniq, i)
			}
			prev = i
		}
		o.Stdoutln(len(uniq))
	}, []*ecmodels.Execution{
		{
			Args: []string{"1"},
			Want: "1",
		},
		{
			Args: []string{"5000"},
			Want: "2868868",
			Estimate: 175,
		},
	})
}

func generatePoints165(n int) []*point.RationalLineSegment {
	if n < 3 {
		return []*point.RationalLineSegment{
			point.NewRationalLineSegment(point.NewRationalPointI(27, 44), point.NewRationalPointI(12, 32)),
			point.NewRationalLineSegment(point.NewRationalPointI(46, 53), point.NewRationalPointI(17, 62)),
			point.NewRationalLineSegment(point.NewRationalPointI(46, 70), point.NewRationalPointI(22, 40)),
		}
	}
	ss := []int{290797}
	var ts []int

	for i := 0; i < 4*n; i++ {
		sn := (ss[i] * ss[i]) % 50515093
		ss = append(ss, sn)
		ts = append(ts, sn%500)
	}

	var ls []*point.RationalLineSegment
	for i := 0; i < n; i++ {
		ls = append(ls, point.NewRationalLineSegment(
			point.NewRationalPointI(ts[4*i], ts[4*i+1]),
			point.NewRationalPointI(ts[4*i+2], ts[4*i+3]),
		))
	}
	return ls
}
