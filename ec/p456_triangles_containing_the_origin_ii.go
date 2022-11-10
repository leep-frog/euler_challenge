package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/slices"
)

type square456 struct {
	points []*point.Point2D[int]
	ch     *point.ConvexHull2D[int]

	id int

	cornersCalculated bool
	minX              int
	minY              int
	maxX              int
	maxY              int

	actualMinX int
	actualMinY int
	actualMaxX int
	actualMaxY int
}

func (s *square456) addPoint(x, y int) {
	s.points = append(s.points, point.New2D(x, y))
}

func (s *square456) contains(p *point.Point2D[int]) bool {
	if s.ch == nil {
		s.ch = point.ConvexHull2DFromPoints(
			point.New2D(s.actualMinX, s.actualMinY),
			point.New2D(s.actualMinX, s.actualMaxY),
			point.New2D(s.actualMaxX, s.actualMinY),
			point.New2D(s.actualMaxX, s.actualMaxY),
		)
	}
	return s.ch.Contains(p)
}

func (s *square456) corners() []*point.Point2D[int] {
	if !s.cornersCalculated {
		s.minX, s.minY = s.points[0].X, s.points[0].Y
		s.maxX, s.maxY = s.points[0].X, s.points[0].Y
		for _, p := range s.points {
			s.minX = maths.Min(s.minX, p.X)
			s.maxX = maths.Max(s.maxX, p.X)
			s.minY = maths.Min(s.minY, p.Y)
			s.maxY = maths.Max(s.maxY, p.Y)
		}
	}

	return []*point.Point2D[int]{
		point.New2D(s.minX, s.minY),
		point.New2D(s.minX, s.maxY),
		point.New2D(s.maxX, s.minY),
		point.New2D(s.maxX, s.maxY),
	}
}

func (s *square456) cornerTriangles(sq2, sq3 *square456) []*point.Triangle2D[int] {
	var r []*point.Triangle2D[int]
	for _, c1 := range s.corners() {
		for _, c2 := range sq2.corners() {
			for _, c3 := range sq3.corners() {
				if !c1.Eq(c2) && !c1.Eq(c3) && !c2.Eq(c3) {
					r = append(r, point.NewTriangle2D(c1, c2, c3))
				}
			}
		}
	}
	return r
}

func (s *square456) triangles(sq2, sq3 *square456) []*point.Triangle2D[int] {
	var r []*point.Triangle2D[int]
	for _, c1 := range s.points {
		for _, c2 := range sq2.points {
			for _, c3 := range sq3.points {
				if !c1.Eq(c2) && !c1.Eq(c3) && !c2.Eq(c3) {
					r = append(r, point.NewTriangle2D(c1, c2, c3))
				}
			}
		}
	}
	return r
}

/*func (s *square456) contains(x, y int) bool {
	return x >= s.actualMinX && x < s.actualMaxX && y >= s.actualMinY && y < s.actualMaxY
}

func (s *square456) corners() [][]int {
	if s.minX == s.maxX {
		if s.minY == s.maxY {
			return [][]int{
				{s.minX, s.minY},
			}
		}
		return [][]int{
			{s.minX, s.minY},
			{s.minX, s.maxY},
		}
	}

	if s.minY == s.maxY {
		return [][]int{
			{s.minX, s.minY},
			{s.maxX, s.minY},
		}
	}

	return [][]int{
		{s.minX, s.minY},
		{s.minX, s.maxY},
		{s.maxX, s.minY},
		{s.maxX, s.maxY},
	}
}*/

/*func (s *square456) twoCornerTriangles(sq2 *square456) []*triangle456 {
	var r []*triangle456
	s1Corners := s.corners()
	for i, c1 := range s1Corners {
		for j := i + 1; j < len(s1Corners); j++ {
			c2 := s1Corners[j]
			for _, c3 := range sq2.corners() {
				r = append(r, &triangle456{
					point.New(c1[0], c1[1], 0),
					point.New(c2[0], c2[1], 0),
					point.New(c3[0], c3[1], 0),
				})
			}
		}
	}
	return r
}*/

/*func (s *square456) convexHull(sq2, sq3 *square456) *point.ConvexHull2D[int] {
	points := append(append(s.points, sq2.points...), sq3.points...)
	return point.ConvexHull2DFromPoints(points...)
	/*var r []*triangle456
	for _, c1 := range s.corners() {
		for _, c2 := range sq2.corners() {
			if c1[0] == c2[0] && c1[1] == c2[1] {
				continue
			}
			for _, c3 := range sq3.corners() {
				if c1[0] == c3[0] && c1[1] == c3[1] {
					continue
				}
				if c3[0] == c2[0] && c3[1] == c2[1] {
					continue
				}
				r = append(r, &triangle456{
					point.New(c1[0], c1[1], 0),
					point.New(c2[0], c2[1], 0),
					point.New(c3[0], c3[1], 0),
				})
			}
		}
	}
	return r
}*/

