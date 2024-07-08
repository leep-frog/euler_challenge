package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/slices"
)

func P165() *problem {
	return intInputNode(165, func(o command.Output, n int) {
		ls := generatePoints165(n)
		var intersections []*point.RationalPoint
		for i, l := range ls {
			for j := i + 1; j < len(ls); j++ {
				o := ls[j]
				if intersect := l.Copy().Intersect(o.Copy()); intersect != nil {
					intersections = append(intersections, intersect)
				}
			}
		}

		// Sort the points
		slices.SortFunc(intersections, func(p, q *point.RationalPoint) bool {
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
	}, []*execution{
		{
			args: []string{"1"},
			want: "1",
		},
		{
			args:     []string{"5000"},
			want:     "2868868",
			estimate: 200,
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
