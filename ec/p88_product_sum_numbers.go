package eulerchallenge

import (
	"sort"
	"fmt"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
	"github.com/leep-frog/euler_challenge/maths"
)

func P88() *problem {
	return intInputNode(88, func(o command.Output, n int) {

		var sum int
		got := map[int]bool{}
		for k := 2; k <= n; k++ {
			initStates := []*p88{{
				m: map[int]int{
					1: k,
				},
				product: 1,
				sum: k,
				k: k,
			}}
			ps, _ := bfs.ShortestWeightedPath(initStates, 0)
			final := ps[0].product
			fmt.Println(k, final)
			if !got[final] {
				sum += final
				got[final] = true
			}
		}
		o.Stdoutln(sum)
	})
}

type p88 struct {
	// map from int to count
	m       map[int]int
	product int
	sum     int
	k int
}

func (p *p88) Code(*bfs.Context[int, *p88]) string {
	return p.String()
}

func (p *p88) String() string {
	var ks []int
	for k := range p.m {
		ks = append(ks, k)
	}
	sort.Ints(ks)
	var sl []string
	for _, k := range ks {
		sl = append(sl, fmt.Sprintf("%d:%d", k, p.m[k]))
	}
	return "{" + fmt.Sprintf("%d", p.product) + " " + strings.Join(sl, ", ") + "}"
}

func (p *p88) Done(*bfs.Context[int, *p88]) bool {
	return p.sum == p.product
}

func (p *p88) Distance(*bfs.Context[int, *p88]) int {
	return maths.Max(p.product*p.product, p.sum*p.sum)
	//return maths.Max(p.product, p.sum)
	//return p.sum
}

func (p *p88) AdjacentStates(*bfs.Context[int, *p88]) []*p88 {
	// Increment each value
	var states []*p88
	for k, v := range p.m {
		c := maths.CopyMap(p.m)
		if v == 1 {
			delete(c, k)
		} else {
			c[k]--
		}
		c[k+1]++
		
		new := &p88{
			m: c,
			product: (p.product / k) * (k + 1),
			sum: p.sum + 1,
			k: p.k,
		}
		if new.product > p.k*p.k {
			continue
		}
		states = append(states, new)
	}
	return states
}
