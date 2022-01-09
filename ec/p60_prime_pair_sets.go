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
	adjPairs map[int][]*bfs.AdjacentState[*primePair]
	edges map[int]map[int]bool
}

func P60() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=60"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)
			_ = n

			// Get all pairs and then find cycle!
			p := generator.Primes()
			pairs := map[int][]*bfs.AdjacentState[*primePair]{}
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
					if spn == 31 && npn == 391 {
						o.Stdoutln("hello")
					}
					if generator.IsPrime(r, p) && generator.IsPrime(l, p) {
						pairs[spn] = append(pairs[spn], &bfs.AdjacentState[*primePair]{&primePair{npn, n}, npn})
						pairs[npn] = append(pairs[npn], &bfs.AdjacentState[*primePair]{&primePair{spn, n}, spn})
						//pairs[spn] = append(pairs[spn], &primePair{npn, n})
						//pairs[npn] = append(pairs[npn], &primePair{spn, n})
						maths.Set(edges, spn, npn, true)
						maths.Set(edges, npn, spn, true)
					}
				}
			}

			ctx := &pairCtx{pairs, edges}
			path := bfs.CompleteSets(primes, ctx, n)
			o.Stdoutln(maths.SumType(path), path)

			/*for i := 1; i < 2; i++ {
				fmt.Println("===========")
				// TODO: return all states that were checked?
				cycle, length := bfs.CyclePath(&primePair{p.Nth(i), n}, pairs)
				if length > 0 {
					o.Stdoutln("yippee", length, cycle)
					for cur := range cycle {
						o.Stdoutln(cur)
					}
					return
				}
			}*/

			/*pairs := map[int]map[int]bool{}
			p := generator.Primes()
			for start := 0; p.Nth(start) < 10_000; start++ {
				for next := start + 1; p.Nth(next) < 10_000; next++ {
					spn, npn := p.Nth(start), p.Nth(next)
					sp := strconv.Itoa(spn)
					np := strconv.Itoa(npn)
					r, l := parse.Atoi(sp+np), parse.Atoi(np+sp)
					if generator.IsPrime(r, p) && generator.IsPrime(l, p) {
						maths.Set(pairs, spn, npn, true)
						maths.Set(pairs, npn, spn, true)
					}
				}
			}
			o.Stdoutln(pairs)*/
		}),
	)
}

type primePair struct {
	prime int
	n int
}

func (p *primePair) ToInt() int {
	return p.prime
}

func (p *primePair) Value() int {
	return p.prime
}

func (p *primePair) String() string {
	return fmt.Sprintf("primePair: %d", p.prime)
}

func (p *primePair) Code(_ *bfs.Context[*pairCtx, *primePair]) string {
	return strconv.Itoa(p.prime)
}

/*func (p *primePair) AdjacentStates(ctx *bfs.Context[map[int][]*primePair, *primePair]) []*primePair {
	return ctx.GlobalContext[p.prime]
}*/

func (p *primePair) AdjacentStates(ctx *bfs.Context[*pairCtx, *primePair]) []*bfs.AdjacentState[*primePair] {
	return ctx.GlobalContext.adjPairs[p.prime]
}

func (p *primePair) BiggerThan(that *primePair) bool {
	return p.prime > that.prime
}

func (p *primePair) HasEdge(ctx *bfs.Context[*pairCtx, *primePair], that *primePair) bool {
	return ctx.GlobalContext.edges[p.prime][that.prime]
}

func (p *primePair) DoneCycle(ctx *bfs.Context[map[int][]*primePair, *primePair]) bool {
	// Iterate back until we get the current prime
	var count int
	fmt.Println("DC START")
	if ctx == nil || ctx.StateValue == nil {
		return false
	}
	for cur := ctx.StateValue.Prev(); cur != nil; cur = cur.Prev() {
		count++
		fmt.Println(cur)
		if cur.State().prime == p.prime {
			fmt.Println("DC END", count, p.n)
			return count >= p.n
		}
	}
	return false
}

/*func (p *primePair) inList(ctx *Context[map[int]int, *primePair]) bool {
	for cur := ctx.StateValue; cur != nil; cur = cur.Prev() {
		if cur.State().prime == p.prime {
			return true
		}
	}
	return false
}

func (p *primePair) Code(ctx *Context[map[int]int, T]) string {
	// Always return a unique number because we are looking for cycles
	ctx.GlobalContext[p.prime]++
	return fmt.Sprintf("%d_%d", p.prime, ctx.GlobalContext[p.prime])
}

func (p *primePair) Done(*Context[M, T]) bool {

}

func (p *primePair) AdjacentStates(*Context[M, T]) bool {
	// Only get adjacent states if this is the first instance of intera
}
*/
