package p120

import (
	"fmt"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/maths"
)

func P120() *ecmodels.Problem {
	return ecmodels.NoInputNode(120, func(o command.Output) {
		var sum int
		for a := 3; a <= 1000; a++ {
			a2 := a * a
			max := maths.Largest[int, int]()
			has := map[string]bool{}
			for left, right := (a - 1), (a + 1); ; left, right = (left*(a-1))%a2, (right*(a+1))%a2 {
				code := fmt.Sprintf("%d:%d", left, right)
				if has[code] {
					break
				}
				has[code] = true
				max.Check((left + right) % a2)
			}
			sum += max.Best()
		}
		o.Stdoutln(sum)
	}, &ecmodels.Execution{
		Want:     "333082500",
		Estimate: 0.5,
	})
}
