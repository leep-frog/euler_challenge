package p142

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P142() *ecmodels.Problem {
	return ecmodels.IntInputNode(142, func(o command.Output, n int) {
		m := map[int][][]int{}
		spg := generator.SmallPowerGenerator(2)
		for iter, s1 := generator.SmallPowerGenerator(2).Start(1); ; s1 = iter.Next() {
			for iter, s2 := generator.SmallPowerGenerator(2).Start(1); s2 < s1; s2 = iter.Next() {
				if (s1%2 == 0) != (s2%2 == 0) {
					continue
				}
				x := (s1 + s2) / 2
				y := x - s2
				for _, opt := range m[x] {
					z := x - opt[1]
					if spg.Contains(y+z) && spg.Contains(maths.Abs(y-z)) {
						o.Stdoutln(x + y + z)
						return
					}
				}
				m[x] = append(m[x], []int{s1, s2})
			}
		}
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1"},
			Want:     "1006193",
			Estimate: 0.25,
		},
	})
}
