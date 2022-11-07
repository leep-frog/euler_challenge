package eulerchallenge

import (
	"fmt"
	"math/big"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func P152() *problem {
	return intInputNode(152, func(o command.Output, n int) {
		// First, get all inverse square sums (with at most 4 elements)
		// that add up to the same value.
		ctx := &sumEquivalentContext{n, 4, map[string][]*sumEquivalent{}}
		bfs.DFS([]*sumEquivalent{{nil, newRat(0, 1)}}, ctx)

		// Construct edges between equivalent integer sets.
		var edges []*seEdge
		for _, vs := range ctx.m {
			for i, iVals := range vs {
				iSet := maths.NewSimpleSet(iVals.ints...)
				for j := i + 1; j < len(vs); j++ {
					jVals := vs[j]
					jSet := maths.NewSimpleSet(jVals.ints...)
					var overlap bool
					for iv := range iSet {
						if jSet[iv] {
							overlap = true
							break
						}
					}
					if overlap {
						continue
					}
					// Create both directed edges between each equivalent pair.
					edges = append(edges,
						&seEdge{slices.Clone(iVals.ints), slices.Clone(jVals.ints), maps.Clone(iSet), maps.Clone(jSet)},
						&seEdge{slices.Clone(jVals.ints), slices.Clone(iVals.ints), maps.Clone(jSet), maps.Clone(iSet)},
					)
				}
			}
		}

		// Using all of the constructed edges above, find all nodes
		// that are equivalent to the provided example.
		init := []*node152{{
			[]int{3, 4, 5, 7, 12, 15, 20, 28, 35},
		}}
		ctx2 := &ctx152{0, edges}
		bfs.DFS(init, ctx2)
		o.Stdoutln(ctx2.count)
	}, []*execution{
		{
			args:     []string{"80"},
			want:     "301",
			estimate: 5,
		},
	})
}

// Below are DFS types for finding all inverse square sums
// with the same value (1/2).
type node152 struct {
	ints []int
}

type ctx152 struct {
	count int
	edges []*seEdge
}

func (n *node152) Code(*ctx152) string {
	return n.String()
}

func (n *node152) String() string {
	return fmt.Sprintf("%v", n.ints)
}

func (n *node152) Done(ctx *ctx152) bool {
	ctx.count++
	return false
}

func (n *node152) AdjacentStates(ctx *ctx152) []*node152 {
	var r []*node152
	for _, edge := range ctx.edges {
		if next, ok := edge.Convert(n.ints); ok {
			r = append(r, &node152{next})
		}
	}
	return r
}

type seEdge struct {
	from    []int
	to      []int
	fromSet map[int]bool
	toSet   map[int]bool
}

func (se *seEdge) Convert(input []int) ([]int, bool) {
	remaining := maps.Clone(se.fromSet)
	for _, v := range input {
		delete(remaining, v)
		if se.toSet[v] {
			return nil, false
		}
	}

	if len(remaining) != 0 {
		return nil, false
	}

	r := slices.Clone(se.to)
	for _, v := range input {
		if !se.fromSet[v] {
			r = append(r, v)
		}
	}
	slices.Sort(r)
	return r, true
}

func (se *seEdge) String() string {
	return fmt.Sprintf("%v %v", se.from, se.to)
}

// Below are DFS contexts and nodes for finding all equivalent sets
// of inverse square sums.
type sumEquivalentContext struct {
	n      int
	maxLen int
	m      map[string][]*sumEquivalent
}

type sumEquivalent struct {
	ints []int
	sum  *big.Rat
}

func (se *sumEquivalent) Code(*sumEquivalentContext) string {
	return se.String()
}

func (se *sumEquivalent) String() string {
	return fmt.Sprintf("%v", se.ints)
}

func (se *sumEquivalent) Done(ctx *sumEquivalentContext) bool {
	ctx.m[se.sum.String()] = append(ctx.m[se.sum.String()], se)
	return false
}

func (se *sumEquivalent) AdjacentStates(ctx *sumEquivalentContext) []*sumEquivalent {
	if len(se.ints) >= ctx.maxLen {
		return nil
	}
	var r []*sumEquivalent
	// sum from x=2 to 80 of 1/x^2 = 0.63251
	// sum from x=2 to 80 of 1/x^2 = 0.63251 - 1/4 = 0.63251 - 0.25
	//                               0.38251
	// Therefore, we always need a 2 so we can just look for
	// sum equivalents starting at 3
	start := 3
	if len(se.ints) > 0 {
		start = se.ints[len(se.ints)-1] + 1
	}
	for k := start; k <= ctx.n; k++ {
		r = append(r, &sumEquivalent{
			append(slices.Clone(se.ints), k),
			ratAdd(se.sum, newRat(1, k*k)),
		})
	}
	return r
}
