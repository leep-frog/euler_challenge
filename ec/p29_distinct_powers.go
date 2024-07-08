package eulerchallenge

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
)

func P29() *problem {
	return intInputNode(29, func(o command.Output, n int) {
		p := generator.Primes()

		unique := map[string]bool{}
		for a := 2; a <= n; a++ {
			for b := 2; b <= n; b++ {
				scaled := map[int]int{}
				for k, v := range p.PrimeFactors(a) {
					scaled[k] = v * b
				}
				unique[polyCode(scaled)] = true
			}
		}

		o.Stdoutln(len(unique))
	}, []*execution{
		{
			args: []string{"100"},
			want: "9183",
		},
		{
			args: []string{"5"},
			want: "15",
		},
	})
}

func polyCode(m map[int]int) string {
	var a []int
	for k := range m {
		a = append(a, k)
	}
	sort.Ints(a)

	var r []string
	for _, k := range a {
		r = append(r, fmt.Sprintf("(%d, %d)", k, m[k]))
	}
	return strings.Join(r, ", ")
}
