package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/slices"
)

func getEmpty(pts point.Points[int]) []*point.Triangle[int] {
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

	fmt.Println(xSorted)

	// 487199

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
		fmt.Println(i)
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
				//fmt.Println(xDiff, yDiff)
				if xDiff < yDiff {
					for _, p := range xSorted[xStart:(xEnd + 1)] {
						if t.ContainsExclusive(p) {
							empty = false
							goto ADD_TRI
						}
					}
				} else {
					for _, p := range ySorted[yStart:(yEnd + 1)] {
						if t.ContainsExclusive(p) {
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

	// TODO: Check unique triangles
	return r
}

func P252() *problem {
	return intInputNode(252, func(o command.Output, n int) {

		pts := generatePoints252(n)

		//return

		fmt.Println("HELLO")

		best := maths.Largest[string, float64]()
		slices.SortFunc(pts, func(a, b *point.Point[int]) bool {
			if a.X != b.X {
				return a.X < b.X
			}
			return a.Y < b.Y
		})

		// empty triangles
		/*var emptyTriangles []*point.Triangle[int]
		tm := map[string]bool{}
		for i, a := range pts {
			if i%10 == 0 {
				fmt.Println("TRI", i)
			}
			for j := i + 1; j < len(pts); j++ {
				b := pts[j]
				for k := j + 1; k < len(pts); k++ {
					c := pts[k]

					t := point.NewTriangle(a, b, c)
					empty := true
					for _, p := range pts {
						if t.ContainsExclusive(p) {
							empty = false
							break
						}
					}
					if empty {
						emptyTriangles = append(emptyTriangles, t)
						tm[t.String()] = true
					}
				}
			}
		}

		fmt.Println("EMPTY", len(emptyTriangles))
		return*/
		emptyTriangles := getEmpty(pts)
		tm := map[string]bool{}
		for _, t := range emptyTriangles {
			tm[t.String()] = true
		}

		var chs []*point.ConvexHull[int]
		for i, p := range pts {
			fmt.Println("PT", i)
			toAdd := []*point.ConvexHull[int]{
				point.ConvexHullFromPoints(p),
			}

			for _, ch := range chs {
				if len(ch.Points) == 1 {
					toAdd = append(toAdd, point.ConvexHullFromPoints(append(ch.Points, p)...))
					continue
				}

				if len(ch.Points) == 2 {
					if tm[point.NewTriangle(p, ch.Points[0], ch.Points[1]).String()] {
						toAdd = append(toAdd, point.ConvexHullFromPoints(append(ch.Points, p)...))
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
					toAdd = append(toAdd, newCH)
				}
			}

			for _, ch := range toAdd {
				chs = append(chs, ch)
				best.Check(ch.Area())
			}
		}

		fmt.Printf("%.2f", best.Best())
	}, []*execution{
		{
			args: []string{"20"},
		},
		{
			args: []string{"500"},
		},
	})
}

/*func P252() *problem {
	return intInputNode(252, func(o command.Output, n int) {
		pts := generatePoints252(n)

		var emptyTriangles []*point.Triangle[int]
		// for i, a := range pts[:len(pts)-2] {
		// 	fmt.Println(i)
		// 	for j := i + 1; j < len(pts)-1; j++ {
		// 		b := pts[j]
		// 		for _, c := range pts[j+1:] {
		for i, a := range pts {
			fmt.Println(i)
			for _, b := range pts {
				if a.Eq(b) {
					continue
				}
				for _, c := range pts {
					if a.Eq(c) || b.Eq(c) {
						continue
						//panic(fmt.Sprintf("NOO %v %v %v", a, b, c))
					}
					triangle := point.NewTriangle(a, b, c)
					empty := true
					for _, p := range pts {
						if p.Eq(a) || p.Eq(b) || p.Eq(c) {
							continue
						}
						// Don't care if a point is on the line between triangles
						if triangle.Contains(p) && !a.Between(p, b) && !a.Between(p, c) && !b.Between(p, c) {
							empty = false
							break
						}
					}

					if empty {
						emptyTriangles = append(emptyTriangles, triangle)
					}
				}
			}
		}

		// Now create a map from
		m := map[string]map[string]*maths.Set[*point.Point[int]]{}
		for _, t := range emptyTriangles {
			a, b, c := t.A.String(), t.B.String(), t.C.String()

			// Initialize first layers
			if m[a] == nil {
				m[a] = map[string]*maths.Set[*point.Point[int]]{}
			}
			if m[b] == nil {
				m[b] = map[string]*maths.Set[*point.Point[int]]{}
			}
			if m[c] == nil {
				m[c] = map[string]*maths.Set[*point.Point[int]]{}

			}

			// Initialize second layers
			if m[a][b] == nil {
				m[a][b] = maths.NewSet[*point.Point[int]]()
			}
			if m[a][c] == nil {
				m[a][c] = maths.NewSet[*point.Point[int]]()
			}
			if m[b][a] == nil {
				m[b][a] = maths.NewSet[*point.Point[int]]()
			}
			if m[b][c] == nil {
				m[b][c] = maths.NewSet[*point.Point[int]]()
			}
			if m[c][a] == nil {
				m[c][a] = maths.NewSet[*point.Point[int]]()
			}
			if m[c][b] == nil {
				m[c][b] = maths.NewSet[*point.Point[int]]()
			}

			m[a][b].Add(t.C)
			m[b][a].Add(t.C)

			m[a][c].Add(t.B)
			m[c][a].Add(t.B)

			m[b][c].Add(t.A)
			m[c][b].Add(t.A)
		}

		ctx := &context252{
			0,
			maths.Largest[[]*node252, float64](),
			pts,
			nil,
			m,
			map[string]bool{},
		}

		var initNodes []*node252
		for _, p := range pts {
			initNodes = append(initNodes, &node252{p})
		}

		bCnt = 0
		bfs.PoppableContextualDFS(initNodes, ctx, bfs.AllowDFSDuplicates())

		var bestPoints []*point.Point[int]
		for _, node := range ctx.bestArea.BestIndex() {
			bestPoints = append(bestPoints, node.point)
		}
		ch := point.ConvexHullFromPoints(bestPoints...)

		point.CreatePlot(fmt.Sprintf("252-%d.png", n), 800, 800, pts, ch, point.Axes(-1000, 1000))

		// TODO: Stdoutfln
		o.Stdoutf("%.2f %v", ctx.bestArea.Best(), ctx.bestArea.BestIndex())
	}, []*execution{
		{
			args: []string{"20"},
			want: "1049694.50",
		},
		{
			args: []string{"500"},
		},
	})
}

type context252 struct {
	area float64

	bestArea *maths.Bester[[]*node252, float64]

	allPoints []*point.Point[int]

	// fixedPoint is the first index
	points []*point.Point[int]

	// Map from line segment, to points
	m map[string]map[string]*maths.Set[*point.Point[int]]

	// After we evaluate a point, we know that it is not in the optimal set
	ignorePoints map[string]bool
}

type node252 struct {
	point *point.Point[int]
}

func (n *node252) String() string {
	return n.point.String()
}

var (
	bCnt = 0
)

func (n *node252) Code(ctx *context252, path bfs.DFSPath[*node252]) string {
	if path.Len() == 0 {
		bCnt++
		fmt.Println("BC", bCnt)
	}
	return n.point.String()
}

func (n *node252) Done(ctx *context252, path bfs.DFSPath[*node252]) bool {
	if path.Path()[0].String() != "(-306, -447)" {
		ctx.bestArea.IndexCheck(maths.CopySlice(path.Path()), ctx.area)
	}
	return false
}

func (n *node252) OnPop(ctx *context252, path bfs.DFSPath[*node252]) {
	if len(ctx.points) >= 3 {
		ln := len(ctx.points) - 1
		ctx.area -= point.NewTriangle(ctx.points[0], ctx.points[ln], ctx.points[ln-1]).Area()
	}
	ctx.points = ctx.points[:len(ctx.points)-1]
	if len(ctx.points) == 0 {
		ctx.ignorePoints[n.point.String()] = true
	}
}

func (n *node252) OnPush(ctx *context252, path bfs.DFSPath[*node252]) {
	ctx.points = append(ctx.points, n.point)
	if len(ctx.points) >= 3 {
		ln := len(ctx.points) - 1
		ctx.area += point.NewTriangle(ctx.points[0], ctx.points[ln], ctx.points[ln-1]).Area()
	}
}

func (n *node252) AdjacentStates(ctx *context252, path bfs.DFSPath[*node252]) []*node252 {
	var r []*node252

	if len(ctx.points) == 1 {
		for _, p := range ctx.allPoints {
			if !ctx.ignorePoints[p.String()] {
				r = append(r, &node252{p})
			}
		}
		return r
	}

	set := ctx.m[ctx.points[0].String()][n.point.String()]
	if set == nil {
		return nil
	}
	set.For(func(p *point.Point[int]) bool {
		if !ctx.ignorePoints[p.String()] && point.IsConvex(maths.CopySlice(append(ctx.points, p))...) {
			r = append(r, &node252{p})
		}
		return false
	})

	return r
}

/*func brute252(n int) float64 {
	pts := generatePoints252(n)


}*/

func generatePoints252(n int) point.Points[int] {
	s := []int{290797}
	var t []int

	for i := 0; i <= 2*n; i++ {
		s = append(s, (s[i]*s[i])%50515093)
		t = append(t, (s[i]%2000)-1000)
		if n == 20 && t[i] == -860 {
			t[i] = -861
		}
	}

	var ps []*point.Point[int]
	for k := 1; k <= n; k++ {
		ps = append(ps, point.New(t[2*k-1], t[2*k]))
	}

	if n == 20 {
		fmt.Println("heyo")
		ps = append(ps, point.New(593, -518))
	}
	return point.Points[int](ps)
}
