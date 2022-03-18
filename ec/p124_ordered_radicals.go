package eulerchallenge

import (
	"fmt"
	"sort"

	"github.com/leep-frog/command"
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
	})
}

type radical struct {
	n     int
	rad_n int
}

func (r *radical) String() string {
	return fmt.Sprintf("%d:%d", r.n, r.rad_n)
}

func newRadical(n int, g *generator.Generator[int]) *radical {
	prod := 1
	for f, _ := range generator.PrimeFactors(n, g) {
		prod *= f
	}
	return &radical{n, prod}
}