func (s *square456) String() string {
	return fmt.Sprintf("id=%d: (%d,%d) to (%d,%d)", s.id, s.actualMinX, s.actualMinY, s.actualMaxX, s.actualMaxY)
}

/*func (s *square456) EffectiveString() string {
	return fmt.Sprintf("id=%d: (%d,%d) to (%d,%d)", s.id, s.minX, s.minY, s.maxX, s.maxY)
}*/

func (s *square456) pointTriangles(sq2, sq3 *square456) []*point.Triangle2D[int] {
	var r []*point.Triangle2D[int]
	for _, c1 := range s.points {
		for _, c2 := range sq2.points {
			for _, c3 := range sq3.points {
				r = append(r, point.NewTriangle2D(c1, c2, c3))
			}
		}
	}
	return r
}

func (s *square456) twoPointTriangles(sq2 *square456) []*point.Triangle2D[int] {
	var r []*point.Triangle2D[int]
	for i, c1 := range s.points {
		for j := i + 1; j < len(s.points); j++ {
			c2 := s.points[j]
			for _, c3 := range sq2.points {
				r = append(r, point.NewTriangle2D(c1, c2, c3))
			}
		}
	}
	return r
}

/*func originBetween(x1, y1, x2, y2 int) bool {
	// true if origin is a linear combination of the provided points
	if x1 == 0 && x2 == 0 {
		return y1 == 0 || y2 == 0 || ((y1 > 0) != (y2 > 0))
	}
	if x1 == 0 {
		return y1 == 0
	}
	if x2 == 0 {
		return y2 == 0
	}

	if y1 == 0 && y2 == 0 {
		return x1 == 0 || x2 == 0 || ((x1 > 0) != (x2 > 0))
	}
	if y1 == 0 {
		return x1 == 0
	}
	if y2 == 0 {
		return x2 == 0
	}

	// Nothing equals zero

	// y = mx + b (true if b == 0)
	// y1 = m*x1,  y2 = m*x2
	// y1/x1 = y2/x2
	// y1*x2 = y2*x1
	return y1*x2 == y2*x1 && ((x1 > 0) != (x2 > 0)) && ((y1 > 0) != (y2 > 0))
}*/

/*func (t *triangle456) containsTheOrigin() bool {
	// First check if the origin lies on an edge
	btwnAB := originBetween(t.a.X, t.a.Y, t.b.X, t.b.Y)
	btwnBC := originBetween(t.c.X, t.c.Y, t.b.X, t.b.Y)
	btwnAC := originBetween(t.a.X, t.a.Y, t.c.X, t.c.Y)
	if btwnAB || btwnBC || btwnAC {
		return true
	}

	// Then check that the origin falls on the same side of every line.
	ab := simpleCross(t.b.Minus(t.a), (point.Origin().Minus(t.a))) > 0
	bc := simpleCross(t.c.Minus(t.b), (point.Origin().Minus(t.b))) > 0
	ca := simpleCross(t.a.Minus(t.c), (point.Origin().Minus(t.c))) > 0
	// The origin is on the same side of every line.
	return ab == bc && ab == ca
}*/

func generatePoints(n int) [][]int {
	var points [][]int
	xp, yp := 1, 1
	for i := 0; i < n; i++ {
		xp = (xp * 1248) % 32323
		yp = (yp * 8421) % 30103
		points = append(points, []int{xp - 16161, yp - 15051})
	}
	return points
}

