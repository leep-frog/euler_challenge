package p456

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/point"
	"github.com/leep-frog/functional"
)

/* Solution: If we have two points, draw a line that intersects those points
and the origin. The area on the other side of the unit circle contains all
slopes that make a triangle that contains the origin.

This runs in O(n) time by starting with a line from radian=0 to radian=360
and keeping track of the number of triangles that can be made with the set of
'B' points and 'C' points (where all triangles made are A-B-C and sorted by radians
such that radians(A) < radians(B) < radians(C))

Note: interior of triangle does not include boundaries (as problem 184 demonstrates)
*/

func generatePoints456(n int) []*point.Point[int] {
	var r []*point.Point[int]
	xp, yp := 1, 1
	for i := 0; i < n; i++ {
		// Create point
		xp = (xp * 1248) % 32323
		yp = (yp * 8421) % 30103
		r = append(r, point.New(xp-16161, yp-15051))
	}
	return r
}

func generateSlopeGroups(pts []*point.Point[int]) []*slopeGroup {
	// Map (actually list) from quadrant to slope to slope group the point belongs.
	m := []map[string]*slopeGroup{
		{},
		{},
		{},
		{},
	}

	primes := generator.Primes()
	sgID := 0
	for _, p := range pts {

		// Simplify slope
		f := fraction.Simplify(p.Y, p.X, primes)
		if p.X == 0 {
			// Set the fraction to the guaranteed highest slope (simulate infinite slope)
			f = fraction.New(17_000, 1)
		}

		q := p.Quadrant()
		oq := (q + 2) % 4
		if m[q][f.String()] == nil {
			reg := &slopeGroup{f, 0, 0, q, sgID, nil}
			sgID++
			op := &slopeGroup{f, 0, 0, oq, sgID, reg}
			sgID++
			reg.opposite = op
			m[q][f.String()] = reg
			m[oq][f.String()] = op
		}
		m[q][f.String()].cnt++
	}

	// Now sort slopeGroups by radians
	var sgs []*slopeGroup
	for _, sd := range m {
		for _, d := range sd {
			sgs = append(sgs, d)
		}
	}
	functional.SortFunc(sgs, func(this, that *slopeGroup) bool { return this.LT(that) })

	// Set cumulative number of points between (-1, 0) and each slope group.
	sum := 0
	for _, slopeGroup := range sgs {
		sum += slopeGroup.cnt
		slopeGroup.cum = sum
	}

	return sgs
}

type slopeGroup struct {
	slopeFrac *fraction.Fraction
	cnt       int
	cum       int
	quadrant  int
	id        int
	opposite  *slopeGroup
}

// Comparing radians of slopeGroups.
func (this *slopeGroup) LT(that *slopeGroup) bool {
	if this.quadrant != that.quadrant {
		return this.quadrant < that.quadrant
	}

	if this.slopeFrac.D == 0 || that.slopeFrac.D == 0 {
		panic("Unexpected denominator")
	}

	return !this.slopeFrac.LT(that.slopeFrac)
}

func OriginTriangles456(pts []*point.Point[int]) int {
	// Generate point slopes
	sgs := generateSlopeGroups(pts)

	// Compute the number of bs and cs (where every triangle made is a-b-c ordered by radians)
	bs, cs := 0, 0
	var firstQ2 *slopeGroup
	for _, d := range sgs {
		if d.quadrant < 2 {
			bs += d.cnt
		} else {
			if firstQ2 == nil {
				firstQ2 = d
			}
			cs += d.cnt
		}
	}

	// k is the number of triangles that can be made when the first point is on the vector (-1, 0)
	k := 0
	for _, b := range sgs {
		if b.quadrant == 2 {
			break
		}
		k += b.cnt * ((b.opposite.cum - b.opposite.cnt) - firstQ2.cum + firstQ2.cnt)
	}

	// Iterate through slopeGroups, incrementing the number of triangles that can be made (k)
	// when the first point of the triangle, A, is on the slope group.
	triCnt := 0
	for _, d := range sgs {
		if d.quadrant > 1 {
			break
		}

		// The point on the slope group is no longer in the set of bs
		bs -= d.cnt

		// The opposite slope group is no longer in cs. Remove the cnt and the
		// number of triangles it makes as a C point.
		k -= d.opposite.cnt * bs
		cs -= d.opposite.cnt

		// Incrememnt the number of triangles that can be made with the slopeGroup as A
		triCnt += d.cnt * k

		// Move the opposite slope group into the B group, and add the triangles
		// that can be made with it as a B point.
		k += d.opposite.cnt * cs
		bs += d.opposite.cnt
	}
	return triCnt
}

func P456() *ecmodels.Problem {
	return ecmodels.IntInputNode(456, func(o command.Output, n int) {
		o.Stdoutln(OriginTriangles456(generatePoints456(n)))
	}, []*ecmodels.Execution{
		{
			Args: []string{"8"},
			Want: "20",
		},
		{
			Args: []string{"600"},
			Want: "8950634",
		},
		{
			Args: []string{"40000"},
			Want: "2666610948988",
		},
		{
			Args:     []string{"2000000"},
			Want:     "333333208685971546",
			Estimate: 25,
		},
	})
}
