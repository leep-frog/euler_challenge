package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P87() *problem {
	return intInputNode(87, func(o command.Output, n int) {
		p := generator.Primes()
		unique := map[int]bool{}
		for a := 0; ; a++ {
			av := p.Nth(a) * p.Nth(a)
			if av >= n {
				break
			}
			for b := 0; ; b++ {
				bv := p.Nth(b)*p.Nth(b)*p.Nth(b) + av
				if bv >= n {
					break
				}
				for c := 0; ; c++ {
					cv := p.Nth(c)*p.Nth(c)*p.Nth(c)*p.Nth(c) + bv
					if cv > n {
						break
					}
					unique[cv] = true
				}
			}
		}
		o.Stdoutln(len(unique))
	}, []*execution{
		{
			args:     []string{"50000000"},
			want:     "1097343",
			estimate: 0.3,
		},
		{
			args: []string{"50"},
			want: "4",
		},
	})
}
