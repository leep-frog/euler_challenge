package eulerchallenge

import (
	"fmt"
	"sort"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P118() *problem {
	return intInputNode(118, func(o command.Output, n int) {
		ctx := &context118{map[string]bool{}, generator.Primes()}
		perms := maths.Permutations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		for _, p := range perms {
			bfs.DFS([]*pandigitalOrdering{{p, nil}}, ctx)
		}
		o.Stdoutln(len(ctx.validSets))
	}, []*execution{
		{
			args:     []string{"1"},
			want:     "44680",
			estimate: 5,
		},
	})
}

type pandigitalOrdering struct {
	perm   []int
	breaks []int
}

type context118 struct {
	validSets map[string]bool
	primes    *generator.Generator[int]
}

func (po *pandigitalOrdering) copy() *pandigitalOrdering {
	return &pandigitalOrdering{
		maths.CopySlice(po.perm),
		maths.CopySlice(po.breaks),
	}
}

func (po *pandigitalOrdering) parts() []int {
	var r []int
	var start int
	for _, b := range po.breaks {
		r = append(r, maths.FromDigits(po.perm[start:b]))
		start = b
	}
	return append(r, maths.FromDigits(po.perm[start:]))
}

func (po *pandigitalOrdering) Code(*context118) string {
	return fmt.Sprintf("%v", po.parts())
}

func (po *pandigitalOrdering) Done(ctx *context118) bool {
	parts := po.parts()
	for _, p := range parts {
		if !generator.IsPrime(p, ctx.primes) {
			return false
		}
	}
	sort.Ints(parts)
	ctx.validSets[fmt.Sprintf("%v", parts)] = true
	return false
}

func (po *pandigitalOrdering) AdjacentStates(ctx *context118) []*pandigitalOrdering {
	var r []*pandigitalOrdering
	start := 0
	if len(po.breaks) > 0 {
		start = po.breaks[len(po.breaks)-1]
	}
	for i := start + 1; i <= len(po.perm)-1; i++ {
		cp := po.copy()
		cp.breaks = append(cp.breaks, i)
		if generator.IsPrime(maths.FromDigits(cp.perm[start:i]), ctx.primes) {
			r = append(r, cp)
		}
	}
	return r
}