func elegant456(points [][]int) (int, []*point.Triangle2D[int]) {
	var xBuckets, yBuckets []int
	splits := 8
	for i := 0; i <= splits; i++ {
		xBuckets = append(xBuckets, (i*32323/splits)-16161)
		yBuckets = append(yBuckets, (i*30103/splits)-15051)
	}

	var squares []*square456
	for xi, xBucket := range xBuckets[:splits] {
		// xOffset should be 1 for last bucket
		xOffset := xi / (splits - 1)
		for yi, yBucket := range yBuckets[:len(yBuckets)-1] {
			yOffset := yi / (splits - 1)
			squares = append(squares, &square456{
				id:         xi*splits + yi,
				actualMinX: xBucket,
				actualMaxX: xBuckets[xi+1] + xOffset,
				actualMinY: yBucket,
				actualMaxY: yBuckets[yi+1] + yOffset,
			})
		}
	}

	// TODO: this can be improved with sorting
	for _, p := range points {
		var added bool
		for _, sq := range squares {
			sq.contains(point.New2D(p[0], p[1]))
			if sq.contains(point.New2D(p[0], p[1])) {
				added = true
				sq.addPoint(p[0], p[1])
				break
			}
		}
		if p[0] == 0 && p[1] == 0 {
			return -1, nil
		}
		if !added {
			fmt.Println(p)
			return -2, nil
		}
	}

	var filteredSquares []*square456
	for _, sq := range squares {
		if len(sq.points) > 0 {
			filteredSquares = append(filteredSquares, sq)
		}
	}
	squares = filteredSquares

	originTriangleCount := 0
	r := []*point.Triangle2D[int]{}
	for i, sq1 := range squares {
		fmt.Println("SQ1", i)
		for j := i; j < len(squares); j++ {
			//fmt.Println("SQ2", j)
			sq2 := squares[j]
			for k := j; k < len(squares); k++ {
				if i == j && j == k {
					continue
				}
				sq3 := squares[k]

				originCount := 0
				cornerTris := sq1.cornerTriangles(sq2, sq3)
				if len(cornerTris) == 0 {
					// Two of the squares are the same and the square has only one point.
					continue
				}
				for _, t := range cornerTris {
					if t.Contains(point.Origin2D[int]()) {
						originCount++
					} else if originCount > 0 {
						break
					}
				}

				// All triangles contain the origin
				if originCount == len(cornerTris) {
					if i == j || j == k || i == k {
						fmt.Println("bleh", i, j, k)
					}
					originTriangleCount += len(sq1.points) * len(sq2.points) * len(sq3.points)
					//fmt.Println("OTC", sq1.id, sq2.id, sq3.id, "|", len(cornerTris), len(sq1.points)*len(sq2.points)*len(sq3.points), len(sq1.triangles(sq2, sq3)))
					g := map[string]bool{}
					for _, t := range sq1.triangles(sq2, sq3) {
						if g[t.String()] {
							continue
						}
						g[t.String()] = true
						r = append(r, t)
					}
					continue
				}

				mightContain := originCount != 0
				if !mightContain {
					ch := point.ConvexHull2DFromPoints(append(append(sq1.points, sq2.points...), sq3.points...)...)
					mightContain = ch.Contains(point.Origin2D[int]())
				}

				if !mightContain {
					continue
				}

				tc := 0
				g := map[string]bool{}
				for _, t := range sq1.triangles(sq2, sq3) {
					if g[t.String()] {
						continue
					}
					g[t.String()] = true
					got := map[string]bool{}
					if got[t.String()] {
						continue
					}
					got[t.String()] = true
					if t.Contains(point.Origin2D[int]()) {
						tc++
						r = append(r, t)
					}
				}
				if i == j || j == k || i == k {
					//tc /= 2
				}
				originTriangleCount += tc
			}
		}
	}

	// Now check cases where there are two points in the same square
	return originTriangleCount, r
}

func uniqTris(ts []*point.Triangle2D[int]) []string {
	m := map[string]bool{}
	var r []string
	for _, t := range ts {
		if !m[t.String()] {
			r = append(r, t.String())
			m[t.String()] = true
		}
	}
	slices.Sort(r)
	return r
}

func P456() *problem {
	return intInputNode(456, func(o command.Output, n int) {

		points := generatePoints(n)

		points = points[:140]
		//fmt.Println("LAST POINT", points[len(points)-1])

		/*points = [][]int{
			{-1691, 12703},
			{-3392, 11910},
			{2799, -12852},
			{1178, -11583},
			{3766, -11954},
		}*/

		/*points = [][]int{
			{-16054, 5971},
			{-11917, 5571},
			{-7405, -4902},
			{-6558, -4433},
			{6429, -2875},
			{11702, 8707},
			{15181, 9283},
			{15588, 11540},
		}*/

		fmt.Println("EC")
		eCnt, ets := elegant456(points)
		fmt.Println("BC")
		bCnt, bts := brute456(points)
		_ = bts
		_ = ets

		/*if diff := cmp.Diff(uniqTris(ets), uniqTris(bts)); diff != "" {
			fmt.Printf("Yes diff:\n%s", diff)
		}*/

		fmt.Println("ECounts:", eCnt, len(ets), len(uniqTris(ets)))
		fmt.Println("BCounts:", bCnt, len(bts), len(uniqTris(bts)))
		//o.Stdoutln(eCnt, bCnt)
	}, []*execution{
		/*{
			args: []string{"8"},
			want: "20",
		}, /**/
		{
			args: []string{"600"},
			want: "8950634",
		},
		/*{
			args: []string{"40000"},
			want: "2666610948988",
		},
		/*{
			args: []string{"2000000"},
			want: "",
		},*/
	})
}

func brute456(points [][]int) (int, []*point.Triangle2D[int]) {
	cnt := 0
	r := []*point.Triangle2D[int]{}
	for i, p1 := range points {
		fmt.Println(i)
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			for k := j + 1; k < len(points); k++ {
				p3 := points[k]
				t := point.NewTriangle2D(
					point.New2D(p1[0], p1[1]),
					point.New2D(p2[0], p2[1]),
					point.New2D(p3[0], p3[1]),
				)
				if t.Contains(point.Origin2D[int]()) {
					cnt++
					r = append(r, t)
				}
			}
		}
	}
	return len(r), r
}
