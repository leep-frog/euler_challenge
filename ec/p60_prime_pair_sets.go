package eulerchallenge

import (
	"fmt"
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

type pairCtx struct {
	adjPairs map[int][]*primePair
	edges    map[int]map[int]bool
}

func P60() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=60"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

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
					if generator.IsPrime(r, p) && generator.IsPrime(l, p) {
						pairs[spn] = append(pairs[spn], &primePair{npn, n})
						pairs[npn] = append(pairs[npn], &primePair{spn, n})
						maths.Set(edges, spn, npn, true)
						maths.Set(edges, npn, spn, true)
					}
				}
			}

			ctx := &pairCtx{pairs, edges}
			path := bfs.CompleteSets(primes, ctx, n)
			o.Stdoutln(maths.SumType(path), path)
		}),
	)
}

type primePair struct {
	prime int
	n     int
}

func (p *primePair) ToInt() int {
	return p.prime
}

func (p *primePair) Offset(_ *bfs.Context[*pairCtx, *primePair]) int {
	return p.prime
}

func (p *primePair) String() string {
	return fmt.Sprintf("%d", p.prime)
}

func (p *primePair) Code(_ *bfs.Context[*pairCtx, *primePair]) string {
	return strconv.Itoa(p.prime)
}

func (p *primePair) AdjacentStates(ctx *bfs.Context[*pairCtx, *primePair]) []*primePair {
	return ctx.GlobalContext.adjPairs[p.prime]
}

func (p *primePair) BiggerThan(that *primePair) bool {
	return p.prime > that.prime
}

func (p *primePair) HasEdge(ctx *bfs.Context[*pairCtx, *primePair], that *primePair) bool {
	return ctx.GlobalContext.edges[p.prime][that.prime]
}
