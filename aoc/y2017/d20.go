package y2017

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/rgx"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Day20() aoc.Day {
	return &day20{}
}

type day20 struct{}

type particle struct {
	idx int
	// In order of px, py, pz
	ps []int
	// In order of vx, vy, vz
	vs []int
	// In order of ax, ay, az
	as []int
}

func (p *particle) String() string {
	return fmt.Sprintf("{idx=%d; ps=%v; vs=%v; as=%v}", p.idx, p.ps, p.vs, p.as)
}

func (p *particle) bruteIntersects(q *particle) []int {
	var ntrs map[int]bool
	for i := 0; i < 3; i++ {
		p_p, v_p, a_p := p.ps[i], p.vs[i], p.as[i]
		p_q, v_q, a_q := q.ps[i], q.vs[i], q.as[i]
		var ts []int
		for t := 0; t < 1_000; t++ {
			if p_p == p_q {
				ts = append(ts, t)
			}
			v_p += a_p
			v_q += a_q
			p_p += v_p
			p_q += v_q
		}
		if ntrs == nil {
			ntrs = maths.NewSimpleSet(ts...)
		} else {
			ntrs = maths.Intersection(ntrs, maths.NewSimpleSet(ts...))
		}
	}
	ks := maps.Keys(ntrs)
	return ks
}

/*
THIS DOES NOT WORK
Because the formulas for continuous acceleration don't apply easily
to systems with discrete acceleration changes
*/
func (p *particle) intersects(q *particle) []int {
	// Check if x intersects
	// Formula p_pt = p_p + v_p * t + a_p * t * t / 2
	// WAIT: above is formula for continuous acceleration. Since
	// the problem uses discrete formula, we need to use the following
	// term for acceleration: a_p * (t)*(t+1)/2
	//                      = a_p * (t^2 + t) / 2
	//                      = a_p * t / 2 + a_p * t*t / 2
	// p_pt = p_p + v_p * t + a_p * t * t / 2
	// p_pt = p_p + t * (v_p + a_p/2) + t^2 * (a_p / 2)

	// Check when x_t for p equals x_t q
	// p_p + t * (v_p + a_p/2) + t^2 * (a_p / 2) = p_q + t * (v_qp + a_q/2) + t^2 * (a_q / 2)
	// 0 = (p_p - p_q) + t * [(v_p + a_p/2)-(v_q + a_q/2)] + t*t * [(a_p / 2) - (a_q / 2)]
	// Multiple the whole thing by 2 to avoid fractions
	// 0 = 2*(p_p - p_q) + t * [(2*v_p + a_p)-(2*v_q + a_q)] + t*t * [a_p - a_q]
	// var ts map[int]bool
	var ntrs []map[int]bool
	for i := 2; i >= 0; i-- {
		// (-b += sqrt(b*b - 4 * a * c)) / 2 * a
		p_p, v_p, a_p := p.ps[i], p.vs[i], p.as[i]
		p_q, v_q, a_q := q.ps[i], q.vs[i], q.as[i]

		// c, b, a := 2*(p_q-p_p), 2*(v_q-v_p), (a_q - a_p)
		c, b, a := 2*(p_p-p_q), (2*v_p + a_p - 2*v_q - a_q), (a_p - a_q)

		if a == 0 && b == 0 {
			// Will always be at separate spot
			if c != 0 {
				return nil
			}
			// Otherwise, c is zero and they will be at the same spot forever
			continue
		}

		if a == 0 {
			// Then just a linear equation
			// 0 = (p_q - p_p) + (v_q - v_p) * t
			// t = (p_p - p_q) / (v_q - v_p)
			t := (p_p - p_q) / (v_q - v_p)
			// Make sure no fractions
			if ((p_q - p_p) + (v_q-v_p)*t) != 0 {
				return nil
			}
			ntrs = append(ntrs, map[int]bool{
				t: true,
			})
		} else {
			// Otherwise, a quadratic equation
			determinant := b*b - 4*a*c
			if determinant < 0 {
				return nil
			}
			if !maths.IsSquare(determinant) {
				return nil
			}

			// (-b += sqrt(b*b - 4 * a * c)) / 2 * a
			sq := maths.IntSquareRoot(determinant)
			ntrs = append(ntrs, map[int]bool{
				(-b + sq) / (2 * a): true,
				(-b - sq) / (2 * a): true,
			})
		}
	}

	var r []int
	for t := range maths.Intersection(ntrs...) {
		if t >= 0 {
			r = append(r, t)
		}
	}
	return r
}

type intersection struct {
	pi, qi, t int
}

func (d *day20) Solve(lines []string, o command.Output) {
	r := rgx.New(`^p=< *([-0-9]+), *([-0-9]+), *([-0-9]+)>, v=< *([-0-9]+), *([-0-9]+), *([-0-9]+)>, a=< *([-0-9]+), *([-0-9]+), *([-0-9]+)>$`)
	best := maths.Smallest[int, int]()
	var particles []*particle
	for i, line := range lines {
		m := r.MustMatch(line)
		px, py, pz := parse.Atoi(m[0]), parse.Atoi(m[1]), parse.Atoi(m[2])
		vx, vy, vz := parse.Atoi(m[3]), parse.Atoi(m[4]), parse.Atoi(m[5])
		ax, ay, az := parse.Atoi(m[6]), parse.Atoi(m[7]), parse.Atoi(m[8])

		best.IndexCheck(i, maths.Abs(ax)+maths.Abs(ay)+maths.Abs(az))

		np := &particle{
			i,
			[]int{px, py, pz},
			[]int{vx, vy, vz},
			[]int{ax, ay, az},
		}
		particles = append(particles, np)
	}

	// Calculate all of the intersections
	intersections := map[int][]*intersection{}
	for i, p := range particles {
		for _, q := range particles[i+1:] {
			// bts := p.bruteIntersects(q)
			ts := p.intersects(q)
			// slices.Sort(bts)
			// slices.Sort(ts)
			// if !slices.Equal(bts, ts) {
			// fmt.Println("MISMATCH", p, q)
			// fmt.Println(bts)
			// fmt.Println(ts)
			// return
			// }
			if len(ts) > 0 {
				for _, t := range ts {
					intersections[t] = append(intersections[t], &intersection{p.idx, q.idx, t})
				}
			}
		}
	}

	keys := maps.Keys(intersections)
	slices.Sort(keys)

	destroyed := map[int]bool{}
	for _, t := range keys {
		toDestory := map[int]bool{}
		for _, intr := range intersections[t] {
			if destroyed[intr.pi] || destroyed[intr.qi] {
				continue
			}
			// Both exist
			toDestory[intr.pi] = true
			toDestory[intr.qi] = true
		}

		for id := range toDestory {
			destroyed[id] = true
		}
	}

	o.Stdoutln(best.BestIndex(), len(particles)-len(destroyed))
}

func (d *day20) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"0 1",
			},
		},
		{
			ExpectedOutput: []string{
				"125 461",
			},
		},
	}
}
