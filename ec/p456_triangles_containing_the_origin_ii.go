package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/point"
	"golang.org/x/exp/slices"
)

func generateDns(n int) []*dn {
	xp, yp := 1, 1

	dnID := 0
	m := []map[string]*dn{
		{},
		{},
		{},
		{},
	}

	primes := generator.Primes()
	for i := 0; i < n; i++ {
		// Create point
		xp = (xp * 1248) % 32323
		yp = (yp * 8421) % 30103
		p := point.New(xp-16161, yp-15051)

		// Simplify slope
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

	return dns
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

func qk(n int) int {
	dns := generateDns(n)

	// BREAKER
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

	triCnt := 0
	for _, d := range dns {
		if d.quad > 1 {
			break
		}
		if d.cnt > 0 && d.opposite.cnt == 0 {
			// No longer a 'B' point
			bs -= d.cnt
			triCnt += d.cnt * k
		} else if d.cnt == 0 && d.opposite.cnt > 0 {
			// No longer a 'C' point
			k -= d.opposite.cnt * bs
			cs -= d.opposite.cnt

			// Now a 'B' point
			k += d.opposite.cnt * cs
			bs += d.opposite.cnt
		} else {
			bs -= d.cnt
			k -= d.opposite.cnt * bs
			cs -= d.opposite.cnt

			triCnt += d.cnt * k

			k += d.opposite.cnt * cs
			bs += d.opposite.cnt
		}
	}
	return triCnt
}

func P456() *problem {
	return intInputNode(456, func(o command.Output, n int) {
		o.Stdoutln(qk(n))
	}, []*execution{
		{
			args: []string{"8"},
			want: "20",
		},
		{
			args: []string{"600"},
			want: "8950634",
		},
		{
			args: []string{"40000"},
			want: "2666610948988",
		},
		{
			args:     []string{"2000000"},
			want:     "333333208685971546",
			estimate: 25,
		},
	})
}
