package eulerchallenge

import (
	"github.com/leep-frog/command"
)

func P60() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=60"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			//n := d.Int(N)

			// Get all pairs and then find cycle!
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
