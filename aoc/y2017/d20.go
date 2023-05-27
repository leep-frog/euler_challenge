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

func (p *particle) intersects(q *particle) []int {
	// Check if x intersects
	// Formula x_t = px + vx * t + ax * ax * t / 2
	// Check when x_t for p equals x_t q
	// px_q + t * vx_q + t^2 * ax_q / 2 = px_p + t * vx_p + t^2 * ax_p / 2
	// 0 = (px_q - px_p) + (vx_q - vx_p) * t + (ax_q - ax_p)/2 * t^2
	var ts map[int]bool
	for i := 0; i < 3; i++ {
		// (-b += sqrt(b*b - 4 * a * c)) / 2 * a
		px_p, vx_p, ax_p := p.ps[i], p.vs[i], p.as[i]
		px_q, vx_q, ax_q := q.ps[i], q.vs[i], q.as[i]

		// a is the
		c, b, a := 2*(px_q-px_p), 2*(vx_q-vx_p), (ax_q - ax_p)

		if a == 0 && b == 0 {
			// Will always be at separate spot
			if c != 0 {
				return nil
			}
			// Otherwise, c is zero and they will be at the same spot forever
			continue
		}

		var roots []int
		if a == 0 {
			// Then just a linear equation
			// 0 = (px_q - px_p) + (vx_q - vx_p) * t
			// t = (px_p - px_q) / (vx_q - vx_p)
			t := (px_p - px_q) / (vx_q - vx_p)
			// Make sure no fractions
			if ((px_q - px_p) + (vx_q-vx_p)*t) != 0 {
				return nil
			}
			roots = append(roots, t)
		} else {
			// quadratic
			determinant := b*b - 4*a*c
			if determinant < 0 {
				return nil
			}
			if !maths.IsSquare(determinant) {
				return nil
			}
			// (-b += sqrt(b*b - 4 * a * c)) / 2 * a
			sq := maths.IntSquareRoot(determinant)
			roots = append(roots, (-b+sq)/(2*a), (-b-sq)/(2*a))
		}

		if ts == nil {
			ts = maths.NewSimpleSet(roots...)
		} else {
			if len(ts) == 1 && len(roots) == 1 && ts[roots[0]] {
				fmt.Println("YUP", p, q, ts)
			}
			ts = maths.Intersection(ts, maths.NewSimpleSet(roots...))
		}
	}

	var r []int
	for t := range ts {
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
	o.Stdoutln(best.BestIndex())

	particles = []*particle{particles[346], particles[349]}

	// Calculate all of the intersections
	intersections := map[int][]*intersection{}
	for i, p := range particles {
		for _, q := range particles[i+1:] {
			ts := p.intersects(q)
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

	o.Stdoutln(len(particles) - len(destroyed))
}

func (d *day20) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}

/*
{idx=346; ps=[-939 -4363 2031]; vs=[35 140 -73]; as=[0 1 0]}
{idx=349; ps=[2393 1125 2717]; vs=[-84 -56 -83]; as=[0 1 -1]}
2293 - 84 * 28
*/
