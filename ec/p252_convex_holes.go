package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/slices"
)

func getEmptyTriangles(pts point.Points[int]) []*point.Triangle[int] {
	xSorted := maths.CopySlice(pts)
	ySorted := maths.CopySlice(pts)

	slices.SortFunc(xSorted, func(p, q *point.Point[int]) bool {
		if p.X != q.X {
			return p.X < q.X
		}
		return p.Y < q.Y
	})

	slices.SortFunc(ySorted, func(p, q *point.Point[int]) bool {
		if p.Y != q.Y {
			return p.Y < q.Y
		}
		return p.X < q.X
	})

	// Map from point to index
	xm, ym := map[int]int{}, map[int]int{}
	for xi, xp := range xSorted {
		xm[xp.X] = xi
	}
	for yi, yp := range ySorted {
		ym[yp.Y] = yi
	}

	var r []*point.Triangle[int]
	for i, a := range xSorted {
		for j := i + 1; j < len(xSorted); j++ {
			b := xSorted[j]
			for k := j + 1; k < len(xSorted); k++ {
				c := xSorted[k]
				t := point.NewTriangle(a, b, c)

				minX := maths.Min(a.X, b.X, c.X)
				maxX := maths.Max(a.X, b.X, c.X)
				minY := maths.Min(a.Y, b.Y, c.Y)
				maxY := maths.Max(a.Y, b.Y, c.Y)

				xStart := xm[minX]
				xEnd := xm[maxX]
				xDiff := xEnd - xStart

				yStart := ym[minY]
				yEnd := ym[maxY]
				yDiff := yEnd - yStart

				empty := true
				if xDiff < yDiff {
					for _, p := range xSorted[xStart:(xEnd + 1)] {
						if p.Y >= minY && p.Y <= maxY && t.ContainsExclusive(p) {
							empty = false
							goto ADD_TRI
						}
					}
				} else {
					for _, p := range ySorted[yStart:(yEnd + 1)] {
						if p.X >= minX && p.X <= maxX && t.ContainsExclusive(p) {
							empty = false
							goto ADD_TRI
						}
					}
				}

			ADD_TRI:
				if empty {
					r = append(r, t)
				}
			}
		}
	}

	return r
}

func P252() *problem {
	return intInputNode(252, func(o command.Output, n int) {

		pts := generatePoints252(n)

		best := maths.Largest[string, float64]()
		slices.SortFunc(pts, func(a, b *point.Point[int]) bool {
			if a.X != b.X {
				return a.X < b.X
			}
			return a.Y < b.Y
		})

		emptyTriangles := getEmptyTriangles(pts)
		tm := map[string]bool{}
		for _, t := range emptyTriangles {
			tm[t.String()] = true
		}

		var chs []*point.ConvexHull[int]
		for _, p := range pts {
			var newCHs []*point.ConvexHull[int]
			for _, ch := range chs {
				newCHs = append(newCHs, ch)
				if len(ch.Points) == 1 {
					newCHs = append(newCHs, point.ConvexHullFromPoints(append(ch.Points, p)...))
					continue
				}

				if len(ch.Points) == 2 {
					if tm[point.NewTriangle(p, ch.Points[0], ch.Points[1]).String()] {
						newCHs = append(newCHs, point.ConvexHullFromPoints(append(ch.Points, p)...))
					}
					continue
				}

				newCH := point.ConvexHullFromPoints(append(ch.Points, p)...)
				if len(newCH.Points) != len(ch.Points)+1 {
					continue
				}

				empty := true
				for _, a := range ch.Points {
					for _, b := range ch.Points {
						if a.Eq(b) {
							continue
						}
						if !tm[point.NewTriangle(a, b, p).String()] {
							empty = false
							goto POINT_LOOP
						}
					}
				}
			POINT_LOOP:
				if empty {
					newCHs = append(newCHs, newCH)
					best.Check(newCH.Area())
				}
			}

			chs = append(newCHs, point.ConvexHullFromPoints(p))
		}

		o.Stdoutf("%.1f\n", best.Best())
	}, []*execution{
		{
			args: []string{"20"},
			want: "1049694.5",
		},
		{
			args:     []string{"500"},
			want:     "104924.0",
			estimate: 700,
		},
	})
}

func generatePoints252(n int) point.Points[int] {
	s := []int{290797}
	var t []int

	for i := 0; i <= 2*n; i++ {
		s = append(s, (s[i]*s[i])%50515093)
		t = append(t, (s[i]%2000)-1000)
	}

	var ps []*point.Point[int]
	for k := 1; k <= n; k++ {
		ps = append(ps, point.New(t[2*k-1], t[2*k]))
	}

	return point.Points[int](ps)
}
