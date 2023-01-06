package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/slices"
)

// Previously had a line sweep algorithm that sorted points by x
// and kept track of all valid convex hulls, but that took > 10 minutes.

// This solution
// - gets all empty triangles (which was also needed for the previous solution),
// - constructs a slice of points sorted by x
// - constructs a set of slices of points sorted by y where the key is a pair of x values;
//   the points returned all within those two x values
// - Iterate over all sets of points keeping track of valid points (options) that
//   can still be added.

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

func getEmptyTriangles(n int) (point.Points[int], map[int]map[int]map[int]bool) {
	pts := generatePoints252(n)

	m := map[int]map[int]map[int]bool{}
	xSorted := bread.Copy(pts)

	slices.SortFunc(xSorted, func(p, q *point.Point[int]) bool {
		if p.X != q.X {
			return p.X < q.X
		}
		return p.Y < q.Y
	})

	// Map from xStartPoint.X to xEndPoint.X to list of points between the two inclusive
	ySortedByXs := map[int]map[int][]*point.Point[int]{}
	// Map from xStartPoint.X to xEndPoint.X to Y coordinate to index of that Y coordinate in ySortedByXs
	yIndicesSortedByXs := map[int]map[int]map[int]int{}
	prevX := xSorted[0].X - 1
	for xStart, xPt := range xSorted {
		if xPt.X == prevX {
			continue
		}
		prevX = xPt.X

		ySorted := []*point.Point[int]{xPt}
		for xEnd := xStart + 1; xEnd < len(xSorted); xEnd++ {
			ePt := xSorted[xEnd]
			ySorted = append(ySorted, ePt)
			cpy := bread.Copy(ySorted)
			slices.SortFunc(cpy, func(p, q *point.Point[int]) bool {
				if p.Y != q.Y {
					return p.Y < q.Y
				}
				return p.X < q.X
			})

			indices := map[int]int{}
			for i, p := range cpy {
				indices[p.Y] = i
			}
			maths.Insert(ySortedByXs, xPt.X, ePt.X, cpy)
			maths.Insert(yIndicesSortedByXs, xPt.X, ePt.X, indices)
		}
	}

	// Map from point to index
	xm := map[int]int{}
	for xi, xp := range xSorted {
		xm[xp.X] = xi
	}

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

				ySorted := ySortedByXs[minX][maxX]
				ym := yIndicesSortedByXs[minX][maxX]

				yStart := ym[minY]
				yEnd := ym[maxY]

				for _, p := range ySorted[yStart : yEnd+1] {
					if t.ContainsExclusive(p) {
						goto DONT_ADD_TRI
					}
				}
				maths.DeepInsert(m, i, j, k, true)
			DONT_ADD_TRI:
			}
		}
	}
	return xSorted, m
}

// Previously had a line
func rec252(options []int, points []*point.Point[int], triIndices []int, triMap map[int]map[int]map[int]bool, best *maths.Bester[int, float64]) {
	if len(options) == 0 {
		var pointSet []*point.Point[int]
		for _, index := range triIndices {
			pointSet = append(pointSet, points[index])
		}
		ch := point.ConvexHullFromPoints(pointSet...)
		best.Check(ch.Area())
	}

	for i, pointIdxA := range options {
		var newOptions []int
		for o, pointIdxB := range options {
			if o <= i {
				continue
			}

			valid := true
			for _, ti := range triIndices {
				indices := []int{pointIdxA, pointIdxB, ti}
				slices.Sort(indices)
				if !triMap[indices[0]][indices[1]][indices[2]] {
					valid = false
					break
				}
			}
			if valid {
				newOptions = append(newOptions, pointIdxB)
			}
		}
		rec252(newOptions, points, append(triIndices, pointIdxA), triMap, best)
	}
}

func P252() *problem {
	return intInputNode(252, func(o command.Output, n int) {

		pts, triangleIndexMap := getEmptyTriangles(n)

		var slopts []int
		for i := range pts {
			slopts = append(slopts, i)
		}

		best := maths.Largest[int, float64]()
		rec252(slopts, pts, nil, triangleIndexMap, best)
		o.Stdoutf("%.1f\n", best.Best())
	}, []*execution{
		{
			args: []string{"20"},
			want: "1049694.5",
		},
		{
			args:     []string{"500"},
			want:     "104924.0",
			estimate: 80,
		},
	})
}
