package eulerchallenge

import (
	"fmt"
	"math"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/slices"
)

type sq struct {
	subSqs []*sq
	// the hull at each depth
	//hull   []*point.ConvexHull[int]
	points []*point.Point[int]

	box *point.Rectangle[int]
}

func newSq(minX, minY, maxX, maxY int) *sq {
	return &sq{nil, nil, point.NewRectangle(minX, minY, maxX, maxY)}
}

func (s *sq) String() string {
	return s.box.String()
}

func (s *sq) originBetween(b *sq) bool {
	return point.ConvexHullFromPoints(append(s.points, b.points...)...).Contains(point.Origin[int]())
}

func (a *sq) originInTriangle(b, c *sq) bool {
	return point.ConvexHullFromPoints(append(append(a.points, b.points...), c.points...)...).Contains(point.Origin[int]())
}

func (s *sq) size() int {
	return len(s.points)
}

func (s *sq) split() []*sq {
	if len(s.subSqs) == 0 {
		xMid := (s.box.MinX + s.box.MaxX) / 2
		yMid := (s.box.MinY + s.box.MaxY) / 2
		s.subSqs = []*sq{
			newSq(s.box.MinX, s.box.MinY, xMid, yMid),
			newSq(s.box.MinX, yMid, xMid, s.box.MaxY),
			newSq(xMid, s.box.MinY, s.box.MaxX, yMid),
			newSq(xMid, yMid, s.box.MaxX, s.box.MaxY),
		}

		populate(s.points, s.subSqs)
	}
	return s.subSqs
}

const (
	thsh  = 500
	thsh2 = 50_000
)

func calcTrisTwo(a, b *sq) int {
	// TODO: Cache originBetween (by ID)
	if !a.originBetween(b) {
		return 0
	}

	if a.size()*(a.size()-1)*b.size() < thsh2 {
		cnt := 0
		for i, pa := range a.points {
			for j := i + 1; j < len(a.points); j++ {
				pb := a.points[j]
				for _, pc := range b.points {
					if point.NewTriangle(pa, pb, pc).Contains(point.Origin[int]()) {
						cnt++
					}
				}
			}
		}
		return cnt
	}

	sum := 0
	splt := a.split()
	for i, subA := range splt {
		for j := i + 1; j < len(splt); j++ {
			subB := splt[j]
			sum += calcTris(subA, subB, b)
			sum += calcTrisTwo(subA, b)
			sum += calcTrisTwo(subB, b)
		}
	}
	return sum
}

func calcTris(a, b, c *sq) int {

	if a.size()*b.size()*c.size() < thsh2 {
		cnt := 0
		for _, pa := range a.points {
			for _, pb := range b.points {
				for _, pc := range c.points {
					if point.NewTriangle(pa, pb, pc).Contains(point.Origin[int]()) {
						cnt++
					}
				}
			}
		}
		//fmt.Println("OKAY", cnt, a, b, c)
		//fmt.Println("LNS", a.size(), b.size(), c.size())
		return cnt
	}

	if a.size() == 0 || b.size() == 0 || c.size() == 0 {
		panic(fmt.Sprintln("WAIT WHAT"))
		return 0
	}

	// Only split them if they're big enough to care
	if a.originBetween(b) {
		//fmt.Println("SPLIT A")
		sum := 0
		for _, subA := range a.split() {
			for _, subB := range b.split() {
				sum += calcTris(subA, subB, c)
			}
		}
		return sum
	}

	if a.originBetween(c) {
		//fmt.Println("SPLIT B")
		sum := 0
		for _, subA := range a.split() {
			for _, subC := range c.split() {
				sum += calcTris(subA, b, subC)
			}
		}
		return sum
	}

	if b.originBetween(c) {
		//fmt.Println("SPLIT C")
		sum := 0
		for _, subB := range b.split() {
			for _, subC := range c.split() {
				sum += calcTris(a, subB, subC)
			}
		}
		return sum
	}

	// If we are here then no things contain the origin
	if a.originInTriangle(b, c) {
		return a.size() * b.size() * c.size()
	}

	return 0
}

func populate(pts []*point.Point[int], sqs []*sq) {
	for _, p := range pts {
		added := false
		for _, s := range sqs {
			if s.box.Contains(p) {
				s.points = append(s.points, p)
				added = true
				break
			}
		}
		if !added {
			panic(fmt.Sprintln("ARGHY", p, sqs))
		}
	}
}

type dn struct {
	f        *fraction.Fraction[int]
	cnt      int
	cum      int
	quad     int
	id       int
	opposite *dn
	pts      []*point.Point[int]
}

func (d *dn) String() string {
	return fmt.Sprintf("{(%d): %d, %d, %d, %v}", d.id, d.quad, d.cnt, d.cum, d.f)
}

func (this *dn) LT(that *dn) bool {
	if this.quad != that.quad {
		return this.quad < that.quad
	}

	if this.f.D == 0 || that.f.D == 0 {
		panic("NOOOO")
	}

	switch this.quad {
	case 0, 2:
		return !this.f.LT(that.f)
	case 1, 3:
		return !this.f.LT(that.f)
	}

	panic("NOPE")
}

