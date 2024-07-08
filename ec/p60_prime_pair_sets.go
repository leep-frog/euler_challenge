package eulerchallenge

import (
	"fmt"
	"strconv"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

type pairCtx struct {
	adjPairs map[int][]*primePair
	edges    map[int]map[int]bool
}

func P60() *problem {
	return intInputNode(60, func(o command.Output, n int) {
		// Get all pairs and then find cycle!
		p := generator.Primes()
		pairs := map[int][]*primePair{}
		edges := map[int]map[int]bool{}
		primes := []*primePair{}
		for start := 0; p.Nth(start) < 10000; start++ {
			spn := p.Nth(start)
			primes = append(primes, &primePair{spn, n})
			for next := start + 1; p.Nth(next) < 10000; next++ {
				npn := p.Nth(next)
				sp := strconv.Itoa(spn)
				np := strconv.Itoa(npn)
				r, l := parse.Atoi(sp+np), parse.Atoi(np+sp)
				if p.Contains(r) && p.Contains(l) {
					pairs[spn] = append(pairs[spn], &primePair{npn, n})
					pairs[npn] = append(pairs[npn], &primePair{spn, n})
					maths.Insert(edges, spn, npn, true)
					maths.Insert(edges, npn, spn, true)
				}
			}
		}

		ctx := &pairCtx{pairs, edges}
		_, dist := bfs.ContextDistancePathSearch[string, bfs.Int](ctx, primes, bfs.CheckDuplicates())
		o.Stdoutln(dist)
	}, []*execution{
		{
			args:     []string{"4"},
			want:     "792",
			estimate: 0.5,
		},
		{
			args:     []string{"5"},
			want:     "26033",
			estimate: 1,
		},
	})
}

type primePair struct {
	prime int
	n     int
}

func (p *primePair) ToInt() int {
	return p.prime
}

func (p *primePair) Distance(*pairCtx, bfs.Path[*primePair]) bfs.Int {
	return bfs.Int(p.prime)
}

func (p *primePair) String() string {
	return fmt.Sprintf("%d", p.prime)
}

func (p *primePair) Done(m *pairCtx, path bfs.Path[*primePair]) bool {
	if path.Len() != p.n {
		return false
	}
	// Done if we circle back to the front
	return m.edges[p.prime][path.Fetch()[0].prime]
}

func (p *primePair) Code(*pairCtx, bfs.Path[*primePair]) string {
	return strconv.Itoa(p.prime)
}

func (p *primePair) AdjacentStates(m *pairCtx, path bfs.Path[*primePair]) []*primePair {
	if path.Len() >= p.n {
		return nil
	}

	var r []*primePair
	ps := path.Fetch()
	for _, pp := range m.adjPairs[p.prime] {
		add := true
		for _, parent := range ps {
			if !m.edges[pp.prime][parent.prime] {
				add = false
				break
			}
		}
		if add {
			r = append(r, pp)
		}
	}
	return r
}
