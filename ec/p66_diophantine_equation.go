package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P66() *problem {
	// See https://mathworld.wolfram.com/PellEquation.html for math info
	return intInputNode(66, func(o command.Output, n int) {
		best := maths.LargestT[int, *maths.Int]()
		for D := 2; D <= n; D++ {
			start, period := maths.SquareRootPeriod(D)
			if len(period) == 0 {
				continue
			}
			as := maths.Biggify(append([]int{start}, period...))
			r := len(as) - 2
			as = append(as, as[1:]...) // Needed because we need 2*r + 1 index for odd periods

			p := []*maths.Int{as[0], as[0].Times(as[1]).Plus(maths.One())}
			q := []*maths.Int{maths.One(), as[1]}
			for idx := 2; idx < len(as); idx++ {
				p = append(p, as[idx].Times(p[idx-1]).Plus(p[idx-2]))
				q = append(q, as[idx].Times(q[idx-1]).Plus(q[idx-2]))
			}

			var x, y *maths.Int
			if ((len(as)+1)/2)%2 == 1 {
				// r is odd
				x, y = p[r], q[r]
			} else {
				x, y = p[2*r+1], q[2*r+1]
			}

			if x.Times(x).Minus(y.Times(y).Times(maths.NewInt(D))).NEQ(maths.One()) {
				o.Terminatef("does not satisfy equation")
			}

			best.IndexCheck(D, x)
		}
		o.Stdoutln(best.BestIndex())
	}, []*execution{
		{
			args: []string{"1000"},
			want: "661",
		},
		{
			args: []string{"7"},
			want: "5",
		},
	})
}
