package eulerchallenge

import (
	"strconv"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P51() *problem {
	return intInputNode(51, func(o command.Output, n int) {
		primes := generator.Primes()
		m := map[string]map[int]bool{}
		for i, pn := 0, primes.Nth(0); ; i, pn = i+1, primes.Nth(i+1) {
			if pn < 10 {
				continue
			}
			checked := map[int]bool{}
			digits := maths.Digits(pn)
			for _, d := range digits {
				if checked[d] {
					continue
				}
				checked[d] = true

				var positions []int
				for i, d2 := range digits {
					if d2 == d {
						positions = append(positions, i)
					}
				}

				pnStr := strings.Split(strconv.Itoa(pn), "")
				cp := make([]string, len(pnStr))
				copy(cp, pnStr)
				for _, s := range maths.Sets(positions) {
					for _, pos := range s {
						pnStr[pos] = "_"
					}
					coded := strings.Join(pnStr, "")
					if m[coded] == nil {
						m[coded] = map[int]bool{}
					}
					m[coded][d] = true
					if len(m[coded]) >= n {
						min := 10
						for k := range m[coded] {
							min = maths.Min(min, k)
						}
						o.Stdoutln(coded, strings.ReplaceAll(coded, "_", strconv.Itoa(min)))
						return
					}

					for _, pos := range s {
						pnStr[pos] = cp[pos]
					}
				}
			}
		}
	}, []*execution{
		{
			args:     []string{"8"},
			want:     "_2_3_3 121313",
			estimate: 1,
		},
		{
			args: []string{"7"},
			want: "56__3 56003",
		},
		{
			args: []string{"6"},
			want: "_3 13",
		},
	})
}