func dnzo(pts []*point.Point[int]) int {
	dnID := 0
	//fraction.New()
	/*m := map[int]map[string]*dn{
		0: {},
		1: {},
		2: {},
		3: {},
	}*/
	// Quadrant to fraction
	m := []map[string]*dn{
		{},
		{},
		{},
		{},
	}

	primes := generator.Primes()
	for _, p := range pts {
		//uf := fraction.New(p.Y, p.X)
		f := fraction.Simplify(p.Y, p.X, primes)
		if p.X == 0 {
			// Set the fraction to the guaranteed highest slope (simulate infinite slope)
			f = fraction.New(17_000, 1)
		}

		q := p.Quadrant()
		oq := (q + 2) % 4
		if m[q][f.String()] == nil {
			reg := &dn{f, 0, 0, q, dnID, nil, nil}
			dnID++
			op := &dn{f, 0, 0, oq, dnID, reg, nil}
			dnID++
			reg.opposite = op
			m[q][f.String()] = reg
			m[oq][f.String()] = op
		}
		curD := m[q][f.String()]
		curD.cnt++
		curD.pts = append(curD.pts, p)
	}

	// Now sort
	var dns []*dn
	for _, sd := range m {
		for _, d := range sd {
			dns = append(dns, d)
		}
	}
	slices.SortFunc(dns, func(this, that *dn) bool { return this.LT(that) })

	// Cumulate
	sum := 0
	for _, dn := range dns {
		sum += dn.cnt
		dn.cum = sum
	}

	max := dns[len(dns)-1].cum
	fmt.Println("LDN", len(dns))

	// Now do every pair
	triCnt := 0
	for i, dn1 := range dns {
		if dn1.quad >= 2 {
			break
		}
		if dn1.f.N == 0 && dn1.quad > 0 {
			// TODO: Remove this check (taken care of by quad)
			fmt.Println("OOPS", i)
			break
		}
		if dn1.cnt == 0 {
			continue
		}
		op1 := dn1.opposite

		var stopIt bool
		for j := i + 1; !stopIt && j < len(dns); j++ {
			dn2 := dns[j]
			stopIt = dn2.id == op1.id
			if dn2.cnt == 0 {
				continue
			}
			if stopIt && dn2.quad >= 2 {
				break
			}

			op2 := dn2.opposite

			d3cnt := op2.cum - op1.cum - op2.cnt
			if dn2.quad >= 2 {
				d3cnt = max - op1.cum
			}
			//second += op1.cnt

			//fmt.Println("TCP", dn1, dn2, o1, o2, second, first)

			/*if second-first <= 0 {
				fmt.Println("NOOp", second, first, max)
			}*/
			v := dn1.cnt * dn2.cnt * d3cnt

			triCnt += v
		}
	}

	p := point.NewPlot()
	if err := p.Add(point.Points[int](pts)); err != nil {
		panic(fmt.Sprintf("NO: %V", err))
	}

	for _, dn := range dns {
		/*x, y := 1, 1
		if dn.quad == 0 || dn.quad == 3 {
			x = -1
		}
		if dn.quad == 2 || dn.quad == 3 {
			y = -1
		}*/
		p.Add(point.NewLineSegment(point.Origin[int](), point.New(dn.f.D, dn.f.N)))
		p.Add(point.NewLineSegment(point.Origin[int](), point.New(-dn.f.D, -dn.f.N)))
	}

	//p.Add(point.Axes(15000, 15000))
	p.Save(800, 800, "456.png")

	fmt.Println("DNZO", triCnt, math.MaxInt)
	//
	return triCnt
}

func P4560() *problem {
	return intInputNode(4560, func(o command.Output, n int) {

		ps := []*point.Point[int]{
			point.New(0, 3),
			point.New(0, -3),
			point.New(3, 0),
			point.New(-3, 0),
		}
		/**/

		ps = generatePoints456(n)

		dnzo(ps)
		b := 0
		//b, _ = brute456(ps)
		fmt.Println("B", b)

		return
		for i := 0; ; i++ {
			if i%1_000_000_000 == 0 {
				fmt.Println(i)
			}
		}
		pts := generatePoints456(n)

		var minX, minY, maxX, maxY int
		for _, p := range pts {
			minX = maths.Min(minX, p.X)
			minY = maths.Min(minY, p.Y)
			maxX = maths.Max(maxX, p.X)
			maxY = maths.Max(maxY, p.Y)
		}

		sqs := []*sq{
			newSq(minX, minY, 0, 0),
			newSq(minX, 0, 0, maxY),
			newSq(0, 0, maxX, maxY),
			newSq(0, minY, maxX, 0),
		}

		//fmt.Println("MASTPOP")
		populate(pts, sqs)

		sum := 0
		for i, a := range sqs {
			fmt.Println("SQ", i)
			for j := i + 1; j < len(sqs); j++ {
				b := sqs[j]
				for k := j + 1; k < len(sqs); k++ {
					c := sqs[k]
					sum += calcTris(a, b, c)
				}
			}
		}

		for _, a := range sqs {
			for _, b := range sqs {
				if a.box.Eq(b.box) {
					continue
				}

				sum += calcTrisTwo(a, b)
			}
		}
		fmt.Println(sum)

		// TODO: Plot
		point.CreatePlot("456.png", 800, 800, point.Points[int](pts), point.Axes(-33333, 33333))
		//fmt.Println(pts)
	}, []*execution{
		{
			args: []string{"8"},
		},
		/*{
			args: []string{"8"},
		},
		/*{
			args: []string{"40000"},
		},*/
	})
}
