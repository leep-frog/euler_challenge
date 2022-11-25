package eulerchallenge

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/generator"
)

type context243 struct {
	f   *fraction.Fraction
	p   *generator.Generator[int]
	max int
}

type node243 struct {
	numer int
	denom int
}

func (n *node243) Code(ctx *context243) string {
	return n.String()
}

func (n *node243) String() string {
	return fmt.Sprintf("%d", n.denom+1)
}

func (n *node243) Done(ctx *context243) bool {
	if n.denom > ctx.max {
		fmt.Println(n.denom)
		ctx.max += 10_000
	}
	return fraction.New(n.numer, n.denom).LT(ctx.f)
}

func (n *node243) Distance(ctx *context243) bfs.Int {
	return bfs.Int(n.denom + 1)
}

func (n *node243) AdjacentStates(ctx *context243) []*node243 {
	var r []*node243
	for i := 0; i < 100; i++ {
		k := (n.denom + 1) * ctx.p.Nth(i)

		r = append(r, &node243{cnt243(k, ctx.p), k - 1})
	}
	return r
}

func cnt243(n int, p *generator.Generator[int]) int {
	cnt := 1
	for i := 2; i < n; i++ {
		if !generator.Coprimes(n, i, p) {
			//fmt.Println(j, i)
			cnt++
		}
	}
	return cnt
}

func P243() *problem {
	return intInputNode(243, func(o command.Output, n int) {
		p := generator.Primes()

		mf := fraction.New(4, 10)
		if n > 1 {
			mf = fraction.New(15499, 94744)
		}

		// After implementing generator.ResilienceCount, I noticed the following numbers
		// were the incremental bests:
		// (3 values)  2, 4, 6
		// (5 values)  6, 12, 18, 24, 30
		// (7 values)  30, 60, 90, 120, 150, 180, 210
		// (11 values) 210, 420, 630, 840, 1050, 1260, 1470, 1680, 1890, 2100, 2310
		// (13 values) 2310, ...

		// So, we can exclusively check numbers that are in this pattern.
		i := 2
		for primeIdx := 1; ; primeIdx++ {
			prime := p.Nth(primeIdx)
			offset := i
			for j := 0; j < prime-1; j++ {
				i += offset
				f := fraction.New(generator.ResilienceCount(i, p), i-1)
				if f.LT(mf) {
					o.Stdoutln(i)
					return
				}
			}
		}
	}, []*execution{
		{
			args: []string{"1"},
			want: "12",
		},
		{
			args: []string{"1"},
			want: "892371480",
		},
	})
}
