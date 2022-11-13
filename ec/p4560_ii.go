package eulerchallenge

import (
	"fmt"

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

	good1, good2 int
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
			reg := &dn{f, 0, 0, q, dnID, nil, 0, 0}
			dnID++
			op := &dn{f, 0, 0, oq, dnID, reg, 0, 0}
			dnID++
			reg.opposite = op
			m[q][f.String()] = reg
			m[oq][f.String()] = op
		}
		curD := m[q][f.String()]
		curD.cnt++
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

	qk(dns)
	plot456(pts, nil)
	return 1

	max := dns[len(dns)-1].cum

	// Now do every pair
	triCnt := 0
	for i, dn1 := range dns {
		if dn1.quad >= 2 {
			break
		}
		if dn1.f.N == 0 && dn1.quad > 0 {
			// TODO: Remove this check (taken care of by quad)
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

			v := dn1.cnt * dn2.cnt * d3cnt

			triCnt += v
			qk(dns)
			plot456(pts, dns)
			return 1
			goto AHA
		}
	}
AHA:

	//
	return triCnt
}

func plot456(pts []*point.Point[int], dns []*dn) {
	p := point.NewPlot()
	if err := p.Add(point.Points[int](pts)); err != nil {
		panic(fmt.Sprintf("NO: %V", err))
	}

	for _, pt := range pts {
		p.Add(point.NewLineSegment(point.Origin[int](), pt))
		p.Add(point.NewLineSegment(point.Origin[int](), point.New(-pt.X, -pt.Y)))
	}

	p.Add(point.Axes(-16000, 16000))

	/*for _, dn := range dns {
		/*x, y := 1, 1
		if dn.quad == 0 || dn.quad == 3 {
			x = -1
		}
		if dn.quad == 2 || dn.quad == 3 {
			y = -1
		}* /
		p.Add(point.NewLineSegment(point.Origin[int](), point.New(dn.f.D, dn.f.N)))
		p.Add(point.NewLineSegment(point.Origin[int](), point.New(-dn.f.D, -dn.f.N)))
	}*/

	//p.Add(point.Axes(15000, 15000))
	if len(pts) < 40 {
		fmt.Println("PLOTTING")
		p.Save(800, 800, "456.png")
	}
}

func qk(dns []*dn) {
	k, bs, cs := 0, 0, 0
	var firstQ2 *dn
	for _, d := range dns {
		if d.quad < 2 {
			bs += d.cnt
		} else {
			if firstQ2 == nil {
				firstQ2 = d
			}
			cs += d.cnt
		}
	}

	for _, b := range dns {
		if b.quad == 2 {
			break
		}
		k += b.cnt * ((b.opposite.cum - b.opposite.cnt) - firstQ2.cum + firstQ2.cnt)
	}

	fmt.Println("START", k, bs, cs)
	triCnt := 0
	for i, d := range dns {
		if i%10000 == 0 {
			fmt.Println("I", i)
		}
		if d.quad > 1 {
			break
		}
		if d.cnt > 0 && d.opposite.cnt == 0 {
			bs -= d.cnt
			triCnt += d.cnt * k
		} else if d.cnt == 0 && d.opposite.cnt > 0 {
			//fmt.Println("OP", d)
			// No longer a 'C'
			k -= d.opposite.cnt * bs // remove triangles
			cs -= d.opposite.cnt

			// Now a b
			k += d.opposite.cnt * cs
			bs += d.opposite.cnt
			//fmt.Println("M", k, cs, bs)
			//return
		} else {
			bs -= d.cnt
			k -= d.opposite.cnt * bs // remove triangles
			cs -= d.opposite.cnt

			triCnt += d.cnt * k

			k += d.opposite.cnt * cs
			bs += d.opposite.cnt
		}
	}
	fmt.Println("TRICNT", triCnt)
}

func qk2(dns []*dn) {

	k, bs, cs := 0, 0, 0

	first := dns[0]
	op1 := first.opposite
	var inCs bool
	for _, dn2 := range dns[1:] {
		if dn2.id == op1.id {
			inCs = true
			bs += dn2.cnt
			continue
		}

		// Increment b and c counts
		if !inCs {
			bs += dn2.cnt
		} else {
			cs += dn2.cnt
		}
	}

	var stopIt bool
	max := dns[len(dns)-1].cum
	for j := 1; !stopIt && j < len(dns); j++ {
		dn2 := dns[j]
		stopIt = dn2.id == op1.id
		if dn2.cnt == 0 {
			continue
		}
		if stopIt { //&& dn2.quad >= 2 {
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
		fmt.Println("HAHA", dn2.cnt*d3cnt)
		k += dn2.cnt * d3cnt
	}
	triCnt := first.cnt * k

	/*collinear := 0

	cur := dns[0]*/
	/*for _, dn2 := range dns[1:] {
		if dn2.id == cur.opposite.id {
			collinear += dn2.cnt
			break
		}

		if dn2.quad >= 2 {
			break
		}

		if dn2.quad < 2 {
			bs += dn2.cnt
		} else {
			cs += dn2.cnt
		}
	}*/

	fmt.Println("init", triCnt, first, k, bs, cs)
	//return

	poppedCs := 0

	bs += first.opposite.cnt
	cs -= first.opposite.cnt

	for _, dn1 := range dns[1:] {
		if dn1.quad >= 2 {
			break
		}

		if dn1.cnt > 0 && dn1.opposite.cnt == 0 {
			triCnt += dn1.cnt * k
			bs -= dn1.cnt
			fmt.Println("HELLO")
			/*k -= dn1.cnt * poppedCs
			triCnt += dn1.cnt * k
			bs -= dn1.cnt*/
		} else if dn1.cnt == 0 && dn1.opposite.cnt > 0 {
			//poppedCs += dn1.opposite.cnt
			cs -= dn1.opposite.cnt
			k -= dn1.opposite.cnt * bs

			bs += dn1.opposite.cnt
			k += dn1.opposite.cnt * cs

		} else if dn1.cnt > 0 && dn1.opposite.cnt > 0 {
			panic("not Yet")
			cs -= dn1.opposite.cnt
			k -= dn1.opposite.cnt * bs

			k -= dn1.cnt * poppedCs
			triCnt += dn1.cnt * k
			bs -= dn1.cnt

			bs += dn1.opposite.cnt
			k += dn1.opposite.cnt * cs
		} else {
			panic("AAAHAHAH")
		}
		fmt.Println(triCnt, dn1, k, bs, cs)
	}

	fmt.Println("QK", triCnt)

	/*for i := 0

	for _, dn2 := range dns {

	}*/
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

		//ps = append(generatePoints456(n), point.New(3370, -2510))
		ps = generatePoints456(n)

		dnzo(ps)
		//b := 0
		//b, _ = brute456(ps)
		//fmt.Println("B", b)

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
