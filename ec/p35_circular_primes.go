package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func P35() *problem {
	return intInputNode(35, func(o command.Output, n int) {
		checked := map[string]bool{}
		unique := map[string]bool{}
		p := generator.Primes()
		for i := 0; p.Nth(i) < n; i++ {
			pn := p.Nth(i)
			spn := strconv.Itoa(pn)
			if checked[spn] {
				continue
			}
			var digits []string
			for j := 0; j < len(spn); j++ {
				digits = append(digits, spn[j:j+1])
			}

			allPrime := true
			rots := maths.Rotations(digits)
			for _, rot := range rots {
				checked[rot] = true
				if !p.Contains(parse.Atoi(rot)) {
					allPrime = false
				}
			}
			if allPrime {
				for _, rot := range rots {
					unique[rot] = true
				}
			}
		}
		o.Stdoutln(len(unique))
	}, []*execution{
		{
			args:     []string{"1000000"},
			want:     "55",
			estimate: 0.25,
		},
		{
			args: []string{"100"},
			want: "13",
		},
	})
}
