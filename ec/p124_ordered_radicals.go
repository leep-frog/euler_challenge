package eulerchallenge

import (
	"fmt"
	"sort"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P124() *problem {
	return noInputNode(124, func(o command.Output) {
		var rs []*radical
		g := generator.Primes()
		for i := 1; i <= 100_000; i++ {
			rs = append(rs, newRadical(i, g))
		}
		sort.SliceStable(rs, func(i, j int) bool {
			if rs[i].rad_n != rs[j].rad_n {
				return rs[i].rad_n < rs[j].rad_n
			}
			return rs[i].n < rs[j].n
		})
		o.Stdoutln(rs[9999].n)
	}, &execution{
		want:     "21417",
		estimate: 0.5,
	})
}

type radical struct {
	n     int
	rad_n int
}

func (r *radical) String() string {
	return fmt.Sprintf("%d:%d", r.n, r.rad_n)
}

var (
	radicalCache = []int{}
	pnc          = false
)

func calcRadical(n int, g *generator.Prime) int {
	for len(radicalCache) < n {
		prod := 1
		for f := range g.PrimeFactors(len(radicalCache) + 1) {
			prod *= f
		}
		radicalCache = append(radicalCache, prod)
	}
	return radicalCache[n-1]
}

func newRadical(n int, g *generator.Prime) *radical {
	return &radical{n, calcRadical(n, g)}
}
