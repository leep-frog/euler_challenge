package eulerchallenge

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P123() *problem {
	return intInputNode(123, func(o command.Output, n int) {
		minRem := maths.Pow(10, n)
		g := generator.Primes()
		for k := 0; ; k++ {
			pn := g.Nth(k)
			pn2 := pn * pn
			if pn2 <= minRem {
				continue
			}
			left, right := 1, 1
			for i := 0; i <= k; i++ {
				left = (left * (pn - 1)) % pn2
				right = (right * (pn + 1)) % pn2
			}
			if (left+right)%pn2 > minRem {
				o.Stdoutln(k + 1)
				return
			}
		}
	}, []*execution{
		{
			args:     []string{"10"},
			want:     "21035",
			estimate: 1.5,
		},
		{
			args:     []string{"9"},
			want:     "7037",
			estimate: 0.25,
		},
	})
}
