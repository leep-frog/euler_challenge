package p75

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P75() *ecmodels.Problem {
	return ecmodels.IntInputNode(75, func(o command.Output, L int) {
		counts := map[int]int{}
		tg := generator.RightTriangleGenerator().Iterator()
		for t := tg.Next(); t.GuaranteedMinimumPerimeter() <= L; t = tg.Next() {
			p := t.Perimeter()
			for l := p; l <= L; l += p {
				counts[l]++
			}
		}
		var count int
		for _, v := range counts {
			if v == 1 {
				count++
			}
		}
		o.Stdoutln(count)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1500000"},
			Want:     "161667",
			Estimate: 0.5,
		},
		{
			Args: []string{"48"},
			Want: "6",
		},
	})
}
