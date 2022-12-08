package eulerchallenge

import (
	"strconv"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P108() *problem {
	return intInputNode(108, func(o command.Output, n int) {
		best := maths.Smallest[int, int]()
		initStates := []*diophantineReciprocals{{
			[]int{1},
			generator.Primes(),
			n,
		}}
		bfs.ContextualShortestPath[bfs.Int](initStates, best)
		o.Stdoutln(best.Best())
	}, []*execution{
		{
			args: []string{"4000000"},
			want: "9350130049860600",
		},
		{
			args: []string{"100"},
			want: "1260",
		},
		{
			args: []string{"1000"},
			want: "180180",
		},
	})
}

type diophantineReciprocals struct {
	// parts is a decreasing set of positive integers
	parts []int
	g     *generator.Prime
	n     int
}

func (dr *diophantineReciprocals) intValue() int {
	v := 1
	for idx, p := range dr.parts {
		prime := dr.g.Nth(idx) // idx-th prime
		v *= maths.Pow(prime, p)
	}
	return v
}

func (dr *diophantineReciprocals) numFractions() int {
	v := 1
	for _, p := range dr.parts {
		v *= 2*p + 1
	}
	return (v + 1) / 2
}

func (dr *diophantineReciprocals) String() string {
	var r []string
	for _, p := range dr.parts {
		r = append(r, strconv.Itoa(p))
	}
	return strings.Join(r, "_")
}

func (dr *diophantineReciprocals) Code(*maths.Bester[int, int]) string {
	return dr.String()
}

func (dr *diophantineReciprocals) Done(*maths.Bester[int, int]) bool {
	// AdjacentStates will eventually return nothing, so we just check all states
	return false
}

func (dr *diophantineReciprocals) Distance(best *maths.Bester[int, int]) bfs.Int {
	nf := dr.numFractions()
	iv := dr.intValue()
	if nf >= dr.n {
		best.IndexCheck(nf, iv)
	}
	return bfs.Int(-1_000 * nf / maths.Sqrt(iv))
}

func (dr *diophantineReciprocals) AdjacentStates(best *maths.Bester[int, int]) []*diophantineReciprocals {
	if best.Set() {
		iv := dr.intValue()
		if iv > best.Best() {
			return nil
		}
	}
	var neighbors []*diophantineReciprocals
	for i, p := range dr.parts {
		if i != 0 && p+1 > dr.parts[i-1] {
			continue
		}
		arr := make([]int, len(dr.parts), len(dr.parts))
		copy(arr, dr.parts)
		arr[i]++
		neighbors = append(neighbors, &diophantineReciprocals{arr, dr.g, dr.n})
	}

	arr := make([]int, len(dr.parts), len(dr.parts))
	copy(arr, dr.parts)
	arr = append(arr, 1)
	neighbors = append(neighbors, &diophantineReciprocals{arr, dr.g, dr.n})
	return neighbors
}

/*func (dr *diophantineReciprocals) Code(*bfs.Context[*maths.Bester[int, int], *diophantineReciprocals]) string {
	return dr.String()
}

func (dr *diophantineReciprocals) Done(ctx *bfs.Context[*maths.Bester[int, int], *diophantineReciprocals]) bool {
	// AdjacentStates will eventually return nothing, so we just check all states
	return false
}

func (dr *diophantineReciprocals) Distance(ctx *bfs.Context[*maths.Bester[int, int], *diophantineReciprocals]) int {
	nf := dr.numFractions()
	iv := dr.intValue()
	if nf >= dr.n {
		ctx.GlobalContext.IndexCheck(nf, iv)
	}
	return -1_000*nf/maths.Sqrt(iv)
}

func (dr *diophantineReciprocals) AdjacentStates(ctx *bfs.Context[*maths.Bester[int, int], *diophantineReciprocals]) []*diophantineReciprocals {
	if ctx.GlobalContext.Set() {
		iv := dr.intValue()
		if iv > ctx.GlobalContext.Best() {
			return nil
		}
	}
	var neighbors []*diophantineReciprocals
	for i, p := range dr.parts {
		if i != 0 && p + 1 > dr.parts[i-1] {
			continue
		}
		arr := make([]int, len(dr.parts), len(dr.parts))
		copy(arr, dr.parts)
		arr[i]++
		neighbors = append(neighbors, &diophantineReciprocals{arr, dr.g, dr.n})
	}

	arr := make([]int, len(dr.parts), len(dr.parts))
	copy(arr, dr.parts)
	arr = append(arr, 1)
	neighbors = append(neighbors, &diophantineReciprocals{arr, dr.g, dr.n})
	return neighbors
}
*/
